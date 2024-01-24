package monitor

import (
	"server/global"
	modelMonitor "server/model/monitor"
	monitorReq "server/model/monitor/request"
)

type OperationLogService struct{}

// CreateOperationLog 创建操作日志
func (o *OperationLogService) CreateOperationLog(operationRecord modelMonitor.OperationLogModel) error {
	return global.TD27_DB.Create(&operationRecord).Error
}

// GetOperationLogList 分页获取操作日志
func (o *OperationLogService) GetOperationLogList(orSp monitorReq.OrSearchParams) ([]modelMonitor.OperationLogModel, int64, error) {
	limit := orSp.PageSize
	offset := orSp.PageSize * (orSp.Page - 1)
	db := global.TD27_DB.Model(&modelMonitor.OperationLogModel{})
	var olList []modelMonitor.OperationLogModel

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
		return olList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if orSp.Asc {
			err = db.Find(&olList).Error
		} else {
			err = db.Order("id desc").Find(&olList).Error
		}
	}
	return olList, total, err
}

// DeleteOperationLog 删除操作日志
func (o *OperationLogService) DeleteOperationLog(id uint) error {
	return global.TD27_DB.Unscoped().Delete(&modelMonitor.OperationLogModel{}, id).Error
}

// DeleteOperationLogByIds 批量删除操作日志
func (o *OperationLogService) DeleteOperationLogByIds(ids []uint) error {
	return global.TD27_DB.Unscoped().Delete(&[]modelMonitor.OperationLogModel{}, ids).Error
}
