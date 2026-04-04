package configs

// Casbin RBAC配置
type Casbin struct {
	CacheTTL          int  `mapstructure:"cache-ttl" json:"cache-ttl" yaml:"cache-ttl"`                     // 缓存TTL（秒）
	AutoLoadInterval  int  `mapstructure:"auto-load-interval" json:"auto-load-interval" yaml:"auto-load-interval"` // 自动加载策略间隔（秒）
	EnableRoleHierarchy bool `mapstructure:"enable-role-hierarchy" json:"enable-role-hierarchy" yaml:"enable-role-hierarchy"` // 启用角色继承
	EnableDataPermission bool `mapstructure:"enable-data-permission" json:"enable-data-permission" yaml:"enable-data-permission"` // 启用数据权限
}
