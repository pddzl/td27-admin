package authority

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"server/internal/model/common"
	"server/internal/pkg"
)

type UserEntity interface {
	FindOne(context.Context, *FindOneUserReq) (*UserModel, error)
	List(context.Context, *common.PageInfo) ([]*UserResp, int64, error)
	Delete(context.Context, uint) error
	Create(context.Context, *AddUserReq) error
	Update(context.Context, *UpdateUserReq) (*UserModel, error)
	GetUserInfo(context.Context, uint) (userResults *UserResp, err error)
	ModifyPasswd(context.Context, *ModifyPasswdReq) error
	SwitchActive(context.Context, *SwitchActiveReq) error
}

type defaultUserEntity struct {
	conn *gorm.DB
}

func NewDefaultUserEntity(conn *gorm.DB) UserEntity {
	return &defaultUserEntity{conn: conn}
}

func (e *defaultUserEntity) FindOne(ctx context.Context, req *FindOneUserReq) (*UserModel, error) {
	db := e.conn.WithContext(ctx)
	query := db.Model(&UserModel{})
	// OR conditions
	if req.ID != 0 && req.RoleModelID != 0 {
		query = query.Where("id = ? OR role_model_id = ?", req.ID, req.RoleModelID)
	} else if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	} else {
		query = query.Where("role_model_id = ?", req.RoleModelID)
	}

	var userModel UserModel
	if err := query.First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &userModel, nil
}

func (e *defaultUserEntity) GetUserInfo(ctx context.Context, userId uint) (listUserResp *UserResp, err error) {
	var resp UserResp

	tx := e.conn.
		WithContext(ctx).
		Table("authority_user").
		Select(`
			authority_user.id,
			authority_user.created_at,
			authority_user.username,
			authority_user.phone,
			authority_user.email,
			authority_user.active,
			authority_user.role_model_id,
			authority_role.role_name
		`).
		Joins("JOIN authority_role ON authority_user.role_model_id = authority_role.id").
		Where("authority_user.id = ?", userId).
		Scan(&resp)

	if err = tx.Error; err != nil {
		return nil, fmt.Errorf("get user info failed (id=%d): %w", userId, err)
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &resp, nil
}

func (e *defaultUserEntity) List(ctx context.Context, pageInfo *common.PageInfo) ([]*UserResp, int64, error) {
	var users []*UserResp
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
	db := e.conn.WithContext(ctx).Model(&UserModel{})

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

func (e *defaultUserEntity) Delete(ctx context.Context, id uint) (err error) {
	result := e.conn.WithContext(ctx).Where("id = ?", id).Unscoped().Delete(&UserModel{})

	if err = result.Error; err != nil {
		return fmt.Errorf("delete user failed, id=%d: %w", id, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (e *defaultUserEntity) Create(ctx context.Context, req *AddUserReq) error {
	var userModel UserModel
	userModel.Username = req.Username
	userModel.Password = pkg.MD5V([]byte(req.Password))
	userModel.Phone = req.Phone
	userModel.Email = req.Email
	userModel.Active = req.Active
	userModel.RoleModelID = req.RoleModelID

	return e.conn.WithContext(ctx).Create(&userModel).Error
}

func (e *defaultUserEntity) Update(ctx context.Context, req *UpdateUserReq) (*UserModel, error) {
	db := e.conn.WithContext(ctx)

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

	// Fetch updated record
	var user UserModel
	if err := db.Where("id = ?", req.ID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (e *defaultUserEntity) ModifyPasswd(ctx context.Context, req *ModifyPasswdReq) (err error) {
	db := e.conn.WithContext(ctx)

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
func (e *defaultUserEntity) SwitchActive(ctx context.Context, req *SwitchActiveReq) (err error) {
	tx := e.conn.WithContext(ctx).Model(&UserModel{}).Where("id = ?", req.ID).Update("active", req.Active)

	if err = tx.Error; err != nil {
		return fmt.Errorf("switch user active failed (id=%d): %w", req.ID, err)
	}

	if tx.RowsAffected == 0 {
		return errors.New("记录不存在")
	}

	return nil
}
