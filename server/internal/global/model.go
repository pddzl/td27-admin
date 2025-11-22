package global

import (
	"time"
	
	"gorm.io/gorm"
)

type TD27_MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"`                                          // 主键ID
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;default:null"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;default:null"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                                                // 删除时间
}
