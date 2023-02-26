package request

type Role struct {
	RoleName string `json:"roleName" validate:"required"`
}

type EditRole struct {
	ID       uint   `json:"id" validate:"required"`
	RoleName string `json:"roleName" validate:"required"`
}
