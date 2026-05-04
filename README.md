<div>
  <h1>TD27 Admin</h1>
  <h4>基于Gin+Vue3前后端分离的Golang快速开发框架</h4>
  <span><a href="./README.en.md">English</a> | 中文</span>
</div>

## 平台简介

* 前端技术栈 TypeScript、Vue3、Element-Plus、Vite、Pinia
* 后端技术栈 Golang、Gin、Gorm、MySQL、Redis、Casbin

## 内置功能

- 用户管理：提供系统账号的完整 CRUD 与生命周期管理，包括角色绑定和状态切换。

- 角色管理：权限实体，用于将角色映射到菜单和 API 接口。

- 菜单管理：基于角色的动态路由与菜单生成。

- 接口管理：对后端 API 路由进行角色级的访问控制。

- 操作日志：记录用户操作与请求链路，便于审计与追踪。

- 定时任务：支持通过前端界面进行类 Cron 的任务管理。

- 文件管理：实现后台文件存储，提供上传、读取与删除接口。

- 字典管理：支持嵌套字典定义，用于统一数据映射与前端渲染。

## 运行

默认账号密码

`admin/123456`

克隆项目
`git clone https://github.com/pddzl/td27-admin.git`

### 前端

```bash
# 配置
1. 一键安装 .vscode 目录中推荐的插件
2. node 版本 22+
3. pnpm 版本 8.x

# 进入项目目录
cd web

# 安装依赖
pnpm i

# 启动服务
pnpm dev

# 预览预发布环境
pnpm preview:stage

# 预览正式环境
pnpm preview:prod

# 构建预发布环境
pnpm build:stage

# 构建正式环境
pnpm build

# 代码格式化
pnpm lint
```

### 后端

```bash
# 配置
1. Go >= 1.25

# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server cmd/server/main.go

# 运行二进制
./server
```

#### 目录结构
```shell
├── cmd                      # Main application entry points (one folder per binary)
│   └── server               # Main HTTP server entry (main.go)
│
├── configs                  # Configuration files (YAML/JSON), config templates
│
├── docs                     # API docs, Swagger files, architecture docs
│
├── internal                 # Private application code (not for import by other modules)
│   ├── api                  # Request handlers (Gin handlers / controllers)
│   ├── core                 # Core startup logic (config load, logger, DB, server setup)
│   ├── global               # Global variables (DB, Redis, Config, Logger, etc.)
│   ├── initialize           # Init functions (router setup, config init, cron init)
│   ├── middleware           # Gin middleware
│   ├── model                # Data models: entity, request, response, VO, DTO
│   ├── pkg                  # Shared utilities (tools, common helpers, not business logic)
│   ├── router               # Router groups & route registration
│   └── service              # Business logic & database operations (service layer)
│
├── log                      # Application logs
│
├── resource                 # Static resources (images, attachments, templates)
│   └── upload               # File upload target directory
│
└── scripts                  # Shell scripts (build, deploy, maintenance)

```

**<u>如果选择手动部署，需要创建数据库 `td27` 并导入初始化数据。sql文件位置：`./docker-compose/pgsql/init/init.sql` </u>**

## swagger

```bash
cd server
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

`浏览`

```bash
http://localhost:8888/swagger/index.html
```

## 一键安装

Docker Compose 版本需要大于 V1

```bash
git clone https://github.com/pddzl/td27-admin
cd td27-admin
docker compose -f docker-compose/compose.yml build
docker compose -f docker-compose/compose.yml up -d
```

浏览器打开 `http://ip:8500`

## 版本说明
TD27 Admin 版本号遵循 x.y.z

+ x 为重大架构变更
+ y 为功能新增
+ z 为 Bug 修复。

## 项目预览图

<table>
  <tr>
    <td><img src="./img/dashboard.png"/></td>
    <td><img src="./img/personal.png"/></td>
  </tr>
  <tr>
    <td><img src="./img/p1.png"/></td>
    <td><img src="./img/p2.png"/></td>
  </tr>
  <tr>
    <td><img src="./img/menu.png"/></td>
    <td><img src="./img/multi-menu.png"/></td>
  </tr>
  <tr>
    <td><img src="./img/api.png"/></td>
    <td><img src="./img/oplog.png"/></td>
  </tr>
</table>

## 致谢
+ 项目前端脚手架 [v3-admin-vite](https://github.com/un-pany/v3-admin-vite)

## 📄 License

[MIT](./LICENSE)

Copyright (c) 2022-present [pddzl](https://github.com/pddzl)
