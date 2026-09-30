package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestLimiter(rate float64, burst int) *IPRateLimiter {
	l := NewIPRateLimiter(rate, burst)
	l.maxIdle = time.Minute // keep eviction out of allow/deny assertions
	return l
}

func TestIPRateLimiter_BurstThenDeny(t *testing.T) {
	l := newTestLimiter(100, 3) // high refill rate: no refill noise within the test

	assert.True(t, l.Allow("1.1.1.1"))
	assert.True(t, l.Allow("1.1.1.1"))
	assert.True(t, l.Allow("1.1.1.1"))
	assert.False(t, l.Allow("1.1.1.1"), "must deny once burst is drained")
	assert.False(t, l.Allow("1.1.1.1"))
}

func TestIPRateLimiter_PerIPIsolation(t *testing.T) {
	l := newTestLimiter(100, 1)

	assert.True(t, l.Allow("1.1.1.1"))
	assert.False(t, l.Allow("1.1.1.1"))
	assert.True(t, l.Allow("2.2.2.2"), "another IP must have its own bucket")
}

func TestIPRateLimiter_Refill(t *testing.T) {
	l := newTestLimiter(10, 1)

	assert.True(t, l.Allow("1.1.1.1"))
	assert.False(t, l.Allow("1.1.1.1"))

	// Simulate 1s passing since last refill: rate 10/s => +10 tokens (capped at burst)
	l.mu.Lock()
	v := l.visitors["1.1.1.1"]
	v.last = time.Now().Add(-time.Second)
	l.mu.Unlock()

	assert.True(t, l.Allow("1.1.1.1"), "tokens must refill over time")
}

func TestIPRateLimiter_EvictsIdleVisitors(t *testing.T) {
	l := newTestLimiter(100, 5)
	require.True(t, l.Allow("1.1.1.1"))

	// Visitor idle beyond maxIdle, GC overdue
	l.mu.Lock()
	l.visitors["1.1.1.1"].lastSeen = time.Now().Add(-2 * l.maxIdle)
	l.lastGC = time.Now().Add(-2 * gcInterval)
	l.mu.Unlock()

	assert.True(t, l.Allow("2.2.2.2")) // triggers GC
	l.mu.Lock()
	_, exists := l.visitors["1.1.1.1"]
	count := len(l.visitors)
	l.mu.Unlock()
	assert.False(t, exists, "idle visitor must be evicted")
	assert.Equal(t, 1, count)
}

func TestIPRateLimiter_DefaultsOnBadConfig(t *testing.T) {
	l := NewIPRateLimiter(0, 0)
	assert.Equal(t, defaultRate, l.rate)
	assert.Equal(t, float64(defaultBurst), l.burst)
}

func TestRateLimitMiddleware_429(t *testing.T) {
	gin.SetMode(gin.TestMode)
	l := newTestLimiter(100, 1)
	handler := l.Middleware()

	r := gin.New()
	r.POST("/ping", handler, func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

	// First request passes
	w1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodPost, "/ping", nil)
	req1.RemoteAddr = "9.9.9.9:1234"
	r.ServeHTTP(w1, req1)
	assert.Equal(t, http.StatusOK, w1.Code)

	// Second request from same IP is limited
	w2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/ping", nil)
	req2.RemoteAddr = "9.9.9.9:1234"
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusTooManyRequests, w2.Code)
	assert.Equal(t, "1", w2.Header().Get("Retry-After"))

	var body struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	require.NoError(t, json.Unmarshal(w2.Body.Bytes(), &body))
	assert.Equal(t, 7, body.Code)
	assert.NotEmpty(t, body.Msg)
}
