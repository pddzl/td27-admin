package sysMonitor

import (
	"context"
	"errors"
	"server/internal/global"

	"gorm.io/gorm"
)

type OperationLogRepository interface {
	Create(context.Context, *OperationLogModel) error
	List(context.Context, *OrListReq) ([]*OperationLogModel, int64, error)
	Delete(context.Context, uint) error
	DeleteByIds(context.Context, []uint) (int64, error)
}

type operationLogEntity struct {
	conn *gorm.DB
}

func NewOperationLogEntity(conn *gorm.DB) OperationLogRepository {
	return &operationLogEntity{conn: conn}
}

func (e *operationLogEntity) Create(ctx context.Context, req *OperationLogModel) error {
	if req == nil {
		return errors.New("operation log is nil")
	}

	// result := e.conn.WithContext(ctx).Model(&OperationLogModel{}).Create(req)
	// middleware use it by global variable, so it must use global.TD27_DB
	// instead of e.conn
	result := global.TD27_DB.WithContext(ctx).Model(&OperationLogModel{}).Create(req)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil
}

func (e *operationLogEntity) List(ctx context.Context, req *OrListReq) ([]*OperationLogModel, int64, error) {
	req.Normalize()

	var operationLogs []*OperationLogModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&OperationLogModel{})

	// Filters
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}

	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}

	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}

	// Count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Order
	if req.Asc {
		db = db.Order("id asc")
	} else {
		db = db.Order("id desc")
	}

	// Pagination
	if err := db.Limit(req.PageSize).Offset(req.Offset()).Find(&operationLogs).Error; err != nil {
		return nil, 0, err
	}

	return operationLogs, total, nil
}

func (e *operationLogEntity) Delete(ctx context.Context, id uint) error {
	return e.conn.WithContext(ctx).Unscoped().Delete(&OperationLogModel{}, id).Error
}

func (e *operationLogEntity) DeleteByIds(ctx context.Context, ids []uint) (int64, error) {
	// Safety Guard
	if len(ids) == 0 {
		return 0, errors.New("ids is empty")
	}

	result := e.conn.WithContext(ctx).Where("id IN ?", ids).Delete(&OperationLogModel{})

	return result.RowsAffected, result.Error
}
