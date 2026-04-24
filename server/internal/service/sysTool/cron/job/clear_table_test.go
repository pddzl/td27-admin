package job

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	"server/internal/testutil"
)

func TestClearTableJob(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db

	// Create a test table
	err := db.Exec(`CREATE TABLE test_logs (
		id INTEGER PRIMARY KEY,
		created_at DATETIME
	)`).Error
	require.NoError(t, err)

	now := time.Now()
	// Insert old and new rows
	old := now.Add(-48 * time.Hour)
	recent := now.Add(-1 * time.Hour)
	err = db.Exec("INSERT INTO test_logs (created_at) VALUES (?), (?), (?)", old, old, recent).Error
	require.NoError(t, err)

	job := &ClearTableJob{
		Configs: []ClearTableConfig{
			{TableName: "test_logs", CompareField: "created_at", Interval: "24h"},
		},
	}
	assert.Equal(t, "clearTable", job.Name())

	err = job.Run(context.Background())
	assert.NoError(t, err)

	var count int64
	db.Raw("SELECT COUNT(*) FROM test_logs").Scan(&count)
	assert.Equal(t, int64(1), count)
}

func TestClearTableJob_InvalidInterval(t *testing.T) {
	job := &ClearTableJob{
		Configs: []ClearTableConfig{
			{TableName: "t", CompareField: "c", Interval: "bad"},
		},
	}
	err := job.Run(context.Background())
	assert.Error(t, err)
}
