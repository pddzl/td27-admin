package sysSet

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	modelSysSet "server/model/sysSet"
	sysSetReq "server/model/sysSet/request"
)

type DictDetailService struct{}

func (dds *DictDetailService) GetDictDetail(searchParams sysSetReq.DictDetailSearchParams) ([]modelSysSet.DictDetailModel, int64, error) {
	limit := searchParams.PageSize
	offset := searchParams.PageSize * (searchParams.Page - 1)
	db := global.TD27_DB.Model(&modelSysSet.DictDetailModel{}).Where("dict_id = ?", searchParams.DictID)
	var dictDetailList []modelSysSet.DictDetailModel
	var total int64
	err := db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&dictDetailList).Error
	return dictDetailList, total, err

}

func (dds *DictDetailService) AddDictDetail(instance *modelSysSet.DictDetailModel) (*modelSysSet.DictDetailModel, error) {
	var existing modelSysSet.DictDetailModel

	// check duplicate by label
	if err := global.TD27_DB.Where("dict_id = ? AND label = ?", instance.DictModelID, instance.Label).First(&existing).Error; err == nil {
		return nil, errors.New("duplicate label in the same dict")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// check duplicate by value
	if err := global.TD27_DB.Where("dict_id = ? AND value = ?", instance.DictModelID, instance.Value).First(&existing).Error; err == nil {
		return nil, errors.New("duplicate value in the same dict")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// safe to create
	if err := global.TD27_DB.Create(instance).Error; err != nil {
		return nil, err
	}
	return instance, nil
}

func (dds *DictDetailService) DelDictDetail(id uint) (err error) {
	var dictDetailModel modelSysSet.DictDetailModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&dictDetailModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("not exist DictDetail")
	}

	err = global.TD27_DB.Unscoped().Delete(&dictDetailModel).Error
	if err != nil {
		return err
	}

	return
}

func (dds *DictDetailService) EditDictDetail(instance *modelSysSet.DictDetailModel) (*modelSysSet.DictDetailModel, error) {
	var existing, checkDuplicate modelSysSet.DictDetailModel

	// 1. check if record exists
	if errors.Is(global.TD27_DB.First(&existing, instance.ID).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not exist")
	}

	// 2. check duplicate label (exclude current record)
	if err := global.TD27_DB.
		Where("dict_id = ? AND label = ? AND `id` <> ?", instance.DictModelID, instance.Label, instance.ID).
		First(&checkDuplicate).Error; err == nil {
		return nil, errors.New("duplicate label in the same dict")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 3. check duplicate value (exclude current record)
	if err := global.TD27_DB.Where("dict_id = ? AND value = ? AND id <> ?", instance.DictModelID, instance.Value, instance.ID).
		First(&checkDuplicate).Error; err == nil {
		return nil, errors.New("duplicate value in the same dict")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 4. update
	if err := global.TD27_DB.Model(&existing).Updates(instance).Error; err != nil {
		return nil, err
	}

	return instance, nil
}
