---
name: logging
description: Implement structured logging in Go using slog with consistent format and best practices
---

## Purpose
Provide a standardized logging implementation using Go slog for the TD27 Admin backend.

## Instructions

### Core Setup (server/internal/core/logger.go)
- Use `log/slog` standard library — no external dependency
- Global default logger via `slog.SetDefault(logger)`
- **Console**: stdout with configured minimum level (respects `logger.level` config)
- **Per-level log files**: output split by exact level into separate lumberjack-rotated files:
  - `log/debug.log` — only `slog.LevelDebug`
  - `log/info.log` — only `slog.LevelInfo`
  - `log/warn.log` — only `slog.LevelWarn`
  - `log/error.log` — only `slog.LevelError`
- **levelFilter** handler wrapper: routes records to the correct file by exact level match
- **Important**: `fileOpts` MUST set `Level: slog.LevelDebug` explicitly. Go 1.25's `slog.HandlerOptions{}` with nil `Level` defaults to `LevelInfo`, which silently drops debug records before `levelFilter` can route them.
- Static attributes (`service`, `env`) injected directly into `multiHandler.attrs`, added to every record before dispatch (no `staticAttrHandler` wrapper — removed to avoid routing issues with levelFilter)
- `parseLevel` defaults to `slog.LevelInfo` for unrecognized values (safe production default)
- `show-line: true` enables `AddSource` for caller file/line in debug builds

### Log Levels
- `debug` — verbose internal state (token validation, JWT operations, service token API calls)
- `info` — startup, shutdown, successful operations, normal request handled, Gorm SQL
- `warn` — 4xx client errors, broken pipe (connection reset), degraded operations
- `error` — 5xx server errors, panics, recoverable failures

### Access Logging (server/internal/middleware/log/access.go)
`GinLogger` middleware emits ONE log line per request:
- **Status-based level**: `>= 500` → `Error`, `>= 400` → `Warn`, otherwise `Info`
- **Fields**: `status`, `method`, `path`, `query`, `ip`, `user-agent`, `errors`, `cost`

```go
// 200 OK → logger.Info("request handled", ...)
// 404 → logger.Warn("client error", ...)
// 500 → logger.Error("server error", ...)
```

### Recovery Middleware (server/internal/middleware/log/error.go)
- Panics logged at `Error` level (not Debug — won't be lost in production)
- Broken pipe / connection reset → `Warn` level (not a real error)
- Optional stack trace when `system.stack: true`

### Gorm Logger (server/internal/initialize/gorm.go)
- Custom `writer.Printf` that converts Gorm output to structured slog
- Logged as: `slog.Info("gorm", "sql", formattedSQL)`
- Level: `logger.Warn`, slow threshold: 200ms, ignores `ErrRecordNotFound`
- Only active when `pgsql.log: true`

### Handler Logging
- Use `global.TD27_LOG.Error()` / `Info()` etc. with key-value pairs
- Always include `"error", err` when logging errors
- Keep message text short and descriptive

```go
global.TD27_LOG.Error("获取失败", "error", err)
global.TD27_LOG.Info("注册表成功")
```

### Config (server/configs/logger.go + config.yaml)
```yaml
logger:
  level: debug          # debug | info | warn | error
  format: text          # text | json
  service: td27-admin   # injected as static attr on every log line
  director: log         # log file directory
  show-line: true       # include caller file:line
  log-in-console: true  # also write to stdout
```

## Constraints
- **No PII in logs** — do not include user IDs, usernames, tokens in log messages
- Log keys must be English (`"error"` not `"错误"`) — message text may be Chinese
- Access log already captures `ip`, `path`, `method`, `status` — don't duplicate
- Gorm SQL logs use structured `"gorm" "sql"` attributes, not raw fmt.Sprintf
- File rotation: lumberjack config in `configs.rotateLogs` (max size, backups, age, compress)

## Key Files
| File | Purpose |
|------|---------|
| `server/internal/core/logger.go` | Logger init, `multiHandler`, `levelFilter`, `parseLevel` |
| `server/internal/middleware/log/access.go` | `GinLogger` — request access log |
| `server/internal/middleware/log/error.go` | `GinRecovery` — panic recovery |
| `server/internal/initialize/gorm.go` | Gorm `writer.Printf` → structured slog |
| `server/configs/logger.go` | `Logger` config struct |
| `server/configs/config.yaml` | Runtime logger config |
| `server/internal/global/global.go` | `TD27_LOG *slog.Logger` global var |
