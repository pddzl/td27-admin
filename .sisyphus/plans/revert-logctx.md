# Revert: Remove logctx.FromGin + Request ID Middleware

## TL;DR
> Reverts the `logctx.FromGin` helper, `RequestID` middleware, and handler-level logctx references.
> Keeps all other logging improvements (recovery level, dead config cleanup, static attrs, status-based GinLogger, multiHandler fix).
>
> Rationale: logctx decodes JWT claims to stuff PII into log lines (security liability),
> and request ID without frontend propagation is noise.

## Context

The user rightly pointed out:
1. Frontend doesn't send `X-Request-ID` → backend-generated UUIDs are untraceable noise
2. Decoding JWT claims in every `logctx.FromGin(c)` call adds overhead + PII exposure for zero debugging value
3. The access log already has `ip`, `path`, `method`, `status` — sufficient for correlation

## What We Keep (already applied, do NOT touch)
- `server/internal/core/logger.go` — `staticAttrHandler` (service/env), `parseLevel` default `Info`, `multiHandler` error handling
- `server/configs/logger.go` — removed dead fields, added `Service`
- `server/configs/config.yaml` — cleaned up logger section
- `server/internal/middleware/log/error.go` — panic level `Debug`→`Error`, broken pipe `Debug`→`Warn`
- `server/internal/middleware/log/access.go` — status-based level (4xx→Warn, 5xx→Error) — **but currently broken** (imports removed logctx)
- `server/internal/initialize/gorm.go` — structured slog `"gorm" "sql"`

## What to Revert

### 1. Remove `server/internal/pkg/logctx/` directory (entire package)
```
rm -rf server/internal/pkg/logctx
```

### 2. Fix `server/internal/middleware/log/access.go` — remove logctx import + request_id block
Current (broken):
```go
import (
    "log/slog"
    "time"
    "github.com/gin-gonic/gin"
    // "server/internal/pkg/logctx" — MISSING, was removed
)
```
Fix: Remove the now-orphaned import, remove the request_id `c.GetString` block.

Edits:
```go
// Remove logctx import line
// Remove the block:
//   if rid := c.GetString(logctx.RequestIDKey); rid != "" {
//       attrs = append(attrs, "request_id", rid)
//   }
```

### 3. Fix `server/internal/api/sysManagement/claims_utils.go` — revert imports + calls
Revert to:
```go
import (
    "github.com/gin-gonic/gin"
    "server/internal/global"
    modelSysManagement "server/internal/model/sysManagement"
    pkgJwt "server/internal/pkg/jwt"
)
```
Replace `logctx.FromGin(c).Error(...)` → `global.TD27_LOG.Error(...)` (both occurrences)

### 4. Fix `server/internal/api/sysManagement/login.go` — revert imports + calls
Remove `"server/internal/pkg/logctx"` from imports.
Replace all `logctx.FromGin(c).Error(...)` → `global.TD27_LOG.Error(...)` (6 occurrences)

### 5. Fix `server/internal/middleware/casbin.go` — revert imports + calls
Remove `"server/internal/pkg/logctx"` from imports.
Replace `logctx.FromGin(c).Error/Warn(...)` → `global.TD27_LOG.Error/Warn(...)` (4 occurrences)

### 6. Fix `server/internal/middleware/data_permission.go` — revert imports + calls
Remove `"server/internal/pkg/logctx"` from imports.
Replace `logctx.FromGin(c).Error(...)` → `global.TD27_LOG.Error(...)` (1 occurrence)

### 7. Fix `server/internal/middleware/request_id.go` — delete file
```
rm -f server/internal/middleware/request_id.go
```

### 8. Fix `server/internal/initialize/router.go` — remove RequestID middleware
Remove `middleware.RequestID()` from the middleware chain.

### 9. Fix `server/internal/middleware/casbin.go` — fix formatting
The file has lost its indentation in the `if` branches. Restore original indentation so the `if err != nil` / `if !success` blocks inside `isServiceToken` and JWT sections are properly nested.

## Verification

```bash
cd server && go build ./...
```
Expect: clean build, no errors.

```bash
cd server && go test ./internal/pkg/... ./internal/pkg/cron/...
```
Expect: ALL PASS (pre-existing test failures in `service/` packages are unrelated).

## Edge Cases

- **access.go**: currently has a dangling import for `logctx` — the file won't compile. This is the most urgent fix.
- **casbin.go**: the indentation is slightly off (loss of nesting in the `if err` blocks) — cosmetic but the code still works. Fixing it is nice-to-have.

## Success Criteria

- [ ] `go build ./...` passes
- [ ] `server/internal/pkg/logctx` directory deleted
- [ ] `server/internal/middleware/request_id.go` deleted
- [ ] `access.go` compiles clean (no dangling import)
- [ ] `claims_utils.go`, `login.go`, `casbin.go`, `data_permission.go` all use `global.TD27_LOG` again
- [ ] `router.go` no longer calls `middleware.RequestID()`
