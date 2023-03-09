package system

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"

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

// GetApisTree 获取所有api tree
// element-plus el-tree的数据格式
func (a *ApiService) GetApisTree(roleId uint) (list []systemModel.ApiTree, checkedKey []string, err error) {
	var apiModels []systemModel.ApiModel
	err = global.TD27_DB.Find(&apiModels).Error
	if err != nil {
		return nil, nil, fmt.Errorf("getApisTree: find -> %v", err)
	}

	var apiGroup []string
	err = global.TD27_DB.Model(&systemModel.ApiModel{}).Distinct().Pluck("api_group", &apiGroup).Error
	if err != nil {
		return nil, nil, fmt.Errorf("getApisTree: apiGroup -> %v", err)
	}

	// 前端 el-tree data
	treeData := make(map[string][]systemModel.Children, len(apiModels))
	for _, model := range apiModels {
		var children systemModel.Children
		sPath := strings.Split(model.Path, fmt.Sprintf("%s/", model.ApiGroup))
		var tPath string
		if len(sPath) == 2 {
			tPath = sPath[1]
		}
		children.Key = fmt.Sprintf("%s,%s", model.Path, model.Method)
		children.ApiGroup = fmt.Sprintf("%s -> %s", tPath, model.Description)
		children.Path = model.Path
		children.Method = model.Method
		children.Description = model.Description
		treeData[model.ApiGroup] = append(treeData[model.ApiGroup], children)
	}

	for _, value := range apiGroup {
		var apiTree systemModel.ApiTree
		apiTree.ApiGroup = value
		apiTree.Children = treeData[value]
		list = append(list, apiTree)
	}

	// 前端 el-tree default-checked-keys
	e := CasbinServiceApp.Casbin()
	authorityId := strconv.Itoa(int(roleId))
	cData := e.GetFilteredPolicy(0, authorityId)
	for _, v := range cData {
		checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", v[1], v[2]))
	}

	return
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
