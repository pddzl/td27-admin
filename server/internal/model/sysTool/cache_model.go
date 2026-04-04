package sysTool

import (
	"time"

	"server/internal/model/common"
)

// CacheModel PostgreSQL 缓存表模型
type CacheModel struct {
	common.Td27Model
	Key        string    `json:"key" gorm:"uniqueIndex;size:255;comment:缓存键"`
	Value      string    `json:"value" gorm:"type:text;comment:缓存值"`
	ExpiresAt  time.Time `json:"expiresAt" gorm:"index;comment:过期时间"`
}

func (CacheModel) TableName() string {
	return "sys_tool_cache"
}
