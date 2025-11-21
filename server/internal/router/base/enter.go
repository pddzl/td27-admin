package base

type RouterGroup struct {
	*CasbinRouter
	*LogRegRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		CasbinRouter: NewCasbinRouter(),
		LogRegRouter: NewLogRegRouter(),
	}
}
