package sysTool

import (
	"context"

	"server/internal/global"
	"server/internal/model/common"
	"server/internal/model/sysTool"
)

type CronService struct {
	repository sysTool.CronRepository
	ctx        context.Context
}

func NewCronService() *CronService {
	return &CronService{
		repository: sysTool.NewCronRepository(global.TD27_DB),
		ctx:        context.Background(),
	}
}

// List 分页获取cron
func (s *CronService) List(req *common.PageInfo) ([]*sysTool.CronModel, int64, error) {
	list, count, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

func (s *CronService) Create(req *sysTool.CronModel) (*sysTool.CronModel, error) {
	create, err := s.repository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (s *CronService) Delete(id uint) error {
	err := s.repository.Delete(s.ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CronService) DeleteByIds(ids []uint) error {
	err := s.repository.DeleteByIds(s.ctx, ids)
	if err != nil {
		return err
	}
	return nil
}

func (s *CronService) Update(req *sysTool.CronModel) (*sysTool.CronModel, error) {
	update, err := s.repository.Update(s.ctx, req)
	if err != nil {
		return nil, err
	}
	return update, nil
}

// SwitchOpen 切换cron活跃状态
func (s *CronService) SwitchOpen(id uint, open bool) (int, error) {
	switchOpen, err := s.repository.SwitchOpen(s.ctx, id, open)
	if err != nil {
		return 0, err
	}
	return switchOpen, nil
}
