package sysSet

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	modelSysSet "server/model/sysSet"
)

type DictService struct{}

func (ds *DictService) GetDict() ([]modelSysSet.DictModel, error) {
	var dictList []modelSysSet.DictModel
	err := global.TD27_DB.Find(&dictList).Error

	return dictList, err
}

func (ds *DictService) AddDict(instance *modelSysSet.DictModel) (*modelSysSet.DictModel, error) {
	err := global.TD27_DB.Create(instance).Error
	return instance, err

}

func (ds *DictService) DelDict(id uint) (err error) {
	var dictModel modelSysSet.DictModel

	// load dict and details in one query
	if err = global.TD27_DB.Preload("DictDetails").First(&dictModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("dict not found")
		}
		return err
	}

	// prevent deletion if details exist
	if len(dictModel.DictDetails) > 0 {
		return errors.New("cannot delete: dict has DictDetails")
	}

	// delete permanently
	if err = global.TD27_DB.Unscoped().Delete(&dictModel).Error; err != nil {
		return err
	}

	return nil
}

func (ds *DictService) EditDict(instance *modelSysSet.DictModel) (*modelSysSet.DictModel, error) {
	var dictModel modelSysSet.DictModel
	if errors.Is(global.TD27_DB.Where("id = ?", instance.ID).First(&dictModel).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not exist Dict")
	}

	err := global.TD27_DB.Model(&dictModel).Update("ch_name", instance.CHName).Error
	return instance, err
}
