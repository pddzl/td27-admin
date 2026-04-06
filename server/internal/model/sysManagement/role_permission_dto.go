package sysManagement

type UpdateRolePermissionReq struct {
	RoleId        uint   `json:"role_id" binding:"required"` // Role ID
	Domain        string `json:"domain" binding:"required,oneof=menu api button data"`
	PermissionIds []uint `json:"permission_ids"` // Permission IDs
}
