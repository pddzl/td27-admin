package sysMonitor

import (
	"server/internal/model/common"
)

type orStruct struct {
	Method string // 请求方法
	Path   string // 请求路径
	Status int    // http code
}

type OrListReq struct {
	orStruct
	common.PageInfo
}
