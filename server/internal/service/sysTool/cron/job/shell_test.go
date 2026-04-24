package job

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellJob(t *testing.T) {
	job := &ShellJob{Command: "echo hello"}
	assert.Equal(t, "shell", job.Name())

	err := job.Run(context.Background())
	assert.NoError(t, err)
}

func TestShellJob_InvalidCommand(t *testing.T) {
	job := &ShellJob{Command: "exit 1"}
	err := job.Run(context.Background())
	assert.Error(t, err)
}
