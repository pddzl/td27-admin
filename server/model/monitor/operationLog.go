package monitor

import (
	"server/global"
)

type OperationLogModel struct {
	global.TD27_MODEL
	Ip        string `json:"ip" gorm:"comment:请求ip"`                       // 请求ip
	Method    string `json:"method" gorm:"comment:请求方法"`                   // 请求方法
	Path      string `json:"path" gorm:"comment:请求路径"`                     // 请求路径
	Status    int    `json:"status" gorm:"comment:请求状态"`                   // 请求状态
	UserAgent string `json:"userAgent"`                                    // http userAgent
	ReqParam  string `json:"reqParam" gorm:"type:text;comment:请求Body"`     // 请求参数
	RespData  string `json:"respData" gorm:"type:mediumtext;comment:响应数据"` // 响应数据
	RespTime  int64  `json:"respTime"`                                     // 响应时间
	UserID    uint   `json:"userID" gorm:"comment:用户id"`                   // 用户id
	UserName  string `json:"userName" gorm:"comment:用户名称"`                 // 用户名称
}

func (ol *OperationLogModel) TableName() string {
	return "monitor_operationLog"
}
