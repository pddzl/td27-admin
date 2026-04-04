package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type DeptRepository interface {
	List(ctx context.Context, req *DeptListReq) ([]*DeptModel, error)
	Create(ctx context.Context, req *CreateDeptReq) error
	Update(ctx context.Context, req *UpdateDeptReq) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*DeptModel, error)
	GetChildren(ctx context.Context, parentID uint) ([]*DeptModel, error)
	GetDescendants(ctx context.Context, deptID uint) ([]*DeptModel, error) // 获取所有后代（使用物化路径）
	HasChildren(ctx context.Context, id uint) (bool, error)
	HasUsers(ctx context.Context, id uint) (bool, error)
}

type deptEntity struct {
	conn *gorm.DB
}

func NewDeptEntity(conn *gorm.DB) DeptRepository {
	return &deptEntity{conn: conn}
}

func (e *deptEntity) List(ctx context.Context, req *DeptListReq) ([]*DeptModel, error) {
	db := e.conn.WithContext(ctx).Model(&DeptModel{})

	if req.DeptName != "" {
		db = db.Where("dept_name LIKE ?", "%"+req.DeptName+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	var depts []*DeptModel
	if err := db.Find(&depts).Error; err != nil {
		return nil, fmt.Errorf("list depts failed: %w", err)
	}

	return depts, nil
}

func (e *deptEntity) Create(ctx context.Context, req *CreateDeptReq) error {
	// 计算物化路径
	path, err := e.buildPath(ctx, req.ParentID)
	if err != nil {
		return fmt.Errorf("build path failed: %w", err)
	}

	dept := DeptModel{
		DeptName: req.DeptName,
		ParentID: req.ParentID,
		Path:     path,
		Sort:     req.Sort,
		Status:   req.Status,
	}

	if err := e.conn.WithContext(ctx).Create(&dept).Error; err != nil {
		return fmt.Errorf("create dept failed: %w", err)
	}

	// 更新路径（包含自己的ID）
	dept.Path = dept.GetFullPath()
	if err := e.conn.WithContext(ctx).Model(&dept).Update("path", dept.Path).Error; err != nil {
		return fmt.Errorf("update dept path failed: %w", err)
	}

	return nil
}

// buildPath 根据父ID构建物化路径
func (e *deptEntity) buildPath(ctx context.Context, parentID uint) (string, error) {
	if parentID == 0 {
		return "/", nil
	}

	var parent DeptModel
	if err := e.conn.WithContext(ctx).First(&parent, parentID).Error; err != nil {
		return "", fmt.Errorf("get parent dept failed: %w", err)
	}

	return parent.GetFullPath(), nil
}

func (e *deptEntity) Update(ctx context.Context, req *UpdateDeptReq) error {
	// 如果父部门改变，需要重新计算路径
	var oldDept DeptModel
	if err := e.conn.WithContext(ctx).First(&oldDept, req.ID).Error; err != nil {
		return fmt.Errorf("get old dept failed: %w", err)
	}

	updates := map[string]interface{}{
		"dept_name": req.DeptName,
		"parent_id": req.ParentID,
		"sort":      req.Sort,
		"status":    req.Status,
	}

	// 如果父部门改变，更新路径
	if oldDept.ParentID != req.ParentID {
		newPath, err := e.buildPath(ctx, req.ParentID)
		if err != nil {
			return fmt.Errorf("build new path failed: %w", err)
		}
		updates["path"] = newPath + string(rune(req.ID)) + "/"
	}

	result := e.conn.WithContext(ctx).
		Model(&DeptModel{}).
		Where("id = ?", req.ID).
		Updates(updates)

	if err := result.Error; err != nil {
		return fmt.Errorf("update dept failed: %w", err)
	}

	if result.RowsAffected == 0 {
		return errors.New("部门不存在")
	}

	// 如果父部门改变，更新所有子部门的路径
	if oldDept.ParentID != req.ParentID {
		if err := e.updateChildrenPath(ctx, req.ID, updates["path"].(string)); err != nil {
			return fmt.Errorf("update children path failed: %w", err)
		}
	}

	return nil
}

// updateChildrenPath 递归更新所有子部门的路径
func (e *deptEntity) updateChildrenPath(ctx context.Context, parentID uint, newParentPath string) error {
	var children []DeptModel
	if err := e.conn.WithContext(ctx).Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return err
	}

	for _, child := range children {
		newPath := newParentPath + string(rune(child.ID)) + "/"
		if err := e.conn.WithContext(ctx).Model(&child).Update("path", newPath).Error; err != nil {
			return err
		}
		// 递归更新子部门的子部门
		if err := e.updateChildrenPath(ctx, child.ID, newPath); err != nil {
			return err
		}
	}

	return nil
}

func (e *deptEntity) Delete(ctx context.Context, id uint) error {
	// Check if has children
	hasChildren, err := e.HasChildren(ctx, id)
	if err != nil {
		return err
	}
	if hasChildren {
		return errors.New("该部门下存在子部门，无法删除")
	}

	// Check if has users
	hasUsers, err := e.HasUsers(ctx, id)
	if err != nil {
		return err
	}
	if hasUsers {
		return errors.New("该部门下存在用户，无法删除")
	}

	result := e.conn.WithContext(ctx).Delete(&DeptModel{}, id)
	if err := result.Error; err != nil {
		return fmt.Errorf("delete dept failed: %w", err)
	}

	if result.RowsAffected == 0 {
		return errors.New("部门不存在")
	}

	return nil
}

func (e *deptEntity) GetByID(ctx context.Context, id uint) (*DeptModel, error) {
	var dept DeptModel
	if err := e.conn.WithContext(ctx).First(&dept, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("部门不存在")
		}
		return nil, fmt.Errorf("get dept failed: %w", err)
	}
	return &dept, nil
}

func (e *deptEntity) GetChildren(ctx context.Context, parentID uint) ([]*DeptModel, error) {
	var depts []*DeptModel
	if err := e.conn.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Find(&depts).Error; err != nil {
		return nil, fmt.Errorf("get children depts failed: %w", err)
	}
	return depts, nil
}

// GetDescendants 使用物化路径获取所有后代部门（包含所有层级）
func (e *deptEntity) GetDescendants(ctx context.Context, deptID uint) ([]*DeptModel, error) {
	// 先获取当前部门的路径
	var dept DeptModel
	if err := e.conn.WithContext(ctx).First(&dept, deptID).Error; err != nil {
		return nil, fmt.Errorf("get dept failed: %w", err)
	}

	// 使用物化路径查询所有后代：路径以当前部门路径开头
	var depts []*DeptModel
	if err := e.conn.WithContext(ctx).
		Where("path LIKE ?", dept.GetFullPath()+"%").
		Where("id != ?", deptID). // 排除自己
		Find(&depts).Error; err != nil {
		return nil, fmt.Errorf("get descendants failed: %w", err)
	}
	return depts, nil
}

func (e *deptEntity) HasChildren(ctx context.Context, id uint) (bool, error) {
	var count int64
	if err := e.conn.WithContext(ctx).
		Model(&DeptModel{}).
		Where("parent_id = ?", id).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("check children failed: %w", err)
	}
	return count > 0, nil
}

func (e *deptEntity) HasUsers(ctx context.Context, id uint) (bool, error) {
	var count int64
	if err := e.conn.WithContext(ctx).
		Model(&UserModel{}).
		Where("dept_id = ?", id).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("check users failed: %w", err)
	}
	return count > 0, nil
}

// BuildDeptTree 构建部门树（O(n)算法）
func BuildDeptTree(deptList []*DeptModel) []*DeptResp {
	if len(deptList) == 0 {
		return []*DeptResp{}
	}

	// 使用 map 存储所有部门
	deptMap := make(map[uint]*DeptResp, len(deptList))
	for _, dept := range deptList {
		deptMap[dept.ID] = &DeptResp{
			DeptModel: *dept,
			Children:  nil,
		}
	}

	// 构建树结构
	var rootDepts []*DeptResp
	for _, dept := range deptList {
		node := deptMap[dept.ID]
		if dept.ParentID == 0 {
			rootDepts = append(rootDepts, node)
		} else {
			if parent, ok := deptMap[dept.ParentID]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	// 排序
	sortDeptChildren(rootDepts)

	return rootDepts
}

func sortDeptChildren(depts []*DeptResp) {
	if len(depts) == 0 {
		return
	}
	sort.Slice(depts, func(i, j int) bool {
		return depts[i].Sort < depts[j].Sort
	})
	for _, dept := range depts {
		if len(dept.Children) > 0 {
			sort.Slice(dept.Children, func(i, j int) bool {
				return dept.Children[i].Sort < dept.Children[j].Sort
			})
			sortDeptChildren(dept.Children)
		}
	}
}
