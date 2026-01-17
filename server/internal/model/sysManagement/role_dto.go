package sysManagement

//type Role struct {
//	RoleName string `json:"roleName" binding:"required"` // 角色名称
//}

type UpdateRoleReq struct {
	ID       uint   `json:"id" binding:"required"`       // 角色ID
	RoleName string `json:"roleName" binding:"required"` // 角色名称
}

type UpdateRoleMenuReq struct {
	RoleId uint   `json:"roleId"` // 角色 ID
	Ids    []uint `json:"ids"`    // 菜单 ID
}
