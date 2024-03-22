package sysTool

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"

	"server/global"
)

type Method struct {
	ClearTable string
	Shell      string
}

var CronMethod = Method{
	ClearTable: "clearTable",
	Shell:      "shell",
}

type CronModel struct {
	global.TD27_MODEL
	Name        string      `json:"name" gorm:"column:name;unique;comment:任务名称" binding:"required"`
	Method      string      `json:"method" gorm:"column:method;not null;comment:任务方法" binding:"required"`
	Expression  string      `json:"expression" gorm:"column:expression;not null;comment:表达式" binding:"required"`
	Strategy    string      `json:"strategy" gorm:"column:strategy;type:enum('always','once');default:always;comment:执行策略" binding:"omitempty,oneof=always once"`
	Open        bool        `json:"open" gorm:"column:open;comment:活跃状态"`
	ExtraParams ExtraParams `json:"extraParams" gorm:"column:extraParams;type:json;comment:额外参数"`
	EntryId     int         `json:"entryId" gorm:"column:entryId;comment:cron ID"`
	Comment     string      `json:"comment" gorm:"column:comment;comment:备注"`
}

type ExtraParams struct {
	TableInfo []ClearTable `json:"tableInfo,omitempty"` // for clearTable
	Command   string       `json:"command,omitempty"`   // for shell
}

func (e ExtraParams) Value() (driver.Value, error) {
	b, err := json.Marshal(e)
	return string(b), err
}

func (e *ExtraParams) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), e)
}

type ClearTable struct {
	TableName    string `json:"tableName,omitempty"`
	CompareField string `json:"compareField,omitempty"`
	Interval     string `json:"interval,omitempty"`
}

func (cm *CronModel) TableName() string {
	return "sysTool_cron"
}

func (cm *CronModel) Run() {
	global.TD27_LOG.Info("[CRON]", zap.String("START", cm.Method))
	switch cm.Method {
	case "clearTable":
		for _, v := range cm.ExtraParams.TableInfo {
			duration, err := time.ParseDuration(v.Interval)
			if err != nil {
				global.TD27_LOG.Error(fmt.Sprintf("parse duration err: %v", err))
			}
			if duration < 0 {
				global.TD27_LOG.Error("parse duration < 0")
			}
			err = global.TD27_DB.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", v.TableName, v.CompareField), time.Now().Add(-duration)).Error
			if err != nil {
				global.TD27_LOG.Error(fmt.Sprintf("delete err: %v", err))
			}
		}
	default:
		global.TD27_LOG.Error("unsupport method")
	}
	if cm.Strategy == "once" {
		global.TD27_LOG.Info("[CRON]", zap.String(cm.Name, "stop"))
		global.TD27_CRON.Remove(cron.EntryID(cm.EntryId))
		global.TD27_DB.Model(cm).Updates(map[string]interface{}{"open": false, "entryId": 0})
	}
}
