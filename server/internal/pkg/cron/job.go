package cron

import (
	"context"
	"log/slog"
)

// Job 定时任务接口
type Job interface {
	Name() string
	Run(ctx context.Context) error
}

// JobFunc 允许用函数实现 Job
type JobFunc func(ctx context.Context) error

func (f JobFunc) Run(ctx context.Context) error { return f(ctx) }
func (f JobFunc) Name() string                  { return "func" }

// Runner 包装 Job，提供统一执行入口（日志、错误恢复）
type Runner struct {
	Job Job
}

func NewRunner(j Job) *Runner {
	return &Runner{Job: j}
}

func (r *Runner) Run() {
	logger := slog.With("cron", r.Job.Name())
	logger.Info("[CRON] job start")
	if err := r.Job.Run(context.Background()); err != nil {
		logger.Error("[CRON] job failed", "error", err)
	} else {
		logger.Info("[CRON] job done")
	}
}
