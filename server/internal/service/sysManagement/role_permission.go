package sysManagement

import (
	"context"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
)

type RolePermissionService struct {
	roleRepository           modelSysManagement.RoleRepository
	rolePermissionRepository modelSysManagement.RolePermissionRepository
	casbinService            *CasbinService
	ctx                      context.Context
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{
		roleRepository:           modelSysManagement.NewRoleRepo(global.TD27_DB),
		rolePermissionRepository: modelSysManagement.NewRolePermissionRepository(global.TD27_DB),
		casbinService:            NewCasbinService(),
		ctx:                      context.Background(),
	}
}

func (s *RolePermissionService) Rebuild(req *modelSysManagement.RebuildRolePermissionReq) error {
	// check role existence
	_, err := s.roleRepository.FindOne(s.ctx, req.RoleId)
	if err != nil {
		return err
	}

	// update sys_management_role_permissions
	// returns the actual permissions that were inserted
	permissions, err := s.rolePermissionRepository.Rebuild(s.ctx, req.RoleId, req.PermissionIds, req.Domain)
	if err != nil {
		return err
	}

	// if domain is API, update casbin policies incrementally
	if req.Domain == string(modelSysManagement.PermissionDomainAPI) {
		if err = s.casbinService.RebuildRolePolicies(req.RoleId, permissions); err != nil {
			return err
		}
	}

	return nil
}
