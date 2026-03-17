package sysManagement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"server/internal/model/common"
)

type DictRepository interface {
	FindOne(context.Context, uint) (*DictModel, error)
	List(context.Context, *common.PageInfo) ([]*DictModel, int64, error)
	Create(context.Context, *DictModel) (*DictModel, error)
	Delete(context.Context, uint) error
	Update(context.Context, *UpdateDictReq) error
}

type dictEntity struct {
	conn *gorm.DB
}

func NewDictEntity(conn *gorm.DB) DictRepository {
	return &dictEntity{conn: conn}
}

func (e *dictEntity) FindOne(ctx context.Context, id uint) (*DictModel, error) {
	var dict DictModel

	err := e.conn.WithContext(ctx).
		Preload("DictDetails").
		First(&dict, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("dict not found: id=%d", id)
		}
		return nil, err
	}

	return &dict, nil
}

func (e *dictEntity) List(ctx context.Context, req *common.PageInfo) ([]*DictModel, int64, error) {
	var list []*DictModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&DictModel{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	req.Normalize()
	if err := db.
		Limit(req.PageSize).
		Offset(req.Offset()).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (e *dictEntity) Create(ctx context.Context, req *DictModel) (*DictModel, error) {
	db := e.conn.WithContext(ctx)

	err := db.
		Where("cn_name = ? OR en_name = ?", req.CNName, req.ENName).
		First(&DictModel{}).Error

	if err == nil {
		return nil, errors.New("dict already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = db.Create(req).Error; err != nil {
		return nil, err
	}

	return req, nil
}

func (e *dictEntity) Delete(ctx context.Context, id uint) (err error) {
	db := e.conn.WithContext(ctx)

	var dictModel DictModel

	// load dict and details in one query
	if err = db.Preload("DictDetails").First(&dictModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("dict not found")
		}
		return err
	}

	// prevent deletion if details exist
	if len(dictModel.DictDetails) > 0 {
		return errors.New("cannot delete: dict has DictDetails")
	}

	result := db.Unscoped().Delete(&DictModel{}, id)

	if result.RowsAffected == 0 {
		return errors.New("dict not found")
	}

	return result.Error
}

func (e *dictEntity) Update(ctx context.Context, req *UpdateDictReq) error {

	return e.conn.WithContext(ctx).Model(&DictModel{}).Where("id = ?", req.ID).Update("cn_name", req.CNName).Error
}
