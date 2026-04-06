package sysManagement

// RolePermissionModel 角色权限关联表（统一权限模型）
type RolePermissionModel struct {
	RoleID       uint   `gorm:"column:role_id;primaryKey"`
	PermissionID uint   `gorm:"column:permission_id;primaryKey"`
	DataScope    string `gorm:"column:data_scope;size:20;default:'all';comment:数据权限范围:all全部|dept部门|self本人|custom自定义"`
	CustomSQL    string `gorm:"column:custom_sql;size:500;comment:自定义数据权限SQL条件"`
}

func (RolePermissionModel) TableName() string {
	return "sys_management_role_permissions"
}
