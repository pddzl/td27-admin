-- MySQL init script for TD27 Admin with unified RBAC permission model
-- All permissions (menu, api, button, data) are stored in sys_management_permission

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for sys_management_dept (for data permission)
-- ----------------------------
DROP TABLE IF EXISTS `sys_management_dept`;
CREATE TABLE `sys_management_dept` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `dept_name` varchar(100) NOT NULL COMMENT '部门名称',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '父部门ID',
  `path` varchar(500) DEFAULT '/' COMMENT '部门路径(物化路径),如:/1/2/3/',
  `sort` bigint unsigned DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_sys_management_dept_path` (`path`),
  KEY `idx_sys_management_dept_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门表';

-- ----------------------------
-- Table structure for sys_management_permission (unified permission model)
-- Stores: menu, api, button, data permissions
-- ----------------------------
DROP TABLE IF EXISTS `sys_management_permission`;
CREATE TABLE `sys_management_permission` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '权限名称',
  `type` varchar(20) NOT NULL COMMENT '权限类型: menu|api|button|data',
  `resource` varchar(200) NOT NULL COMMENT '资源标识 (path for api, route for menu)',
  `action` varchar(20) DEFAULT 'view' COMMENT '操作: view|create|update|delete|all',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父权限ID',
  `sort` bigint unsigned DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标 (for menu)',
  `component` varchar(200) DEFAULT NULL COMMENT '前端组件 (for menu)',
  `redirect` varchar(200) DEFAULT NULL COMMENT '重定向 (for menu)',
  `hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏 (for menu)',
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '缓存 (for menu)',
  `api_group` varchar(50) DEFAULT NULL COMMENT 'API分组 (for api)',
  `method` varchar(10) DEFAULT 'GET' COMMENT 'HTTP方法 (for api)',
  PRIMARY KEY (`id`),
  KEY `idx_sys_management_permission_type` (`type`),
  KEY `idx_sys_management_permission_parent_id` (`parent_id`),
  KEY `idx_sys_management_permission_resource` (`resource`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='统一权限表';

-- ----------------------------
-- Table structure for sys_management_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_management_role`;
CREATE TABLE `sys_management_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(191) DEFAULT NULL COMMENT '角色名称',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID（支持角色继承）',
  `permission_hash` varchar(64) DEFAULT NULL COMMENT '权限哈希，用于缓存失效判断',
  PRIMARY KEY (`id`),
  KEY `idx_sys_management_role_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表（支持层级）';

-- ----------------------------
-- Table structure for sys_management_role_permissions (unified)
-- ----------------------------
DROP TABLE IF EXISTS `sys_management_role_permissions`;
CREATE TABLE `sys_management_role_permissions` (
  `role_id` bigint unsigned NOT NULL,
  `permission_id` bigint unsigned NOT NULL,
  `data_scope` varchar(20) DEFAULT 'all' COMMENT '数据权限范围: all|dept|self|custom',
  `custom_sql` varchar(500) DEFAULT NULL COMMENT '自定义数据权限SQL条件',
  PRIMARY KEY (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- ----------------------------
-- Table structure for sys_management_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_management_user`;
CREATE TABLE `sys_management_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL COMMENT '用户名',
  `password` varchar(191) NOT NULL COMMENT '密码',
  `phone` varchar(191) DEFAULT NULL COMMENT '手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) DEFAULT NULL COMMENT '是否活跃',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '部门ID（用于数据权限）',
  PRIMARY KEY (`id`),
  KEY `idx_sys_management_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for sys_management_user_roles (multi-role support)
-- ----------------------------
DROP TABLE IF EXISTS `sys_management_user_roles`;
CREATE TABLE `sys_management_user_roles` (
  `user_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表（多对多）';

-- ----------------------------
-- Table structure for sys_monitor_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_monitor_operation_log`;
CREATE TABLE `sys_monitor_operation_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ip` varchar(191) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `user_agent` varchar(191) DEFAULT NULL COMMENT 'UserAgent',
  `req_param` text COMMENT '请求Body',
  `resp_data` text COMMENT '响应数据',
  `resp_time` bigint DEFAULT NULL COMMENT '响应时间',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  `user_name` varchar(191) DEFAULT NULL COMMENT '用户名称',
  PRIMARY KEY (`id`),
  KEY `idx_sys_monitor_operation_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- ----------------------------
-- Table structure for sys_tool_cache (Redis replacement)
-- ----------------------------
DROP TABLE IF EXISTS `sys_tool_cache`;
CREATE TABLE `sys_tool_cache` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `key` varchar(255) NOT NULL COMMENT '缓存键',
  `value` text COMMENT '缓存值',
  `expires_at` datetime NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`),
  KEY `idx_sys_tool_cache_expires_at` (`expires_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='缓存表（替代Redis）';

-- ----------------------------
-- Table structure for sys_tool_cron
-- ----------------------------
DROP TABLE IF EXISTS `sys_tool_cron`;
CREATE TABLE `sys_tool_cron` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '任务名称',
  `method` varchar(191) NOT NULL COMMENT '任务方法',
  `expression` varchar(191) NOT NULL COMMENT '表达式',
  `strategy` varchar(20) DEFAULT 'always' COMMENT '执行策略: always|once',
  `open` tinyint(1) DEFAULT NULL COMMENT '活跃状态',
  `extraParams` json DEFAULT NULL COMMENT '额外参数',
  `entryId` bigint DEFAULT NULL COMMENT 'cron ID',
  `comment` varchar(191) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_tool_cron_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='定时任务表';

-- ----------------------------
-- Table structure for sys_tool_file
-- ----------------------------
DROP TABLE IF EXISTS `sys_tool_file`;
CREATE TABLE `sys_tool_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `file_name` varchar(191) DEFAULT NULL COMMENT '文件名',
  `full_path` varchar(191) DEFAULT NULL COMMENT '文件完整路径',
  `mime` varchar(191) DEFAULT NULL COMMENT '文件类型',
  PRIMARY KEY (`id`),
  KEY `idx_sys_tool_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文件表';

-- ----------------------------
-- Insert default data
-- ----------------------------

-- Default department
INSERT INTO `sys_management_dept` (`id`, `dept_name`, `parent_id`, `path`, `sort`, `status`) VALUES
(1, '总公司', 0, '/1/', 1, 1);

-- Default role (root)
INSERT INTO `sys_management_role` (`id`, `role_name`, `parent_id`) VALUES
(1, 'root', NULL);

-- Default user (admin/123456)
INSERT INTO `sys_management_user` (`id`, `username`, `password`, `phone`, `email`, `active`, `dept_id`) VALUES
(1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '', '', 1, 1);

-- User-Role association (admin has root role)
INSERT INTO `sys_management_user_roles` (`user_id`, `role_id`) VALUES
(1, 1);

-- ============================
-- Unified Permission Data
-- ============================

-- Menu Permissions (type='menu')
INSERT INTO `sys_management_permission` (`id`, `name`, `type`, `resource`, `action`, `parent_id`, `sort`, `status`, `icon`, `component`, `redirect`, `hidden`, `keep_alive`) VALUES
(1, '系统管理', 'menu', '/sysManagement', 'view', 0, 1, 1, 'lock', 'Layout', '/sysManagement/user', 0, 0),
(2, '用户管理', 'menu', '/sysManagement/user', 'view', 1, 1, 1, NULL, 'sysManagement/user/index.vue', NULL, 0, 0),
(3, '角色管理', 'menu', '/sysManagement/role', 'view', 1, 2, 1, NULL, 'sysManagement/role/index.vue', NULL, 0, 0),
(4, '菜单管理', 'menu', '/sysManagement/menu', 'view', 1, 3, 1, NULL, 'sysManagement/menu/index.vue', NULL, 0, 0),
(5, '接口管理', 'menu', '/sysManagement/api', 'view', 1, 4, 1, NULL, 'sysManagement/api/index.vue', NULL, 0, 0),
(6, '字典管理', 'menu', '/sysManagement/dict', 'view', 1, 5, 1, NULL, 'sysManagement/dict/index.vue', NULL, 0, 0),
(20, '系统工具', 'menu', '/systool', 'view', 0, 4, 1, 'config', 'Layout', '/systool/cron', 0, 0),
(21, '定时任务', 'menu', '/systool/cron', 'view', 20, 1, 1, NULL, 'sysTool/cron/index.vue', NULL, 0, 0),
(22, '文件管理', 'menu', '/systool/file', 'view', 20, 2, 1, NULL, 'sysTool/file/index.vue', NULL, 0, 0),
(40, '系统监控', 'menu', '/sysMonitor', 'view', 0, 5, 1, 'monitor', 'Layout', '/sysMonitor/operationLog', 0, 0),
(41, '操作日志', 'menu', '/sysMonitor/operationLog', 'view', 40, 1, 1, NULL, 'sysMonitor/operationLog/index.vue', NULL, 0, 0),
(100, '多级菜单', 'menu', '/cenu', 'view', 0, 2, 1, 'menu', 'Layout', '/cenu/cenu1', 0, 0),
(101, 'cenu1', 'menu', '/cenu/cenu1', 'view', 100, 1, 1, NULL, 'cenu/cenu1/index.vue', '/cenu/cenu1/cenu1-1', 0, 0),
(102, 'cenu1-1', 'menu', '/cenu/cenu1/cenu1-1', 'view', 101, 1, 1, NULL, 'cenu/cenu1/cenu1-1/index.vue', NULL, 0, 0),
(103, 'cenu2', 'menu', '/cenu/cenu2', 'view', 100, 2, 1, NULL, 'cenu/cenu2/index.vue', NULL, 0, 0);

-- API Permissions (type='api')
INSERT INTO `sys_management_permission` (`id`, `name`, `type`, `resource`, `action`, `parent_id`, `sort`, `status`, `api_group`, `method`) VALUES
(200, '获取验证码', 'api', '/captcha', 'all', 0, 1, 1, 'base', 'POST'),
(201, '登录', 'api', '/login', 'all', 0, 2, 1, 'base', 'POST'),
(202, '登出', 'api', '/logout', 'all', 0, 3, 1, 'base', 'POST'),
(203, '获取用户信息', 'api', '/user/getUserInfo', 'view', 0, 4, 1, 'user', 'GET'),
(204, '获取所有用户', 'api', '/user/list', 'view', 0, 5, 1, 'user', 'POST'),
(205, '删除用户', 'api', '/user/delete', 'delete', 0, 6, 1, 'user', 'POST'),
(206, '添加用户', 'api', '/user/create', 'create', 0, 7, 1, 'user', 'POST'),
(207, '编辑用户', 'api', '/user/update', 'update', 0, 8, 1, 'user', 'POST'),
(208, '修改用户密码', 'api', '/user/modifyPasswd', 'update', 0, 9, 1, 'user', 'POST'),
(209, '切换用户状态', 'api', '/user/switchActive', 'update', 0, 10, 1, 'user', 'POST'),
(210, '获取所有角色', 'api', '/role/list', 'view', 0, 11, 1, 'role', 'POST'),
(211, '添加角色', 'api', '/role/create', 'create', 0, 12, 1, 'role', 'POST'),
(212, '删除角色', 'api', '/role/delete', 'delete', 0, 13, 1, 'role', 'POST'),
(213, '编辑角色', 'api', '/role/update', 'update', 0, 14, 1, 'role', 'POST'),
(214, '编辑角色菜单', 'api', '/role/updateRoleMenu', 'update', 0, 15, 1, 'role', 'POST'),
(215, '获取所有菜单', 'api', '/menu/list', 'view', 0, 16, 1, 'menu', 'GET'),
(216, '添加菜单', 'api', '/menu/create', 'create', 0, 17, 1, 'menu', 'POST'),
(217, '编辑菜单', 'api', '/menu/update', 'update', 0, 18, 1, 'menu', 'POST'),
(218, '删除菜单', 'api', '/menu/delete', 'delete', 0, 19, 1, 'menu', 'POST'),
(219, '获取菜单树', 'api', '/menu/getElTreeMenus', 'view', 0, 20, 1, 'menu', 'POST'),
(220, '获取所有API', 'api', '/api/list', 'view', 0, 21, 1, 'api', 'POST'),
(221, '添加API', 'api', '/api/create', 'create', 0, 22, 1, 'api', 'POST'),
(222, '编辑API', 'api', '/api/update', 'update', 0, 23, 1, 'api', 'POST'),
(223, '删除API', 'api', '/api/delete', 'delete', 0, 24, 1, 'api', 'POST'),
(224, '获取字典列表', 'api', '/dict/list', 'view', 0, 25, 1, 'dict', 'POST'),
(225, '添加字典', 'api', '/dict/create', 'create', 0, 26, 1, 'dict', 'POST'),
(226, '编辑字典', 'api', '/dict/update', 'update', 0, 27, 1, 'dict', 'POST'),
(227, '删除字典', 'api', '/dict/delete', 'delete', 0, 28, 1, 'dict', 'POST'),
(228, '获取字典详情', 'api', '/dictDetail/list', 'view', 0, 29, 1, 'dictDetail', 'POST'),
(229, '添加字典详情', 'api', '/dictDetail/create', 'create', 0, 30, 1, 'dictDetail', 'POST'),
(230, '编辑字典详情', 'api', '/dictDetail/update', 'update', 0, 31, 1, 'dictDetail', 'POST'),
(231, '删除字典详情', 'api', '/dictDetail/delete', 'delete', 0, 32, 1, 'dictDetail', 'POST'),
(232, '获取文件列表', 'api', '/file/list', 'view', 0, 33, 1, 'file', 'POST'),
(233, '上传文件', 'api', '/file/upload', 'create', 0, 34, 1, 'file', 'POST'),
(234, '下载文件', 'api', '/file/download', 'view', 0, 35, 1, 'file', 'GET'),
(235, '删除文件', 'api', '/file/delete', 'delete', 0, 36, 1, 'file', 'GET'),
(236, '获取定时任务', 'api', '/cron/list', 'view', 0, 37, 1, 'cron', 'POST'),
(237, '添加定时任务', 'api', '/cron/create', 'create', 0, 38, 1, 'cron', 'POST'),
(238, '编辑定时任务', 'api', '/cron/update', 'update', 0, 39, 1, 'cron', 'POST'),
(239, '删除定时任务', 'api', '/cron/delete', 'delete', 0, 40, 1, 'cron', 'POST'),
(240, '获取操作日志', 'api', '/opl/list', 'view', 0, 41, 1, 'opl', 'POST'),
(241, '删除操作日志', 'api', '/opl/delete', 'delete', 0, 42, 1, 'opl', 'POST');

-- Data Permissions (type='data') - for data scope control
INSERT INTO `sys_management_permission` (`id`, `name`, `type`, `resource`, `action`, `parent_id`, `sort`, `status`) VALUES
(300, '用户数据权限', 'data', 'sys_management_user', 'all', 0, 1, 1),
(301, '角色数据权限', 'data', 'sys_management_role', 'all', 0, 2, 1),
(302, '部门数据权限', 'data', 'sys_management_dept', 'all', 0, 3, 1);

-- Assign all permissions to root role
INSERT INTO `sys_management_role_permissions` (`role_id`, `permission_id`, `data_scope`) 
SELECT 1, id, 'all' FROM `sys_management_permission`;

-- Casbin rules (for backward compatibility - root role has all API permissions)
INSERT INTO `casbin_rule` (`ptype`, `v0`, `v1`, `v2`) VALUES
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

SET FOREIGN_KEY_CHECKS = 1;
