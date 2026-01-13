package authority

import (
	"context"

	"server/internal/global"
	"server/internal/model/authority"
	"server/internal/model/common"
)

type UserService struct {
	userRepository authority.UserEntity
	roleRepository authority.RoleEntity
	ctx            context.Context
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: authority.NewDefaultUserEntity(global.TD27_DB),
		roleRepository: authority.NewDefaultRoleEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *UserService) GetUserInfo(userId uint) (*authority.UserResp, error) {
	resp, err := s.userRepository.GetUserInfo(s.ctx, userId)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *UserService) List(req *common.PageInfo) ([]*authority.UserResp, int64, error) {
	list, count, err := s.userRepository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (s *UserService) Delete(id uint) error {
	return s.userRepository.Delete(s.ctx, id)
}

func (s *UserService) Create(req *authority.AddUserReq) error {
	// check role existence
	_, err := s.roleRepository.FindOne(s.ctx, req.RoleModelID)
	if err != nil {
		return err
	}

	return s.userRepository.Create(s.ctx, req)
}

func (s *UserService) Update(req *authority.UpdateUserReq) (*authority.UserResp, error) {
	// check role existence
	role, err := s.roleRepository.FindOne(s.ctx, req.RoleModelID)
	if err != nil {
		return nil, err
	}

	update, err := s.userRepository.Update(s.ctx, req)
	if err != nil {
		return nil, err
	}

	userResp := &authority.UserResp{
		UserModel: *update,
		RoleName:  role.RoleName,
	}

	return userResp, nil
}

// ModifyPasswd 修改用户密码
func (s *UserService) ModifyPasswd(req *authority.ModifyPasswdReq) error {
	return s.userRepository.ModifyPasswd(s.ctx, req)
}

// SwitchActive 切换启用状态
func (s *UserService) SwitchActive(req *authority.SwitchActiveReq) error {
	return s.userRepository.SwitchActive(s.ctx, req)
}
