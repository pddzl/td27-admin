package cron

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	modelSysTool "server/internal/model/sysTool"
	"server/internal/testutil"

	_ "server/internal/service/sysTool/cron/job"
)

func TestScheduler_ScheduleAndRemove(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db
	db.AutoMigrate(&modelSysTool.CronModel{}, &modelSysTool.CacheModel{})

	s := NewScheduler()
	s.Start()
	defer s.Stop()

	m := modelSysTool.CronModel{
		Name:       "test-job",
		Method:     "clearCache",
		Expression: "*/1 * * * * *", // every second
		Strategy:   "always",
		Open:       true,
	}
	db.Create(&m)

	err := s.Schedule(m)
	require.NoError(t, err)

	// Give it time to register
	time.Sleep(100 * time.Millisecond)

	// Remove
	err = s.Remove(m.ID)
	assert.NoError(t, err)

	// Stop job
	err = s.StopJob(m.ID)
	assert.NoError(t, err)
}

func TestScheduler_Trigger(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db

	s := NewScheduler()
	s.Start()
	defer s.Stop()

	m := modelSysTool.CronModel{
		Name:       "trigger-test",
		Method:     "clearCache",
		Expression: "0 0 * * *",
		Strategy:   "always",
	}

	err := s.Trigger(m)
	assert.NoError(t, err)

	// Wait for goroutine
	time.Sleep(200 * time.Millisecond)
}
