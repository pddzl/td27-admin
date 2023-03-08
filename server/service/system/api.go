package system

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
)

type ApiService struct{}

// AddApi 添加api
func (a *ApiService) AddApi(api systemModel.ApiModel) (*systemModel.ApiModel, error) {
	if !errors.Is(global.TD27_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&systemModel.ApiModel{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("存在相同api")
	}

	err := global.TD27_DB.Create(&api).Error

	return &api, err
}

// GetApis 获取所有api
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

// DeleteApi 删除指定api
func (a *ApiService) DeleteApi(id uint) (err error) {
	var apiModel systemModel.ApiModel
	err = global.TD27_DB.Where("id = ?", id).First(&apiModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.TD27_LOG.Error("deleteApi -> 查找id", zap.Error(err))
		return err
	}

	err = global.TD27_DB.Unscoped().Delete(&apiModel).Error
	if err != nil {
		global.TD27_LOG.Error("deleteApi -> 删除id", zap.Error(err))
		return err
	}

	ok := CasbinServiceApp.ClearCasbin(1, apiModel.Path, apiModel.Method)
	if !ok {
		return errors.New(apiModel.Path + ":" + apiModel.Method + "casbin同步清理失败")
	}
	e := CasbinServiceApp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

// EditApi 编辑api
func (a *ApiService) EditApi(eApi systemReq.EditApi) (err error) {
	var oldApiModel systemModel.ApiModel
	err = global.TD27_DB.Where("id = ?", eApi.Id).First(&oldApiModel).Error
	if err != nil {
		return errors.New("editApi: id不存在")
	}

	if oldApiModel.Path != eApi.Path || oldApiModel.Method != eApi.Method {
		if !errors.Is(global.TD27_DB.Where("path = ? AND method = ?", eApi.Path, eApi.Method).First(&systemModel.ApiModel{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("editApi: 存在相同接口")
		}
	}

	err = CasbinServiceApp.UpdateCasbinApi(oldApiModel.Path, eApi.Path, oldApiModel.Method, eApi.Method)
	if err != nil {
		return fmt.Errorf("editApi: 更新casbin rule -> %v", err)
	}

	return global.TD27_DB.Debug().Model(&oldApiModel).Updates(map[string]interface{}{"path": eApi.Path, "method": eApi.Method, "api_group": eApi.ApiGroup, "description": eApi.Description}).Error
}
