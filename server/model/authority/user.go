package authority

import (
	"server/global"
)

type UserModel struct {
	global.TD27_MODEL
	Username    string `json:"username" gorm:"index;unique;comment:用户名" binding:"required"` // 用户名
	Password    string `json:"-"  gorm:"not null;comment:密码"`
	Phone       string `json:"phone"  gorm:"comment:手机号"`                          // 手机号
	Email       string `json:"email"  gorm:"comment:邮箱" binding:"omitempty,email"` // 邮箱
	Active      bool   `json:"active"`                                             // 是否活跃
	RoleModelID uint   `json:"roleId" gorm:"not null" binding:"required"`          // 角色ID
}

func (UserModel) TableName() string {
	return "authority_user"
}
