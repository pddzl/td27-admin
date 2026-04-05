package sysManagement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ListApiReq API列表请求
type ListApiReq struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Path     string `json:"path" form:"path"`
	Method   string `json:"method" form:"method"`
	ApiGroup string `json:"apiGroup" form:"apiGroup"`
}

// CreateApiReq 创建API请求
type CreateApiReq struct {
	ApiName  string `json:"apiName" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	ApiGroup string `json:"apiGroup"`
}

// UpdateApiReq 更新API请求
type UpdateApiReq struct {
	ID       uint   `json:"id" binding:"required"`
	ApiName  string `json:"apiName" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	ApiGroup string `json:"apiGroup"`
}

// ApiTreeNode API树节点
type ApiTreeNode struct {
	ID       uint           `json:"id"`
	ApiName  string         `json:"apiName"`
	Path     string         `json:"path"`
	Method   string         `json:"method"`
	ApiGroup string         `json:"apiGroup"`
	Children []*ApiTreeNode `json:"children,omitempty"`
}

// ApiTreeResp API树响应
type ApiTreeResp struct {
	List       []*ApiTreeNode `json:"list"`
	CheckedKey []string       `json:"checkedKey"`
	CheckedIds []uint         `json:"checkedIds"`
}

// APIRepository API仓库接口
type APIRepository interface {
	List(ctx context.Context, req *ListApiReq) ([]*ApiModel, int64, error)
	Create(ctx context.Context, req *CreateApiReq) (*ApiModel, error)
	Update(ctx context.Context, req *UpdateApiReq) (*ApiModel, error)
	Delete(ctx context.Context, id uint) error
	DeleteByIds(ctx context.Context, ids []uint) error
	FindOne(ctx context.Context, id uint) (*ApiModel, error)
	FindByIds(ctx context.Context, ids []uint) ([]*ApiModel, error)
	GetElTree(ctx context.Context) ([]*ApiTreeNode, error)
	GetAllGroups(ctx context.Context) ([]string, error)
}

type apiEntity struct {
	conn *gorm.DB
}

func NewApiEntity(conn *gorm.DB) APIRepository {
	return &apiEntity{conn: conn}
}

func (e *apiEntity) List(ctx context.Context, req *ListApiReq) ([]*ApiModel, int64, error) {
	var list []*ApiModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&ApiModel{})

	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
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

	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	if err := db.Order("api_group, id").Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, fmt.Errorf("list api failed: %w", err)
	}

	return list, total, nil
}

func (e *apiEntity) Create(ctx context.Context, req *CreateApiReq) (*ApiModel, error) {
	api := &ApiModel{
		ApiName:  req.ApiName,
		Path:     req.Path,
		Method:   req.Method,
		ApiGroup: req.ApiGroup,
	}

	if err := e.conn.WithContext(ctx).Create(api).Error; err != nil {
		return nil, fmt.Errorf("create api failed: %w", err)
	}

	return api, nil
}

func (e *apiEntity) Update(ctx context.Context, req *UpdateApiReq) (*ApiModel, error) {
	updates := map[string]interface{}{
		"api_name":  req.ApiName,
		"path":      req.Path,
		"method":    req.Method,
		"api_group": req.ApiGroup,
	}

	result := e.conn.WithContext(ctx).Model(&ApiModel{}).Where("id = ?", req.ID).Updates(updates)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("update api failed: %w", err)
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("API不存在")
	}

	var api ApiModel
	if err := e.conn.WithContext(ctx).First(&api, req.ID).Error; err != nil {
		return nil, fmt.Errorf("find updated api failed: %w", err)
	}

	return &api, nil
}

func (e *apiEntity) Delete(ctx context.Context, id uint) error {
	result := e.conn.WithContext(ctx).Where("id = ?", id).Delete(&ApiModel{})
	if err := result.Error; err != nil {
		return fmt.Errorf("delete api failed: %w", err)
	}
	if result.RowsAffected == 0 {
		return errors.New("API不存在")
	}
	return nil
}

func (e *apiEntity) DeleteByIds(ctx context.Context, ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	return e.conn.WithContext(ctx).Where("id IN ?", ids).Delete(&ApiModel{}).Error
}

func (e *apiEntity) FindOne(ctx context.Context, id uint) (*ApiModel, error) {
	var api ApiModel
	if err := e.conn.WithContext(ctx).First(&api, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("API不存在")
		}
		return nil, fmt.Errorf("find api failed: %w", err)
	}
	return &api, nil
}

func (e *apiEntity) FindByIds(ctx context.Context, ids []uint) ([]*ApiModel, error) {
	if len(ids) == 0 {
		return []*ApiModel{}, nil
	}
	var apis []*ApiModel
	if err := e.conn.WithContext(ctx).Where("id IN ?", ids).Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("find apis by ids failed: %w", err)
	}
	return apis, nil
}

func (e *apiEntity) GetElTree(ctx context.Context) ([]*ApiTreeNode, error) {
	var apis []*ApiModel
	if err := e.conn.WithContext(ctx).Where("status = ?", true).Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("get all apis failed: %w", err)
	}

	// Group by api_group
	groupMap := make(map[string][]*ApiTreeNode)
	for _, api := range apis {
		node := &ApiTreeNode{
			ID:       api.ID,
			ApiName:  api.ApiName,
			Path:     api.Path,
			Method:   api.Method,
			ApiGroup: api.ApiGroup,
		}
		groupMap[api.ApiGroup] = append(groupMap[api.ApiGroup], node)
	}

	// Build tree with groups as parent nodes
	var tree []*ApiTreeNode
	for group, children := range groupMap {
		tree = append(tree, &ApiTreeNode{
			ApiName:  group,
			ApiGroup: group,
			Children: children,
		})
	}

	return tree, nil
}

func (e *apiEntity) GetAllGroups(ctx context.Context) ([]string, error) {
	var groups []string
	err := e.conn.WithContext(ctx).
		Model(&ApiModel{}).
		Where("status = ?", true).
		Distinct("api_group").
		Pluck("api_group", &groups).Error
	if err != nil {
		return nil, fmt.Errorf("get api groups failed: %w", err)
	}
	return groups, nil
}
