package cron

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobFunc(t *testing.T) {
	t.Run("run success", func(t *testing.T) {
		called := false
		jf := JobFunc(func(ctx context.Context) error {
			called = true
			return nil
		})
		err := jf.Run(context.Background())
		require.NoError(t, err)
		require.True(t, called)
	})

	t.Run("run error", func(t *testing.T) {
		expectedErr := errors.New("job failed")
		jf := JobFunc(func(ctx context.Context) error {
			return expectedErr
		})
		err := jf.Run(context.Background())
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("name", func(t *testing.T) {
		jf := JobFunc(func(ctx context.Context) error {
			return nil
		})
		require.Equal(t, "func", jf.Name())
	})
}

type mockJob struct {
	name    string
	 runErr  error
	 runCalled bool
}

func (m *mockJob) Name() string { return m.name }
func (m *mockJob) Run(ctx context.Context) error {
	m.runCalled = true
	return m.runErr
}

func TestRunner(t *testing.T) {
	t.Run("run success", func(t *testing.T) {
		mj := &mockJob{name: "test-job"}
		r := NewRunner(mj)
		require.NotNil(t, r)
		r.Run()
		require.True(t, mj.runCalled)
	})

	t.Run("run error", func(t *testing.T) {
		mj := &mockJob{name: "fail-job", runErr: errors.New("run failed")}
		r := NewRunner(mj)
		r.Run() // should log error but not panic
		require.True(t, mj.runCalled)
	})
}
