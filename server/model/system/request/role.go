package request

type Role struct {
	RoleName string `json:"roleName" validate:"required"`
}
