package sysSet

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	modelSysSet "server/model/sysSet"
	sysSetReq "server/model/sysSet/request"
)

type DictDetailService struct{}

func (dds *DictDetailService) GetDictDetail(searchParams sysSetReq.DictDetailSearchParams) ([]*modelSysSet.DictDetailModel, int64, error) {
	var roots []modelSysSet.DictDetailModel
	var total int64

	// 1. Count total root nodes
	if err := global.TD27_DB.
		Model(&modelSysSet.DictDetailModel{}).
		Where("dict_id = ? AND parent_id IS NULL", searchParams.DictID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 2. Get paginated root nodes
	if err := global.TD27_DB.
		Where("dict_id = ? AND parent_id IS NULL", searchParams.DictID).
		Order("sort asc").
		Limit(searchParams.PageSize).
		Offset((searchParams.Page - 1) * searchParams.PageSize).
		Find(&roots).Error; err != nil {
		return nil, 0, err
	}

	if len(roots) == 0 {
		return []*modelSysSet.DictDetailModel{}, total, nil
	}

	// 3. Collect root IDs
	rootIDs := make([]uint, 0, len(roots))
	for _, r := range roots {
		rootIDs = append(rootIDs, r.ID)
	}

	// 4. Load all descendants for these roots
	var all []modelSysSet.DictDetailModel
	if err := global.TD27_DB.
		Where("dict_id = ? AND parent_id IS NOT NULL", searchParams.DictID).
		Order("sort asc").
		Find(&all).Error; err != nil {
		return nil, 0, err
	}

	// 5. Build tree
	m := make(map[uint]*modelSysSet.DictDetailModel)
	for i := range roots {
		roots[i].Children = []*modelSysSet.DictDetailModel{}
		m[roots[i].ID] = &roots[i]
	}
	for i := range all {
		all[i].Children = []*modelSysSet.DictDetailModel{}
		m[all[i].ID] = &all[i]
	}

	for i := range all {
		if all[i].ParentID != nil {
			if parent, ok := m[uint(*all[i].ParentID)]; ok {
				parent.Children = append(parent.Children, m[all[i].ID])
			}
		}
	}

	// Convert roots to []*modelSysSet.DictDetailModel
	result := make([]*modelSysSet.DictDetailModel, 0, len(roots))
	for i := range roots {
		result = append(result, &roots[i])
	}

	return result, total, nil
}

func (dds *DictDetailService) AddDictDetail(instance *modelSysSet.DictDetailModel) (*modelSysSet.DictDetailModel, error) {
	var existing modelSysSet.DictDetailModel

	// 🔹 check duplicate by label
	if err := global.TD27_DB.
		Where("dict_id = ? AND label = ?", instance.DictModelID, instance.Label).
		First(&existing).Error; err == nil {
		return nil, errors.New("duplicate label in the same dict")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 🔹 check duplicate by value
	if err := global.TD27_DB.
		Where("dict_id = ? AND value = ?", instance.DictModelID, instance.Value).
		First(&existing).Error; err == nil {
		return nil, errors.New("duplicate value in the same dict")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 🔹 create record
	if err := global.TD27_DB.Create(instance).Error; err != nil {
		return nil, err
	}
	return instance, nil
}

func (dds *DictDetailService) DelDictDetail(id uint) (err error) {
	var dictDetail modelSysSet.DictDetailModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&dictDetail).Error, gorm.ErrRecordNotFound) {
		return errors.New("DictDetail not exist")
	}

	// delete children recursively
	if err = delChildren(id); err != nil {
		return err
	}

	// delete self
	if err = global.TD27_DB.Unscoped().Delete(&dictDetail).Error; err != nil {
		return err
	}

	return nil
}

func delChildren(parentID uint) error {
	var children []modelSysSet.DictDetailModel
	if err := global.TD27_DB.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return err
	}

	for _, child := range children {
		// recursive delete
		if err := delChildren(child.ID); err != nil {
			return err
		}
		if err := global.TD27_DB.Unscoped().Delete(&child).Error; err != nil {
			return err
		}
	}

	return nil
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
	if instance.ParentID == nil {
		// gorm.Updates(instance) won't overwrite the old value with NULL unless you explicitly allow it
		global.TD27_DB.Model(&existing).UpdateColumn("parent_id", instance.ParentID)
	}

	return &existing, nil
}
