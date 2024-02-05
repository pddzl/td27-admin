package sysTool

import (
	"database/sql/driver"
	"encoding/json"
	"server/global"
)

type CronModel struct {
	global.TD27_MODEL
	Name        string      `json:"name" gorm:"column:name;unique;comment:任务名称" validate:"required"`
	Method      string      `json:"method" gorm:"column:method;unique;not null;comment:任务方法" validate:"required"`
	Expression  string      `json:"expression" gorm:"column:expression;not null;comment:表达式" validate:"required"`
	Strategy    string      `json:"strategy" gorm:"column:strategy;type:enum('always', 'once');default:always;comment:执行策略" validate:"omitempty,oneof=always once"`
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
