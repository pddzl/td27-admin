package sysManagement

import "server/internal/model/common"

type ButtonModel struct {
	common.Td27Model
	ButtonCode  string `json:"buttonCode" gorm:"size:100;not null;unique"`
	ButtonName  string `json:"buttonName" gorm:"size:100;not null"`
	Description string `json:"description" gorm:"size:200"`
	PagePath    string `json:"pagePath" gorm:"size:200;not null"`
}

func (ButtonModel) TableName() string {
	return "sys_management_button"
}

type ButtonDto struct {
	ID            uint   `json:"id"`
	ButtonCode    string `json:"buttonCode"`
	ButtonName    string `json:"buttonName"`
	Description   string `json:"description"`
	PagePath      string `json:"pagePath"`
	HasPermission bool   `json:"hasPermission"`
}

type CreateButtonReq struct {
	ButtonCode  string `json:"buttonCode" binding:"required,max=100"`
	ButtonName  string `json:"buttonName" binding:"required,max=100"`
	Description string `json:"description"`
	PagePath    string `json:"pagePath" binding:"required,max=200"`
}

type UpdateButtonReq struct {
	ID          uint   `json:"id" binding:"required"`
	ButtonCode  string `json:"buttonCode" binding:"required,max=100"`
	ButtonName  string `json:"buttonName" binding:"required,max=100"`
	Description string `json:"description"`
	PagePath    string `json:"pagePath" binding:"required,max=200"`
}

type ListButtonReq struct {
	common.PageInfo
	PagePath string `json:"pagePath" form:"pagePath"`
}

type CheckButtonReq struct {
	ButtonCode string `json:"buttonCode" binding:"required"`
}
