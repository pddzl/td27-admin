package sysTool

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"server/global"
	commonReq "server/model/common/request"
	modelSysTool "server/model/sysTool"
	"server/utils"
)

type CronService struct{}

// GetCronList 分页获取cron
func (cs *CronService) GetCronList(pageInfo commonReq.PageInfo) (cronModelList []modelSysTool.CronModel, total int64, err error) {
	db := global.TD27_DB.Model(&modelSysTool.CronModel{})

	// 计算记录数量
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("count err %v", err)
	}

	// 分页
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	if pageInfo.PageSize >= 0 && pageInfo.Page > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&cronModelList).Error
	return
}

// AddCron 添加cron
func (cs *CronService) AddCron(cronModel *modelSysTool.CronModel) (*modelSysTool.CronModel, error) {
	// 开启cron
	if cronModel.Open {
		entryId, err := global.TD27_CRON.AddJob(cronModel.Expression, cronModel)
		if err != nil {
			return nil, err
		} else {
			cronModel.EntryId = int(entryId)
		}
	}
	err := global.TD27_DB.Create(cronModel).Error
	return cronModel, err
}

// DeleteCron 删除cron
func (cs *CronService) DeleteCron(id uint) error {
	var cronModel modelSysTool.CronModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&cronModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录未找到")
	}
	// 删除定时任务
	global.TD27_CRON.Remove(cron.EntryID(cronModel.EntryId))
	// 删除数据库记录
	return global.TD27_DB.Unscoped().Delete(&cronModel).Error
}

// DeleteCronByIds 批量删除cron
func (cs *CronService) DeleteCronByIds(ids []uint) error {
	var cronModels []modelSysTool.CronModel
	global.TD27_DB.Find(&cronModels, ids)
	// 删除定时任务
	for _, value := range cronModels {
		global.TD27_CRON.Remove(cron.EntryID(value.EntryId))
	}
	// 删除数据库记录
	return global.TD27_DB.Unscoped().Delete(&cronModels).Error
}

// EditCron 编辑cron
func (cs *CronService) EditCron(instance *modelSysTool.CronModel) (*modelSysTool.CronModel, error) {
	if errors.Is(global.TD27_DB.Where("id = ?", instance.ID).First(&modelSysTool.CronModel{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("记录不存在")
	}

	// params 拼接
	var extraParams modelSysTool.ExtraParams
	for _, v := range instance.ExtraParams.TableInfo {
		var clearTable modelSysTool.ClearTable
		clearTable.TableName = v.TableName
		clearTable.CompareField = v.CompareField
		clearTable.Interval = v.Interval
		extraParams.TableInfo = append(extraParams.TableInfo, clearTable)
	}
	instance.ExtraParams = extraParams

	if instance.Open {
		if !utils.IsContain(utils.GetEntries(), instance.EntryId) {
			entryId, err := global.TD27_CRON.AddJob(instance.Expression, instance)
			if err != nil {
				return nil, err
			} else {
				instance.Open = true
				instance.EntryId = int(entryId)
			}
		}
	} else {
		if instance.EntryId != 0 {
			global.TD27_CRON.Remove(cron.EntryID(instance.EntryId))
			instance.EntryId = 0
		}
		instance.Open = false
	}
	err := global.TD27_DB.Omit("created_at").Save(instance).Error

	return instance, err
}

// SwitchOpen 切换cron活跃状态
func (cs *CronService) SwitchOpen(id uint, open bool) (resId int, err error) {
	var cronModel modelSysTool.CronModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&cronModel).Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("记录未找到")
	}

	// 判断cron是否已经运行
	if open && !utils.IsContain(utils.GetEntries(), cronModel.EntryId) {
		entryId, err := global.TD27_CRON.AddJob(cronModel.Expression, &cronModel)
		if err != nil {
			return cronModel.EntryId, err
		} else {
			err = global.TD27_DB.Model(&cronModel).Updates(map[string]interface{}{"open": true, "entryId": entryId}).Error
		}
		resId = int(entryId)
	} else {
		if cronModel.EntryId != 0 {
			global.TD27_CRON.Remove(cron.EntryID(cronModel.EntryId))
		}
		err = global.TD27_DB.Model(&cronModel).Updates(map[string]interface{}{"open": false, "entryId": 0}).Error
		resId = 0
	}

	return
}
