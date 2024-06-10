package authority

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"

	"server/global"
	modelAuthority "server/model/authority"
	authorityReq "server/model/authority/request"
	serviceBase "server/service/base"
)

type ApiService struct{}

// AddApi 添加api
func (a *ApiService) AddApi(api *modelAuthority.ApiModel) (*modelAuthority.ApiModel, error) {
	if !errors.Is(global.TD27_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&modelAuthority.ApiModel{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("存在相同api")
	}

	err := global.TD27_DB.Create(api).Error

	return api, err
}

// GetApis 获取所有api
func (a *ApiService) GetApis(apiSp authorityReq.ApiSearchParams) ([]modelAuthority.ApiModel, int64, error) {
	limit := apiSp.PageSize
	offset := apiSp.PageSize * (apiSp.Page - 1)
	db := global.TD27_DB.Model(&modelAuthority.ApiModel{})
	var apiList []modelAuthority.ApiModel

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
			var orderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 4)
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[apiSp.OrderKey] {
				if apiSp.Desc {
					orderStr = apiSp.OrderKey + " desc"
				} else {
					orderStr = apiSp.OrderKey
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", apiSp.OrderKey)
				return apiList, total, err
			}

			err = db.Order(orderStr).Find(&apiList).Error
		} else {
			err = db.Find(&apiList).Error
		}
	}
	return apiList, total, err
}

// GetElTreeApis 获取所有api tree
// element-plus el-tree的数据格式
func (a *ApiService) GetElTreeApis(roleId uint) (list []modelAuthority.ApiTree, checkedKey []string, err error) {
	var apiModels []modelAuthority.ApiModel
	err = global.TD27_DB.Find(&apiModels).Error
	if err != nil {
		return nil, nil, fmt.Errorf("GetElTreeApis: find -> %v", err)
	}

	var apiGroup []string
	err = global.TD27_DB.Model(&modelAuthority.ApiModel{}).Distinct().Pluck("api_group", &apiGroup).Error
	if err != nil {
		return nil, nil, fmt.Errorf("GetElTreeApis: apiGroup -> %v", err)
	}

	// 前端 el-tree data
	treeData := make(map[string][]modelAuthority.Children, len(apiModels))
	for _, model := range apiModels {
		var children modelAuthority.Children
		sPath := strings.Split(model.Path, fmt.Sprintf("%s/", model.ApiGroup))
		var tPath string
		if len(sPath) == 2 {
			tPath = sPath[1]
		}
		children.Key = fmt.Sprintf("%s,%s", model.Path, model.Method)
		children.ApiGroup = fmt.Sprintf("%s | %s", tPath, model.Description)
		children.Path = model.Path
		children.Method = model.Method
		children.Description = model.Description
		treeData[model.ApiGroup] = append(treeData[model.ApiGroup], children)
	}

	for _, value := range apiGroup {
		var apiTree modelAuthority.ApiTree
		apiTree.ApiGroup = value
		apiTree.Children = treeData[value]
		list = append(list, apiTree)
	}

	// 前端 el-tree default-checked-keys
	e := serviceBase.CasbinServiceApp.Casbin()
	authorityId := strconv.Itoa(int(roleId))
	cData := e.GetFilteredPolicy(0, authorityId)
	for _, v := range cData {
		checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", v[1], v[2]))
	}

	return
}

// DeleteApi 删除指定api
func (a *ApiService) DeleteApi(id uint) (err error) {
	var apiModel modelAuthority.ApiModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&apiModel).Error, gorm.ErrRecordNotFound) {
		global.TD27_LOG.Error("deleteApi -> 查找id", zap.Error(err))
		return err
	}

	err = global.TD27_DB.Unscoped().Delete(&apiModel).Error
	if err != nil {
		global.TD27_LOG.Error("deleteApi -> 删除id", zap.Error(err))
		return err
	}

	ok := serviceBase.CasbinServiceApp.ClearCasbin(1, apiModel.Path, apiModel.Method)
	if !ok {
		return errors.New(apiModel.Path + ":" + apiModel.Method + "casbin同步清理失败")
	}

	return nil
}

// DeleteApiById 批量删除API
func (a *ApiService) DeleteApiById(ids []uint) (err error) {
	var apis []modelAuthority.ApiModel
	err = global.TD27_DB.Find(&apis, "id in ?", ids).Unscoped().Delete(&apis).Error
	// 删除对应casbin条目
	if err == nil {
		for _, sysApi := range apis {
			ok := serviceBase.CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
			if !ok {
				global.TD27_LOG.Error(fmt.Sprintf("%s:%s casbin同步清理失败", sysApi.Path, sysApi.Method))
			}
		}
	}

	return
}

// EditApi 编辑api
func (a *ApiService) EditApi(instance *modelAuthority.ApiModel) (err error) {
	var apiModel modelAuthority.ApiModel
	if errors.Is(global.TD27_DB.Where("id = ?", instance.ID).First(&apiModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	if apiModel.Path != instance.Path || apiModel.Method != instance.Method {
		if !errors.Is(global.TD27_DB.Where("path = ? AND method = ?", instance.Path, instance.Method).First(&modelAuthority.ApiModel{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同接口")
		}
	}

	err = serviceBase.CasbinServiceApp.UpdateCasbinApi(apiModel.Path, apiModel.Path, apiModel.Method, apiModel.Method)
	if err != nil {
		return fmt.Errorf("更新casbin rule err: %v", err)
	}

	return global.TD27_DB.Omit("created_at").Save(instance).Error
}
