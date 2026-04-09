package sysManagement

import (
	"context"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
)

type RolePermissionService struct {
	roleRepository           modelSysManagement.RoleRepository
	rolePermissionRepository modelSysManagement.RolePermissionRepository
	ctx                      context.Context
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{
		roleRepository:           modelSysManagement.NewRoleRepo(global.TD27_DB),
		rolePermissionRepository: modelSysManagement.NewRolePermissionRepository(global.TD27_DB),
		ctx:                      context.Background(),
	}
}

func (s *RolePermissionService) Update(req *modelSysManagement.UpdateRolePermissionReq) error {
	// check role existence
	_, err := s.roleRepository.FindOne(s.ctx, req.RoleId)
	if err != nil {
		return err
	}

	// update sys_management_role_permissions
	err = s.rolePermissionRepository.Update(s.ctx, req.RoleId, req.PermissionIds, req.Domain)
	if err != nil {
		return err
	}

	// todo
	// if domain equals to API, refresh casbin

	return nil
}
