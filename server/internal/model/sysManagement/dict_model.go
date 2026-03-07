package sysManagement

import (
	"server/internal/model/common"
)

type DictModel struct {
	common.Td27Model
	CHName      string            `json:"chName" gorm:"column:ch_name;unique" binding:"required"`
	ENName      string            `json:"enName" gorm:"column:en_name;unique" binding:"required"`
	DictDetails []DictDetailModel `json:"dictDetails"`
}

func (dm *DictModel) TableName() string {
	return "sys_management_dict"
}
