package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	"server/global"
	modelSysTool "server/model/sysTool"
)

func ClearTable(db *gorm.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	return db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", tableName, compareField), time.Now().Add(-duration)).Error
}

func AddJob(instance *modelSysTool.CronModel) (entryId int, err error) {
	switch instance.Method {
	case "clearTable":
		var clearTableSlice modelSysTool.ClearTableSlice
		for _, v := range instance.ExtraParams.TableInfo {
			var clearTableModel modelSysTool.ClearTable
			clearTableModel.TableName = v.TableName
			clearTableModel.CompareField = v.CompareField
			clearTableModel.Interval = v.Interval
			clearTableSlice.TableInfo = append(clearTableSlice.TableInfo, clearTableModel)
		}
		tmpId, err := global.TD27_CRON.AddJob(instance.Expression, &clearTableSlice)
		if err != nil {
			return -1, err
		}
		entryId = int(tmpId)
	default:
		return -1, fmt.Errorf("不支持的method")
	}
	return entryId, err
}

func GetEntries() (entrySlice []int) {
	for _, v := range global.TD27_CRON.Entries() {
		entrySlice = append(entrySlice, int(v.ID))
	}

	return
}
