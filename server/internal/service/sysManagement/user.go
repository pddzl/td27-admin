package sysManagement

import (
	"context"

	"server/internal/global"
	"server/internal/model/common"
	"server/internal/model/sysManagement"
)

type UserService struct {
	userRepository sysManagement.UserRepository
	roleRepository sysManagement.RoleRepository
	ctx            context.Context
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: sysManagement.NewUserEntity(global.TD27_DB),
		roleRepository: sysManagement.NewRoleEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *UserService) GetUserInfo(userId uint) (*sysManagement.UserResp, error) {
	resp, err := s.userRepository.GetUserInfo(s.ctx, userId)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *UserService) List(req *common.PageInfo, currentUserID uint) ([]*sysManagement.UserResp, int64, error) {
	// 获取当前用户的数据权限
	dataPermService := NewDataPermissionService()
	dataPerm, err := dataPermService.GetUserDataPermission(s.ctx, currentUserID, "sys_management_user")
	if err != nil {
		// 如果获取失败，默认只能看自己的数据
		dataPerm = &sysManagement.DataPermission{
			Scope:  sysManagement.DataScopeSelf,
			UserID: currentUserID,
		}
	}

	list, count, err := s.userRepository.List(s.ctx, req, dataPerm)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (s *UserService) Delete(id uint) error {
	// 删除用户前清除缓存
	jwtService := NewJwtService()
	if err := jwtService.DeleteUserCache(id); err != nil {
		// 缓存删除失败不影响主流程
	}
	return s.userRepository.Delete(s.ctx, id)
}

func (s *UserService) Create(req *sysManagement.AddUserReq) error {
	// check roles existence
	for _, roleID := range req.RoleIDs {
		_, err := s.roleRepository.FindOne(s.ctx, roleID)
		if err != nil {
			return err
		}
	}

	_, err := s.userRepository.Create(s.ctx, req)
	return err
}

func (s *UserService) Update(req *sysManagement.UpdateUserReq) (*sysManagement.UserResp, error) {
	// check roles existence
	var primaryRole *sysManagement.RoleModel
	for _, roleID := range req.RoleIDs {
		role, err := s.roleRepository.FindOne(s.ctx, roleID)
		if err != nil {
			return nil, err
		}
		if primaryRole == nil {
			primaryRole = role
		}
	}

	update, err := s.userRepository.Update(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// 更新用户后清除缓存，下次请求会重新加载
	jwtService := NewJwtService()
	if cacheErr := jwtService.DeleteUserCache(req.ID); cacheErr != nil {
		// 缓存删除失败不影响主流程
	}

	userResp := &sysManagement.UserResp{
		UserModel: *update,
	}
	if primaryRole != nil {
		userResp.RoleName = primaryRole.RoleName
		userResp.RoleID = primaryRole.ID
	}

	return userResp, nil
}

// ModifyPasswd 修改用户密码
func (s *UserService) ModifyPasswd(req *sysManagement.ModifyPasswdReq) error {
	err := s.userRepository.ModifyPasswd(s.ctx, req)
	if err != nil {
		return err
	}
	// 修改密码后清除缓存，强制重新登录
	jwtService := NewJwtService()
	jwtService.DeleteUserCache(req.ID)
	return nil
}

// SwitchActive 切换启用状态
func (s *UserService) SwitchActive(req *sysManagement.SwitchActiveReq) error {
	// 切换状态前清除缓存
	jwtService := NewJwtService()
	if err := jwtService.DeleteUserCache(req.ID); err != nil {
		// 缓存删除失败不影响主流程
	}
	return s.userRepository.SwitchActive(s.ctx, req)
}
