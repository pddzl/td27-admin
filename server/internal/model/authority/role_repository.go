package authority

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"server/internal/model/common"
)

type RoleEntity interface {
	FindOne(ctx context.Context, id uint) (*RoleModel, error)
	List(ctx context.Context, req *common.PageInfo) ([]RoleModel, int64, error)
	Create(ctx context.Context, req *RoleModel) (*RoleModel, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *UpdateRoleReq) error
	UpdateRoleMenu(ctx context.Context, menus []*MenuModel) error
}

type defaultRoleEntity struct {
	conn *gorm.DB
}

func NewDefaultRoleEntity(conn *gorm.DB) RoleEntity {
	return &defaultRoleEntity{conn: conn}
}

func (re *defaultRoleEntity) FindOne(ctx context.Context, id uint) (*RoleModel, error) {
	result := re.conn.WithContext(ctx).Find(&RoleModel{}, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return nil, nil
}

func (re *defaultRoleEntity) List(ctx context.Context, req *common.PageInfo) ([]RoleModel, int64, error) {
	var roles []RoleModel
	var total int64

	db := re.conn.WithContext(ctx).Model(&RoleModel{})

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count roles failed: %w", err)
	}

	// Pagination
	page := req.Page
	pageSize := req.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Query data with preload
	if err := db.
		Preload("Menus").
		Limit(pageSize).
		Offset(offset).
		//Order("id DESC").
		Find(&roles).Error; err != nil {
		return nil, 0, fmt.Errorf("list roles failed: %w", err)
	}

	return roles, total, nil
}

func (re *defaultRoleEntity) Create(ctx context.Context, req *RoleModel) (*RoleModel, error) {
	err := re.conn.WithContext(ctx).Create(req).Error
	//if err == nil {
	//	if err = casbinService.EditCasbin(instance.ID, baseReq.DefaultCasbin()); err != nil {
	//		global.TD27_LOG.Error("更新casbin rule失败", zap.Error(err))
	//	}
	//}
	return req, err
}

func (re *defaultRoleEntity) Delete(ctx context.Context, id uint) error {
	// check user
	//if !errors.Is(global.TD27_DB.Where("role_model_id = ?", id).First(&authority2.UserModel{}).Error, gorm.ErrRecordNotFound) {
	//	return errors.New("该角色下面还有所属用户")
	//}

	tx := re.conn.WithContext(ctx)

	result := tx.Unscoped().Delete(&RoleModel{}, id)

	if err := result.Error; err != nil {
		return fmt.Errorf("delete role failed, id=%d: %w", id, err)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// 清空menus关联
	//err = global.TD27_DB.Model(&roleModel).Association("Menus").Clear()
	//if err != nil {
	//	return fmt.Errorf("删除role关联menus err: %v", err)
	//}

	// 删除对应casbin rule
	//authorityId := strconv.Itoa(int(roleModel.ID))
	//ok := casbinService.ClearCasbin(0, authorityId)
	//if !ok {
	//	global.TD27_LOG.Warn("删除role关联casbin_rule失败")
	//}
	return nil
}

func (re *defaultRoleEntity) Update(ctx context.Context, req *UpdateRoleReq) error {
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
func (re *defaultRoleEntity) UpdateRoleMenu(ctx context.Context, menus []*MenuModel) error {
	var roleModel RoleModel
	//if errors.Is(re.conn.WithContext(ctx).Where("id = ?", roleId).First(&roleModel).Error, gorm.ErrRecordNotFound) {
	//	return errors.New("记录不存在")
	//}

	//var menuModel []menu.MenuModel
	//err = global.TD27_DB.Where("id in ?", ids).Find(&menuModel).Error
	//if err != nil {
	//	global.TD27_LOG.Error("EditRoleMenu 查询menu", zap.Error(err))
	//	return err
	//}

	err := re.conn.WithContext(ctx).Model(&roleModel).Association("Menus").Replace(menus)
	if err != nil {
		//global.TD27_LOG.Error("EditRoleMenu 替换menu", zap.Error(err))
		return fmt.Errorf("edit menu failed: %w", err)
	}

	return nil
}
