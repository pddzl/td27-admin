package job

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"


	pkgCron "server/internal/pkg/cron"
	"log/slog"
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
	slog.Info("[CRON] shell done", "cmd", j.Command, "output", strings.TrimSpace(string(out)))
	return nil
}
