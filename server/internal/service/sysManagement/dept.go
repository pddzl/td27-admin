package sysManagement

import (
	"context"

	"server/internal/global"
	"server/internal/model/sysManagement"
)

type DeptService struct {
	deptRepository sysManagement.DeptRepository
	ctx            context.Context
}

func NewDeptService() *DeptService {
	return &DeptService{
		deptRepository: sysManagement.NewDeptEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

// List 获取部门列表（树形结构）
func (s *DeptService) List(req *sysManagement.DeptListReq) ([]*sysManagement.DeptResp, error) {
	depts, err := s.deptRepository.List(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// 构建部门树
	tree := sysManagement.BuildDeptTree(depts)
	return tree, nil
}

// Create 创建部门
func (s *DeptService) Create(req *sysManagement.CreateDeptReq) error {
	return s.deptRepository.Create(s.ctx, req)
}

// Update 更新部门
func (s *DeptService) Update(req *sysManagement.UpdateDeptReq) error {
	return s.deptRepository.Update(s.ctx, req)
}

// Delete 删除部门
func (s *DeptService) Delete(id uint) error {
	return s.deptRepository.Delete(s.ctx, id)
}

// GetElTreeDepts 获取部门树（用于el-tree选择器）
func (s *DeptService) GetElTreeDepts() ([]*sysManagement.DeptTreeResp, []uint, error) {
	req := &sysManagement.DeptListReq{}
	depts, err := s.deptRepository.List(s.ctx, req)
	if err != nil {
		return nil, nil, err
	}

	// 构建树
	tree := buildDeptTreeForElTree(depts)

	// 获取所有部门ID
	var allIDs []uint
	for _, dept := range depts {
		allIDs = append(allIDs, dept.ID)
	}

	return tree, allIDs, nil
}

// buildDeptTreeForElTree 构建el-tree需要的部门树
func buildDeptTreeForElTree(deptList []*sysManagement.DeptModel) []*sysManagement.DeptTreeResp {
	if len(deptList) == 0 {
		return []*sysManagement.DeptTreeResp{}
	}

	deptMap := make(map[uint]*sysManagement.DeptTreeResp, len(deptList))
	for _, dept := range deptList {
		deptMap[dept.ID] = &sysManagement.DeptTreeResp{
			ID:       dept.ID,
			DeptName: dept.DeptName,
			ParentID: dept.ParentID,
			Sort:     dept.Sort,
			Status:   dept.Status,
			Children: nil,
		}
	}

	var rootDepts []*sysManagement.DeptTreeResp
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

	return rootDepts
}
