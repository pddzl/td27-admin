package system

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	systemModel "server/model/system"
)

type ApiService struct{}

func (a *ApiService) AddApi(api systemModel.ApiModel) (err error) {
	if !errors.Is(global.TD27_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&systemModel.ApiModel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.TD27_DB.Create(&api).Error
}
