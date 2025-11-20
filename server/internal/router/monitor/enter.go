package monitor

type RouterGroup struct {
	*OperationLogRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		OperationLogRouter: NewOperationLogRouter(),
	}
}
