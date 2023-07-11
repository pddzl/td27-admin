package system

import (
	"server/global"
	"time"
)

type OperationRecord struct {
	global.TD27_MODEL
	Ip        string        `json:"ip" gorm:"column:ip;comment:请求ip"`                          // 请求ip
	Method    string        `json:"method" gorm:"column:method;comment:请求方法"`                  // 请求方法
	Path      string        `json:"path" gorm:"column:path;comment:请求路径"`                      // 请求路径
	Status    int           `json:"status" gorm:"column:status;comment:请求状态"`                  // 请求状态
	UserAgent string        `json:"userAgent" gorm:"column:user_agent"`                        // http userAgent
	ReqParam  string        `json:"reqParam" gorm:"type:text;column:req_param;comment:请求Body"` // 请求参数
	RespData  string        `json:"respData" gorm:"type:text;column:resp_data;comment:响应数据"`   // 响应数据
	RespTime  time.Duration `json:"respTime" gorm:"column:resp_time"`                          // 响应时间
	UserID    uint          `json:"userID" gorm:"column:user_id;comment:用户id"`                 // 用户id
	UserName  string        `json:"userName" gorm:"user_name;comment:用户名称"`                    // 用户名称
}

func (o *OperationRecord) TableName() string {
	return "sys_operation_record"
}
