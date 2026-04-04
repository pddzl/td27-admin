# TD27 Admin - AI Agent Guide

## Project Overview

TD27 Admin is a full-stack admin dashboard framework based on **Gin + Vue3** with a Go backend and TypeScript/Vue3 frontend.

- **Frontend**: TypeScript, Vue3, Element-Plus, Vite, Pinia, UnoCSS
- **Backend**: Golang, Gin, Gorm, MySQL, Redis, Casbin (RBAC)
- **Default Credentials**: `admin/123456`

## Project Structure

```
├── web/                     # Frontend (Vue3 + TypeScript)
│   ├── src/
│   │   ├── api/            # API request modules (organized by feature)
│   │   ├── common/         # Shared utilities, components, composables
│   │   ├── layouts/        # Page layout components
│   │   ├── pages/          # Page components (organized by feature)
│   │   ├── pinia/          # Pinia stores (state management)
│   │   ├── plugins/        # Plugin registration
│   │   └── router/         # Vue Router configuration
│   ├── types/              # TypeScript type definitions
│   └── public/             # Static assets
│
├── server/                  # Backend (Go)
│   ├── cmd/server/         # Application entry point
│   ├── configs/            # Configuration files (YAML)
│   ├── internal/           # Private application code
│   │   ├── api/            # HTTP handlers (Gin controllers)
│   │   ├── core/           # Core initialization (config, logger, DB)
│   │   ├── global/         # Global variables
│   │   ├── initialize/     # Initialization functions
│   │   ├── middleware/     # Gin middleware (JWT, CORS, Casbin, logging)
│   │   ├── model/          # Data models (entity, request, response)
│   │   ├── pkg/            # Shared utilities
│   │   ├── router/         # Route registration
│   │   └── service/        # Business logic layer
│   ├── docs/               # Swagger documentation
│   ├── log/                # Application logs
│   ├── resource/upload/    # File upload directory
│   └── scripts/            # Build/deployment scripts
│
└── docker-compose/          # Docker Compose configuration
    ├── docker-compose.yml
    ├── mysql/
    └── redis/
```

## Technology Stack

### Frontend
- **Framework**: Vue 3.5.17 (Composition API)
- **Language**: TypeScript 5.8.3
- **Build Tool**: Vite 7.0.4
- **UI Library**: Element-Plus 2.10.4
- **State Management**: Pinia 3.0.3
- **Routing**: Vue-Router 4.5.1
- **CSS Framework**: UnoCSS 66.3.3
- **HTTP Client**: Axios 1.10.0
- **Table Component**: vxe-table 4.6.25
- **Linting**: ESLint 9.31.0 with @antfu/eslint-config

### Backend
- **Language**: Go 1.25+
- **Web Framework**: Gin 1.11.0
- **ORM**: Gorm 1.31.1
- **Database**: MySQL 8.0 (via Gorm MySQL driver)
- **Cache**: Redis (go-redis/v9)
- **Auth**: JWT (golang-jwt/jwt/v4) + Casbin 2.134.0 (RBAC)
- **Logging**: Zap (go.uber.org/zap)
- **Cron Jobs**: robfig/cron/v3
- **Config**: spf13/viper
- **Docs**: Swagger (swaggo/swag)
- **Captcha**: mojocn/base64Captcha

## Build and Development Commands

### Frontend (web/)

```bash
cd web

# Install dependencies (requires pnpm 8.x)
pnpm install

# Development server (port 8080)
pnpm dev

# Build for production
pnpm build

# Build for staging
pnpm build:staging

# Preview production build
pnpm preview

# Code linting and formatting
pnpm lint

# Run tests
pnpm test
```

**Requirements:**
- Node.js 18+
- pnpm 8.x

### Backend (server/)

```bash
cd server

# Install dependencies
go mod download
# or
go generate

# Build binary
make build
# or
go build -o server cmd/server/main.go

# Run development server
make run
# or
go run cmd/server/main.go

# Build for Linux (deployment)
make build-linux

# Run tests
make test

# Code formatting
make fmt

# Run linter
make lint

# Generate Swagger docs
make swag
# or
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

**Requirements:**
- Go >= 1.25

## Configuration

### Frontend Environment Variables

Files: `.env.development`, `.env.production`, `.env.staging`

```bash
# Backend API base URL
VITE_BASE_URL = '/api'

# Public path for deployment
VITE_PUBLIC_PATH = '/'

# Frontend dev server port
VITE_CLI_PORT = 8080

# Backend address (development)
VITE_BASE_PATH = http://127.0.0.1

# Backend port
VITE_SERVER_PORT = 8888
```

### Backend Configuration

File: `server/configs/config.yaml`

Key sections:
- `jwt`: JWT signing key and expiration settings
- `zap`: Logger configuration
- `system`: Server host/port and environment
- `mysql`: Database connection settings
- `redis`: Redis connection settings
- `captcha`: Captcha image dimensions
- `cors`: CORS whitelist configuration
- `crontab`: Scheduled task configuration

## Code Style Guidelines

### Frontend

- **ESLint Config**: `@antfu/eslint-config` with custom overrides
- **Indent**: 2 spaces
- **Quotes**: Double quotes
- **Semicolons**: Disabled (no semicolons)
- **Vue Block Order**: `script`, `template`, `style`
- **Path Aliases**:
  - `@/` → `src/`
  - `@@/` → `src/common/`

### Backend

- Follow standard Go conventions (`go fmt`)
- Use `make fmt` to format code
- Package naming: lowercase, no underscores
- Internal packages under `internal/` are private

## API Conventions

### Backend API Structure

- **Base Path**: `/api` (configurable via `router.prefix`)
- **Handler Pattern**: `internal/api/{module}/{feature}.go`
- **Service Pattern**: `internal/service/{module}/{feature}.go`
- **Model Pattern**: `internal/model/{module}/{entity}.go`

### Modules

- `sysManagement`: User, Role, Menu, API, Dictionary management
- `sysMonitor`: Operation logs
- `sysTool`: File management, scheduled tasks

### Response Format

```go
// Success
{
  "code": 0,
  "data": {...},
  "msg": "success"
}

// Error
{
  "code": 7,
  "data": null,
  "msg": "error message"
}
```

## Database Setup

### Manual Deployment

1. Create MySQL database named `td27`
2. Import initialization SQL: `docker-compose/mysql/init/init.sql`

### Docker Compose

```bash
# Build and start all services
docker-compose -f docker-compose/docker-compose.yml build
docker-compose -f docker-compose/docker-compose.yml up -d

# Access the application at http://localhost:8500
```

## Testing

### Frontend

```bash
cd web
pnpm test
```

Uses Vitest with happy-dom environment.

### Backend

```bash
cd server
make test
# or
go test ./... -v
```

## Security Considerations

- **Authentication**: JWT-based with configurable expiration
- **Authorization**: Casbin RBAC with database-backed policies
- **CORS**: Configurable whitelist mode (see `configs/config.yaml`)
- **Password**: Hashed before storage
- **Captcha**: Base64 captcha for login protection
- **Operation Logging**: All user actions are logged for audit

## Deployment

### Docker

Both frontend and backend have Dockerfiles:
- `web/Dockerfile`: Nginx-based static file serving
- `server/Dockerfile`: Multi-stage Go build

### Environment Variables for Production

Ensure these are properly configured in `server/configs/config.yaml`:
- Database credentials
- JWT signing key (change from default)
- Redis password (if used)
- CORS whitelist

## Common Tasks

### Adding a New API Endpoint

1. Define model in `server/internal/model/{module}/`
2. Implement service in `server/internal/service/{module}/`
3. Create handler in `server/internal/api/{module}/`
4. Register route in `server/internal/router/{module}/`
5. Add frontend API in `web/src/api/{module}/`
6. Create frontend page in `web/src/pages/{module}/`

### Adding a New Table

1. Define Gorm model in `server/internal/model/`
2. Register auto-migration in `server/internal/initialize/gorm.go`
3. Restart server to auto-migrate

## Swagger API Documentation

Generate and access API docs:

```bash
cd server
make swag
```

Access at: `http://localhost:8888/swagger/index.html`

## Notes for AI Agents

- Frontend uses Chinese comments in configuration files
- Backend follows Go standard project layout
- All API responses follow the `common.Response` structure
- Use `global.TD27_LOG` for logging in backend
- Use Pinia stores for state management in frontend
- Path aliases `@/` and `@@/` are preconfigured
- UnoCSS provides atomic CSS utilities
