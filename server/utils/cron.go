package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	"server/global"
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

func GetEntries() (entrySlice []int) {
	for _, v := range global.TD27_CRON.Entries() {
		entrySlice = append(entrySlice, int(v.ID))
	}

	return
}
