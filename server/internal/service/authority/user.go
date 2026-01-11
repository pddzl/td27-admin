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

func (us *UserService) GetUserInfo(userId uint) (*authority.UserResp, error) {
	resp, err := us.userRepository.GetUserInfo(us.ctx, userId)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (us *UserService) List(req *common.PageInfo) ([]*authority.UserResp, int64, error) {
	list, count, err := us.userRepository.List(us.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (us *UserService) Delete(id uint) error {
	return us.userRepository.Delete(us.ctx, id)
}

func (us *UserService) Create(req *authority.AddUserReq) error {
	// check role existence
	_, err := us.roleRepository.FindOne(us.ctx, req.RoleModelID)
	if err != nil {
		return err
	}

	return us.userRepository.Create(us.ctx, req)
}

func (us *UserService) Update(req *authority.UpdateUserReq) (*authority.UserResp, error) {
	// check role existence
	role, err := us.roleRepository.FindOne(us.ctx, req.RoleModelID)
	if err != nil {
		return nil, err
	}

	update, err := us.userRepository.Update(us.ctx, req)
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
func (us *UserService) ModifyPasswd(req *authority.ModifyPasswdReq) error {
	return us.userRepository.ModifyPasswd(us.ctx, req)
}

// SwitchActive 切换启用状态
func (us *UserService) SwitchActive(req *authority.SwitchActiveReq) error {
	return us.userRepository.SwitchActive(us.ctx, req)
}
