package job

import (
	"context"
	"fmt"
	"time"


	"server/internal/global"
	pkgCron "server/internal/pkg/cron"
	"log/slog"
)

func init() {
	pkgCron.Register("clearTable", func(meta map[string]interface{}) (pkgCron.Job, error) {
		infos, ok := meta["tableInfo"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("clearTable: missing tableInfo")
		}
		var configs []ClearTableConfig
		for _, v := range infos {
			m, _ := v.(map[string]interface{})
			configs = append(configs, ClearTableConfig{
				TableName:    getString(m, "tableName"),
				CompareField: getString(m, "compareField"),
				Interval:     getString(m, "interval"),
			})
		}
		return &ClearTableJob{Configs: configs}, nil
	})
}

type ClearTableConfig struct {
	TableName    string
	CompareField string
	Interval     string
}

type ClearTableJob struct {
	Configs []ClearTableConfig
}

func (j *ClearTableJob) Name() string { return "clearTable" }

func (j *ClearTableJob) Run(ctx context.Context) error {
	for _, cfg := range j.Configs {
		duration, err := time.ParseDuration(cfg.Interval)
		if err != nil {
			return fmt.Errorf("parse duration %q: %w", cfg.Interval, err)
		}
		if duration <= 0 {
			return fmt.Errorf("duration must be > 0")
		}
		cutoff := time.Now().Add(-duration)
		result := global.TD27_DB.WithContext(ctx).
			Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", cfg.TableName, cfg.CompareField), cutoff)
		if result.Error != nil {
			return fmt.Errorf("delete from %s: %w", cfg.TableName, result.Error)
		}
		slog.Info("[CRON] clearTable", "table", cfg.TableName, "rows", result.RowsAffected)
	}
	return nil
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
