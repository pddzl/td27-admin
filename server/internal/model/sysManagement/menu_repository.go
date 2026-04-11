package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type MenuRepository interface {
	List(ctx context.Context, roleIDs []uint) ([]MenuResp, error)
	Create(ctx context.Context, req *Menu) (*MenuModel, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *UpdateMenuReq) error
	ElTree(ctx context.Context, roleId uint) ([]MenuResp, []uint, error)
	FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error)
	GetAll(ctx context.Context) ([]*MenuModel, error)
}

type menuRepo struct {
	conn *gorm.DB
}

func NewMenuRepo(conn *gorm.DB) MenuRepository {
	return &menuRepo{conn: conn}
}

// toMenuResp converts MenuModel to MenuResp
func toMenuResp(m *MenuModel) MenuResp {
	return MenuResp{
		ID:         m.ID,
		MenuName:   m.MenuName,
		Icon:       m.Icon,
		Path:       m.Path,
		Component:  m.Component,
		Redirect:   m.Redirect,
		ParentID:   m.ParentID,
		Sort:       m.Sort,
		Hidden:     m.Hidden,
		KeepAlive:  m.KeepAlive,
		Affix:      m.Affix,
		AlwaysShow: m.AlwaysShow,
		Title:      m.Title,
	}
}

// buildMenuTree 构建菜单树
func buildMenuTree(menus []MenuResp) []MenuResp {
	if len(menus) == 0 {
		return []MenuResp{}
	}

	// Use map to store menus by ID
	menuMap := make(map[uint]*MenuResp, len(menus))
	for i := range menus {
		//menuCopy := menus[i] // create a copy
		//menuMap[menus[i].ID] = &menuCopy
		menuMap[menus[i].ID] = &menus[i]
	}

	// Build tree structure
	var rootMenus []*MenuResp
	for _, menu := range menuMap {
		if menu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			if parent, ok := menuMap[menu.ParentID]; ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	// Sort root menus
	sort.Slice(rootMenus, func(i, j int) bool {
		return rootMenus[i].Sort < rootMenus[j].Sort
	})

	// Convert back to value slice
	result := make([]MenuResp, len(rootMenus))
	for i, rm := range rootMenus {
		result[i] = *rm
	}

	return result
}

func (e *menuRepo) List(ctx context.Context, roleIDs []uint) ([]MenuResp, error) {
	if len(roleIDs) == 0 {
		return []MenuResp{}, nil
	}

	var menus []*MenuModel

	subQuery := e.conn.
		Table("sys_management_permission p").
		Select("1").
		Joins("JOIN sys_management_role_permissions rp ON rp.permission_id = p.id").
		Where("p.domain_id = sys_management_menu.id").
		Where("p.domain = ?", "menu").
		Where("rp.role_id IN ?", roleIDs)

	err := e.conn.WithContext(ctx).
		Model(&MenuModel{}).
		Where("EXISTS (?)", subQuery).
		Find(&menus).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role ids err: %w", err)
	}

	menuResps := make([]MenuResp, 0, len(menus))
	for _, m := range menus {
		menuResps = append(menuResps, toMenuResp(m))
	}

	return buildMenuTree(menuResps), nil
}

func (e *menuRepo) Create(ctx context.Context, req *Menu) (*MenuModel, error) {
	menu := &MenuModel{
		MenuName:   req.MenuName,
		Icon:       req.Icon,
		Path:       req.Path,
		Component:  req.Component,
		Redirect:   req.Redirect,
		ParentID:   req.ParentId,
		Sort:       req.Sort,
		Hidden:     req.Hidden,
		KeepAlive:  req.KeepAlive,
		Affix:      req.Affix,
		AlwaysShow: req.AlwaysShow,
		Title:      req.Title,
	}

	if err := e.conn.WithContext(ctx).Create(menu).Error; err != nil {
		return nil, fmt.Errorf("create menu failed: %v", err)
	}

	return menu, nil
}

func (e *menuRepo) Update(ctx context.Context, req *UpdateMenuReq) error {
	updates := map[string]interface{}{
		"menu_name":   req.Title,
		"icon":        req.Icon,
		"path":        req.Path,
		"component":   req.Component,
		"redirect":    req.Redirect,
		"parent_id":   req.ParentId,
		"sort":        req.Sort,
		"hidden":      req.Hidden,
		"keep_alive":  req.KeepAlive,
		"affix":       req.Affix,
		"always_show": req.AlwaysShow,
		"title":       req.Title,
	}

	result := e.conn.WithContext(ctx).
		Model(&MenuModel{}).
		Where("id = ?", req.ID).
		Updates(updates)

	if err := result.Error; err != nil {
		return fmt.Errorf("update menu failed, id=%d: %w", req.ID, err)
	}

	if result.RowsAffected == 0 {
		return errors.New("菜单不存在")
	}

	return nil
}

func (e *menuRepo) Delete(ctx context.Context, id uint) error {
	result := e.conn.WithContext(ctx).
		Where("id = ?", id).
		Delete(&MenuModel{})

	if err := result.Error; err != nil {
		return fmt.Errorf("delete menu failed: %w", err)
	}

	if result.RowsAffected == 0 {
		return errors.New("菜单不存在")
	}

	return nil
}

func (e *menuRepo) ElTree(ctx context.Context, roleId uint) ([]MenuResp, []uint, error) {
	db := e.conn.WithContext(ctx)

	var allMenus []*MenuModel
	if err := db.Find(&allMenus).Error; err != nil {
		return nil, nil, fmt.Errorf("query menus failed: %w", err)
	}

	var menuResps []MenuResp
	for _, m := range allMenus {
		menuResps = append(menuResps, toMenuResp(m))
	}

	rootMenus := buildMenuTree(menuResps)

	// Query role-permission relations
	var relations []RolePermissionModel
	if err := db.
		Model(&RolePermissionModel{}).
		Select("permission_id").
		Where("role_id = ?", roleId).
		Scan(&relations).Error; err != nil {
		return nil, nil, fmt.Errorf("query role menus failed: %w", err)
	}

	permSet := make(map[uint]struct{}, len(relations))
	for _, r := range relations {
		permSet[r.PermissionID] = struct{}{}
	}

	// element-plus el-tree default-checked-keys
	// find the checked node without children
	// the child checked, parent auto checked, if you checked parent, its all child will be checked, it's wrong
	checkedIDs := make([]uint, 0)
	for _, m := range menuResps {
		if _, ok := permSet[m.ID]; !ok {
			continue
		}

		isParent := false
		for _, c := range menuResps {
			if c.ParentID == m.ID {
				//if _, ok := permSet[c.ID]; ok {
				//	isParent = true
				//	break
				//}
				isParent = true
				break
			}
		}

		if !isParent {
			checkedIDs = append(checkedIDs, m.ID)
		}
	}

	return rootMenus, checkedIDs, nil
}

func (e *menuRepo) FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error) {
	if len(ids) == 0 {
		return []*MenuModel{}, nil
	}

	var menus []*MenuModel
	err := e.conn.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("find menus by ids failed: %v", err)
	}

	return menus, nil
}

func (e *menuRepo) GetAll(ctx context.Context) ([]*MenuModel, error) {
	var menus []*MenuModel
	err := e.conn.WithContext(ctx).
		Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("get all menus failed: %w", err)
	}
	return menus, nil
}
