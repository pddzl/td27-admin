package sysManagement

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	Update(ctx context.Context, roleId uint, permissionIds []uint, domain string) error
}

type rolePermissionRepo struct {
	conn *gorm.DB
}

func NewRolePermissionRepository(conn *gorm.DB) RolePermissionRepository {
	return &rolePermissionRepo{conn: conn}
}

// Update 编辑角色权限（使用统一权限表）
func (r *rolePermissionRepo) Update(ctx context.Context, roleId uint, permissionIds []uint, domain string) error {
	//if len(permissionIds) == 0 {
	//	return nil
	//}

	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// delete exists
		if err := tx.
			Where("role_id = ?", roleId).
			Where("permission_id IN (?)",
				tx.Model(&PermissionModel{}).
					Select("id").
					Where("domain = ?", domain),
			).
			Delete(&RolePermissionModel{}).Error; err != nil {
			return fmt.Errorf("delete role permissions failed: %w", err)
		}

		// query permissions
		var permissions []PermissionModel
		if err := tx.
			Where("id IN ? AND domain = ?", permissionIds, domain).
			Find(&permissions).Error; err != nil {
			return fmt.Errorf("find permissions failed: %w", err)
		}

		// batch insert
		rps := make([]RolePermissionModel, 0, len(permissions))
		for _, perm := range permissions {
			rps = append(rps, RolePermissionModel{
				RoleID:       roleId,
				PermissionID: perm.ID,
				DataScope:    "all",
			})
		}

		if len(rps) > 0 {
			if err := tx.Create(&rps).Error; err != nil {
				return fmt.Errorf("create role permissions failed: %w", err)
			}
		}

		return nil
	})
}
