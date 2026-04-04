package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type MenuRepository interface {
	List(ctx context.Context, roleId uint) ([]*MenuModel, error)
	ListByRoleIDs(ctx context.Context, roleIDs []uint) ([]*MenuModel, error)
	Create(ctx context.Context, req *Menu) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *UpdateMenuReq) error
	GetElTreeMenus(ctx context.Context, roleId uint) ([]*MenuModel, []uint, error)
	FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error)
}

type menuEntity struct {
	conn *gorm.DB
}

func NewMenuEntity(conn *gorm.DB) MenuRepository {
	return &menuEntity{conn: conn}
}

// permissionToMenuModel converts PermissionModel to MenuModel
func permissionToMenuModel(p *PermissionModel) *MenuModel {
	var pid uint
	if p.ParentID != nil {
		pid = *p.ParentID
	}
	return &MenuModel{
		Td27Model: p.Td27Model,
		Name:      p.Name,
		Path:      p.Resource, // resource stores the path for menus
		Component: p.Component,
		Redirect:  p.Redirect,
		Pid:       pid,
		Sort:      p.Sort,
		Meta: Meta{
			Title:      p.Name,
			SvgIcon:    p.Icon,
			Hidden:     p.Hidden,
			KeepAlive:  p.KeepAlive,
			AlwaysShow: false,
			Affix:      false,
		},
	}
}

// buildMenuTreeOptimized 使用 map 优化构建菜单树，时间复杂度 O(n)
func buildMenuTreeOptimized(menuList []*MenuModel) []*MenuModel {
	if len(menuList) == 0 {
		return []*MenuModel{}
	}

	// 使用 map 存储所有菜单，实现 O(1) 查找
	menuMap := make(map[uint]*MenuModel, len(menuList))
	for _, menu := range menuList {
		menuMap[menu.ID] = menu
		// 清空子菜单，避免重复
		menu.Children = nil
	}

	// 构建树结构
	var rootMenus []*MenuModel
	for _, menu := range menuList {
		if menu.Pid == 0 {
			// 根菜单
			rootMenus = append(rootMenus, menu)
		} else {
			// 子菜单，找到父菜单并添加
			if parent, ok := menuMap[menu.Pid]; ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	// 对根菜单排序
	sort.Slice(rootMenus, func(i, j int) bool {
		return rootMenus[i].Sort < rootMenus[j].Sort
	})

	// 递归排序子菜单
	sortMenuChildren(rootMenus)

	return rootMenus
}

// sortMenuChildren 递归排序子菜单
func sortMenuChildren(menus []*MenuModel) {
	if len(menus) == 0 {
		return
	}
	for _, menu := range menus {
		if len(menu.Children) > 0 {
			sort.Slice(menu.Children, func(i, j int) bool {
				return menu.Children[i].Sort < menu.Children[j].Sort
			})
			sortMenuChildren(menu.Children)
		}
	}
}

func (e *menuEntity) List(ctx context.Context, roleId uint) ([]*MenuModel, error) {
	// Query menu permissions from unified permission table
	var permissions []*PermissionModel
	err := e.conn.WithContext(ctx).
		Model(&PermissionModel{}).
		Joins("JOIN sys_management_role_permissions rp ON rp.permission_id = sys_management_permission.id").
		Where("rp.role_id = ? AND sys_management_permission.type = 'menu'", roleId).
		Where("sys_management_permission.status = ?", true).
		Find(&permissions).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role err: %w", err)
	}

	// Convert to MenuModel
	var menuModels []*MenuModel
	for _, p := range permissions {
		menuModels = append(menuModels, permissionToMenuModel(p))
	}

	// 使用优化的 O(n) 算法构建菜单树
	return buildMenuTreeOptimized(menuModels), nil
}

func (e *menuEntity) ListByRoleIDs(ctx context.Context, roleIDs []uint) ([]*MenuModel, error) {
	if len(roleIDs) == 0 {
		return []*MenuModel{}, nil
	}

	// Query menu permissions from unified permission table
	var permissions []*PermissionModel
	err := e.conn.WithContext(ctx).
		Model(&PermissionModel{}).
		Joins("JOIN sys_management_role_permissions rp ON rp.permission_id = sys_management_permission.id").
		Where("rp.role_id IN ? AND sys_management_permission.type = 'menu'", roleIDs).
		Where("sys_management_permission.status = ?", true).
		Group("sys_management_permission.id").
		Find(&permissions).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role ids err: %w", err)
	}

	// Convert to MenuModel
	var menuModels []*MenuModel
	for _, p := range permissions {
		menuModels = append(menuModels, permissionToMenuModel(p))
	}

	// 使用优化的 O(n) 算法构建菜单树
	return buildMenuTreeOptimized(menuModels), nil
}

func (e *menuEntity) Create(ctx context.Context, req *Menu) error {
	// Create as permission with type='menu'
	permission := PermissionModel{
		Name:     req.Name,
		Type:     "menu",
		Resource: req.Path, // path is stored in resource
		Action:   "view",
		ParentID: func() *uint { p := uint(req.Pid); return &p }(),
		Sort:     req.Sort,
		Status:   true,
		Icon:     req.Meta.Icon,
		Component: req.Component,
		Redirect:  req.Redirect,
		Hidden:    req.Meta.Hidden,
		KeepAlive: req.Meta.KeepAlive,
	}

	if err := e.conn.WithContext(ctx).Create(&permission).Error; err != nil {
		return fmt.Errorf("create menu failed: %v", err)
	}

	return nil
}

func (e *menuEntity) Update(ctx context.Context, req *UpdateMenuReq) error {
	// Find the permission by ID (menu ID is permission ID)
	var permission PermissionModel
	if err := e.conn.WithContext(ctx).First(&permission, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("菜单不存在")
		}
		return fmt.Errorf("find menu failed: %w", err)
	}

	// Check if it's a menu type permission
	if permission.Type != "menu" {
		return errors.New("该ID不是菜单类型权限")
	}

	updates := PermissionModel{
		Name:     req.Name,
		Resource: req.Path,
		Component: req.Component,
		Redirect:  req.Redirect,
		ParentID:  func() *uint { p := uint(req.Pid); return &p }(),
		Sort:      req.Sort,
		Icon:      req.Meta.Icon,
		Hidden:    req.Meta.Hidden,
		KeepAlive: req.Meta.KeepAlive,
	}

	result := e.conn.WithContext(ctx).
		Model(&PermissionModel{}).
		Where("id = ? AND type = 'menu'", req.ID).
		Updates(&updates)

	if err := result.Error; err != nil {
		return fmt.Errorf("update menu failed, id=%d: %w", req.ID, err)
	}

	if result.RowsAffected == 0 {
		return errors.New("菜单不存在")
	}

	return nil
}

func (e *menuEntity) Delete(ctx context.Context, id uint) error {
	// Delete the permission with type='menu'
	result := e.conn.WithContext(ctx).
		Where("id = ? AND type = 'menu'", id).
		Delete(&PermissionModel{})

	if err := result.Error; err != nil {
		return fmt.Errorf("delete menu failed: %w", err)
	}

	if result.RowsAffected == 0 {
		return errors.New("菜单不存在")
	}

	return nil
}

// GetElTreeMenus 获取所有menu
func (e *menuEntity) GetElTreeMenus(ctx context.Context, roleId uint) ([]*MenuModel, []uint, error) {
	db := e.conn.WithContext(ctx)

	// Query all menu permissions
	var permissions []*PermissionModel
	if err := db.Where("type = 'menu' AND status = ?", true).Find(&permissions).Error; err != nil {
		return nil, nil, fmt.Errorf("query menus failed: %w", err)
	}

	// Convert to MenuModel
	var allMenus []*MenuModel
	for _, p := range permissions {
		allMenus = append(allMenus, permissionToMenuModel(p))
	}

	// Build menu tree using optimized O(n) algorithm
	rootMenus := buildMenuTreeOptimized(allMenus)

	// Query role-permission relations for menus
	var relations []RolePermission
	if err := db.
		Model(&RolePermission{}).
		Select("permission_id").
		Where("role_id = ?", roleId).
		Scan(&relations).Error; err != nil {
		return nil, nil, fmt.Errorf("query role menus failed: %w", err)
	}

	// Build permission ID set
	permSet := make(map[uint]struct{}, len(relations))
	for _, r := range relations {
		permSet[r.PermissionID] = struct{}{}
	}

	// el-tree leaf selection
	checkedIDs := make([]uint, 0)
	for _, m := range allMenus {
		if _, ok := permSet[m.ID]; !ok {
			continue
		}

		// if has child selected → skip parent
		isParent := false
		for _, c := range allMenus {
			if c.Pid == m.ID {
				if _, ok := permSet[c.ID]; ok {
					isParent = true
					break
				}
			}
		}

		if !isParent {
			checkedIDs = append(checkedIDs, m.ID)
		}
	}

	return rootMenus, checkedIDs, nil
}

func (e *menuEntity) FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error) {
	// Guard: avoid full table scan
	if len(ids) == 0 {
		return []*MenuModel{}, nil
	}

	// Query menu permissions by IDs
	var permissions []*PermissionModel
	err := e.conn.WithContext(ctx).
		Where("id IN ? AND type = 'menu'", ids).
		Find(&permissions).Error
	if err != nil {
		return nil, fmt.Errorf("find menus by ids failed: %v", err)
	}

	// Convert to MenuModel
	var menus []*MenuModel
	for _, p := range permissions {
		menus = append(menus, permissionToMenuModel(p))
	}

	return menus, nil
}
