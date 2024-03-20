package request

type Role struct {
	RoleName string `json:"roleName" binding:"required"` // 角色名称
}

type EditRole struct {
	ID       uint   `json:"id" binding:"required"`       // 角色ID
	RoleName string `json:"roleName" binding:"required"` // 角色名称
}
