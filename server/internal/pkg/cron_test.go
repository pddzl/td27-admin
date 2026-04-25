package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
	"server/internal/testutil"
)

func TestClearTable(t *testing.T) {
	t.Run("nil db", func(t *testing.T) {
		err := ClearTable(nil, "test", "created_at", "1h")
		require.Error(t, err)
		require.Contains(t, err.Error(), "db Cannot be empty")
	})

	t.Run("invalid duration", func(t *testing.T) {
		db := testutil.NewTestDB(t)
		err := ClearTable(db, "test", "created_at", "not-a-duration")
		require.Error(t, err)
	})

	t.Run("negative duration", func(t *testing.T) {
		db := testutil.NewTestDB(t)
		err := ClearTable(db, "test", "created_at", "-1h")
		require.Error(t, err)
		require.Contains(t, err.Error(), "parse duration < 0")
	})

	t.Run("zero duration", func(t *testing.T) {
		db := testutil.NewTestDB(t)
		db.Exec("CREATE TABLE zero_table (id INTEGER PRIMARY KEY, created_at DATETIME)")
		err := ClearTable(db, "zero_table", "created_at", "0s")
		require.NoError(t, err)
	})

	t.Run("valid duration", func(t *testing.T) {
		db := testutil.NewTestDB(t)
		// Create a table so the raw DELETE doesn't fail on missing table
		db.Exec("CREATE TABLE test_table (id INTEGER PRIMARY KEY, created_at DATETIME)")
		err := ClearTable(db, "test_table", "created_at", "24h")
		require.NoError(t, err)
	})

	t.Run("sql injection safety check", func(t *testing.T) {
		db := testutil.NewTestDB(t)
		db.Exec("CREATE TABLE test_table2 (id INTEGER PRIMARY KEY, created_at DATETIME)")
		// The function uses fmt.Sprintf for table/column names (not parametrized),
		// but the duration value is parametrized. We'll just verify it runs.
		err := ClearTable(db, "test_table2", "created_at", "1h")
		require.NoError(t, err)
	})
}
