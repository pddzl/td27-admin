# Fix: RegisterTables error handling and exit codes

## Context
`RegisterTables` in `server/internal/initialize/gorm.go` has several issues:
1. `os.Exit(0)` on failure — should be `1`
2. `"register table success"` logged unconditionally (misleading when disabled or failed)
3. No error return — caller can't detect failure
4. Fragile string-based error matching (PostgreSQL error codes)
5. main.go calls it without checking result

## Files
- `server/internal/initialize/gorm.go` — fix RegisterTables + error helpers
- `server/cmd/server/main.go` — check return value

## Changes

### 1. Make RegisterTables return an error
```go
func RegisterTables(db *gorm.DB) error {
```
This lets the caller decide how to handle failure.

### 2. Fix success log → only on actual success
Move `"register table success"` inside the success path, guarded by `DisableAutoMigrate`.

### 3. Fix exit code
Remove `os.Exit` from RegisterTables entirely (caller's responsibility). Return error instead.

### 4. Update main.go
```go
if err := initialize.RegisterTables(global.TD27_DB); err != nil {
    global.TD27_LOG.Error("auto migrate failed", "error", err)
    os.Exit(1)
}
```

### 5. (Optional) Add pgconn error code checks for robustness
Instead of `strings.Contains`, use `github.com/jackc/pgerrcode` or `pgconn` to check structured PostgreSQL error codes. This avoids locale/version fragility.

Actually, removing `strings.Contains` in favor of `pgconn` adds a dependency. A simpler improvement: keep string matching but annotate it with a comment. The better fix is to log the actual error type for debugging.

## Verification
```bash
cd server && go build ./...
```
The existing behaviour test — `go test ./internal/initialize/...` (if any).

## Edge cases
- `DisableAutoMigrate: true` → function returns nil (not an error), logs nothing about migration success
- `AutoMigrate` partially fails → error returned with details logged
- Missing `init.sql` → same error path as before, just cleanly returned

## Success criteria
- [ ] `os.Exit(0)` removed from gorm.go
- [ ] `RegisterTables` returns `error`
- [ ] `"register table success"` only printed on actual success
- [ ] `main.go` checks the error and exits with code `1` on failure
- [ ] `go build ./...` passes
