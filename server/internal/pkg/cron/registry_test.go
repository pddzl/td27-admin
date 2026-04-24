package cron

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterAndGet(t *testing.T) {
	// Register a factory
	Register("testMethod", func(meta map[string]interface{}) (Job, error) {
		return JobFunc(func(_ context.Context) error { return nil }), nil
	})

	f, ok := Get("testMethod")
	assert.True(t, ok)
	assert.NotNil(t, f)

	_, ok = Get("nonExistent")
	assert.False(t, ok)
}

func TestBuild(t *testing.T) {
	Register("buildable", func(meta map[string]interface{}) (Job, error) {
		if meta["fail"] == true {
			return nil, errors.New("intentional fail")
		}
		return JobFunc(func(_ context.Context) error { return nil }), nil
	})

	job, err := Build("buildable", map[string]interface{}{})
	assert.NoError(t, err)
	assert.NotNil(t, job)

	_, err = Build("buildable", map[string]interface{}{"fail": true})
	assert.Error(t, err)

	_, err = Build("unknown", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not registered")
}

func TestList(t *testing.T) {
	// Clear and re-register for deterministic test
	Register("a", func(_ map[string]interface{}) (Job, error) { return nil, nil })
	Register("b", func(_ map[string]interface{}) (Job, error) { return nil, nil })

	methods := List()
	assert.GreaterOrEqual(t, len(methods), 2)
	assert.Contains(t, methods, "a")
	assert.Contains(t, methods, "b")
}
