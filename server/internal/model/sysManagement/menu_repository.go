package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

// Menu 菜单请求结构（前端传入）
type Menu struct {
	ParentId   uint   `json:"parentId"`                     // 默认0 根目录
	MenuName   string `json:"menu_name"`                    // 名称（路由名）
	Path       string `json:"path" binding:"required"`      // 路径
	Redirect   string `json:"redirect"`                     // 重定向
	Component  string `json:"component" binding:"required"` // 前端组件
	Sort       uint   `json:"sort" binding:"required"`      // 排序
	Hidden     bool   `json:"hidden"`                       // 隐藏菜单
	Title      string `json:"title"`                        // 菜单名
	Icon       string `json:"icon"`                         // 图标
	Affix      bool   `json:"affix"`                        // 固定
	KeepAlive  bool   `json:"keepAlive"`                    // 缓存
	AlwaysShow bool   `json:"alwaysShow"`                   // 一直显示根路由
}

// UpdateMenuReq 更新菜单请求
type UpdateMenuReq struct {
	ID uint `json:"id" binding:"required"` // 菜单 ID
	Menu
}

// MenuResp 菜单响应（扁平化结构，直接返回给前端）
type MenuResp struct {
	ID         uint       `json:"id"`
	MenuName   string     `json:"menu_name"`
	Icon       string     `json:"icon"`
	Path       string     `json:"path"`
	Component  string     `json:"component"`
	Redirect   string     `json:"redirect"`
	ParentID   uint       `json:"parentId"`
	Sort       uint       `json:"sort"`
	Hidden     bool       `json:"hidden"`
	KeepAlive  bool       `json:"keepAlive"`
	Affix      bool       `json:"affix,omitempty"`
	AlwaysShow bool       `json:"alwaysShow,omitempty"`
	Title      string     `json:"title,omitempty"`
	Children   []MenuResp `json:"children,omitempty"`
}

// MenuElTreeResp el-tree菜单响应
type MenuElTreeResp struct {
	List    []MenuResp `json:"list"`
	MenuIds []uint     `json:"menuIds"`
}

type MenuRepository interface {
	List(ctx context.Context, roleId uint) ([]MenuResp, error)
	ListByRoleIDs(ctx context.Context, roleIDs []uint) ([]MenuResp, error)
	Create(ctx context.Context, req *Menu) (*MenuModel, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *UpdateMenuReq) error
	GetElTreeMenus(ctx context.Context, roleId uint) ([]MenuResp, []uint, error)
	FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error)
	GetAll(ctx context.Context) ([]*MenuModel, error)
}

type menuEntity struct {
	conn *gorm.DB
}

func NewMenuEntity(conn *gorm.DB) MenuRepository {
	return &menuEntity{conn: conn}
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
		menuCopy := menus[i] // create a copy
		menuMap[menus[i].ID] = &menuCopy
	}

	// Build tree structure
	var rootMenus []*MenuResp
	for _, menu := range menuMap {
		if menu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			if parent, ok := menuMap[menu.ParentID]; ok {
				parent.Children = append(parent.Children, *menu)
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

func (e *menuEntity) List(ctx context.Context, roleId uint) ([]MenuResp, error) {
	// Query menus through role_permissions -> permissions -> menus
	var menus []*MenuModel
	err := e.conn.WithContext(ctx).
		Model(&MenuModel{}).
		Joins("JOIN sys_management_permission p ON p.domain_id = sys_management_menu.id AND p.type = 'menu'").
		Joins("JOIN sys_management_role_permissions rp ON rp.permission_id = p.id").
		Where("rp.role_id = ?", roleId).
		Where("sys_management_menu.status = ?", true).
		Find(&menus).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role err: %w", err)
	}

	var menuResps []MenuResp
	for _, m := range menus {
		menuResps = append(menuResps, toMenuResp(m))
	}

	return buildMenuTree(menuResps), nil
}

func (e *menuEntity) ListByRoleIDs(ctx context.Context, roleIDs []uint) ([]MenuResp, error) {
	if len(roleIDs) == 0 {
		return []MenuResp{}, nil
	}

	var menus []*MenuModel
	err := e.conn.WithContext(ctx).
		Model(&MenuModel{}).
		Joins("JOIN sys_management_permission p ON p.domain_id = sys_management_menu.id AND p.type = 'menu'").
		Joins("JOIN sys_management_role_permissions rp ON rp.permission_id = p.id").
		Where("rp.role_id IN ?", roleIDs).
		Group("sys_management_menu.id").
		Find(&menus).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role ids err: %w", err)
	}

	var menuResps []MenuResp
	for _, m := range menus {
		menuResps = append(menuResps, toMenuResp(m))
	}

	return buildMenuTree(menuResps), nil
}

func (e *menuEntity) Create(ctx context.Context, req *Menu) (*MenuModel, error) {
	menu := &MenuModel{
		MenuName:   req.Title,
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

func (e *menuEntity) Update(ctx context.Context, req *UpdateMenuReq) error {
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

func (e *menuEntity) Delete(ctx context.Context, id uint) error {
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

func (e *menuEntity) GetElTreeMenus(ctx context.Context, roleId uint) ([]MenuResp, []uint, error) {
	db := e.conn.WithContext(ctx)

	var allMenus []*MenuModel
	if err := db.Where("status = ?", true).Find(&allMenus).Error; err != nil {
		return nil, nil, fmt.Errorf("query menus failed: %w", err)
	}

	var menuResps []MenuResp
	for _, m := range allMenus {
		menuResps = append(menuResps, toMenuResp(m))
	}

	rootMenus := buildMenuTree(menuResps)

	// Query role-permission relations
	var relations []RolePermission
	if err := db.
		Model(&RolePermission{}).
		Select("permission_id").
		Where("role_id = ?", roleId).
		Scan(&relations).Error; err != nil {
		return nil, nil, fmt.Errorf("query role menus failed: %w", err)
	}

	permSet := make(map[uint]struct{}, len(relations))
	for _, r := range relations {
		permSet[r.PermissionID] = struct{}{}
	}

	checkedIDs := make([]uint, 0)
	for _, m := range menuResps {
		if _, ok := permSet[m.ID]; !ok {
			continue
		}

		isParent := false
		for _, c := range menuResps {
			if c.ParentID == m.ID {
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

func (e *menuEntity) GetAll(ctx context.Context) ([]*MenuModel, error) {
	var menus []*MenuModel
	err := e.conn.WithContext(ctx).
		Where("status = ?", true).
		Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("get all menus failed: %w", err)
	}
	return menus, nil
}
