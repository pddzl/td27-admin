package sysManagement

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"server/internal/global"
)

// PermissionRepository Permission仓库接口
type PermissionRepository interface {
	List(ctx context.Context, roleId uint, pd PermissionDomain) ([]PermissionModel, error)
}

type permissionRepo struct {
	conn *gorm.DB
}

func NewPermissionRepo(conn *gorm.DB) PermissionRepository {
	return &permissionRepo{conn: conn}
}

func (i *permissionRepo) List(ctx context.Context, roleId uint, pd PermissionDomain) ([]PermissionModel, error) {
	var permissions []PermissionModel

	err := global.TD27_DB.Raw(`
		SELECT p.*
		FROM sys_management_permission p
		JOIN sys_management_role_permissions rp ON p.id = rp.permission_id
		WHERE rp.role_id = ? AND p.domain = ?
	`, roleId, pd).Scan(&permissions).Error

	if err != nil {
		return nil, fmt.Errorf("get role permissions failed: %w", err)
	}

	return permissions, nil
}
