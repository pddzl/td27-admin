package sysManagement

import (
	"context"

	"server/internal/global"
	"server/internal/model/common"
	"server/internal/model/sysManagement"
)

type DictService struct {
	repository sysManagement.DictRepository
	ctx        context.Context
}

func NewDictService() *DictService {
	return &DictService{
		repository: sysManagement.NewDictRepository(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (s *DictService) List(req *common.PageInfo) ([]*sysManagement.DictModel, int64, error) {
	list, count, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

func (s *DictService) Create(req *sysManagement.DictModel) (*sysManagement.DictModel, error) {
	create, err := s.repository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (s *DictService) Delete(id uint) error {
	return s.repository.Delete(s.ctx, id)
}

func (s *DictService) Update(req *sysManagement.DictModel) (*sysManagement.DictModel, error) {
	update, err := s.repository.Update(s.ctx, req)
	if err != nil {
		return nil, err
	}
	return update, nil
}
