package system

import (
	"server/global"
)

type JwtBlacklist struct {
	global.TD27_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
