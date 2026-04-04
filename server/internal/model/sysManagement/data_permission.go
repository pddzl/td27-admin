package sysManagement

import (
	"gorm.io/gorm"
)

// DataScope 数据权限范围
type DataScope string

const (
	DataScopeAll    DataScope = "all"    // 全部数据
	DataScopeDept   DataScope = "dept"   // 本部门数据
	DataScopeSelf   DataScope = "self"   // 仅本人数据
	DataScopeCustom DataScope = "custom" // 自定义SQL
)

// DataPermission 数据权限配置
type DataPermission struct {
	Scope     DataScope // 数据范围
	DeptID    uint      // 部门ID（当Scope为dept时使用）
	UserID    uint      // 用户ID（当Scope为self时使用）
	CustomSQL string    // 自定义SQL（当Scope为custom时使用）
}

// ApplyDataScope 应用数据权限到查询
// db: GORM查询实例
// perm: 数据权限配置
// tableAlias: 表别名（可选）
// deptColumn: 部门ID字段名
// userColumn: 用户ID字段名
func ApplyDataScope(db *gorm.DB, perm *DataPermission, tableAlias, deptColumn, userColumn string) *gorm.DB {
	if perm == nil {
		return db
	}

	if tableAlias != "" {
		tableAlias = tableAlias + "."
	}

	switch perm.Scope {
	case DataScopeAll:
		// 全部数据，不添加过滤
		return db
	case DataScopeDept:
		// 本部门数据
		if deptColumn != "" {
			return db.Where(tableAlias+deptColumn+" = ?", perm.DeptID)
		}
		return db
	case DataScopeSelf:
		// 仅本人数据
		if userColumn != "" {
			return db.Where(tableAlias+userColumn+" = ?", perm.UserID)
		}
		return db
	case DataScopeCustom:
		// 自定义SQL
		if perm.CustomSQL != "" {
			return db.Where(perm.CustomSQL)
		}
		return db
	default:
		// 默认只能看自己的
		if userColumn != "" {
			return db.Where(tableAlias+userColumn+" = ?", perm.UserID)
		}
		return db
	}
}
