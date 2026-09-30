package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"server/internal/global"
	"server/internal/model/common"
)

const (
	defaultRate    = 50.0             // requests per second when config is unset
	defaultBurst   = 100              // bucket capacity when config is unset
	gcInterval     = 1 * time.Minute  // how often stale visitors are evicted
	visitorMaxIdle = 10 * time.Minute // evict visitors silent longer than this
)

// visitor holds a per-IP token bucket.
type visitor struct {
	tokens   float64   // current tokens (fractional refill)
	last     time.Time // last refill time
	lastSeen time.Time // last activity, used for eviction
}

// IPRateLimiter is an in-memory per-IP token-bucket limiter.
// Buckets start full, refill at `rate` tokens per second, and cap at `burst`.
// Idle visitors are evicted lazily (no background goroutine) so stale IPs
// cannot grow the map unboundedly.
type IPRateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	rate     float64
	burst    float64
	maxIdle  time.Duration
	lastGC   time.Time
}

// NewIPRateLimiter creates a limiter. Non-positive values fall back to
// generous defaults (50 req/s, burst 100) so a misconfigured section
// never hard-blocks all traffic.
func NewIPRateLimiter(requestsPerSec float64, burst int) *IPRateLimiter {
	if requestsPerSec <= 0 {
		requestsPerSec = defaultRate
	}
	if burst <= 0 {
		burst = defaultBurst
	}
	return &IPRateLimiter{
		visitors: make(map[string]*visitor),
		rate:     requestsPerSec,
		burst:    float64(burst),
		maxIdle:  visitorMaxIdle,
		lastGC:   time.Now(),
	}
}

// Allow consumes one token for the given IP.
func (l *IPRateLimiter) Allow(ip string) bool {
	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()

	if now.Sub(l.lastGC) >= gcInterval {
		l.gcLocked(now)
		l.lastGC = now
	}

	v, ok := l.visitors[ip]
	if !ok {
		v = &visitor{tokens: l.burst, last: now}
		l.visitors[ip] = v
	}
	v.lastSeen = now

	// Refill tokens accrued since the last visit
	if elapsed := now.Sub(v.last).Seconds(); elapsed > 0 {
		v.tokens = min(l.burst, v.tokens+elapsed*l.rate)
		v.last = now
	}

	if v.tokens < 1 {
		return false
	}
	v.tokens--
	return true
}

// gcLocked evicts visitors idle beyond maxIdle. Eviction resets the bucket,
// so maxIdle must stay well above the refill window for burst capacity.
func (l *IPRateLimiter) gcLocked(now time.Time) {
	for ip, v := range l.visitors {
		if now.Sub(v.lastSeen) > l.maxIdle {
			delete(l.visitors, ip)
		}
	}
}

// Middleware returns a gin handler enforcing the limiter with HTTP 429.
func (l *IPRateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !l.Allow(ip) {
			if global.TD27_LOG != nil {
				global.TD27_LOG.Warn("rate limit exceeded", "ip", ip, "path", c.Request.URL.Path)
			}
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, common.Response{
				Code: common.ERROR_RES,
				Data: gin.H{},
				Msg:  "请求过于频繁，请稍后再试",
			})
			return
		}
		c.Next()
	}
}
