-- PostgreSQL init script for TD27 Admin with unified RBAC permission model
-- All permissions (menu, api, button, data) are stored in sys_management_permission

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Drop tables if exist (for clean init)
DROP TABLE IF EXISTS casbin_rule CASCADE;
DROP TABLE IF EXISTS sys_management_role_permissions CASCADE;
DROP TABLE IF EXISTS sys_management_user_roles CASCADE;
DROP TABLE IF EXISTS sys_management_permission CASCADE;
DROP TABLE IF EXISTS sys_management_user CASCADE;
DROP TABLE IF EXISTS sys_management_role CASCADE;
DROP TABLE IF EXISTS sys_management_dept CASCADE;
DROP TABLE IF EXISTS sys_monitor_operation_log CASCADE;
DROP TABLE IF EXISTS sys_tool_cache CASCADE;
DROP TABLE IF EXISTS sys_tool_cron CASCADE;
DROP TABLE IF EXISTS sys_tool_file CASCADE;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
CREATE TABLE casbin_rule (
    id BIGSERIAL PRIMARY KEY,
    ptype VARCHAR(100) DEFAULT NULL,
    v0 VARCHAR(100) DEFAULT NULL,
    v1 VARCHAR(100) DEFAULT NULL,
    v2 VARCHAR(100) DEFAULT NULL,
    v3 VARCHAR(100) DEFAULT NULL,
    v4 VARCHAR(100) DEFAULT NULL,
    v5 VARCHAR(100) DEFAULT NULL,
    CONSTRAINT idx_casbin_rule UNIQUE (ptype, v0, v1, v2, v3, v4, v5)
);

-- ----------------------------
-- Table structure for sys_management_dept (for data permission)
-- ----------------------------
CREATE TABLE sys_management_dept (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    dept_name VARCHAR(100) NOT NULL,
    parent_id BIGINT DEFAULT 0,
    path VARCHAR(500) DEFAULT '/',
    sort BIGINT DEFAULT 0,
    status BOOLEAN DEFAULT TRUE
);
CREATE INDEX idx_sys_management_dept_path ON sys_management_dept(path);
CREATE INDEX idx_sys_management_dept_deleted_at ON sys_management_dept(deleted_at);

-- ----------------------------
-- Table structure for sys_management_permission (unified permission model)
-- Stores: menu, api, button, data permissions
-- ----------------------------
CREATE TABLE sys_management_permission (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL, -- menu|api|button|data
    resource VARCHAR(200) NOT NULL, -- path for api, route for menu
    action VARCHAR(20) DEFAULT 'view', -- view|create|update|delete|all
    parent_id BIGINT DEFAULT NULL,
    sort BIGINT DEFAULT 0,
    status BOOLEAN DEFAULT TRUE,
    icon VARCHAR(100) DEFAULT NULL, -- for menu
    component VARCHAR(200) DEFAULT NULL, -- for menu
    redirect VARCHAR(200) DEFAULT NULL, -- for menu
    hidden BOOLEAN DEFAULT FALSE, -- for menu
    keep_alive BOOLEAN DEFAULT FALSE, -- for menu
    api_group VARCHAR(50) DEFAULT NULL, -- for api
    method VARCHAR(10) DEFAULT 'GET' -- for api
);
CREATE INDEX idx_sys_management_permission_type ON sys_management_permission(type);
CREATE INDEX idx_sys_management_permission_parent_id ON sys_management_permission(parent_id);
CREATE INDEX idx_sys_management_permission_resource ON sys_management_permission(resource);

-- ----------------------------
-- Table structure for sys_management_role
-- ----------------------------
CREATE TABLE sys_management_role (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    role_name VARCHAR(191) DEFAULT NULL,
    parent_id BIGINT DEFAULT NULL,
    permission_hash VARCHAR(64) DEFAULT NULL
);
CREATE INDEX idx_sys_management_role_parent_id ON sys_management_role(parent_id);

-- ----------------------------
-- Table structure for sys_management_role_permissions (unified)
-- ----------------------------
CREATE TABLE sys_management_role_permissions (
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    data_scope VARCHAR(20) DEFAULT 'all', -- all|dept|self|custom
    custom_sql VARCHAR(500) DEFAULT NULL,
    PRIMARY KEY (role_id, permission_id)
);

-- ----------------------------
-- Table structure for sys_management_user
-- ----------------------------
CREATE TABLE sys_management_user (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    username VARCHAR(191) DEFAULT NULL,
    password VARCHAR(191) NOT NULL,
    phone VARCHAR(191) DEFAULT NULL,
    email VARCHAR(191) DEFAULT NULL,
    active BOOLEAN DEFAULT NULL,
    dept_id BIGINT DEFAULT NULL
);
CREATE INDEX idx_sys_management_user_deleted_at ON sys_management_user(deleted_at);

-- ----------------------------
-- Table structure for sys_management_user_roles (multi-role support)
-- ----------------------------
CREATE TABLE sys_management_user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

-- ----------------------------
-- Table structure for sys_monitor_operation_log
-- ----------------------------
CREATE TABLE sys_monitor_operation_log (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    ip VARCHAR(191) DEFAULT NULL,
    method VARCHAR(191) DEFAULT NULL,
    path VARCHAR(191) DEFAULT NULL,
    status BIGINT DEFAULT NULL,
    user_agent VARCHAR(191) DEFAULT NULL,
    req_param TEXT,
    resp_data TEXT,
    resp_time BIGINT DEFAULT NULL,
    user_id BIGINT DEFAULT NULL,
    user_name VARCHAR(191) DEFAULT NULL
);
CREATE INDEX idx_sys_monitor_operation_log_deleted_at ON sys_monitor_operation_log(deleted_at);

-- ----------------------------
-- Table structure for sys_tool_cache (Redis replacement)
-- ----------------------------
CREATE TABLE sys_tool_cache (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    key VARCHAR(255) NOT NULL UNIQUE,
    value TEXT,
    expires_at TIMESTAMP NOT NULL
);
CREATE INDEX idx_sys_tool_cache_expires_at ON sys_tool_cache(expires_at);

-- ----------------------------
-- Table structure for sys_tool_cron
-- ----------------------------
CREATE TABLE sys_tool_cron (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name VARCHAR(191) DEFAULT NULL,
    method VARCHAR(191) NOT NULL,
    expression VARCHAR(191) NOT NULL,
    strategy VARCHAR(20) DEFAULT 'always',
    open BOOLEAN DEFAULT NULL,
    "extraParams" JSONB DEFAULT NULL,
    "entryId" BIGINT DEFAULT NULL,
    comment VARCHAR(191) DEFAULT NULL
);
CREATE INDEX idx_sys_tool_cron_deleted_at ON sys_tool_cron(deleted_at);

-- ----------------------------
-- Table structure for sys_tool_file
-- ----------------------------
CREATE TABLE sys_tool_file (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    file_name VARCHAR(191) DEFAULT NULL,
    full_path VARCHAR(191) DEFAULT NULL,
    mime VARCHAR(191) DEFAULT NULL
);
CREATE INDEX idx_sys_tool_file_deleted_at ON sys_tool_file(deleted_at);

-- ----------------------------
-- Insert default data
-- ----------------------------

-- Default department
INSERT INTO sys_management_dept (id, dept_name, parent_id, path, sort, status) VALUES
(1, '总公司', 0, '/1/', 1, TRUE);

-- Default role (root)
INSERT INTO sys_management_role (id, role_name, parent_id) VALUES
(1, 'root', NULL);

-- Default user (admin/123456)
INSERT INTO sys_management_user (id, username, password, phone, email, active, dept_id) VALUES
(1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '', '', TRUE, 1);

-- User-Role association (admin has root role)
INSERT INTO sys_management_user_roles (user_id, role_id) VALUES
(1, 1);

-- ============================
-- Legacy Menu Data (for backward compatibility)
-- ============================

-- ============================
-- Unified Permission Data
-- ============================

-- Menu Permissions (type='menu')
INSERT INTO sys_management_permission (id, name, type, resource, action, parent_id, sort, status, icon, component, redirect, hidden, keep_alive) VALUES
(1, '系统管理', 'menu', '/sysManagement', 'view', 0, 1, TRUE, 'lock', 'Layout', '/sysManagement/user', FALSE, FALSE),
(2, '用户管理', 'menu', '/sysManagement/user', 'view', 1, 1, TRUE, NULL, 'sysManagement/user/index.vue', NULL, FALSE, FALSE),
(3, '角色管理', 'menu', '/sysManagement/role', 'view', 1, 2, TRUE, NULL, 'sysManagement/role/index.vue', NULL, FALSE, FALSE),
(4, '菜单管理', 'menu', '/sysManagement/menu', 'view', 1, 3, TRUE, NULL, 'sysManagement/menu/index.vue', NULL, FALSE, FALSE),
(5, '接口管理', 'menu', '/sysManagement/api', 'view', 1, 4, TRUE, NULL, 'sysManagement/api/index.vue', NULL, FALSE, FALSE),
(6, '字典管理', 'menu', '/sysManagement/dict', 'view', 1, 5, TRUE, NULL, 'sysManagement/dict/index.vue', NULL, FALSE, FALSE),
(7, '部门管理', 'menu', '/sysManagement/dept', 'view', 1, 6, TRUE, NULL, 'sysManagement/dept/index.vue', NULL, FALSE, FALSE),
(20, '系统工具', 'menu', '/systool', 'view', 0, 4, TRUE, 'config', 'Layout', '/systool/cron', FALSE, FALSE),
(21, '定时任务', 'menu', '/systool/cron', 'view', 20, 1, TRUE, NULL, 'sysTool/cron/index.vue', NULL, FALSE, FALSE),
(22, '文件管理', 'menu', '/systool/file', 'view', 20, 2, TRUE, NULL, 'sysTool/file/index.vue', NULL, FALSE, FALSE),
(40, '系统监控', 'menu', '/sysMonitor', 'view', 0, 5, TRUE, 'monitor', 'Layout', '/sysMonitor/operationLog', FALSE, FALSE),
(41, '操作日志', 'menu', '/sysMonitor/operationLog', 'view', 40, 1, TRUE, NULL, 'sysMonitor/operationLog/index.vue', NULL, FALSE, FALSE),
(100, '多级菜单', 'menu', '/cenu', 'view', 0, 2, TRUE, 'menu', 'Layout', '/cenu/cenu1', FALSE, FALSE),
(101, 'cenu1', 'menu', '/cenu/cenu1', 'view', 100, 1, TRUE, NULL, 'cenu/cenu1/index.vue', '/cenu/cenu1/cenu1-1', FALSE, FALSE),
(102, 'cenu1-1', 'menu', '/cenu/cenu1/cenu1-1', 'view', 101, 1, TRUE, NULL, 'cenu/cenu1/cenu1-1/index.vue', NULL, FALSE, FALSE),
(103, 'cenu2', 'menu', '/cenu/cenu2', 'view', 100, 2, TRUE, NULL, 'cenu/cenu2/index.vue', NULL, FALSE, FALSE);

-- API Permissions (type='api')
INSERT INTO sys_management_permission (id, name, type, resource, action, parent_id, sort, status, api_group, method) VALUES
(200, '获取验证码', 'api', '/captcha', 'all', 0, 1, TRUE, 'base', 'POST'),
(201, '登录', 'api', '/login', 'all', 0, 2, TRUE, 'base', 'POST'),
(202, '登出', 'api', '/logout', 'all', 0, 3, TRUE, 'base', 'POST'),
(203, '获取用户信息', 'api', '/user/getUserInfo', 'view', 0, 4, TRUE, 'user', 'GET'),
(204, '获取所有用户', 'api', '/user/list', 'view', 0, 5, TRUE, 'user', 'POST'),
(205, '删除用户', 'api', '/user/delete', 'delete', 0, 6, TRUE, 'user', 'POST'),
(206, '添加用户', 'api', '/user/create', 'create', 0, 7, TRUE, 'user', 'POST'),
(207, '编辑用户', 'api', '/user/update', 'update', 0, 8, TRUE, 'user', 'POST'),
(208, '修改用户密码', 'api', '/user/modifyPasswd', 'update', 0, 9, TRUE, 'user', 'POST'),
(209, '切换用户状态', 'api', '/user/switchActive', 'update', 0, 10, TRUE, 'user', 'POST'),
(210, '获取所有角色', 'api', '/role/list', 'view', 0, 11, TRUE, 'role', 'POST'),
(211, '添加角色', 'api', '/role/create', 'create', 0, 12, TRUE, 'role', 'POST'),
(212, '删除角色', 'api', '/role/delete', 'delete', 0, 13, TRUE, 'role', 'POST'),
(213, '编辑角色', 'api', '/role/update', 'update', 0, 14, TRUE, 'role', 'POST'),
(214, '编辑角色菜单', 'api', '/role/updateRoleMenu', 'update', 0, 15, TRUE, 'role', 'POST'),
(215, '获取所有菜单', 'api', '/menu/list', 'view', 0, 16, TRUE, 'menu', 'GET'),
(216, '添加菜单', 'api', '/menu/create', 'create', 0, 17, TRUE, 'menu', 'POST'),
(217, '编辑菜单', 'api', '/menu/update', 'update', 0, 18, TRUE, 'menu', 'POST'),
(218, '删除菜单', 'api', '/menu/delete', 'delete', 0, 19, TRUE, 'menu', 'POST'),
(219, '获取菜单树', 'api', '/menu/getElTreeMenus', 'view', 0, 20, TRUE, 'menu', 'POST'),
(220, '获取所有API', 'api', '/api/list', 'view', 0, 21, TRUE, 'api', 'POST'),
(221, '添加API', 'api', '/api/create', 'create', 0, 22, TRUE, 'api', 'POST'),
(222, '编辑API', 'api', '/api/update', 'update', 0, 23, TRUE, 'api', 'POST'),
(223, '删除API', 'api', '/api/delete', 'delete', 0, 24, TRUE, 'api', 'POST'),
(224, '获取字典列表', 'api', '/dict/list', 'view', 0, 25, TRUE, 'dict', 'POST'),
(225, '添加字典', 'api', '/dict/create', 'create', 0, 26, TRUE, 'dict', 'POST'),
(226, '编辑字典', 'api', '/dict/update', 'update', 0, 27, TRUE, 'dict', 'POST'),
(227, '删除字典', 'api', '/dict/delete', 'delete', 0, 28, TRUE, 'dict', 'POST'),
(228, '获取字典详情', 'api', '/dictDetail/list', 'view', 0, 29, TRUE, 'dictDetail', 'POST'),
(229, '添加字典详情', 'api', '/dictDetail/create', 'create', 0, 30, TRUE, 'dictDetail', 'POST'),
(230, '编辑字典详情', 'api', '/dictDetail/update', 'update', 0, 31, TRUE, 'dictDetail', 'POST'),
(231, '删除字典详情', 'api', '/dictDetail/delete', 'delete', 0, 32, TRUE, 'dictDetail', 'POST'),
(232, '获取文件列表', 'api', '/file/list', 'view', 0, 33, TRUE, 'file', 'POST'),
(233, '上传文件', 'api', '/file/upload', 'create', 0, 34, TRUE, 'file', 'POST'),
(234, '下载文件', 'api', '/file/download', 'view', 0, 35, TRUE, 'file', 'GET'),
(235, '删除文件', 'api', '/file/delete', 'delete', 0, 36, TRUE, 'file', 'GET'),
(236, '获取定时任务', 'api', '/cron/list', 'view', 0, 37, TRUE, 'cron', 'POST'),
(237, '添加定时任务', 'api', '/cron/create', 'create', 0, 38, TRUE, 'cron', 'POST'),
(238, '编辑定时任务', 'api', '/cron/update', 'update', 0, 39, TRUE, 'cron', 'POST'),
(239, '删除定时任务', 'api', '/cron/delete', 'delete', 0, 40, TRUE, 'cron', 'POST'),
(240, '获取操作日志', 'api', '/opl/list', 'view', 0, 41, TRUE, 'opl', 'POST'),
(241, '删除操作日志', 'api', '/opl/delete', 'delete', 0, 42, TRUE, 'opl', 'POST');

-- Data Permissions (type='data') - for data scope control
INSERT INTO sys_management_permission (id, name, type, resource, action, parent_id, sort, status) VALUES
(300, '用户数据权限', 'data', 'sys_management_user', 'all', 0, 1, TRUE),
(301, '角色数据权限', 'data', 'sys_management_role', 'all', 0, 2, TRUE),
(302, '部门数据权限', 'data', 'sys_management_dept', 'all', 0, 3, TRUE);

-- Assign all permissions to root role
INSERT INTO sys_management_role_permissions (role_id, permission_id, data_scope)
SELECT 1, id, 'all' FROM sys_management_permission;

-- Casbin rules (for backward compatibility - root role has all API permissions)
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES
('p', '1', '/captcha', 'POST'),
('p', '1', '/login', 'POST'),
('p', '1', '/logout', 'POST'),
('p', '1', '/user/getUserInfo', 'GET'),
('p', '1', '/user/list', 'POST'),
('p', '1', '/user/delete', 'POST'),
('p', '1', '/user/create', 'POST'),
('p', '1', '/user/update', 'POST'),
('p', '1', '/user/modifyPasswd', 'POST'),
('p', '1', '/user/switchActive', 'POST'),
('p', '1', '/role/list', 'POST'),
('p', '1', '/role/create', 'POST'),
('p', '1', '/role/delete', 'POST'),
('p', '1', '/role/update', 'POST'),
('p', '1', '/role/updateRoleMenu', 'POST'),
('p', '1', '/menu/list', 'GET'),
('p', '1', '/menu/create', 'POST'),
('p', '1', '/menu/update', 'POST'),
('p', '1', '/menu/delete', 'POST'),
('p', '1', '/menu/getElTreeMenus', 'POST'),
('p', '1', '/api/list', 'POST'),
('p', '1', '/api/create', 'POST'),
('p', '1', '/api/update', 'POST'),
('p', '1', '/api/delete', 'POST'),
('p', '1', '/dict/list', 'POST'),
('p', '1', '/dict/create', 'POST'),
('p', '1', '/dict/update', 'POST'),
('p', '1', '/dict/delete', 'POST'),
('p', '1', '/dictDetail/list', 'POST'),
('p', '1', '/dictDetail/create', 'POST'),
('p', '1', '/dictDetail/update', 'POST'),
('p', '1', '/dictDetail/delete', 'POST'),
('p', '1', '/file/list', 'POST'),
('p', '1', '/file/upload', 'POST'),
('p', '1', '/file/download', 'GET'),
('p', '1', '/file/delete', 'GET'),
('p', '1', '/cron/list', 'POST'),
('p', '1', '/cron/create', 'POST'),
('p', '1', '/cron/update', 'POST'),
('p', '1', '/cron/delete', 'POST'),
('p', '1', '/opl/list', 'POST'),
('p', '1', '/opl/delete', 'POST');
