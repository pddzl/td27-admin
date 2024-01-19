package base

import (
	"server/global"
)

type JwtBlackListModel struct {
	global.TD27_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (jl *JwtBlackListModel) TableName() string {
	return "base_jwtBlackList"
}
