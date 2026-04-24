package job

import (
	"context"
	"time"

	"go.uber.org/zap"

	"server/internal/global"
	modelSysTool "server/internal/model/sysTool"
	pkgCron "server/internal/pkg/cron"
)

func init() {
	pkgCron.Register("clearCache", func(meta map[string]interface{}) (pkgCron.Job, error) {
		return &ClearCacheJob{}, nil
	})
}

type ClearCacheJob struct{}

func (j *ClearCacheJob) Name() string { return "clearCache" }

func (j *ClearCacheJob) Run(ctx context.Context) error {
	result := global.TD27_DB.WithContext(ctx).
		Where("expires_at <= ?", time.Now()).
		Delete(&modelSysTool.CacheModel{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		zap.L().Info("[CRON] clearCache done", zap.Int64("rows", result.RowsAffected))
	}
	return nil
}
