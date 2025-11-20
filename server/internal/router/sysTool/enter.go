package sysTool

type RouterGroup struct {
	*CronRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{CronRouter: NewCronRouter()}
}
