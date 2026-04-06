package sysManagement

//	type Role struct {
//		RoleName string `json:"roleName" binding:"required"` // 角色名称
//	}

//type RoleIdReq struct {
//	RoleId int `json:"role_id" binding:"required"`
//}

type UpdateRoleReq struct {
	ID       uint   `json:"id" binding:"required"`       // 角色ID
	RoleName string `json:"roleName" binding:"required"` // 角色名称
}
