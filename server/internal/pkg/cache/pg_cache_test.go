package cache

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	modelSysTool "server/internal/model/sysTool"
	"server/internal/testutil"
)

func newCacheTestDB(t *testing.T) {
	t.Helper()
	db := testutil.NewTestDB(t)
	global.TD27_DB = db
	global.TD27_LOG = slog.New(slog.NewTextHandler(io.Discard, nil))
	require.NoError(t, db.AutoMigrate(&modelSysTool.CacheModel{}))
}

func TestPGCache_SetGet(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "k1", "v1", time.Minute))
	v, err := c.Get(ctx, "k1")
	require.NoError(t, err)
	assert.Equal(t, "v1", v)

	// Upsert: same key must update in place without constraint errors
	require.NoError(t, c.Set(ctx, "alice", "k1", "v2", time.Minute))
	v, err = c.Get(ctx, "k1")
	require.NoError(t, err)
	assert.Equal(t, "v2", v)

	var count int64
	require.NoError(t, global.TD27_DB.Model(&modelSysTool.CacheModel{}).Where("\"key\" = ?", "k1").Count(&count).Error)
	assert.Equal(t, int64(1), count, "upsert must not create duplicate rows")
}

func TestPGCache_ReviveAfterSoftDelete(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "revive-key", "v1", time.Minute))

	// Simulate CleanupExpired: soft delete leaves the row with deleted_at set
	require.NoError(t, global.TD27_DB.Where("\"key\" = ?", "revive-key").Delete(&modelSysTool.CacheModel{}).Error)
	_, err := c.Get(ctx, "revive-key")
	assert.Error(t, err, "soft-deleted row must be invisible")

	// Setting the same key again must revive the row, not fail the insert
	require.NoError(t, c.Set(ctx, "alice", "revive-key", "v2", time.Minute))
	v, err := c.Get(ctx, "revive-key")
	require.NoError(t, err)
	assert.Equal(t, "v2", v)

	var count int64
	require.NoError(t, global.TD27_DB.Unscoped().Model(&modelSysTool.CacheModel{}).Where("\"key\" = ?", "revive-key").Count(&count).Error)
	assert.Equal(t, int64(1), count, "revival must reuse the existing row")
}

func TestPGCache_ExpiredMiss(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "exp-key", "v", 2*time.Millisecond))
	time.Sleep(5 * time.Millisecond)

	_, err := c.Get(ctx, "exp-key")
	assert.Error(t, err)
	assert.False(t, c.Exists(ctx, "exp-key"))
}

func TestPGCache_Del(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "del-1", "v", time.Minute))
	require.NoError(t, c.Set(ctx, "alice", "del-2", "v", time.Minute))

	require.NoError(t, c.Del(ctx, "del-1"))
	_, err := c.Get(ctx, "del-1")
	assert.Error(t, err)

	v, err := c.Get(ctx, "del-2")
	require.NoError(t, err)
	assert.Equal(t, "v", v)

	// DelByUsername removes everything for the user
	require.NoError(t, c.DelByUsername(ctx, "alice"))
	_, err = c.Get(ctx, "del-2")
	assert.Error(t, err)
}

func TestPGCache_TTL(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "ttl-key", "v", time.Minute))

	ttl, err := c.TTL(ctx, "ttl-key")
	require.NoError(t, err)
	assert.Greater(t, ttl, time.Duration(0))
	assert.LessOrEqual(t, ttl, time.Minute)

	_, err = c.TTL(ctx, "missing-key")
	assert.Error(t, err)
}

func TestPGCache_ListKeysByPrefix(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "user:a:1", "v", time.Minute))
	require.NoError(t, c.Set(ctx, "alice", "user:a:2", "v", time.Minute))
	require.NoError(t, c.Set(ctx, "bob", "user:b:1", "v", time.Minute))

	keys, err := c.ListKeysByPrefix(ctx, "user:a:")
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"user:a:1", "user:a:2"}, keys)

	keys, err = c.ListKeysByPrefix(ctx, "no-match:")
	require.NoError(t, err)
	assert.Empty(t, keys)
}

func TestPGCache_CleanupExpired(t *testing.T) {
	newCacheTestDB(t)
	ctx := context.Background()
	c := NewPGCache()

	require.NoError(t, c.Set(ctx, "alice", "clean-expired", "v", time.Millisecond))
	require.NoError(t, c.Set(ctx, "alice", "clean-live", "v", time.Minute))
	time.Sleep(3 * time.Millisecond)

	require.NoError(t, c.CleanupExpired(ctx))

	_, err := c.Get(ctx, "clean-expired")
	assert.Error(t, err)
	v, err := c.Get(ctx, "clean-live")
	require.NoError(t, err)
	assert.Equal(t, "v", v)
}
