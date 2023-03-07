package system

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
)

type ApiService struct{}

func (a *ApiService) AddApi(api systemModel.ApiModel) (err error) {
	if !errors.Is(global.TD27_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&systemModel.ApiModel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.TD27_DB.Create(&api).Error
}

func (a *ApiService) GetApis(apiSp systemReq.ApiSearchParams) ([]systemModel.ApiModel, int64, error) {
	limit := apiSp.PageSize
	offset := apiSp.PageSize * (apiSp.Page - 1)
	db := global.TD27_DB.Model(&systemModel.ApiModel{})
	var apiList []systemModel.ApiModel

	if apiSp.Path != "" {
		db = db.Where("path LIKE ?", "%"+apiSp.Path+"%")
	}

	if apiSp.Description != "" {
		db = db.Where("description LIKE ?", "%"+apiSp.Description+"%")
	}

	if apiSp.Method != "" {
		db = db.Where("method = ?", apiSp.Method)
	}

	if apiSp.ApiGroup != "" {
		db = db.Where("api_group = ?", apiSp.ApiGroup)
	}

	var total int64
	err := db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if apiSp.OrderKey != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[OrderStr] {
				if apiSp.Desc {
					OrderStr = apiSp.OrderKey + " desc"
				} else {
					OrderStr = apiSp.OrderKey
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", apiSp.OrderKey)
				return apiList, total, err
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}
