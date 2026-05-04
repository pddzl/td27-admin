# TD27 Admin Logging System Improvement Session
## Date: 2026-05-03

## Final Changes Implemented:
### 1. Core slog Hardening (`server/internal/core/logger.go`)
- ✅ `staticAttrHandler`: Injects `service=td27-admin` + `env=<dev/prod>` static attributes on EVERY log line
- ✅ `multiHandler`: No longer swallows file write failures; falls back to console logger if file writing fails
- ✅ Safe level default: Unrecognized log level values default to `Info` instead of `Debug` to avoid production log floods
- ✅ Removed dead `Prefix`, `EncodeLevel`, `StacktraceKey` config fields

### 2. Panic Recovery Safety (`server/internal/middleware/log/error.go`)
- ✅ Panics now logged at `Error` level (was `Debug` — no longer lost in production where log level is set to `info`)
- ✅ Broken pipe / connection reset: logged at `Warn` level (no longer `Debug`)

### 3. Access Log Improvements (`server/internal/middleware/log/access.go`)
- ✅ Status-based log levels:
  - 5xx server errors → `Error`
  - 4xx client errors → `Warn`
  - Everything else → `Info` (was all `Info` before)
- ✅ `cost` duration now consistently human-readable (`"12.345ms"`) in BOTH text and JSON formats (was raw nanoseconds in text)

### 4. Structured Gorm Logging (`server/internal/initialize/gorm.go`)
- ✅ Gorm SQL logs now use structured key-value format: `slog.Info("gorm", "sql", ...)` instead of unstructured `fmt.Sprintf`

### 5. Reverted (per your request)
- ❌ `logctx` package: removed to avoid PII exposure and unnecessary overhead
- ❌ `RequestID` middleware: removed since frontend doesn't propagate it
- ❌ Handler-level `logctx.FromGin(c)` calls: reverted back to `global.TD27_LOG`

## Final Commits on `slog` Branch:
```
b7398ce fix(log): consistent cost duration formatting across text/json handlers
cd3583d chore: track AGENTS.md — remove from gitignore
4efc975 improve(log): harden slog setup — recovery levels, static attrs, structured Gorm, status-based access logs
```

## Verification Status:
- ✅ `go build ./...` passes clean
- ✅ `go test ./internal/pkg/...` all pass
- ✅ No PII in logs
- ✅ Panics captured in production
- ✅ Status-based access logs for easier error spotting
