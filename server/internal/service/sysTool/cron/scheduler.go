package cron

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"

	"server/internal/global"
	pkgCron "server/internal/pkg/cron"
	modelSysTool "server/internal/model/sysTool"
	"log/slog"
)

// Scheduler 定时任务调度器
type Scheduler struct {
	engine *cron.Cron
	mu     sync.RWMutex
	jobs   map[uint]cron.EntryID // cron model id -> entry id
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		engine: cron.New(cron.WithSeconds()),
		jobs:   make(map[uint]cron.EntryID),
	}
}

func (s *Scheduler) Start() {
	s.engine.Start()
}

func (s *Scheduler) Stop() {
	ctx := s.engine.Stop()
	<-ctx.Done()
}

// LoadFromDB 从数据库加载所有启用的任务
func (s *Scheduler) LoadFromDB() error {
	var list []modelSysTool.CronModel
	if err := global.TD27_DB.Where("open = ?", true).Find(&list).Error; err != nil {
		return err
	}
	for _, m := range list {
		if err := s.Schedule(m); err != nil {
			slog.Error("[CRON] load failed", "name", m.Name, "error", err)
		}
	}
	return nil
}

// Schedule 调度单个任务
func (s *Scheduler) Schedule(m modelSysTool.CronModel) error {
	// 先停止旧任务
	_ = s.Remove(m.ID)

	job, err := pkgCron.Build(m.Method, map[string]interface{}{
		"tableInfo": m.ExtraParams.TableInfo,
		"command":   m.ExtraParams.Command,
	})
	if err != nil {
		return fmt.Errorf("build job: %w", err)
	}

	runner := pkgCron.NewRunner(job)
	entryID, err := s.engine.AddJob(m.Expression, runner)
	if err != nil {
		return fmt.Errorf("add cron job: %w", err)
	}

	s.mu.Lock()
	s.jobs[m.ID] = entryID
	s.mu.Unlock()

	// 回写 entryId
	global.TD27_DB.Model(&m).Update("entryId", entryID)

	slog.Info("[CRON] scheduled", "name", m.Name, "expr", m.Expression)
	return nil
}

// Remove 移除任务
func (s *Scheduler) Remove(id uint) error {
	s.mu.Lock()
	entryID, ok := s.jobs[id]
	delete(s.jobs, id)
	s.mu.Unlock()

	if ok {
		s.engine.Remove(entryID)
	}
	return nil
}

// StopJob 停止任务并更新数据库状态（用于 once 策略或手动停止）
func (s *Scheduler) StopJob(id uint) error {
	_ = s.Remove(id)
	return global.TD27_DB.Model(&modelSysTool.CronModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"open": false, "entryId": 0}).Error
}

// Trigger 立即触发一次任务（不修改调度）
func (s *Scheduler) Trigger(m modelSysTool.CronModel) error {
	job, err := pkgCron.Build(m.Method, map[string]interface{}{
		"tableInfo": m.ExtraParams.TableInfo,
		"command":   m.ExtraParams.Command,
	})
	if err != nil {
		return err
	}
	go pkgCron.NewRunner(job).Run()
	return nil
}
