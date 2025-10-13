package sysSet

import "server/global"

type DictDetailModel struct {
	global.TD27_MODEL
	Label       string             `json:"label" gorm:"column:label" binding:"required"`
	Value       string             `json:"value" gorm:"column:value" binding:"required"`
	Sort        int                `json:"sort" gorm:"column:sort"`
	DictModelID int                `json:"dictId" gorm:"column:dict_id" binding:"required"`
	ParentID    *int               `json:"parentId" gorm:"column:parent_id"` // new
	Children    []*DictDetailModel `json:"children" gorm:"-"`
	Description string             `json:"description" gorm:"column:description"`
}

func (ddm *DictDetailModel) TableName() string {
	return "sysSet_dictDetail"
}
