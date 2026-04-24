package configs

type Server struct {
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Logger     Logger     `mapstructure:"logger" json:"logger" yaml:"logger"`
	RotateLogs RotateLogs `mapstructure:"rotateLogs" json:"rotateLogs" yaml:"rotateLogs"`
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	File       File       `mapstructure:"file" json:"file" yaml:"file"`
	Pgsql      Pgsql      `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`    // PostgreSQL 配置
	Casbin     Casbin     `mapstructure:"casbin" json:"casbin" yaml:"casbin"` // Casbin RBAC配置
	Captcha    Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors       CORS       `mapstructure:"cors" json:"cors" yaml:"cors"` // 跨域配置
}
