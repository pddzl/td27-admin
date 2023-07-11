package system

import (
	"server/global"
	"server/model/system"
)

type OperationRecordService struct{}

// CreateSysOperationRecord 创建记录
func (o *OperationRecordService) CreateSysOperationRecord(operationRecord system.OperationRecord) (err error) {
	err = global.TD27_DB.Create(&operationRecord).Error
	return err
}
