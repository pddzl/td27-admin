package sysManagement

import (
	"server/internal/model/common"
)

type ApiModel struct {
	common.Td27Model
	Path        string `json:"path" gorm:"not null;comment:api路径" binding:"required"`             // API 路径
	Description string `json:"description" gorm:"not null;comment:api中文描述" binding:"required"`    // API 中文描述
	ApiGroup    string `json:"apiGroup" gorm:"not null;comment:api组" binding:"required"`          // API 组
	Method      string `json:"method" gorm:"not null;default:POST;comment:方法" binding:"required"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (ApiModel) TableName() string {
	return "authority_api"
}
