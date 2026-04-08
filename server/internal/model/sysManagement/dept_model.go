package sysManagement

import "server/internal/model/common"

// DeptModel 部门模型（用于数据权限）
type DeptModel struct {
	common.Td27Model
	DeptName string `json:"deptName" gorm:"size:100;not null;comment:部门名称"`
	ParentID uint   `json:"parentId" gorm:"index;comment:父部门ID"`
	Path     string `json:"path" gorm:"size:500;index;comment:部门路径(物化路径),如:/1/2/3/"`
	Sort     uint   `json:"sort" gorm:"default:0"`
	Status   bool   `json:"status" gorm:"default:true"`
}

// GetFullPath 获取完整路径（包含自己）
func (d *DeptModel) GetFullPath() string {
	if d.Path == "" {
		return "/" + string(rune(d.ID))
	}
	return d.Path + string(rune(d.ID)) + "/"
}

// IsAncestorOf 检查当前部门是否是目标部门的祖先
func (d *DeptModel) IsAncestorOf(targetPath string) bool {
	fullPath := d.GetFullPath()
	return len(targetPath) > len(fullPath) && targetPath[:len(fullPath)] == fullPath
}

// IsDescendantOf 检查当前部门是否是目标部门的后代
func (d *DeptModel) IsDescendantOf(ancestorPath string) bool {
	return len(d.Path) >= len(ancestorPath) && d.Path[:len(ancestorPath)] == ancestorPath
}

func (DeptModel) TableName() string {
	return "sys_management_dept"
}
