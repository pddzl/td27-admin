package sysManagement

import (
	"server/internal/model/common"
)

// ApiModel API表（独立领域表）
type ApiModel struct {
	common.Td27Model
	ApiName  string `json:"apiName" gorm:"size:100;not null;comment:API名称"`
	Path     string `json:"path" gorm:"size:200;not null;comment:API路径"`
	Method   string `json:"method" gorm:"size:10;not null;comment:HTTP方法"`
	ApiGroup string `json:"apiGroup" gorm:"size:50;comment:API分组"`
	Status   bool   `json:"status" gorm:"default:true;comment:状态"`
	// 关联的权限ID
	PermissionID uint `json:"permissionId" gorm:"index;comment:关联权限ID"`
}

func (ApiModel) TableName() string {
	return "sys_management_api"
}
