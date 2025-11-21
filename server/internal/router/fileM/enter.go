package fileM

type RouterGroup struct {
	*FileRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		FileRouter: NewFileRouter(),
	}
}
