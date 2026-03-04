package sysManagement

import (
	"context"

	"server/internal/global"
	"server/internal/model/sysManagement"
)

type DictDetailService struct {
	repository sysManagement.DictDetailRepository
	ctx        context.Context
}

func NewDictDetailService() *DictDetailService {
	return &DictDetailService{
		repository: sysManagement.NewDictDetailEntity(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (s *DictDetailService) List(req *sysManagement.ListDictDetailReq) ([]*sysManagement.DictDetailModel, int64, error) {
	list, count, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

func (s *DictDetailService) Flat(dictId uint) ([]*sysManagement.DictDetailModel, error) {
	flat, err := s.repository.Flat(s.ctx, dictId)
	if err != nil {
		return nil, err
	}
	return flat, nil
}

func (s *DictDetailService) Create(req *sysManagement.DictDetailModel) (*sysManagement.DictDetailModel, error) {
	create, err := s.repository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (s *DictDetailService) Delete(id uint) error {
	return s.repository.Delete(s.ctx, id)
}

func (s *DictDetailService) Update(instance *sysManagement.DictDetailModel) (*sysManagement.DictDetailModel, error) {
	update, err := s.repository.Update(s.ctx, instance)
	if err != nil {
		return nil, err
	}
	return update, nil
}
