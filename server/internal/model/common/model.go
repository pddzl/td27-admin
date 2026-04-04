package common

import (
	"time"

	"gorm.io/gorm"
)

// Td27Model 基础模型
// 注意：时间字段不指定 type，让 GORM 根据数据库类型自动选择
// MySQL: datetime, PostgreSQL: timestamp
type Td27Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`                           // 主键 ID
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;default:null"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;default:null"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                                 // 删除时间
}
