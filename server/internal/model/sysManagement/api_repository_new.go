package sysManagement

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// APIRepository 使用统一权限表的API仓库接口
type APIRepository interface {
	List(ctx context.Context, req *ListApiReq) ([]*PermissionModel, int64, error)
	Create(ctx context.Context, req *PermissionModel) (*PermissionModel, error)
	Update(ctx context.Context, req *PermissionModel) (*PermissionModel, error)
	Delete(ctx context.Context, id uint) error
	DeleteByIds(ctx context.Context, ids []uint) error
	FindOne(ctx context.Context, id uint) (*PermissionModel, error)
	FindByIds(ctx context.Context, ids []uint) ([]*PermissionModel, error)
	GetElTree(ctx context.Context) ([]*PermissionTree, error)
}

type apiEntity struct {
	conn *gorm.DB
}

func NewApiEntity(conn *gorm.DB) APIRepository {
	return &apiEntity{conn: conn}
}

func (e *apiEntity) List(ctx context.Context, req *ListApiReq) ([]*PermissionModel, int64, error) {
	var list []*PermissionModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&PermissionModel{}).
		Where("type = 'api'")

	if req.Path != "" {
		db = db.Where("resource LIKE ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.ApiGroup != "" {
		db = db.Where("api_group = ?", req.ApiGroup)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count api failed: %w", err)
	}

	if err := db.
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Order("id DESC").
		Find(&list).Error; err != nil {
		return nil, 0, fmt.Errorf("list api failed: %w", err)
	}

	return list, total, nil
}

func (e *apiEntity) Create(ctx context.Context, req *PermissionModel) (*PermissionModel, error) {
	req.Type = "api" // 强制设置为API类型
	if err := e.conn.WithContext(ctx).Create(req).Error; err != nil {
		return nil, fmt.Errorf("create api failed: %w", err)
	}
	return req, nil
}

func (e *apiEntity) Update(ctx context.Context, req *PermissionModel) (*PermissionModel, error) {
	result := e.conn.WithContext(ctx).
		Model(&PermissionModel{}).
		Where("id = ? AND type = 'api'", req.ID).
		Updates(map[string]interface{}{
			"name":      req.Name,
			"resource":  req.Resource,
			"method":    req.Method,
			"action":    req.Action,
			"api_group": req.ApiGroup,
			"status":    req.Status,
		})

	if result.Error != nil {
		return nil, fmt.Errorf("update api failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("API not found")
	}

	return req, nil
}

func (e *apiEntity) Delete(ctx context.Context, id uint) error {
	result := e.conn.WithContext(ctx).
		Where("id = ? AND type = 'api'", id).
		Delete(&PermissionModel{})

	if result.Error != nil {
		return fmt.Errorf("delete api failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("API not found")
	}

	return nil
}

func (e *apiEntity) DeleteByIds(ctx context.Context, ids []uint) error {
	if err := e.conn.WithContext(ctx).
		Where("id IN ? AND type = 'api'", ids).
		Delete(&PermissionModel{}).Error; err != nil {
		return fmt.Errorf("delete apis failed: %w", err)
	}
	return nil
}

func (e *apiEntity) FindOne(ctx context.Context, id uint) (*PermissionModel, error) {
	var api PermissionModel
	if err := e.conn.WithContext(ctx).
		Where("id = ? AND type = 'api'", id).
		First(&api).Error; err != nil {
		return nil, fmt.Errorf("find api failed: %w", err)
	}
	return &api, nil
}

func (e *apiEntity) FindByIds(ctx context.Context, ids []uint) ([]*PermissionModel, error) {
	var apis []*PermissionModel
	if err := e.conn.WithContext(ctx).
		Where("id IN ? AND type = 'api'", ids).
		Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("find apis by ids failed: %w", err)
	}
	return apis, nil
}

func (e *apiEntity) GetElTree(ctx context.Context) ([]*PermissionTree, error) {
	var apis []*PermissionModel
	if err := e.conn.WithContext(ctx).
		Where("type = 'api'").
		Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("get api tree failed: %w", err)
	}

	// 按api_group分组构建树
	groupMap := make(map[string]*PermissionTree)
	for _, api := range apis {
		group := api.ApiGroup
		if group == "" {
			group = "default"
		}

		if _, ok := groupMap[group]; !ok {
			groupMap[group] = &PermissionTree{
				PermissionModel: PermissionModel{
					Name: group,
					Type: "api_group",
				},
			}
		}

		groupMap[group].Children = append(groupMap[group].Children, &PermissionTree{
			PermissionModel: *api,
		})
	}

	// 转换为切片
	var result []*PermissionTree
	for _, tree := range groupMap {
		result = append(result, tree)
	}

	return result, nil
}
