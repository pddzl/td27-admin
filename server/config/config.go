package config

type Server struct {
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	RotateLogs RotateLogs `mapstructure:"rotateLogs" json:"rotateLogs" yaml:"rotateLogs"`
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	File       File       `mapstructure:"file" json:"file" yaml:"file"`
	Router     Router     `mapstructure:"router" json:"router" yaml:"router"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha    Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors       CORS       `mapstructure:"cors" json:"cors" yaml:"cors"`          // 跨域配置
	Crontab    Crontab    `mapstructure:"crontab" json:"crontab" yaml:"crontab"` // 计划任务
}
