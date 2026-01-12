package authority

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type ApiEntity interface {
	FindOne(ctx context.Context, id uint) (*ApiModel, error)
	FindByIds(ctx context.Context, ids []uint) ([]*ApiModel, error)
	List(ctx context.Context, req *ListApiReq) ([]*ApiModel, int64, error)
	Create(ctx context.Context, req *ApiModel) (*ApiModel, error)
	GetElTree(ctx context.Context) ([]*ApiTree, error)
	Delete(ctx context.Context, id uint) error
	DeleteByIds(ctx context.Context, ids []uint) error
	Update(ctx context.Context, req *ApiModel) (*ApiModel, error)
}

type defaultApiEntity struct {
	conn *gorm.DB
}

func NewDefaultApiEntity(conn *gorm.DB) ApiEntity {
	return &defaultApiEntity{conn: conn}
}

func (e *defaultApiEntity) FindOne(ctx context.Context, id uint) (*ApiModel, error) {
	var apiModel ApiModel
	err := e.conn.WithContext(ctx).Where("id = ?", id).First(&apiModel).Error
	if err != nil {
		return nil, err
	}
	return &apiModel, nil
}

func (e *defaultApiEntity) FindByIds(ctx context.Context, ids []uint) ([]*ApiModel, error) {
	if len(ids) == 0 {
		return []*ApiModel{}, nil
	}

	var apiModels []*ApiModel
	err := e.conn.WithContext(ctx).Where("id IN ?", ids).Find(&apiModels).Error

	if err != nil {
		return nil, fmt.Errorf("find apis by ids failed: %w", err)
	}

	return apiModels, nil
}

func (e *defaultApiEntity) List(ctx context.Context, req *ListApiReq) ([]*ApiModel, int64, error) {
	var list []*ApiModel
	var total int64

	// pagination safety
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	db := e.conn.WithContext(ctx).Model(&ApiModel{})

	// filters
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Description != "" {
		db = db.Where("description LIKE ?", "%"+req.Description+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.ApiGroup != "" {
		db = db.Where("api_group LIKE ?", "%"+req.ApiGroup+"%")
	}

	// count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count api failed: %w", err)
	}

	// order by (SQL injection safe)
	if req.OrderKey != "" {
		orderColumns := map[string]string{
			"path":        "path",
			"api_group":   "api_group",
			"description": "description",
			"method":      "method",
		}

		column, ok := orderColumns[req.OrderKey]
		if !ok {
			return nil, total, fmt.Errorf("非法的排序字段: %s", req.OrderKey)
		}

		order := column
		if req.Desc {
			order += " DESC"
		}
		db = db.Order(order)
	}

	// query
	if err := db.
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error; err != nil {
		return nil, total, fmt.Errorf("list api failed: %w", err)
	}

	return list, total, nil
}

func (e *defaultApiEntity) Create(ctx context.Context, req *ApiModel) (*ApiModel, error) {
	db := e.conn.WithContext(ctx)

	// existence check
	var exists int64
	if err := db.
		Model(&ApiModel{}).
		Where("path = ? AND method = ?", req.Path, req.Method).
		Count(&exists).Error; err != nil {
		return nil, fmt.Errorf("check api existence failed: %w", err)
	}

	if exists > 0 {
		return nil, errors.New("存在相同 api")
	}

	if err := db.Create(req).Error; err != nil {
		return nil, fmt.Errorf("create api failed: %w", err)
	}

	return req, nil
}

// GetElTree 获取所有api tree
// element-plus el-tree的数据格式
func (e *defaultApiEntity) GetElTree(ctx context.Context) ([]*ApiTree, error) {
	db := e.conn.WithContext(ctx)

	// query all apis (only required fields)
	var apis []*ApiModel
	if err := db.
		Select("path", "method", "description", "api_group").
		Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("GetElTree: query apis failed: %w", err)
	}

	// build tree map
	treeMap := make(map[string][]Children)
	for _, api := range apis {
		children := Children{
			Key:         api.Path + "," + api.Method,
			ApiGroup:    buildTreeLabel(api.ApiGroup, api.Path, api.Description),
			Path:        api.Path,
			Method:      api.Method,
			Description: api.Description,
		}
		treeMap[api.ApiGroup] = append(treeMap[api.ApiGroup], children)
	}

	// convert map → slice (stable & predictable order)
	list := make([]*ApiTree, 0, len(treeMap))
	for group, children := range treeMap {
		list = append(list, &ApiTree{
			ApiGroup: group,
			Children: children,
		})
	}

	// 前端 el-tree default-checked-keys
	//e := casbinService.Casbin()
	//authorityId := strconv.Itoa(int(roleId))
	//cData, _ := e.GetFilteredPolicy(0, authorityId)
	//for _, v := range cData {
	//	checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", v[1], v[2]))
	//}

	return list, nil
}

func buildTreeLabel(group, path, desc string) string {
	prefix := group + "/"
	if strings.HasPrefix(path, prefix) {
		return path[len(prefix):] + " | " + desc
	}
	return path + " | " + desc
}

func (e *defaultApiEntity) Delete(ctx context.Context, id uint) error {
	db := e.conn.WithContext(ctx)

	result := db.Unscoped().Delete(&ApiModel{}, id)
	if result.Error != nil {
		return fmt.Errorf("delete api failed: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("api not found, id=%d", id)
	}
	return nil

	// optional: casbin cleanup
	// if ok := casbinService.ClearCasbin(1, api.Path, api.Method); !ok {
	//     return fmt.Errorf("casbin cleanup failed: %s:%s", api.Path, api.Method)
	// }
}

func (e *defaultApiEntity) DeleteByIds(ctx context.Context, ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	db := e.conn.WithContext(ctx)

	result := db.Unscoped().Where("id IN ?", ids).Delete(&ApiModel{})

	if result.Error != nil {
		return fmt.Errorf("delete apis failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
	// 删除对应casbin条目
	//if err == nil {
	//	for _, sysApi := range apis {
	//		ok := casbinService.ClearCasbin(1, sysApi.Path, sysApi.Method)
	//		if !ok {
	//			global.TD27_LOG.Error(fmt.Sprintf("%s:%s casbin同步清理失败", sysApi.Path, sysApi.Method))
	//		}
	//	}
	//}
}

func (e *defaultApiEntity) Update(ctx context.Context, req *ApiModel) (*ApiModel, error) {
	db := e.conn.WithContext(ctx)

	// Ensure record exists
	var existing ApiModel
	if err := db.First(&existing, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("记录不存在")
		}
		return nil, fmt.Errorf("find api failed: %w", err)
	}

	// Check uniqueness only if path or method changed
	if existing.Path != req.Path || existing.Method != req.Method {
		var count int64
		if err := db.Model(&ApiModel{}).
			Where("path = ? AND method = ? AND id <> ?", req.Path, req.Method, req.ID).
			Count(&count).Error; err != nil {
			return nil, fmt.Errorf("check api uniqueness failed: %w", err)
		}

		if count > 0 {
			return nil, errors.New("存在相同接口")
		}
	}

	// Update only allowed fields
	err := db.Model(&existing).Updates(map[string]interface{}{
		"path":        req.Path,
		"method":      req.Method,
		"api_group":   req.ApiGroup,
		"description": req.Description,
	}).Error

	if err != nil {
		return nil, fmt.Errorf("update api failed: %w", err)
	}

	return &existing, nil
}
