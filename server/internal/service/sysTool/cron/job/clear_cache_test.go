package job

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	modelSysTool "server/internal/model/sysTool"
	"server/internal/testutil"
)

func TestClearCacheJob(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db

	// Auto-migrate the cache table
	err := db.AutoMigrate(&modelSysTool.CacheModel{})
	require.NoError(t, err)

	now := time.Now()

	// Seed data: 2 expired, 1 valid
	caches := []modelSysTool.CacheModel{
		{Username: "u1", Key: "k1", Value: "v1", ExpiresAt: now.Add(-1 * time.Hour)},
		{Username: "u2", Key: "k2", Value: "v2", ExpiresAt: now.Add(-2 * time.Hour)},
		{Username: "u3", Key: "k3", Value: "v3", ExpiresAt: now.Add(1 * time.Hour)},
	}
	for i := range caches {
		err := db.Create(&caches[i]).Error
		require.NoError(t, err)
	}

	job := &ClearCacheJob{}
	assert.Equal(t, "clearCache", job.Name())

	err = job.Run(context.Background())
	assert.NoError(t, err)

	var remaining int64
	db.Model(&modelSysTool.CacheModel{}).Count(&remaining)
	assert.Equal(t, int64(1), remaining)
}
