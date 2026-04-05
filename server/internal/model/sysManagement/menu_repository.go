package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

// Menu 菜单请求结构
type Menu struct {
	Pid       uint   `json:"pid"`                          // 默认0 根目录
	Name      string `json:"name"`                         // 名称
	Path      string `json:"path" binding:"required"`      // 路径
	Redirect  string `json:"redirect"`                     // 重定向
	Component string `json:"component" binding:"required"` // 前端组件
	Sort      uint   `json:"sort" binding:"required"`      // 排序
	Meta      Meta   `json:"meta"`
}

// Meta 菜单元信息
type Meta struct {
	Hidden     bool   `json:"hidden"`     // 隐藏菜单
	Title      string `json:"title"`      // 菜单名
	Icon       string `json:"icon"`       // element 图标
	Affix      bool   `json:"affix"`      // 组件固定
	KeepAlive  bool   `json:"keepAlive"`  // 组件缓存
	AlwaysShow bool   `json:"alwaysShow"` // 是否一直显示根路由
}

// UpdateMenuReq 更新菜单请求
type UpdateMenuReq struct {
	ID uint `json:"id" binding:"required"` // 菜单 ID
	Menu
}

// MenuResp 菜单响应
type MenuResp struct {
	List    interface{} `json:"list"`
	MenuIds []uint      `json:"menuIds"`
}

// MenuItem 菜单项（用于树形结构）
type MenuItem struct {
	ID        uint       `json:"id"`
	MenuName  string     `json:"menuName"`
	Icon      string     `json:"icon"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	Redirect  string     `json:"redirect"`
	Pid       uint       `json:"pid"`
	Sort      uint       `json:"sort"`
	Hidden    bool       `json:"hidden"`
	KeepAlive bool       `json:"keepAlive"`
	Children  []MenuItem `json:"children,omitempty"`
}

// MenuData 前端菜单数据格式
type MenuData struct {
	ID        uint       `json:"id"`
	Pid       uint       `json:"pid"`
	Name      string     `json:"name"`      // 路由名称
	Path      string     `json:"path"`      // 路由路径
	Redirect  string     `json:"redirect"`  // 重定向
	Component string     `json:"component"` // 组件路径
	Sort      uint       `json:"sort"`      // 排序
	Meta      MenuMeta   `json:"meta"`      // 元信息
	Children  []MenuData `json:"children,omitempty"`
}

// MenuMeta 菜单元信息
type MenuMeta struct {
	Hidden     bool   `json:"hidden,omitempty"`
	Title      string `json:"title,omitempty"`
	SvgIcon    string `json:"svgIcon,omitempty"`
	ElIcon     string `json:"elIcon,omitempty"`
	Affix      bool   `json:"affix,omitempty"`
	KeepAlive  bool   `json:"keepAlive,omitempty"`
	AlwaysShow bool   `json:"alwaysShow,omitempty"`
}

// toMenuData converts MenuItem to frontend MenuData format
func toMenuData(item *MenuItem) *MenuData {
	return &MenuData{
		ID:        item.ID,
		Pid:       item.Pid,
		Name:      item.MenuName,
		Path:      item.Path,
		Redirect:  item.Redirect,
		Component: item.Component,
		Sort:      item.Sort,
		Meta: MenuMeta{
			Hidden:    item.Hidden,
			Title:     item.MenuName,
			SvgIcon:   item.Icon,
			KeepAlive: item.KeepAlive,
		},
	}
}

// ConvertToMenuDataList converts MenuItem tree to MenuData tree
func ConvertToMenuDataList(items []*MenuItem) []MenuData {
	var result []MenuData
	for _, item := range items {
		if item == nil {
			continue
		}
		menuData := *toMenuData(item)
		if len(item.Children) > 0 {
			children := make([]*MenuItem, len(item.Children))
			for i := range item.Children {
				children[i] = &item.Children[i]
			}
			menuData.Children = ConvertToMenuDataList(children)
		}
		result = append(result, menuData)
	}
	return result
}

type MenuRepository interface {
	List(ctx context.Context, roleId uint) ([]*MenuItem, error)
	ListByRoleIDs(ctx context.Context, roleIDs []uint) ([]*MenuItem, error)
	Create(ctx context.Context, req *Menu) (*MenuModel, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req *UpdateMenuReq) error
	GetElTreeMenus(ctx context.Context, roleId uint) ([]*MenuItem, []uint, error)
	FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error)
	GetAll(ctx context.Context) ([]*MenuModel, error)
}

type menuEntity struct {
	conn *gorm.DB
}

func NewMenuEntity(conn *gorm.DB) MenuRepository {
	return &menuEntity{conn: conn}
}

// toMenuItem converts MenuModel to MenuItem
func toMenuItem(m *MenuModel) *MenuItem {
	return &MenuItem{
		ID:        m.ID,
		MenuName:  m.MenuName,
		Icon:      m.Icon,
		Path:      m.Path,
		Component: m.Component,
		Redirect:  m.Redirect,
		Pid:       m.ParentID,
		Sort:      m.Sort,
		Hidden:    m.Hidden,
		KeepAlive: m.KeepAlive,
	}
}

// buildMenuTree 构建菜单树
func buildMenuTree(menus []*MenuItem) []*MenuItem {
	if len(menus) == 0 {
		return []*MenuItem{}
	}

	menuMap := make(map[uint]*MenuItem, len(menus))
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	var rootMenus []*MenuItem
	for _, menu := range menus {
		if menu.Pid == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			if parent, ok := menuMap[menu.Pid]; ok {
				parent.Children = append(parent.Children, *menu)
			}
		}
	}

	sort.Slice(rootMenus, func(i, j int) bool {
		return rootMenus[i].Sort < rootMenus[j].Sort
	})

	return rootMenus
}

func (e *menuEntity) List(ctx context.Context, roleId uint) ([]*MenuItem, error) {
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

	var menuItems []*MenuItem
	for _, m := range menus {
		menuItems = append(menuItems, toMenuItem(m))
	}

	return buildMenuTree(menuItems), nil
}

func (e *menuEntity) ListByRoleIDs(ctx context.Context, roleIDs []uint) ([]*MenuItem, error) {
	if len(roleIDs) == 0 {
		return []*MenuItem{}, nil
	}

	var menus []*MenuModel
	err := e.conn.WithContext(ctx).
		Model(&MenuModel{}).
		Joins("JOIN sys_management_permission p ON p.domain_id = sys_management_menu.id AND p.type = 'menu'").
		Joins("JOIN sys_management_role_permissions rp ON rp.permission_id = p.id").
		Where("rp.role_id IN ?", roleIDs).
		Where("sys_management_menu.status = ?", true).
		Group("sys_management_menu.id").
		Find(&menus).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role ids err: %w", err)
	}

	var menuItems []*MenuItem
	for _, m := range menus {
		menuItems = append(menuItems, toMenuItem(m))
	}

	return buildMenuTree(menuItems), nil
}

func (e *menuEntity) Create(ctx context.Context, req *Menu) (*MenuModel, error) {
	menu := &MenuModel{
		MenuName:  req.Meta.Title,
		Icon:      req.Meta.Icon,
		Path:      req.Path,
		Component: req.Component,
		Redirect:  req.Redirect,
		ParentID:  req.Pid,
		Sort:      req.Sort,
		Hidden:    req.Meta.Hidden,
		KeepAlive: req.Meta.KeepAlive,
	}

	if err := e.conn.WithContext(ctx).Create(menu).Error; err != nil {
		return nil, fmt.Errorf("create menu failed: %v", err)
	}

	return menu, nil
}

func (e *menuEntity) Update(ctx context.Context, req *UpdateMenuReq) error {
	updates := map[string]interface{}{
		"menu_name":  req.Meta.Title,
		"icon":       req.Meta.Icon,
		"path":       req.Path,
		"component":  req.Component,
		"redirect":   req.Redirect,
		"parent_id":  req.Pid,
		"sort":       req.Sort,
		"hidden":     req.Meta.Hidden,
		"keep_alive": req.Meta.KeepAlive,
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

func (e *menuEntity) GetElTreeMenus(ctx context.Context, roleId uint) ([]*MenuItem, []uint, error) {
	db := e.conn.WithContext(ctx)

	var allMenus []*MenuModel
	if err := db.Where("status = ?", true).Find(&allMenus).Error; err != nil {
		return nil, nil, fmt.Errorf("query menus failed: %w", err)
	}

	var menuItems []*MenuItem
	for _, m := range allMenus {
		menuItems = append(menuItems, toMenuItem(m))
	}

	rootMenus := buildMenuTree(menuItems)

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
	for _, m := range menuItems {
		if _, ok := permSet[m.ID]; !ok {
			continue
		}

		isParent := false
		for _, c := range menuItems {
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
