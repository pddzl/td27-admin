package sysManagement

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"server/internal/model/common"
)

type RoleRepository interface {
	FindOne(ctx context.Context, id uint) (*RoleModel, error)
	List(ctx context.Context, req *common.PageInfo) ([]RoleModel, int64, error)
	Create(ctx context.Context, req *RoleModel) (*RoleModel, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *UpdateRoleReq) error
	UpdateRoleMenu(ctx context.Context, roleId uint, menuIds []uint) error
	DeleteRoleMenu(ctx context.Context, roleId uint) error
}

type roleEntity struct {
	conn *gorm.DB
}

func NewRoleEntity(conn *gorm.DB) RoleRepository {
	return &roleEntity{conn: conn}
}

func (e *roleEntity) FindOne(ctx context.Context, id uint) (*RoleModel, error) {
	var roleModel RoleModel
	result := e.conn.WithContext(ctx).Find(&roleModel, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &roleModel, nil
}

func (e *roleEntity) List(ctx context.Context, req *common.PageInfo) ([]RoleModel, int64, error) {
	req.Normalize()

	var roles []RoleModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&RoleModel{})

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count roles failed: %w", err)
	}

	// Query data
	if err := db.
		Limit(req.PageSize).
		Offset(req.Offset()).
		Find(&roles).Error; err != nil {
		return nil, 0, fmt.Errorf("list roles failed: %w", err)
	}

	return roles, total, nil
}

func (e *roleEntity) Create(ctx context.Context, req *RoleModel) (*RoleModel, error) {
	err := e.conn.WithContext(ctx).Create(req).Error

	return req, err
}

func (e *roleEntity) Delete(ctx context.Context, id uint) error {
	tx := e.conn.WithContext(ctx)

	result := tx.Unscoped().Delete(&RoleModel{}, id)

	if err := result.Error; err != nil {
		return fmt.Errorf("delete role failed, id=%d: %w", id, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (e *roleEntity) Update(ctx context.Context, req *UpdateRoleReq) error {
	result := e.conn.WithContext(ctx).
		Model(&RoleModel{}).
		Where("id = ?", req.ID).
		Update("role_name", req.RoleName)

	if err := result.Error; err != nil {
		return fmt.Errorf("update role failed, id=%d: %w", req.ID, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateRoleMenu 编辑角色的菜单权限（使用统一权限表）
func (e *roleEntity) UpdateRoleMenu(ctx context.Context, roleId uint, menuIds []uint) error {
	// First, delete existing menu permissions for this role
	//if err := e.DeleteRoleMenu(ctx, roleId); err != nil {
	//	return err
	//}

	// If no menu IDs, just return
	if len(menuIds) == 0 {
		return nil
	}

	// Get permission IDs for the menu IDs (menu ID = permission ID)
	var permissions []PermissionModel
	if err := e.conn.WithContext(ctx).
		Where("id IN ? AND type = 'menu'", menuIds).
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
		if err := e.conn.WithContext(ctx).Create(&rp).Error; err != nil {
			return fmt.Errorf("create role permission failed: %w", err)
		}
	}

	return nil
}

func (e *roleEntity) DeleteRoleMenu(ctx context.Context, roleId uint) error {
	// Delete from role_permissions where permission is of type 'menu'
	return e.conn.WithContext(ctx).
		Exec(`
			DELETE FROM sys_management_role_permissions 
			WHERE role_id = ? AND permission_id IN (
				SELECT id FROM sys_management_permission WHERE type = 'menu'
			)
		`, roleId).Error
}
