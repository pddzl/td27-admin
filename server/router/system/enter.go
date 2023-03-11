package system

type RouterGroup struct {
	BaseRouter
	UserRouter
	RoleRouter
	MenuRouter
	ApiRouter
	CasbinRouter
	JwtRouter
}
