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

func (e *menuEntity) List(ctx context.Context, roleId uint) ([]*MenuModel, error) {
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

func (e *menuEntity) Create(ctx context.Context, req *Menu) error {
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

func (e *menuEntity) Update(ctx context.Context, req *UpdateMenuReq) error {
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

func (e *menuEntity) Delete(ctx context.Context, id uint) error {
	var menuModel MenuModel
	db := e.conn.WithContext(ctx)

	if errors.Is(db.Where("id = ?", id).First(&menuModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单不存在")
	}

	err := db.Unscoped().Select("Roles").Delete(&menuModel).Error

	return err
}

// GetElTreeMenus 获取所有menu
func (e *menuEntity) GetElTreeMenus(ctx context.Context, roleId uint) ([]*MenuModel, []uint, error) {
	db := e.conn.WithContext(ctx)

	// Query all menus
	var allMenus []*MenuModel
	if err := db.Find(&allMenus).Error; err != nil {
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
	if err := db.
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

func (e *menuEntity) FindByIds(ctx context.Context, ids []uint) ([]*MenuModel, error) {
	// Guard: avoid full table scan
	if len(ids) == 0 {
		return []*MenuModel{}, nil
	}

	var menus []*MenuModel
	err := e.conn.WithContext(ctx).Where("id IN ?", ids).Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("find menus by ids failed: %v", err)
	}
	return menus, nil
}
