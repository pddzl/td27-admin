package authority

type RouterGroup struct {
	*UserRouter
	*RoleRouter
	*MenuRouter
	*ApiRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		UserRouter: NewUserRouter(),
		RoleRouter: NewRoleRouter(),
		MenuRouter: NewMenuRouter(),
		ApiRouter:  NewApiRouter(),
	}
}
