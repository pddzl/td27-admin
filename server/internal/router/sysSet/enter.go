package sysSet

type RouterGroup struct {
	*DictRouter
	*DictDetailRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		DictRouter:       NewDictRouter(),
		DictDetailRouter: NewDictDetailRouter(),
	}
}
