package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
