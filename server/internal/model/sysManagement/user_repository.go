package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"server/internal/model/common"
	"server/internal/pkg"
)

type UserRepository interface {
	FindOne(context.Context, *FindOneUserReq) (*UserModel, error)
	FindOneWithRoles(context.Context, uint) (*UserModel, error)
	List(context.Context, *common.PageInfo, *DataPermission) ([]*UserResp, int64, error)
	Delete(context.Context, string) error
	Create(context.Context, *AddUserReq) (*UserModel, error)
	Update(context.Context, *UpdateUserReq) (*UserModel, error)
	GetUserInfo(context.Context, uint) (userResults *UserResp, err error)
	ModifyPasswd(context.Context, *ModifyPasswdReq) error
	SwitchActive(context.Context, *SwitchActiveReq) error
	CountUsersByRole(context.Context, uint, *int64) error
}

type userEntity struct {
	conn *gorm.DB
}

func NewUserEntity(conn *gorm.DB) UserRepository {
	return &userEntity{conn: conn}
}

func (e *userEntity) FindOne(ctx context.Context, req *FindOneUserReq) (*UserModel, error) {
	db := e.conn.WithContext(ctx)
	query := db.Model(&UserModel{})

	var userModel UserModel
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}

	if err := query.First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &userModel, nil
}

func (e *userEntity) FindOneWithRoles(ctx context.Context, userId uint) (*UserModel, error) {
	var user UserModel
	err := e.conn.WithContext(ctx).
		Preload("Roles").
		Where("id = ?", userId).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("find user with roles failed (id=%d): %w", userId, err)
	}

	return &user, nil
}

func (e *userEntity) GetUserInfo(ctx context.Context, userId uint) (listUserResp *UserResp, err error) {
	var resp UserResp

	// 查询用户及其角色
	var user UserModel
	err = e.conn.WithContext(ctx).
		Preload("Roles").
		Where("id = ?", userId).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("get user info failed (id=%d): %w", userId, err)
	}

	// 填充响应
	resp.UserModel = user
	if len(user.Roles) > 0 {
		resp.RoleName = user.Roles[0].RoleName
		resp.RoleID = user.Roles[0].ID
	}

	// 获取部门名称
	if user.DeptID > 0 {
		var dept DeptModel
		if err := e.conn.WithContext(ctx).First(&dept, user.DeptID).Error; err == nil {
			resp.DeptName = dept.DeptName
		}
	}

	return &resp, nil
}

func (e *userEntity) List(ctx context.Context, pageInfo *common.PageInfo, dataPerm *DataPermission) ([]*UserResp, int64, error) {
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

	// Base query
	db := e.conn.WithContext(ctx).Model(&UserModel{})

	// Apply data permission filter
	if dataPerm != nil {
		db = ApplyDataScope(db, dataPerm, "", "dept_id", "id")
	}

	// Count total (with data permission filter)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count users failed: %w", err)
	}

	// Query list with roles and dept
	var userModels []UserModel
	err := db.
		Preload("Roles").
		Limit(pageSize).
		Offset(offset).
		Find(&userModels).Error

	if err != nil {
		return nil, total, fmt.Errorf("list users failed: %w", err)
	}

	// Get dept names for users with dept_id
	var deptIDs []uint
	for _, user := range userModels {
		if user.DeptID > 0 {
			deptIDs = append(deptIDs, user.DeptID)
		}
	}

	deptMap := make(map[uint]string)
	if len(deptIDs) > 0 {
		var depts []DeptModel
		if err := e.conn.WithContext(ctx).Where("id IN ?", deptIDs).Find(&depts).Error; err == nil {
			for _, dept := range depts {
				deptMap[dept.ID] = dept.DeptName
			}
		}
	}

	// Convert to response
	for _, user := range userModels {
		userResp := UserResp{
			UserModel: user,
			DeptName:  deptMap[user.DeptID],
		}
		if len(user.Roles) > 0 {
			userResp.RoleName = user.Roles[0].RoleName
			userResp.RoleID = user.Roles[0].ID
		}
		users = append(users, &userResp)
	}

	return users, total, nil
}

func (e *userEntity) Delete(ctx context.Context, username string) (err error) {
	result := e.conn.WithContext(ctx).Where("username = ?", username).Unscoped().Delete(&UserModel{})

	if err = result.Error; err != nil {
		return fmt.Errorf("delete user failed, username=%s: %w", username, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (e *userEntity) Create(ctx context.Context, req *AddUserReq) (*UserModel, error) {
	var userModel UserModel
	userModel.Username = req.Username
	userModel.Password = pkg.MD5V([]byte(req.Password))
	userModel.Phone = req.Phone
	userModel.Email = req.Email
	userModel.Active = req.Active
	userModel.DeptID = req.DeptID

	// 开始事务
	tx := e.conn.WithContext(ctx).Begin()

	// 创建用户
	if err := tx.Create(&userModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 关联角色
	if len(req.RoleIDs) > 0 {
		var roles []RoleModel
		if err := tx.Where("id IN ?", req.RoleIDs).Find(&roles).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(&userModel).Association("Roles").Append(roles); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &userModel, nil
}

func (e *userEntity) Update(ctx context.Context, req *UpdateUserReq) (*UserModel, error) {
	db := e.conn.WithContext(ctx)

	// 开始事务
	tx := db.Begin()

	// Rebuild user
	updates := map[string]interface{}{
		"username":   req.Username,
		"phone":      req.Phone,
		"email":      req.Email,
		"active":     req.Active,
		"dept_id":    req.DeptID,
		"updated_at": time.Now(),
	}

	if err := tx.Model(&UserModel{}).Where("id = ?", req.ID).Updates(updates).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("update user failed: %w", err)
	}

	// 更新角色关联
	if len(req.RoleIDs) > 0 {
		var user UserModel
		if err := tx.Where("id = ?", req.ID).First(&user).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// 清除旧角色
		if err := tx.Model(&user).Association("Roles").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		// 添加新角色
		var roles []RoleModel
		if err := tx.Where("id IN ?", req.RoleIDs).Find(&roles).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(&user).Association("Roles").Append(roles); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	// Fetch updated record with roles
	var user UserModel
	if err := db.Preload("Roles").Where("id = ?", req.ID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (e *userEntity) ModifyPasswd(ctx context.Context, req *ModifyPasswdReq) (err error) {
	db := e.conn.WithContext(ctx)

	// Verify old password
	var user UserModel
	if err = db.Where("id = ? AND password = ?", req.ID, pkg.MD5V([]byte(req.OldPassword))).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("旧密码错误")
		}
		return err
	}

	// Rebuild password
	if err = db.Model(&UserModel{}).Where("id = ?", req.ID).Update("password", pkg.MD5V([]byte(req.NewPassword))).Error; err != nil {
		return fmt.Errorf("update password failed: %w", err)
	}

	return nil
}

// SwitchActive 切换启用状态
func (e *userEntity) SwitchActive(ctx context.Context, req *SwitchActiveReq) (err error) {
	tx := e.conn.WithContext(ctx).Model(&UserModel{}).Where("username = ?", req.Username).Update("active", req.Active)

	if err = tx.Error; err != nil {
		return fmt.Errorf("switch user active failed (username=%s): %w", req.Username, err)
	}

	if tx.RowsAffected == 0 {
		return errors.New("记录不存在")
	}

	return nil
}

// CountUsersByRole 统计具有指定角色的用户数量
func (e *userEntity) CountUsersByRole(ctx context.Context, roleID uint, count *int64) error {
	return e.conn.WithContext(ctx).
		Table("sys_management_user_roles").
		Where("role_id = ?", roleID).
		Count(count).Error
}
