package sysTool

import (
	"context"
	"errors"
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"server/internal/global"
	"server/internal/model/common"
	"server/internal/pkg"
)

type CronRepository interface {
	List(context.Context, *common.PageInfo) ([]*CronModel, int64, error)
	Create(context.Context, *CronModel) (*CronModel, error)
	Delete(context.Context, uint) error
	DeleteByIds(context.Context, []uint) error
	Update(context.Context, *CronModel) (*CronModel, error)
	SwitchOpen(context.Context, uint, bool) (int, error)
}

type cronEntity struct {
	conn *gorm.DB
}

func NewCronRepository(conn *gorm.DB) CronRepository {
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

	// Pagination
	if err := db.Limit(req.PageSize).Offset(req.Offset()).Find(&crons).Error; err != nil {
		return nil, 0, err
	}

	return crons, total, nil
}

func (e *cronEntity) Create(ctx context.Context, req *CronModel) (*CronModel, error) {
	// 开启 cron
	if req.Open {
		entryId, err := global.TD27_CRON.AddJob(req.Expression, req)
		if err != nil {
			return nil, err
		}

		req.EntryId = int(entryId)
	}
	err := e.conn.WithContext(ctx).Create(req).Error
	return req, err
}

// Delete 删除cron
func (e *cronEntity) Delete(ctx context.Context, id uint) error {
	db := e.conn.WithContext(ctx)

	var cronModel CronModel
	if errors.Is(db.Where("id = ?", id).First(&cronModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录未找到")
	}
	// 删除定时任务
	global.TD27_CRON.Remove(cron.EntryID(cronModel.EntryId))
	// 删除数据库记录
	return db.Unscoped().Delete(&cronModel).Error
}

// DeleteByIds 批量删除cron
func (e *cronEntity) DeleteByIds(ctx context.Context, ids []uint) error {
	var cronModels []CronModel
	db := e.conn.WithContext(ctx)
	db.Find(&cronModels, ids)
	// 删除定时任务
	for _, value := range cronModels {
		global.TD27_CRON.Remove(cron.EntryID(value.EntryId))
	}
	// 删除数据库记录
	return db.Unscoped().Delete(&cronModels).Error
}

// Update 编辑cron
func (e *cronEntity) Update(ctx context.Context, req *CronModel) (*CronModel, error) {
	db := e.conn.WithContext(ctx)

	// Load existing record
	var old CronModel
	if err := db.Where("id = ?", req.ID).First(&old).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("记录不存在")
		}
		return nil, fmt.Errorf("查询定时任务失败: %w", err)
	}

	// Normalize extra params (pure transformation)
	req.ExtraParams = normalizeExtraParams(req.ExtraParams)

	// Cron state transition
	newEntryID := old.EntryId

	if req.Open {
		// Need to start cron
		if old.EntryId == 0 {
			entryID, err := global.TD27_CRON.AddJob(req.Expression, req)
			if err != nil {
				return nil, fmt.Errorf("添加定时任务失败: %w", err)
			}
			newEntryID = int(entryID)
		}
	} else {
		// Need to stop cron
		if old.EntryId != 0 {
			global.TD27_CRON.Remove(cron.EntryID(old.EntryId))
			newEntryID = 0
		}
	}

	// Persist update
	update := map[string]interface{}{
		"name":        req.Name,
		"method":      req.Method,
		"expression":  req.Expression,
		"strategy":    req.Strategy,
		"open":        req.Open,
		"extraParams": req.ExtraParams,
		"entryId":     newEntryID,
		"comment":     req.Comment,
	}

	if err := db.Model(&CronModel{}).
		Where("id = ?", req.ID).
		Updates(update).Error; err != nil {
		return nil, fmt.Errorf("更新定时任务失败: %w", err)
	}

	req.EntryId = newEntryID
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

// SwitchOpen 切换cron活跃状态
func (e *cronEntity) SwitchOpen(ctx context.Context, id uint, open bool) (resId int, err error) {
	db := e.conn.WithContext(ctx)
	var cronModel CronModel
	if errors.Is(db.Where("id = ?", id).First(&cronModel).Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("记录未找到")
	}

	// 判断 cron是否已经运行
	if open && !pkg.IsContain(pkg.GetEntries(), cronModel.EntryId) {
		entryId, err := global.TD27_CRON.AddJob(cronModel.Expression, &cronModel)
		if err != nil {
			return cronModel.EntryId, err
		}
		err = db.Model(&cronModel).Updates(map[string]interface{}{"open": true, "entryId": entryId}).Error
		resId = int(entryId)
	} else {
		if cronModel.EntryId != 0 {
			global.TD27_CRON.Remove(cron.EntryID(cronModel.EntryId))
		}
		err = db.Model(&cronModel).Updates(map[string]interface{}{"open": false, "entryId": 0}).Error
		resId = 0
	}

	return
}
