package sysManagement

import (
	"server/internal/model/common"
)

// PermissionType 权限类型
const (
	PermissionTypeMenu   = "menu"   // 菜单权限
	PermissionTypeAPI    = "api"    // API权限
	PermissionTypeButton = "button" // 按钮权限
	PermissionTypeData   = "data"   // 数据权限
)

// PermissionAction 权限操作
const (
	PermissionActionView   = "view"   // 查看
	PermissionActionCreate = "create" // 创建
	PermissionActionUpdate = "update" // 更新
	PermissionActionDelete = "delete" // 删除
	PermissionActionAll    = "all"    // 全部
)

// PermissionModel 统一权限模型
type PermissionModel struct {
	common.Td27Model
	Name        string `json:"name" gorm:"size:100;not null;comment:权限名称"`
	Type        string `json:"type" gorm:"size:20;not null;comment:权限类型:menu|api|button|data"`
	Resource    string `json:"resource" gorm:"size:200;not null;comment:资源标识，如:/api/user或menu:users"`
	Action      string `json:"action" gorm:"size:20;default:'view';comment:操作:view|create|update|delete|all"`
	ParentID    *uint  `json:"parentId" gorm:"index;comment:父权限ID"`
	Sort        uint   `json:"sort" gorm:"default:0;comment:排序"`
	Status      bool   `json:"status" gorm:"default:true;comment:状态:true启用|false禁用"`
	// API特有字段
	Method   string `json:"method" gorm:"size:10;comment:HTTP方法"`
	ApiGroup string `json:"apiGroup" gorm:"size:50;comment:API分组"`
	// 菜单特有字段
	Icon      string `json:"icon" gorm:"size:100;comment:图标"`
	Component string `json:"component" gorm:"size:200;comment:前端组件"`
	Redirect  string `json:"redirect" gorm:"size:200;comment:重定向"`
	Hidden    bool   `json:"hidden" gorm:"default:false;comment:是否隐藏"`
	KeepAlive bool   `json:"keepAlive" gorm:"default:false;comment:缓存"`
}

func (PermissionModel) TableName() string {
	return "sys_management_permission"
}

// PermissionTree 权限树结构
type PermissionTree struct {
	PermissionModel
	Children []*PermissionTree `json:"children,omitempty"`
}

// ApiTreeResp API树响应
type ApiTreeResp struct {
	List       []*PermissionTree `json:"list"`
	CheckedKey []string          `json:"checkedKey"`
	CheckedIds []uint            `json:"checkedIds"`
}

// UpdateRoleAPIReq 更新角色API权限请求
type UpdateRoleAPIReq struct {
	RoleID           uint   `json:"roleId" binding:"required"`
	APIPermissionIDs []uint `json:"apiPermissionIds"` // API权限ID列表
}

// CheckPermission 检查是否匹配权限
func (p *PermissionModel) CheckPermission(resource string, action string) bool {
	if p.Resource != resource {
		return false
	}
	if p.Action == PermissionActionAll || p.Action == action {
		return true
	}
	return false
}
