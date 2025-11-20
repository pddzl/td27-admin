package sysSet

import (
	"server/internal/global"
)

type DictModel struct {
	global.TD27_MODEL
	CHName      string            `json:"chName" gorm:"column:ch_name;unique" binding:"required"`
	ENName      string            `json:"enName" gorm:"column:en_name;unique" binding:"required"`
	DictDetails []DictDetailModel `json:"dictDetails"`
}

func (dm *DictModel) TableName() string {
	return "sysSet_dict"
}
