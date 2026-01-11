package authority

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type MenuEntity interface {
	List(context.Context, uint) ([]*MenuModel, error)
	Create(context.Context, *Menu) error
	Delete(context.Context, uint) error
	Update(context.Context, *UpdateMenuReq) error
	GetElTreeMenus(context.Context, uint) ([]*MenuModel, []uint, error)
}

type defaultMenuEntity struct {
	conn *gorm.DB
}

func NewDefaultMenuEntity(conn *gorm.DB) MenuEntity {
	return &defaultMenuEntity{conn: conn}
}

func getTreeMap(menuListFormat []*MenuModel, menuList []*MenuModel) {
	for index, menuF := range menuListFormat {
		for _, menu := range menuList {
			if menuF.ID == menu.Pid {
				// menuF 只是个复制值
				//menuF.Children = append(menuF.Children, menu)
				menuListFormat[index].Children = append(menuListFormat[index].Children, menu)
			}
		}
		if len(menuListFormat[index].Children) > 0 {
			// 排序
			sort.Slice(menuListFormat[index].Children, func(i, j int) bool {
				return menuListFormat[index].Children[i].Sort < menuListFormat[index].Children[j].Sort
			})
			getTreeMap(menuListFormat[index].Children, menuList)
		}
	}
}

func (e *defaultMenuEntity) List(ctx context.Context, roleId uint) ([]*MenuModel, error) {
	var menuModels []*MenuModel
	err := e.conn.WithContext(ctx).
		Table("authority_menu").
		Joins("JOIN role_menus rm ON rm.menu_id = authority_menu.id").
		Where("rm.role_id = ?", roleId).
		Order("authority_menu.sort ASC").
		Find(&menuModels).Error

	if err != nil {
		return nil, fmt.Errorf("list menus by role err: %w", err)
	}

	// 找出第一级路由，（父路由id为0）
	menuListFormat := make([]*MenuModel, 0)
	for _, menu := range menuModels {
		if menu.Pid == 0 {
			menuListFormat = append(menuListFormat, menu)
		}
	}

	// 排序
	sort.Slice(menuListFormat, func(i, j int) bool {
		return menuListFormat[i].Sort < menuListFormat[j].Sort
	})

	// 递归找出一级路由下面的子路由
	getTreeMap(menuListFormat, menuModels)

	return menuListFormat, nil
}

func (e *defaultMenuEntity) Create(ctx context.Context, req *Menu) error {
	var menuModel MenuModel
	menuModel.Name = req.Name
	menuModel.Path = req.Path
	menuModel.Component = req.Component
	menuModel.Redirect = req.Redirect
	menuModel.Pid = req.Pid
	menuModel.Sort = req.Sort
	menuModel.Meta.Title = req.Meta.Title
	menuModel.Meta.SvgIcon = req.Meta.Icon
	menuModel.Meta.Hidden = req.Meta.Hidden
	menuModel.Meta.Affix = req.Meta.Affix
	menuModel.Meta.KeepAlive = req.Meta.KeepAlive
	menuModel.Meta.AlwaysShow = req.Meta.AlwaysShow

	if err := e.conn.WithContext(ctx).Create(&menuModel).Error; err != nil {
		return fmt.Errorf("create menu failed: %v", err)
	}

	return nil
}

func (e *defaultMenuEntity) Update(ctx context.Context, req *UpdateMenuReq) error {
	updates := MenuModel{
		Pid:       req.Pid,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Redirect:  req.Redirect,
		Sort:      req.Sort,
		Meta: Meta{
			SvgIcon:    req.Meta.Icon,
			Title:      req.Meta.Title,
			Hidden:     req.Meta.Hidden,
			Affix:      req.Meta.Affix,
			KeepAlive:  req.Meta.KeepAlive,
			AlwaysShow: req.Meta.AlwaysShow,
		},
	}

	result := e.conn.WithContext(ctx).
		Model(&MenuModel{}).
		Where("id = ?", req.ID).
		Updates(&updates)

	if err := result.Error; err != nil {
		return fmt.Errorf("update menu failed, id=%d: %w", req.ID, err)
	}

	if result.RowsAffected == 0 {
		return errors.New("菜单不存在")
	}

	return nil
}

func (e *defaultMenuEntity) Delete(ctx context.Context, id uint) error {
	tx := e.conn.WithContext(ctx).Unscoped().Select("Roles").Delete(&MenuModel{})

	if err := tx.Error; err != nil {
		return fmt.Errorf("delete menu failed, id=%d: %w", id, err)
	}

	if tx.RowsAffected == 0 {
		return errors.New("菜单不存在")
	}

	return nil
}

// GetElTreeMenus 获取所有menu
func (e *defaultMenuEntity) GetElTreeMenus(ctx context.Context, roleId uint) ([]*MenuModel, []uint, error) {
	// Query all menus
	var allMenus []*MenuModel
	if err := e.conn.WithContext(ctx).
		Order("sort ASC").
		Find(&allMenus).Error; err != nil {
		return nil, nil, fmt.Errorf("query menus failed: %w", err)
	}

	// Build menu tree
	rootMenus := make([]*MenuModel, 0)
	for _, m := range allMenus {
		if m.Pid == 0 {
			rootMenus = append(rootMenus, m)
		}
	}

	getTreeMap(rootMenus, allMenus)

	// Query role-menu relations
	var relations []RoleMenu
	if err := e.conn.WithContext(ctx).
		Table("role_menus").
		Select("menu_model_id").
		Where("role_model_id = ?", roleId).
		Scan(&relations).Error; err != nil {
		return nil, nil, fmt.Errorf("query role menus failed: %w", err)
	}

	// Build menu ID set
	menuSet := make(map[uint]struct{}, len(relations))
	for _, r := range relations {
		menuSet[r.MenuID] = struct{}{}
	}

	// el-tree leaf selection
	checkedIDs := make([]uint, 0)
	for _, m := range allMenus {
		if _, ok := menuSet[m.ID]; !ok {
			continue
		}

		// if has child selected → skip parent
		isParent := false
		for _, c := range allMenus {
			if c.Pid == m.ID {
				if _, ok := menuSet[c.ID]; ok {
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
