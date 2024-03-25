package authority

import "server/global"

type ApiModel struct {
	global.TD27_MODEL
	Path        string `json:"path" gorm:"not null;comment:api路径" binding:"required"`             // api路径
	Description string `json:"description" gorm:"not null;comment:api中文描述" binding:"required"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"not null;comment:api组" binding:"required"`          // api组
	Method      string `json:"method" gorm:"not null;default:POST;comment:方法" binding:"required"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (ApiModel) TableName() string {
	return "authority_api"
}

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
