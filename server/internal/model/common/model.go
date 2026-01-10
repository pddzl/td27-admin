package common

import (
	"time"

	"gorm.io/gorm"
)

type Td27Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`                                          // 主键 ID
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;default:null"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;default:null"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                                                // 删除时间
}
