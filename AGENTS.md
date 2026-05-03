# TD27 Admin — Agent Guide

Gin + Vue3 admin dashboard. Keep this open; refer before acting.

## Entry points

- **Backend**: `server/cmd/server/main.go` — init order: Viper → slog → Gorm → RegisterTables → InitCron → Routers → RunServer
- **Frontend**: `web/src/main.ts`

## Backend

### Architecture
- Gin handlers: `server/internal/api/{module}/{feature}.go`
- Services: `server/internal/service/{module}/{feature}.go`
- Models (entity + dto + repository): `server/internal/model/{module}/{entity}.go`
- Routes: `server/internal/router/{module}/{router}.go` + `enter.go` for registration
- Init modules: `server/internal/initialize/gorm.go` (RegisterTables for auto-migration), `server/internal/initialize/router.go`, `server/internal/initialize/cron.go`

### Conventions
- Logging uses `log/slog` as the underlying implementation. Use `global.TD27_LOG.Info/Error/Debug` with key-value pairs for all logging calls (direct `slog` calls are not recommended).
- Logger config in `server/configs/config.yaml` under `logger` section (level/format/show-line). The `show-line` flag maps to `slog.HandlerOptions.AddSource`. Supports both text and JSON output formats with consistent cost duration formatting across handlers.
- Structured logging is implemented for Gorm operations and status-based HTTP access logs.
- DB auto-migration is `disabled` by default (`disable-auto-migrate: true`). Enable it or run init.sql manually.
- API response codes: `0` = success, `7` = response error, `4` = request error. Helpers in `internal/model/common/response.go`.
- JWT token passed via `x-token` header. Multi-login support with configurable device limit.
- Casbin RBAC: policies in `permissions` table, enforced via `CasbinHandler` middleware. Role hierarchy and data permission flags in config.
- Middleware chain (applied in `router.go`): GinLogger → GinRecovery → JWTAuth → CasbinHandler → OperationLog (+ DataPermissionHandler on specific routes).
- Cron jobs register via `init()` with `pkgCron.Register(name, job)`. The `Scheduler` loads enabled jobs from DB at startup.
- New permissions: use `PermissionDomainAPI`, `PermissionDomainMenu` constants. Subject format for service tokens is `token:{id}`.

### Commands (run from `server/`)
```bash
make run          # go run cmd/server/main.go
make build        # build for host OS
make test         # go test ./... -v
make lint         # golangci-lint run ./...
make fmt          # go fmt ./...
make swag         # swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```
Swagger at `http://localhost:8888/swagger/index.html`.

## Frontend

### Conventions
- Composition API with `<script setup lang="ts">` on all `.vue` files.
- Path aliases: `@/` → `src/`, `@@/` → `src/common/`.
- Vue block order: `<script>`, `<template>`, `<style scoped lang="scss">`.
- ESLint: `@antfu/eslint-config` — double quotes, no semicolons, 2-space indent, no-console allowed.
- State: Pinia with Setup store syntax, stores in `web/src/pinia/`.
- API modules in `web/src/api/{module}/`; page-local APIs go in `@/pages/{page}/apis/`.
- All page components in `web/src/pages/` (kebab-case dirs).
- Auto-imports: vue, vue-router, pinia APIs; Element Plus components.

### Commands (run from `web/`)
```bash
pnpm dev              # dev server on :8080, proxied to backend :8888
pnpm build            # vue-tsc && vite build for production
pnpm build:stage      # build for pre-release environment
pnpm preview:stage    # preview pre-release environment build
pnpm preview:prod     # preview production environment build
pnpm test             # vitest (happy-dom), test files: tests/**/*.test.{ts,js}
pnpm lint             # eslint . --fix
```

## Testing

- **Backend**: `make test` or `go test ./... -v`. Test utilities in `server/internal/testutil/` (DB helpers).
- Test coverage includes: Casbin RBAC, pkg utilities, md5 function, and core business logic.
- Casbin tests in `server/internal/service/sysManagement/casbin_test.go`.

## Full project structure (tl;dr)

```
server/                     # Go backend
├── cmd/server/main.go      # entry point
├── internal/api/           # Gin handlers (sysManagement, sysMonitor, sysTool)
├── internal/service/       # business logic
├── internal/model/         # entities, DTOs, repos
├── internal/router/        # route registration
├── internal/middleware/    # JWT, Casbin, CORS, logging
├── internal/initialize/    # app bootstrap (gorm, router, cron)
├── internal/core/          # config (Viper) + logger (slog)
├── internal/global/        # global vars (DB, Config, LOG)
├── internal/pkg/           # shared (cron, casbin, rbac, cache, jwt, async, utilities)
├── configs/config.yaml     # all runtime config
├── log/                    # Application logs
├── resource/               # Static resources (images, attachments, templates)
│   └── upload/             # File upload target directory
└── scripts/                # Shell scripts (build, deploy, maintenance)

web/                        # Vue3 frontend
├── src/main.ts             # entry point
├── src/api/                # API request modules
├── src/pages/              # page components
├── src/pinia/              # Pinia stores
├── src/router/             # Vue Router
├── src/common/             # shared (components, composables, assets)
├── src/layouts/            # layout components
└── src/http/               # Axios client

docker-compose/             # PostgreSQL + Redis + Nginx
```

## Gotchas

- DB is PostgreSQL (configured under `pgsql` section in config).
- Casbin enforcement is **skipped in dev mode** (`system.env: dev`). Set to `prod` to test permissions.
- The `slog` migration removed `go.uber.org/zap`. Logger init is in `server/internal/core/logger.go` (function `Logger()`). Config struct in `server/configs/logger.go`.
- Pre-commit hook via husky + lint-staged runs `eslint --fix` on all staged files (frontend).
- Default credentials: `admin / 123456`.

## Skills

Specialized skills are available for common tasks:
- `git-commit`: Generate clear and conventional commit messages from git diffs

Invoke skills using the `skill` tool when a task matches the skill description.
