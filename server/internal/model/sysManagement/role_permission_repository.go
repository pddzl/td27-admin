package sysManagement

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	Update(ctx context.Context, roleId uint, permissionIds []uint, domain string) error
	Delete(ctx context.Context, roleId uint) error
}

type rolePermissionRepo struct {
	conn *gorm.DB
}

func NewRolePermissionRepository(conn *gorm.DB) RolePermissionRepository {
	return &rolePermissionRepo{conn: conn}
}

// Update 编辑角色权限（使用统一权限表）
func (r *rolePermissionRepo) Update(ctx context.Context, roleId uint, permissionIds []uint, domain string) error {
	// If no permissionIds, just return
	if len(permissionIds) == 0 {
		return nil
	}

	// delete existing permissions for this role
	if err := r.Delete(ctx, roleId); err != nil {
		return err
	}

	// Get permission IDs for the menu IDs (menu ID = permission ID)
	var permissions []PermissionModel
	if err := r.conn.WithContext(ctx).
		Where("id IN ? AND domain = ?", permissionIds, domain).
		Find(&permissions).Error; err != nil {
		return fmt.Errorf("find menu permissions failed: %w", err)
	}

	// Create role-permission associations
	for _, perm := range permissions {
		rp := RolePermissionModel{
			RoleID:       roleId,
			PermissionID: perm.ID,
			DataScope:    "all", // default data scope for menus
		}
		if err := r.conn.WithContext(ctx).Create(&rp).Error; err != nil {
			return fmt.Errorf("create role permission failed: %w", err)
		}
	}

	return nil
}

func (r *rolePermissionRepo) Delete(ctx context.Context, roleId uint) error {
	// Delete from role_permissions where permission is of type 'menu'
	return r.conn.WithContext(ctx).
		Exec(`
			DELETE FROM sys_management_role_permissions 
			WHERE role_id = ? AND permission_id IN (
				SELECT id FROM sys_management_permission WHERE type = 'menu'
			)
		`, roleId).Error
}
