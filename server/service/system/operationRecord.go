package system

import (
	"server/global"
	modelSystem "server/model/system"
	systemReq "server/model/system/request"
)

type OperationRecordService struct{}

// CreateOperationRecord 创建记录
func (o *OperationRecordService) CreateOperationRecord(operationRecord modelSystem.OperationRecord) error {
	return global.TD27_DB.Create(&operationRecord).Error
}

// GetOperationRecordList 分页获取操作记录
func (o *OperationRecordService) GetOperationRecordList(orSp systemReq.OrSearchParams) ([]modelSystem.OperationRecord, int64, error) {
	limit := orSp.PageSize
	offset := orSp.PageSize * (orSp.Page - 1)
	db := global.TD27_DB.Model(&modelSystem.OperationRecord{})
	var orList []modelSystem.OperationRecord

	if orSp.Path != "" {
		db = db.Where("path LIKE ?", "%"+orSp.Path+"%")
	}

	if orSp.Method != "" {
		db = db.Where("method = ?", orSp.Method)
	}

	if orSp.Status != 0 {
		db = db.Where("status = ?", orSp.Status)
	}

	var total int64
	err := db.Count(&total).Error

	if err != nil {
		return orList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if orSp.Asc {
			err = db.Find(&orList).Error
		} else {
			err = db.Order("id desc").Find(&orList).Error
		}
	}
	return orList, total, err
}

// DeleteOperation 删除操作记录
func (o *OperationRecordService) DeleteOperation(id uint) error {
	return global.TD27_DB.Unscoped().Delete(&modelSystem.OperationRecord{}, id).Error
}

// DeleteOperationByIds 批量删除操作记录
func (o *OperationRecordService) DeleteOperationByIds(ids []uint) error {
	return global.TD27_DB.Unscoped().Delete(&[]modelSystem.OperationRecord{}, ids).Error
}
