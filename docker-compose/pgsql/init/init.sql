-- PostgreSQL init script for TD27 Admin with separate domain tables + unified permission identity

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Drop tables if exist (for clean init)
DROP TABLE IF EXISTS casbin_rule CASCADE;
DROP TABLE IF EXISTS sys_management_role_permissions CASCADE;
DROP TABLE IF EXISTS sys_management_user_roles CASCADE;
DROP TABLE IF EXISTS sys_management_permission CASCADE;
DROP TABLE IF EXISTS sys_management_api CASCADE;
DROP TABLE IF EXISTS sys_management_menu CASCADE;
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
-- Table structure for sys_management_dept
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
-- Table structure for sys_management_menu (domain table)
-- ----------------------------
CREATE TABLE sys_management_menu (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    menu_name VARCHAR(100) NOT NULL,
    icon VARCHAR(100) DEFAULT NULL,
    path VARCHAR(200) NOT NULL,
    component VARCHAR(200) DEFAULT NULL,
    redirect VARCHAR(200) DEFAULT NULL,
    parent_id BIGINT DEFAULT 0,
    sort BIGINT DEFAULT 0,
    hidden BOOLEAN DEFAULT FALSE,
    keep_alive BOOLEAN DEFAULT FALSE,
    status BOOLEAN DEFAULT TRUE,
    permission_id BIGINT DEFAULT NULL
);
CREATE INDEX idx_sys_management_menu_parent_id ON sys_management_menu(parent_id);
CREATE INDEX idx_sys_management_menu_status ON sys_management_menu(status);
CREATE INDEX idx_sys_management_menu_deleted_at ON sys_management_menu(deleted_at);

-- ----------------------------
-- Table structure for sys_management_api (domain table)
-- ----------------------------
CREATE TABLE sys_management_api (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    api_name VARCHAR(100) NOT NULL,
    path VARCHAR(200) NOT NULL,
    method VARCHAR(10) NOT NULL,
    api_group VARCHAR(50) DEFAULT NULL,
    status BOOLEAN DEFAULT TRUE,
    permission_id BIGINT DEFAULT NULL
);
CREATE INDEX idx_sys_management_api_group ON sys_management_api(api_group);
CREATE INDEX idx_sys_management_api_status ON sys_management_api(status);
CREATE INDEX idx_sys_management_api_deleted_at ON sys_management_api(deleted_at);

-- ----------------------------
-- Table structure for sys_management_permission (unified identity)
-- ----------------------------
CREATE TABLE sys_management_permission (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL, -- menu|api|button|data
    resource VARCHAR(200) NOT NULL, -- path for api, route for menu
    action VARCHAR(20) DEFAULT 'all', -- all|view|create|update|delete
    domain_id BIGINT NOT NULL, -- reference to menu.id or api.id
    status BOOLEAN DEFAULT TRUE
);
CREATE INDEX idx_sys_management_permission_type ON sys_management_permission(type);
CREATE INDEX idx_sys_management_permission_domain_id ON sys_management_permission(domain_id);
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
-- Table structure for sys_management_role_permissions
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
-- Table structure for sys_management_user_roles
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
-- Table structure for sys_tool_cache
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

-- User-Role association
INSERT INTO sys_management_user_roles (user_id, role_id) VALUES
(1, 1);

-- ============================
-- Menu Data
-- ============================
INSERT INTO sys_management_menu (id, menu_name, icon, path, component, redirect, parent_id, sort, hidden, keep_alive, status, title) VALUES
(1, '系统管理', 'lock', '/sysManagement', 'Layout', '/sysManagement/user', 0, 1, FALSE, FALSE, TRUE, '系统管理'),
(2, '用户管理', NULL, '/sysManagement/user', 'sysManagement/user/index.vue', NULL, 1, 1, FALSE, FALSE, TRUE, '用户管理'),
(3, '角色管理', NULL, '/sysManagement/role', 'sysManagement/role/index.vue', NULL, 1, 2, FALSE, FALSE, TRUE, '角色管理'),
(4, '菜单管理', NULL, '/sysManagement/menu', 'sysManagement/menu/index.vue', NULL, 1, 3, FALSE, FALSE, TRUE, '菜单管理'),
(5, '接口管理', NULL, '/sysManagement/api', 'sysManagement/api/index.vue', NULL, 1, 4, FALSE, FALSE, TRUE, '接口管理'),
(6, '字典管理', NULL, '/sysManagement/dict', 'sysManagement/dict/index.vue', NULL, 1, 5, FALSE, FALSE, TRUE, '字典管理'),
(7, '部门管理', NULL, '/sysManagement/dept', 'sysManagement/dept/index.vue', NULL, 1, 6, FALSE, FALSE, TRUE, '部门管理'),
(20, '系统工具', 'config', '/systool', 'Layout', '/systool/cron', 0, 4, FALSE, FALSE, TRUE, '系统工具'),
(21, '定时任务', NULL, '/systool/cron', 'sysTool/cron/index.vue', NULL, 20, 1, FALSE, FALSE, TRUE, '定时任务'),
(22, '文件管理', NULL, '/systool/file', 'sysTool/file/index.vue', NULL, 20, 2, FALSE, FALSE, TRUE, '文件管理'),
(40, '系统监控', 'monitor', '/sysMonitor', 'Layout', '/sysMonitor/operationLog', 0, 5, FALSE, FALSE, TRUE, '系统监控'),
(41, '操作日志', NULL, '/sysMonitor/operationLog', 'sysMonitor/operationLog/index.vue', NULL, 40, 1, FALSE, FALSE, TRUE, '操作日志'),
(100, '多级菜单', 'menu', '/cenu', 'Layout', '/cenu/cenu1', 0, 2, FALSE, FALSE, TRUE, '多级菜单'),
(101, 'cenu1', NULL, '/cenu/cenu1', 'cenu/cenu1/index.vue', '/cenu/cenu1/cenu1-1', 100, 1, FALSE, FALSE, TRUE, 'cenu1'),
(102, 'cenu1-1', NULL, '/cenu/cenu1/cenu1-1', 'cenu/cenu1/cenu1-1/index.vue', NULL, 101, 1, FALSE, FALSE, TRUE, 'cenu1-1'),
(103, 'cenu2', NULL, '/cenu/cenu2', 'cenu/cenu2/index.vue', NULL, 100, 2, FALSE, FALSE, TRUE, 'cenu2');

-- ============================
-- API Data
-- ============================
INSERT INTO sys_management_api (id, api_name, path, method, api_group, status) VALUES
(200, '获取验证码', '/captcha', 'POST', 'base', TRUE),
(201, '登录', '/login', 'POST', 'base', TRUE),
(202, '登出', '/logout', 'POST', 'base', TRUE),
(203, '获取用户信息', '/user/getUserInfo', 'GET', 'user', TRUE),
(204, '获取所有用户', '/user/list', 'POST', 'user', TRUE),
(205, '删除用户', '/user/delete', 'POST', 'user', TRUE),
(206, '添加用户', '/user/create', 'POST', 'user', TRUE),
(207, '编辑用户', '/user/update', 'POST', 'user', TRUE),
(208, '修改用户密码', '/user/modifyPasswd', 'POST', 'user', TRUE),
(209, '切换用户状态', '/user/switchActive', 'POST', 'user', TRUE),
(210, '获取所有角色', '/role/list', 'POST', 'role', TRUE),
(211, '添加角色', '/role/create', 'POST', 'role', TRUE),
(212, '删除角色', '/role/delete', 'POST', 'role', TRUE),
(213, '编辑角色', '/role/update', 'POST', 'role', TRUE),
(214, '编辑角色菜单', '/role/updateRoleMenu', 'POST', 'role', TRUE),
(215, '获取所有菜单', '/menu/list', 'GET', 'menu', TRUE),
(216, '添加菜单', '/menu/create', 'POST', 'menu', TRUE),
(217, '编辑菜单', '/menu/update', 'POST', 'menu', TRUE),
(218, '删除菜单', '/menu/delete', 'POST', 'menu', TRUE),
(219, '获取菜单树', '/menu/getElTreeMenus', 'POST', 'menu', TRUE),
(220, '获取所有API', '/api/list', 'POST', 'api', TRUE),
(221, '添加API', '/api/create', 'POST', 'api', TRUE),
(222, '编辑API', '/api/update', 'POST', 'api', TRUE),
(223, '删除API', '/api/delete', 'POST', 'api', TRUE),
(224, '获取字典列表', '/dict/list', 'POST', 'dict', TRUE),
(225, '添加字典', '/dict/create', 'POST', 'dict', TRUE),
(226, '编辑字典', '/dict/update', 'POST', 'dict', TRUE),
(227, '删除字典', '/dict/delete', 'POST', 'dict', TRUE),
(228, '获取字典详情', '/dictDetail/list', 'POST', 'dictDetail', TRUE),
(229, '添加字典详情', '/dictDetail/create', 'POST', 'dictDetail', TRUE),
(230, '编辑字典详情', '/dictDetail/update', 'POST', 'dictDetail', TRUE),
(231, '删除字典详情', '/dictDetail/delete', 'POST', 'dictDetail', TRUE),
(232, '获取文件列表', '/file/list', 'POST', 'file', TRUE),
(233, '上传文件', '/file/upload', 'POST', 'file', TRUE),
(234, '下载文件', '/file/download', 'GET', 'file', TRUE),
(235, '删除文件', '/file/delete', 'GET', 'file', TRUE),
(236, '获取定时任务', '/cron/list', 'POST', 'cron', TRUE),
(237, '添加定时任务', '/cron/create', 'POST', 'cron', TRUE),
(238, '编辑定时任务', '/cron/update', 'POST', 'cron', TRUE),
(239, '删除定时任务', '/cron/delete', 'POST', 'cron', TRUE),
(240, '获取操作日志', '/opl/list', 'POST', 'opl', TRUE),
(241, '删除操作日志', '/opl/delete', 'POST', 'opl', TRUE);

-- ============================
-- Permission Identity (links domain tables)
-- ============================
-- Menu permissions
INSERT INTO sys_management_permission (id, name, type, resource, action, domain_id, status) VALUES
(1, '系统管理', 'menu', '/sysManagement', 'view', 1, TRUE),
(2, '用户管理', 'menu', '/sysManagement/user', 'view', 2, TRUE),
(3, '角色管理', 'menu', '/sysManagement/role', 'view', 3, TRUE),
(4, '菜单管理', 'menu', '/sysManagement/menu', 'view', 4, TRUE),
(5, '接口管理', 'menu', '/sysManagement/api', 'view', 5, TRUE),
(6, '字典管理', 'menu', '/sysManagement/dict', 'view', 6, TRUE),
(7, '部门管理', 'menu', '/sysManagement/dept', 'view', 7, TRUE),
(20, '系统工具', 'menu', '/systool', 'view', 20, TRUE),
(21, '定时任务', 'menu', '/systool/cron', 'view', 21, TRUE),
(22, '文件管理', 'menu', '/systool/file', 'view', 22, TRUE),
(40, '系统监控', 'menu', '/sysMonitor', 'view', 40, TRUE),
(41, '操作日志', 'menu', '/sysMonitor/operationLog', 'view', 41, TRUE),
(100, '多级菜单', 'menu', '/cenu', 'view', 100, TRUE),
(101, 'cenu1', 'menu', '/cenu/cenu1', 'view', 101, TRUE),
(102, 'cenu1-1', 'menu', '/cenu/cenu1/cenu1-1', 'view', 102, TRUE),
(103, 'cenu2', 'menu', '/cenu/cenu2', 'view', 103, TRUE);

-- API permissions
INSERT INTO sys_management_permission (id, name, type, resource, action, domain_id, status) VALUES
(200, '获取验证码', 'api', '/captcha', 'all', 200, TRUE),
(201, '登录', 'api', '/login', 'all', 201, TRUE),
(202, '登出', 'api', '/logout', 'all', 202, TRUE),
(203, '获取用户信息', 'api', '/user/getUserInfo', 'view', 203, TRUE),
(204, '获取所有用户', 'api', '/user/list', 'view', 204, TRUE),
(205, '删除用户', 'api', '/user/delete', 'delete', 205, TRUE),
(206, '添加用户', 'api', '/user/create', 'create', 206, TRUE),
(207, '编辑用户', 'api', '/user/update', 'update', 207, TRUE),
(208, '修改用户密码', 'api', '/user/modifyPasswd', 'update', 208, TRUE),
(209, '切换用户状态', 'api', '/user/switchActive', 'update', 209, TRUE),
(210, '获取所有角色', 'api', '/role/list', 'view', 210, TRUE),
(211, '添加角色', 'api', '/role/create', 'create', 211, TRUE),
(212, '删除角色', 'api', '/role/delete', 'delete', 212, TRUE),
(213, '编辑角色', 'api', '/role/update', 'update', 213, TRUE),
(214, '编辑角色菜单', 'api', '/role/updateRoleMenu', 'update', 214, TRUE),
(215, '获取所有菜单', 'api', '/menu/list', 'view', 215, TRUE),
(216, '添加菜单', 'api', '/menu/create', 'create', 216, TRUE),
(217, '编辑菜单', 'api', '/menu/update', 'update', 217, TRUE),
(218, '删除菜单', 'api', '/menu/delete', 'delete', 218, TRUE),
(219, '获取菜单树', 'api', '/menu/getElTreeMenus', 'view', 219, TRUE),
(220, '获取所有API', 'api', '/api/list', 'view', 220, TRUE),
(221, '添加API', 'api', '/api/create', 'create', 221, TRUE),
(222, '编辑API', 'api', '/api/update', 'update', 222, TRUE),
(223, '删除API', 'api', '/api/delete', 'delete', 223, TRUE),
(224, '获取字典列表', 'api', '/dict/list', 'view', 224, TRUE),
(225, '添加字典', 'api', '/dict/create', 'create', 225, TRUE),
(226, '编辑字典', 'api', '/dict/update', 'update', 226, TRUE),
(227, '删除字典', 'api', '/dict/delete', 'delete', 227, TRUE),
(228, '获取字典详情', 'api', '/dictDetail/list', 'view', 228, TRUE),
(229, '添加字典详情', 'api', '/dictDetail/create', 'create', 229, TRUE),
(230, '编辑字典详情', 'api', '/dictDetail/update', 'update', 230, TRUE),
(231, '删除字典详情', 'api', '/dictDetail/delete', 'delete', 231, TRUE),
(232, '获取文件列表', 'api', '/file/list', 'view', 232, TRUE),
(233, '上传文件', 'api', '/file/upload', 'create', 233, TRUE),
(234, '下载文件', 'api', '/file/download', 'view', 234, TRUE),
(235, '删除文件', 'api', '/file/delete', 'delete', 235, TRUE),
(236, '获取定时任务', 'api', '/cron/list', 'view', 236, TRUE),
(237, '添加定时任务', 'api', '/cron/create', 'create', 237, TRUE),
(238, '编辑定时任务', 'api', '/cron/update', 'update', 238, TRUE),
(239, '删除定时任务', 'api', '/cron/delete', 'delete', 239, TRUE),
(240, '获取操作日志', 'api', '/opl/list', 'view', 240, TRUE),
(241, '删除操作日志', 'api', '/opl/delete', 'delete', 241, TRUE);

-- Update menu.permission_id references
UPDATE sys_management_menu SET permission_id = id;
-- Update api.permission_id references
UPDATE sys_management_api SET permission_id = id;

-- Assign all permissions to root role
INSERT INTO sys_management_role_permissions (role_id, permission_id, data_scope)
SELECT 1, id, 'all' FROM sys_management_permission;

-- Casbin rules for backward compatibility
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
