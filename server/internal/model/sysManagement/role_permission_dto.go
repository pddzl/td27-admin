package sysManagement

type RebuildRolePermissionReq struct {
	RoleId    uint   `json:"role_id" binding:"required"` // Role ID
	Domain    string `json:"domain" binding:"required,oneof=menu api button data"`
	DomainIds []uint `json:"domain_ids"`
}
