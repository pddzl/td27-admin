package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境值
	Host          string `mapstructure:"host" json:"host" yaml:"host"`                               // IP地址
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                               // 端口号
	Stack         bool   `mapstructure:"stack" json:"stack" yaml:"stack"`                            // 是否开启日志栈
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	Upload        string `mapstructure:"upload" json:"upload" yaml:"upload"`                         // 上传目录
}
