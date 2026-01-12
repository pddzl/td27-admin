package authority

type Children struct {
	Key         string `json:"key"`      // for 前端el-tree node-key (path + method)
	ApiGroup    string `json:"apiGroup"` // for 前端el-tree label (path + description)
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type ApiTree struct {
	ApiGroup string     `json:"apiGroup"`
	Children []Children `json:"children"`
}
