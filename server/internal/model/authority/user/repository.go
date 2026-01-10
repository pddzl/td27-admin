package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"server/internal/global"
	"server/internal/model/authority/role"
	"server/internal/model/common"
	"server/internal/pkg"
)

type UserEntity interface {
	List(ctx context.Context, req *common.PageInfo) ([]ListUserResp, int64, error)
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, req *AddUserReq) error
	Update(ctx context.Context, req *EditUserReq) (*ListUserResp, error)
	ModifyPasswd(ctx context.Context, req *ModifyPasswdReq) error
	SwitchActive(ctx context.Context, req *SwitchActiveReq) error
}

type defaultUserEntity struct {
	conn *gorm.DB
}

func NewDefaultUserEntity(conn *gorm.DB) UserEntity {
	return &defaultUserEntity{conn: conn}
}

func (ue *defaultUserEntity) GetUserInfo(userId uint) (userResults *ListUserResp, err error) {
	err = global.TD27_DB.Table("authority_user").Select("authority_user.created_at,authority_user.id,authority_user.username,authority_user.phone,authority_user.email,authority_user.active,authority_user.role_model_id,authority_role.role_name").Joins("inner join authority_role on authority_user.role_model_id = authority_role.id").Where("authority_user.id = ?", userId).Scan(&userResults).Error
	return
}

func (ue *defaultUserEntity) List(ctx context.Context, pageInfo *common.PageInfo) ([]ListUserResp, int64, error) {
	var users []ListUserResp
	var total int64

	// Safety defaults
	page := pageInfo.Page
	pageSize := pageInfo.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Base query (NO joins for count)
	db := ue.conn.WithContext(ctx).Model(&UserModel{})

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count users failed: %w", err)
	}

	// Query list with join
	err := db.
		Select(`
			authority_user.id,
			authority_user.created_at,
			authority_user.updated_at,
			authority_user.username,
			authority_user.phone,
			authority_user.email,
			authority_user.active,
			authority_user.role_model_id,
			authority_role.role_name AS role_name
		`).
		Joins(`
			LEFT JOIN authority_role
			ON authority_user.role_model_id = authority_role.id
		`).
		Limit(pageSize).
		Offset(offset).
		Scan(&users).Error

	if err != nil {
		return nil, total, fmt.Errorf("list users failed: %w", err)
	}

	return users, total, nil
}

func (ue *defaultUserEntity) Delete(ctx context.Context, id uint) (err error) {
	result := ue.conn.WithContext(ctx).Where("id = ?", id).Unscoped().Delete(&UserModel{})

	if err = result.Error; err != nil {
		return fmt.Errorf("delete user failed, id=%d: %w", id, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (ue *defaultUserEntity) Create(ctx context.Context, req *AddUserReq) (err error) {
	if errors.Is(ue.conn.WithContext(ctx).
		Where("id = ?", req.RoleModelID).
		First(&role.RoleModel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在")
	}

	var userModel UserModel
	userModel.Username = req.Username
	userModel.Password = pkg.MD5V([]byte(req.Password))
	userModel.Phone = req.Phone
	userModel.Email = req.Email
	userModel.Active = req.Active
	userModel.RoleModelID = req.RoleModelID

	return global.TD27_DB.Create(&userModel).Error
}

func (ue *defaultUserEntity) Update(ctx context.Context, req *EditUserReq) (*ListUserResp, error) {
	db := ue.conn.WithContext(ctx)

	// Check role existence
	var roleName string
	if err := db.
		Model(&role.RoleModel{}).
		Select("role_name").
		Where("id = ?", req.RoleModelID).
		Take(&roleName).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("角色不存在")
		}
		return nil, fmt.Errorf("check role failed: %w", err)
	}

	// Update user
	updates := map[string]interface{}{
		"username":      req.Username,
		"phone":         req.Phone,
		"email":         req.Email,
		"active":        req.Active,
		"role_model_id": req.RoleModelID,
		"updated_at":    time.Now(),
	}

	tx := db.Model(&UserModel{}).Where("id = ?", req.ID).Updates(updates)

	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("update user failed: %w", err)
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("记录不存在")
	}

	// Build response directly (no extra query)
	resp := &ListUserResp{
		UserModel: UserModel{
			Td27Model: common.Td27Model{
				ID: req.ID,
			},
			Username:    req.Username,
			Phone:       req.Phone,
			Email:       req.Email,
			Active:      req.Active,
			RoleModelID: req.RoleModelID,
		},
		RoleName: roleName,
	}

	return resp, nil
}

func (ue *defaultUserEntity) ModifyPasswd(ctx context.Context, req *ModifyPasswdReq) (err error) {
	db := ue.conn.WithContext(ctx)

	// Verify old password
	tx := db.Model(&UserModel{}).
		Where("id = ? AND password = ?", req.ID, pkg.MD5V([]byte(req.OldPassword)))

	if tx.RowsAffected == 0 {
		return errors.New("旧密码错误")
	}

	// 2. Update password
	if err = tx.Update("password", pkg.MD5V([]byte(req.NewPassword))).Error; err != nil {
		return fmt.Errorf("update password failed: %w", err)
	}

	return nil
}

// SwitchActive 切换启用状态
func (ue *defaultUserEntity) SwitchActive(ctx context.Context, req *SwitchActiveReq) (err error) {
	tx := ue.conn.WithContext(ctx).Model(&UserModel{}).Where("id = ?", req.ID).Update("active", req.Active)

	if err = tx.Error; err != nil {
		return fmt.Errorf("switch user active failed (id=%d): %w", req.ID, err)
	}

	if tx.RowsAffected == 0 {
		return errors.New("记录不存在")
	}

	return nil
}
