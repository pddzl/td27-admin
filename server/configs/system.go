package configs

type System struct {
	Env                string `mapstructure:"env" json:"env" yaml:"env"`                                                    // 环境值
	Host               string `mapstructure:"host" json:"host" yaml:"host"`                                                 // IP地址
	Port               int    `mapstructure:"port" json:"port" yaml:"port"`                                                 // 端口号
	RouterPrefix       string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`                      // global router prefix
	DisableAutoMigrate bool   `mapstructure:"disable-auto-migrate" json:"disable-auto-migrate" yaml:"disable-auto-migrate"` // permit auto migrate table
	Stack              bool   `mapstructure:"stack" json:"stack" yaml:"stack"`                                              // 是否开启日志栈
}
