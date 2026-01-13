package authority

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"server/internal/global"
	modelAuthority "server/internal/model/authority"
)

type ApiService struct {
	repository modelAuthority.ApiEntity
	ctx        context.Context
}

func NewApiService() *ApiService {
	return &ApiService{
		repository: modelAuthority.NewDefaultApiEntity(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (s *ApiService) Create(req *modelAuthority.ApiModel) (*modelAuthority.ApiModel, error) {
	instance, err := s.repository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	return instance, err
}

func (s *ApiService) List(req *modelAuthority.ListApiReq) ([]*modelAuthority.ApiModel, int64, error) {
	list, count, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

// GetElTree 获取所有api tree
// element-plus el-tree的数据格式
func (s *ApiService) GetElTree(roleId uint) ([]*modelAuthority.ApiTree, []string, error) {
	list, err := s.repository.GetElTree(s.ctx)
	if err != nil {
		return nil, nil, err
	}

	// 前端 el-tree default-checked-keys
	checkedKey := make([]string, 0)
	authorityId := strconv.Itoa(int(roleId))
	cData, _ := casbinService.Casbin().GetFilteredPolicy(0, authorityId)
	for _, v := range cData {
		checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", v[1], v[2]))
	}

	return list, checkedKey, nil
}

func (s *ApiService) Delete(id uint) error {
	err := s.repository.Delete(s.ctx, id)
	if err != nil {
		return err
	}

	one, err := s.repository.FindOne(s.ctx, id)
	if err != nil {
		return err
	}

	// clean corresponding casbin entity
	ok := casbinService.ClearCasbin(1, one.Path, one.Method)
	if !ok {
		return errors.New(one.Path + ":" + one.Method + "同步清理 Casbin失败")
	}

	return nil
}

func (s *ApiService) DeleteByIds(ids []uint) error {
	err := s.repository.DeleteByIds(s.ctx, ids)
	if err != nil {
		return err
	}

	apis, err := s.repository.FindByIds(s.ctx, ids)
	if err != nil {
		return err
	}

	// clean corresponding casbin entity
	for _, sysApi := range apis {
		ok := casbinService.ClearCasbin(1, sysApi.Path, sysApi.Method)
		if !ok {
			global.TD27_LOG.Error(fmt.Sprintf("%s:%s 同步清理 casbin失败", sysApi.Path, sysApi.Method))
		}
	}

	return nil
}

func (s *ApiService) Update(req *modelAuthority.ApiModel) error {
	one, err := s.repository.Update(s.ctx, req)
	if err != nil {
		return err
	}

	err = casbinService.UpdateCasbinApi(one.Path, one.Path, one.Method, one.Method)
	if err != nil {
		return fmt.Errorf("更新casbin rule err: %w", err)
	}
	return nil
}
