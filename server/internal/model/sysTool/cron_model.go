package sysTool

import (
	"database/sql/driver"
	"encoding/json"

	"server/internal/model/common"
)

type Method struct {
	ClearTable string
	ClearCache string
	Shell      string
}

var CronMethod = Method{
	ClearTable: "clearTable",
	ClearCache: "clearCache",
	Shell:      "shell",
}

type CronModel struct {
	common.Td27Model
	Name        string      `json:"name" gorm:"column:name;unique;comment:任务名称" binding:"required"`
	Method      string      `json:"method" gorm:"column:method;not null;comment:任务方法" binding:"required"`
	Expression  string      `json:"expression" gorm:"column:expression;not null;comment:表达式" binding:"required"`
	Strategy    string      `json:"strategy" gorm:"column:strategy;size:20;default:'always';comment:执行策略" binding:"omitempty,oneof=always once"`
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
	return "sys_tool_cron"
}
