package sysManagement

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// PermissionRepository Permission仓库接口
type PermissionRepository interface {
	ListByRoleID(ctx context.Context, roleId uint, pd PermissionDomain) ([]PermissionModel, error)
	ListByTokenID(ctx context.Context, tokenID uint, pd PermissionDomain) ([]PermissionModel, error)
	Create(ctx context.Context, permission *PermissionModel) error
	FindByDomainID(ctx context.Context, domainID uint, domain PermissionDomain) (*PermissionModel, error)
	DeleteByDomainID(ctx context.Context, domainID uint, domain PermissionDomain) error
	Update(ctx context.Context, permissionModel *PermissionModel) error
}

type permissionRepo struct {
	conn *gorm.DB
}

func NewPermissionRepo(conn *gorm.DB) PermissionRepository {
	return &permissionRepo{conn: conn}
}

func (r *permissionRepo) ListByRoleID(ctx context.Context, roleId uint, pd PermissionDomain) ([]PermissionModel, error) {
	var permissions []PermissionModel

	err := r.conn.WithContext(ctx).Raw(`
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

func (r *permissionRepo) ListByTokenID(ctx context.Context, tokenID uint, pd PermissionDomain) ([]PermissionModel, error) {
	var permissions []PermissionModel

	err := r.conn.WithContext(ctx).Raw(`
		SELECT p.*
		FROM sys_management_permission p
		JOIN sys_tool_token_permission tp ON p.id = tp.permission_id
		WHERE tp.token_id = ? AND p.domain = ?
	`, tokenID, pd).Scan(&permissions).Error

	if err != nil {
		return nil, fmt.Errorf("get token permissions failed: %w", err)
	}

	return permissions, nil
}

// Create 创建权限
func (r *permissionRepo) Create(ctx context.Context, permission *PermissionModel) error {
	if err := r.conn.WithContext(ctx).Create(permission).Error; err != nil {
		return fmt.Errorf("create permission failed: %w", err)
	}
	return nil
}

// FindByDomainID 根据domain_id和domain查找权限
func (r *permissionRepo) FindByDomainID(ctx context.Context, domainID uint, domain PermissionDomain) (*PermissionModel, error) {
	var perm PermissionModel
	if err := r.conn.WithContext(ctx).Where("domain_id = ? AND domain = ?", domainID, domain).First(&perm).Error; err != nil {
		return nil, fmt.Errorf("find permission failed: %w", err)
	}
	return &perm, nil
}

// DeleteByDomainID 根据domain_id和domain删除权限
func (r *permissionRepo) DeleteByDomainID(ctx context.Context, domainID uint, domain PermissionDomain) error {
	result := r.conn.WithContext(ctx).Unscoped().
		Where("domain_id = ? AND domain = ?", domainID, domain).
		Delete(&PermissionModel{})

	if result.Error != nil {
		return fmt.Errorf("delete permission failed: %w", result.Error)
	}
	return nil
}

func (r *permissionRepo) Update(ctx context.Context, permissionModel *PermissionModel) error {
	if err := r.conn.WithContext(ctx).Save(permissionModel).Error; err != nil {
		return fmt.Errorf("update permission failed: %w", err)
	}
	return nil
}
