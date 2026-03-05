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
	UpdateRoleMenu(ctx context.Context, req []*MenuModel) error
	DeleteRoleMenu(ctx context.Context, id uint) error
}

type roleEntity struct {
	conn *gorm.DB
}

func NewRoleEntity(conn *gorm.DB) RoleRepository {
	return &roleEntity{conn: conn}
}

func (re *roleEntity) FindOne(ctx context.Context, id uint) (*RoleModel, error) {
	var roleModel RoleModel
	result := re.conn.WithContext(ctx).Find(&roleModel, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &roleModel, nil
}

func (re *roleEntity) List(ctx context.Context, req *common.PageInfo) ([]RoleModel, int64, error) {
	req.Normalize()

	var roles []RoleModel
	var total int64

	db := re.conn.WithContext(ctx).Model(&RoleModel{})

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count roles failed: %w", err)
	}

	// Query data with preload
	if err := db.
		Preload("Menus").
		Limit(req.PageSize).
		Offset(req.Offset()).
		//Order("id DESC").
		Find(&roles).Error; err != nil {
		return nil, 0, fmt.Errorf("list roles failed: %w", err)
	}

	return roles, total, nil
}

func (re *roleEntity) Create(ctx context.Context, req *RoleModel) (*RoleModel, error) {
	err := re.conn.WithContext(ctx).Create(req).Error

	return req, err
}

func (re *roleEntity) Delete(ctx context.Context, id uint) error {
	tx := re.conn.WithContext(ctx)

	result := tx.Unscoped().Delete(&RoleModel{}, id)

	if err := result.Error; err != nil {
		return fmt.Errorf("delete role failed, id=%d: %w", id, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (re *roleEntity) Update(ctx context.Context, req *UpdateRoleReq) error {
	result := re.conn.WithContext(ctx).
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

// UpdateRoleMenu 编辑用户menu
func (re *roleEntity) UpdateRoleMenu(ctx context.Context, req []*MenuModel) error {
	var roleModel RoleModel

	err := re.conn.WithContext(ctx).Model(&roleModel).Association("Menus").Replace(req)
	if err != nil {
		return fmt.Errorf("update menu failed: %w", err)
	}

	return nil
}

func (re *roleEntity) DeleteRoleMenu(ctx context.Context, roleId uint) error {
	return re.conn.WithContext(ctx).Where("role_id =?", roleId).Delete(&RoleMenu{}).Error
}
