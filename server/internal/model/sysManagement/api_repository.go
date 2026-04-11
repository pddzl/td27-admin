package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

// APIRepository API仓库接口
type APIRepository interface {
	List(ctx context.Context, req *ListApiReq) ([]*ApiModel, int64, error)
	Create(ctx context.Context, req *CreateApiReq) (*ApiModel, error)
	Update(ctx context.Context, req *UpdateApiReq) (*ApiModel, error)
	Delete(ctx context.Context, id uint) error
	DeleteByIds(ctx context.Context, ids []uint) error
	FindOne(ctx context.Context, id uint) (*ApiModel, error)
	FindByIds(ctx context.Context, ids []uint) ([]*ApiModel, error)
	ElTree(ctx context.Context) ([]*ApiTreeNode, error)
	GetAllGroups(ctx context.Context) ([]string, error)
}

type apiRepo struct {
	conn *gorm.DB
}

func NewApiRepo(conn *gorm.DB) APIRepository {
	return &apiRepo{conn: conn}
}

func (e *apiRepo) List(ctx context.Context, req *ListApiReq) ([]*ApiModel, int64, error) {
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
		db = db.Where("group_en = ?", req.ApiGroup)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count api failed: %w", err)
	}

	req.Normalize()

	if err := db.Order("group_en, id").Limit(req.PageSize).Offset(req.Offset()).Find(&list).Error; err != nil {
		return nil, 0, fmt.Errorf("list api failed: %w", err)
	}

	return list, total, nil
}

func (e *apiRepo) Create(ctx context.Context, req *CreateApiReq) (*ApiModel, error) {
	api := &ApiModel{
		Path:        req.Path,
		Method:      req.Method,
		GroupEN:     req.GroupEN,
		GroupCN:     req.GroupCN,
		Description: req.Description,
	}

	if err := e.conn.WithContext(ctx).Create(api).Error; err != nil {
		return nil, fmt.Errorf("create api failed: %w", err)
	}

	return api, nil
}

func (e *apiRepo) Update(ctx context.Context, req *UpdateApiReq) (*ApiModel, error) {
	updates := map[string]interface{}{
		"path":        req.Path,
		"method":      req.Method,
		"group_en":    req.GroupEN,
		"group_cn":    req.GroupCN,
		"description": req.Description,
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

func (e *apiRepo) Delete(ctx context.Context, id uint) error {
	result := e.conn.WithContext(ctx).Where("id = ?", id).Delete(&ApiModel{})

	if err := result.Error; err != nil {
		return fmt.Errorf("delete api failed: %w", err)
	}

	if result.RowsAffected == 0 {
		return errors.New("API不存在")
	}

	return nil
}

func (e *apiRepo) DeleteByIds(ctx context.Context, ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	return e.conn.WithContext(ctx).Where("id IN ?", ids).Delete(&ApiModel{}).Error
}

func (e *apiRepo) FindOne(ctx context.Context, id uint) (*ApiModel, error) {
	var api ApiModel

	if err := e.conn.WithContext(ctx).First(&api, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("API不存在")
		}

		return nil, fmt.Errorf("find api failed: %w", err)
	}

	return &api, nil
}

func (e *apiRepo) FindByIds(ctx context.Context, ids []uint) ([]*ApiModel, error) {
	if len(ids) == 0 {
		return []*ApiModel{}, nil
	}

	var apis []*ApiModel
	if err := e.conn.WithContext(ctx).Where("id IN ?", ids).Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("find apis by ids failed: %w", err)
	}

	return apis, nil
}

func (e *apiRepo) ElTree(ctx context.Context) ([]*ApiTreeNode, error) {
	var apis []*ApiModel
	if err := e.conn.WithContext(ctx).Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("get all apis failed: %w", err)
	}

	// Group by group_cn
	groupMap := make(map[string][]*ApiTreeNode)
	for _, api := range apis {
		node := &ApiTreeNode{
			Id:          api.ID,
			Key:         strconv.Itoa(int(api.ID)),
			Path:        api.Path,
			Method:      api.Method,
			GroupEN:     api.GroupEN,
			GroupCN:     api.GroupCN,
			Description: api.Description,
		}
		groupMap[api.GroupCN] = append(groupMap[api.GroupCN], node)
	}

	// Build tree with groups as parent nodes
	var tree []*ApiTreeNode
	for group, children := range groupMap {
		tree = append(tree, &ApiTreeNode{
			Key:      group,
			Children: children,
		})
	}

	return tree, nil
}

func (e *apiRepo) GetAllGroups(ctx context.Context) ([]string, error) {
	var groups []string
	err := e.conn.WithContext(ctx).
		Model(&ApiModel{}).
		Distinct("group_en").
		Pluck("group_en", &groups).Error
	if err != nil {
		return nil, fmt.Errorf("get api groups failed: %w", err)
	}
	return groups, nil
}
