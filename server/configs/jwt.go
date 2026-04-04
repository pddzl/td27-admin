package configs

type JWT struct {
	SigningKey      string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`                   // jwt签名
	ExpiresTime     int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`                // 过期时间（秒）
	BufferTime      int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`                   // 缓冲时间（秒）
	Issuer          string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                                  // 签发者
	MultiLogin      bool   `mapstructure:"multi-login" json:"multi-login" yaml:"multi-login"`                   // 是否允许多设备同时登录
	MultiLoginLimit int    `mapstructure:"multi-login-limit" json:"multi-login-limit" yaml:"multi-login-limit"` // 多设备登录数量限制（0表示无限制）
}
