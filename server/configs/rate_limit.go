package configs

// RateLimit 每IP令牌桶限流配置
type RateLimit struct {
	Enabled    bool    `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Rate       float64 `mapstructure:"rate" json:"rate" yaml:"rate"`                      // 全局每IP每秒请求数
	Burst      int     `mapstructure:"burst" json:"burst" yaml:"burst"`                   // 全局每IP突发容量
	LoginRate  float64 `mapstructure:"login-rate" json:"login-rate" yaml:"login-rate"`    // 登录/验证码每IP每秒请求数
	LoginBurst int     `mapstructure:"login-burst" json:"login-burst" yaml:"login-burst"` // 登录/验证码突发容量
}
