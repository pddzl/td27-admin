package sysMonitor

import (
	"context"

	"server/internal/global"
	"server/internal/model/sysMonitor"
)

type OperationLogService struct {
	repository sysMonitor.OperationLogRepository
	ctx        context.Context
}

func NewOperationLogService() *OperationLogService {
	return &OperationLogService{
		repository: sysMonitor.NewOperationLogRepo(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (s *OperationLogService) Create(req *sysMonitor.OperationLogModel) error {
	return s.repository.Create(s.ctx, req)
}

func (s *OperationLogService) List(req *sysMonitor.OrListReq) ([]*sysMonitor.OperationLogModel, int64, error) {
	list, i, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, i, nil
}

func (s *OperationLogService) Delete(id uint) error {
	return s.repository.Delete(s.ctx, id)
}

func (s *OperationLogService) DeleteByIds(ids []uint) (int64, error) {
	rowNums, err := s.repository.DeleteByIds(s.ctx, ids)
	if err != nil {
		return 0, err
	}
	return rowNums, nil
}
