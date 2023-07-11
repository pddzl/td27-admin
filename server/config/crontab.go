package config

type Crontab struct {
	Open bool   `mapstructure:"open" json:"open" yaml:"open"` // 是否启用
	Spec string `mapstructure:"spec" json:"spec" yaml:"spec"` // CRON表达式
	//WithSeconds bool     `mapstructure:"with_seconds" json:"with_seconds" yaml:"with_seconds"` // 是否精确到秒
	Objects []Object `mapstructure:"objects" json:"objects" yaml:"objects"`
}

type Object struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}
