package sysTool

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"server/internal/model/common"
)

type CronRepository interface {
	List(context.Context, *common.PageInfo) ([]*CronModel, int64, error)
	Create(context.Context, *CronModel) (*CronModel, error)
	Delete(context.Context, uint) error
	DeleteByIds(context.Context, []uint) error
	Update(context.Context, *CronModel) (*CronModel, error)
	SwitchOpen(context.Context, uint, bool) (int, error)
	FindByID(context.Context, uint) (*CronModel, error)
}

type cronEntity struct {
	conn *gorm.DB
}

func NewCronEntity(conn *gorm.DB) CronRepository {
	return &cronEntity{conn: conn}
}

func (e *cronEntity) List(ctx context.Context, req *common.PageInfo) ([]*CronModel, int64, error) {
	req.Normalize()

	var crons []*CronModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&CronModel{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Limit(req.PageSize).Offset(req.Offset()).Find(&crons).Error; err != nil {
		return nil, 0, err
	}

	return crons, total, nil
}

func (e *cronEntity) FindByID(ctx context.Context, id uint) (*CronModel, error) {
	var m CronModel
	if err := e.conn.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (e *cronEntity) Create(ctx context.Context, req *CronModel) (*CronModel, error) {
	if err := e.conn.WithContext(ctx).Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (e *cronEntity) Delete(ctx context.Context, id uint) error {
	return e.conn.WithContext(ctx).Unscoped().Delete(&CronModel{}, id).Error
}

func (e *cronEntity) DeleteByIds(ctx context.Context, ids []uint) error {
	return e.conn.WithContext(ctx).Unscoped().Delete(&CronModel{}, ids).Error
}

func (e *cronEntity) Update(ctx context.Context, req *CronModel) (*CronModel, error) {
	if err := e.conn.WithContext(ctx).Model(&CronModel{}).
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"name":        req.Name,
			"method":      req.Method,
			"expression":  req.Expression,
			"strategy":    req.Strategy,
			"open":        req.Open,
			"extraParams": normalizeExtraParams(req.ExtraParams),
			"comment":     req.Comment,
		}).Error; err != nil {
		return nil, fmt.Errorf("更新定时任务失败: %w", err)
	}
	return req, nil
}

func normalizeExtraParams(p ExtraParams) ExtraParams {
	var result ExtraParams
	for _, v := range p.TableInfo {
		result.TableInfo = append(result.TableInfo, ClearTable{
			TableName:    v.TableName,
			CompareField: v.CompareField,
			Interval:     v.Interval,
		})
	}
	return result
}

func (e *cronEntity) SwitchOpen(ctx context.Context, id uint, open bool) (int, error) {
	var cronModel CronModel
	if errors.Is(e.conn.WithContext(ctx).Where("id = ?", id).First(&cronModel).Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("记录未找到")
	}

	entryId := 0
	if open {
		entryId = cronModel.EntryId
	}

	if err := e.conn.WithContext(ctx).Model(&cronModel).
		Updates(map[string]interface{}{"open": open, "entryId": entryId}).Error; err != nil {
		return 0, err
	}

	return entryId, nil
}
