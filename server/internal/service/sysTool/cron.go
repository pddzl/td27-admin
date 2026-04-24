package sysTool

import (
	"context"

	"server/internal/global"
	"server/internal/model/common"
	modelSysTool "server/internal/model/sysTool"
	cronService "server/internal/service/sysTool/cron"
)

type CronService struct {
	repository modelSysTool.CronRepository
	scheduler  *cronService.Scheduler
	ctx        context.Context
}

func NewCronService() *CronService {
	return &CronService{
		repository: modelSysTool.NewCronEntity(global.TD27_DB),
		scheduler:  cronService.NewScheduler(),
		ctx:        context.Background(),
	}
}

func (s *CronService) List(req *common.PageInfo) ([]*modelSysTool.CronModel, int64, error) {
	return s.repository.List(s.ctx, req)
}

func (s *CronService) Create(req *modelSysTool.CronModel) (*modelSysTool.CronModel, error) {
	created, err := s.repository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}
	if req.Open {
		_ = s.scheduler.Schedule(*created)
	}
	return created, nil
}

func (s *CronService) Delete(id uint) error {
	_ = s.scheduler.Remove(id)
	return s.repository.Delete(s.ctx, id)
}

func (s *CronService) DeleteByIds(ids []uint) error {
	for _, id := range ids {
		_ = s.scheduler.Remove(id)
	}
	return s.repository.DeleteByIds(s.ctx, ids)
}

func (s *CronService) Update(req *modelSysTool.CronModel) (*modelSysTool.CronModel, error) {
	updated, err := s.repository.Update(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// Reschedule if the cron is open
	if req.Open {
		_ = s.scheduler.Schedule(*updated)
	} else {
		_ = s.scheduler.Remove(req.ID)
	}
	return updated, nil
}

func (s *CronService) SwitchOpen(id uint, open bool) (int, error) {
	_, err := s.repository.SwitchOpen(s.ctx, id, open)
	if err != nil {
		return 0, err
	}

	m, err := s.repository.FindByID(s.ctx, id)
	if err != nil {
		return 0, err
	}

	if open {
		if err := s.scheduler.Schedule(*m); err != nil {
			return 0, err
		}
		return m.EntryId, nil
	}

	_ = s.scheduler.Remove(id)
	return 0, nil
}

func (s *CronService) RunOnce(id uint) error {
	m, err := s.repository.FindByID(s.ctx, id)
	if err != nil {
		return err
	}
	return s.scheduler.Trigger(*m)
}
