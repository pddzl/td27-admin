package sysTool

import (
	"fmt"
	"server/global"
	"time"
)

//type ClearTable struct {
//	TableName    string `json:"tableName"`
//	CompareField string `json:"compareField"`
//	Interval     string `json:"interval"`
//}

type ClearTableSlice struct {
	TableInfo []ClearTable `json:"tableInfo"`
}

func (c *ClearTableSlice) Run() {
	if global.TD27_DB == nil {
		global.TD27_LOG.Error("db Cannot be empty")
	}

	for _, v := range c.TableInfo {
		duration, err := time.ParseDuration(v.Interval)
		if err != nil {
			global.TD27_LOG.Error(fmt.Sprintf("parse err: %v", err))
			return
		}
		if duration < 0 {
			global.TD27_LOG.Error("parse duration < 0")
			return
		}
		err = global.TD27_DB.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", v.TableName, v.CompareField), time.Now().Add(-duration)).Error
		if err != nil {
			global.TD27_LOG.Error(fmt.Sprintf("exec err: %v", err))
			return
		}
	}
}
