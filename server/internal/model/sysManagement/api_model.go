package sysManagement

import (
	"server/internal/model/common"
)

// ApiModel API表（独立领域表）
type ApiModel struct {
	common.Td27Model
	Path        string `json:"path" gorm:"size:200;not null;comment:API路径"`
	Method      string `json:"method" gorm:"size:10;not null;comment:HTTP方法"`
	GroupEN     string `json:"group_en" gorm:"size:50;not null;comment:API分组(英文)"`
	GroupCN     string `json:"group_cn" gorm:"size:50;not null;comment:API分组(中文)"`
	Description string `json:"description" gorm:"size:100;not null;comment:API描述"`
}

func (ApiModel) TableName() string {
	return "sys_management_api"
}
