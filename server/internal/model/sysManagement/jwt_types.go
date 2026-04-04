package sysManagement

import (
	"github.com/golang-jwt/jwt/v4"
)

// RoleInfo 角色信息
type RoleInfo struct {
	ID       uint   `json:"id"`
	RoleName string `json:"roleName"`
}

// CustomClaims Custom claims structure（支持多角色）
type CustomClaims struct {
	ID         uint       `json:"ID"`
	Username   string     `json:"username"`
	Roles      []RoleInfo `json:"roles"` // 多角色信息
	BufferTime int64      `json:"bufferTime"`
	jwt.RegisteredClaims
}

// GetAllRoleIDs 获取所有角色ID
func (c *CustomClaims) GetAllRoleIDs() []uint {
	ids := make([]uint, 0, len(c.Roles))
	for _, role := range c.Roles {
		ids = append(ids, role.ID)
	}
	return ids
}

// HasRole 检查是否有指定角色
func (c *CustomClaims) HasRole(roleID uint) bool {
	for _, role := range c.Roles {
		if role.ID == roleID {
			return true
		}
	}
	return false
}
