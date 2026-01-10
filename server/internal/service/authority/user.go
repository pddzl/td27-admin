package authority

import (
	"context"

	"server/internal/global"
	"server/internal/model/authority/user"
	"server/internal/model/common"
)

type UserService struct {
	userRepository user.UserEntity
	ctx            context.Context
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: user.NewDefaultUserEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (us *UserService) GetUserInfo(userId uint) (*user.UserResp, error) {
	resp, err := us.userRepository.GetUserInfo(us.ctx, userId)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (us *UserService) List(req *common.PageInfo) ([]user.UserResp, int64, error) {
	list, count, err := us.userRepository.List(us.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (us *UserService) Delete(id uint) error {
	return us.userRepository.Delete(us.ctx, id)
}

func (us *UserService) Create(req *user.AddUserReq) error {
	// todo
	// check role exist

	return us.userRepository.Create(us.ctx, req)
}

func (us *UserService) Update(req *user.UpdateUserReq) (*user.UserResp, error) {
	// todo
	// check role exist

	update, err := us.userRepository.Update(us.ctx, req)
	if err != nil {
		return nil, err
	}
	return update, nil
}

// ModifyPasswd 修改用户密码
func (us *UserService) ModifyPasswd(req *user.ModifyPasswdReq) error {
	return us.userRepository.ModifyPasswd(us.ctx, req)
}

// SwitchActive 切换启用状态
func (us *UserService) SwitchActive(req *user.SwitchActiveReq) error {
	return us.userRepository.SwitchActive(us.ctx, req)
}
