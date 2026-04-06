package sysManagement

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type DictDetailRepository interface {
	List(context.Context, *ListDictDetailReq) ([]*DictDetailModel, int64, error)
	Flat(ctx context.Context, dictId uint) ([]*DictDetailModel, error)
	Create(ctx context.Context, req *DictDetailModel) (*DictDetailModel, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *DictDetailModel) (*DictDetailModel, error)
}

type dictDetailEntity struct {
	conn *gorm.DB
}

func NewDictDetailEntity(conn *gorm.DB) DictDetailRepository {
	return &dictDetailEntity{conn: conn}
}

func (e *dictDetailEntity) List(ctx context.Context, req *ListDictDetailReq) ([]*DictDetailModel, int64, error) {
	req.Normalize() // Page / PageSize safety

	var roots []*DictDetailModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&DictDetailModel{})

	// count root nodes
	if err := db.
		Where("dict_id = ? AND parent_id IS NULL", req.DictID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// load paginated roots
	if err := db.
		Where("dict_id = ? AND parent_id IS NULL", req.DictID).
		Order("sort ASC").
		Limit(req.PageSize).
		Offset(req.Offset()).
		Find(&roots).Error; err != nil {
		return nil, 0, err
	}

	if len(roots) == 0 {
		return []*DictDetailModel{}, total, nil
	}

	// root id set
	rootIDs := make([]uint, 0, len(roots))
	for _, r := range roots {
		r.Children = []*DictDetailModel{}
		rootIDs = append(rootIDs, r.ID)
	}

	// load all descendants (single query)
	var all []*DictDetailModel
	if err := db.
		Where("dict_id = ? AND parent_id IS NOT NULL", req.DictID).
		Order("sort ASC").
		Find(&all).Error; err != nil {
		return nil, 0, err
	}

	// build map[id]*node
	nodeMap := make(map[uint]*DictDetailModel, len(all)+len(roots))
	for _, r := range roots {
		nodeMap[r.ID] = r
	}
	for _, n := range all {
		n.Children = []*DictDetailModel{}
		nodeMap[n.ID] = n
	}

	// attach children
	for _, n := range all {
		if n.ParentID != nil {
			if parent, ok := nodeMap[uint(*n.ParentID)]; ok {
				parent.Children = append(parent.Children, n)
			}
		}
	}

	return roots, total, nil
}

func (e *dictDetailEntity) Flat(ctx context.Context, dictId uint) ([]*DictDetailModel, error) {
	var all []*DictDetailModel

	if err := e.conn.WithContext(ctx).
		Where("dict_id = ?", dictId).
		Order("sort ASC").
		Find(&all).Error; err != nil {
		return nil, err
	}

	if len(all) == 0 {
		return []*DictDetailModel{}, nil
	}

	m := make(map[uint]*DictDetailModel, len(all))
	for _, n := range all {
		m[n.ID] = n
	}

	cache := make(map[uint]string)

	var fullLabel func(n *DictDetailModel) string
	fullLabel = func(n *DictDetailModel) string {
		if v, ok := cache[n.ID]; ok {
			return v
		}
		if n.ParentID == nil {
			cache[n.ID] = n.Label
			return n.Label
		}
		if p, ok := m[uint(*n.ParentID)]; ok {
			cache[n.ID] = fullLabel(p) + " - " + n.Label
			return cache[n.ID]
		}
		cache[n.ID] = n.Label
		return n.Label
	}

	for _, n := range all {
		n.Label = fullLabel(n)
	}

	return all, nil
}

func (e *dictDetailEntity) Create(ctx context.Context, req *DictDetailModel) (*DictDetailModel, error) {
	var count int64

	if err := e.conn.WithContext(ctx).
		Model(&DictDetailModel{}).
		Where("dict_id = ? AND (label = ? OR value = ?)",
			req.DictModelID, req.Label, req.Value).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("duplicate label or value in the same dict")
	}

	if err := e.conn.WithContext(ctx).Create(req).Error; err != nil {
		return nil, err
	}

	return req, nil
}

func (e *dictDetailEntity) Delete(ctx context.Context, id uint) error {
	return e.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var node DictDetailModel
		if err := tx.First(&node, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("DictDetail not exist")
			}
			return err
		}

		return deleteTree(tx, id)
	})
}

func deleteTree(tx *gorm.DB, parentID uint) error {
	var children []DictDetailModel
	if err := tx.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return err
	}

	for _, c := range children {
		if err := deleteTree(tx, c.ID); err != nil {
			return err
		}
	}

	return tx.Unscoped().Delete(&DictDetailModel{}, parentID).Error
}

func (e *dictDetailEntity) Update(ctx context.Context, req *DictDetailModel) (*DictDetailModel, error) {
	var existing DictDetailModel

	if err := e.conn.WithContext(ctx).First(&existing, req.ID).Error; err != nil {
		return nil, errors.New("record not exist")
	}

	var count int64
	if err := e.conn.WithContext(ctx).
		Model(&DictDetailModel{}).
		Where("dict_id = ? AND id <> ? AND (label = ? OR value = ?)",
			req.DictModelID, req.ID, req.Label, req.Value).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("duplicate label or value in the same dict")
	}

	updates := map[string]interface{}{
		"label":       req.Label,
		"value":       req.Value,
		"sort":        req.Sort,
		"description": req.Description,
		"parent_id":   req.ParentID, // NULL-safe
	}

	if err := e.conn.WithContext(ctx).
		Model(&existing).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}
