package job

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"go.uber.org/zap"

	pkgCron "server/internal/pkg/cron"
)

func init() {
	pkgCron.Register("shell", func(meta map[string]interface{}) (pkgCron.Job, error) {
		cmd, _ := meta["command"].(string)
		if strings.TrimSpace(cmd) == "" {
			return nil, fmt.Errorf("shell: empty command")
		}
		return &ShellJob{Command: cmd}, nil
	})
}

type ShellJob struct {
	Command string
}

func (j *ShellJob) Name() string { return "shell" }

func (j *ShellJob) Run(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sh", "-c", j.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("shell exec failed: %w, output: %s", err, string(out))
	}
	zap.L().Info("[CRON] shell done", zap.String("cmd", j.Command), zap.String("output", strings.TrimSpace(string(out))))
	return nil
}
