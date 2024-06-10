# ************************************************************
# Sequel Ace SQL dump
# Version 20067
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.28)
# Database: td27
# Generation Time: 2024-06-10 14:10:04 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table authority_api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `authority_api`;

CREATE TABLE `authority_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`),
  KEY `idx_authority_api_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `authority_api` WRITE;
/*!40000 ALTER TABLE `authority_api` DISABLE KEYS */;

INSERT INTO `authority_api` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`)
VALUES
	(1,'2023-03-10 06:24:36','2024-01-24 08:59:10',NULL,'/logReg/captcha','获取验证码（必选）','logReg','POST'),
	(2,'2023-03-08 06:36:24','2024-01-24 08:59:21',NULL,'/logReg/login','登录（必选）','logReg','POST'),
	(3,'2024-06-09 15:32:30','2024-06-09 15:32:30',NULL,'/logReg/logout','登出（必选）','logReg','POST'),
	(4,'2023-03-10 07:21:37','2023-03-10 07:21:37',NULL,'/casbin/editCasbin','编辑casbin规则','casbin','POST'),
	(5,'2023-03-08 08:56:13','2023-03-10 07:11:53',NULL,'/user/getUserInfo','获取用户信息（必选）','user','GET'),
	(6,'2023-03-08 08:56:54','2023-03-08 08:56:54',NULL,'/user/getUsers','获取所有用户','user','POST'),
	(7,'2023-03-10 06:41:32','2023-03-10 06:41:32',NULL,'/user/deleteUser','删除用户','user','POST'),
	(8,'2023-03-10 06:42:24','2023-03-10 06:42:24',NULL,'/user/addUser','添加用户','user','POST'),
	(9,'2023-03-10 06:47:18','2023-03-10 06:47:18',NULL,'/user/editUser','编辑用户','user','POST'),
	(10,'2023-03-10 06:47:59','2023-03-10 06:47:59',NULL,'/user/modifyPass','修改用户密码','user','POST'),
	(11,'2023-03-10 06:48:43','2023-03-10 06:48:43',NULL,'/user/switchActive','切换用户状态','user','POST'),
	(12,'2023-03-10 06:58:30','2023-03-10 06:58:30',NULL,'/role/getRoles','获取所有角色','role','POST'),
	(13,'2023-03-10 06:59:08','2023-03-10 06:59:08',NULL,'/role/addRole','添加角色','role','POST'),
	(14,'2023-03-10 06:59:54','2023-03-10 06:59:54',NULL,'/role/deleteRole','删除角色','role','POST'),
	(15,'2023-03-10 07:00:14','2023-03-10 07:00:53',NULL,'/role/editRole','编辑角色','role','POST'),
	(16,'2023-03-10 07:01:44','2023-03-10 07:01:44',NULL,'/role/editRoleMenu','编辑角色菜单','role','POST'),
	(17,'2023-03-10 07:14:44','2023-03-10 07:14:44',NULL,'/menu/getMenus','获取所有菜单','menu','GET'),
	(18,'2023-03-10 07:15:25','2023-03-10 07:15:25',NULL,'/menu/addMenu','添加菜单','menu','POST'),
	(19,'2023-03-10 07:15:50','2023-03-10 07:15:50',NULL,'/menu/editMenu','编辑菜单','menu','POST'),
	(20,'2023-03-10 07:16:18','2023-03-10 07:16:18',NULL,'/menu/deleteMenu','删除菜单','menu','POST'),
	(21,'2023-03-10 07:17:13','2023-03-10 07:17:13',NULL,'/menu/getElTreeMenus','获取所有菜单（el-tree结构）','menu','POST'),
	(22,'2023-03-10 07:23:21','2023-03-10 07:33:01',NULL,'/api/addApi','添加api','api','POST'),
	(23,'2023-03-10 07:24:00','2023-03-10 07:24:00',NULL,'/api/getApis','获取所有api','api','POST'),
	(24,'2023-03-10 07:24:33','2023-03-10 07:24:33',NULL,'/api/deleteApi','删除api','api','POST'),
	(25,'2023-03-10 07:26:15','2023-03-10 07:26:15',NULL,'/api/editApi','编辑api','api','POST'),
	(26,'2023-03-10 07:34:08','2023-03-10 07:35:04',NULL,'/api/getElTreeApis','获取所有api（el-tree结构）','api','POST'),
	(27,'2024-01-03 06:20:38','2024-01-03 06:20:38',NULL,'/api/deleteApiById','批量删除API','api','POST'),
	(28,'2023-07-13 02:32:16','2024-01-20 04:50:50',NULL,'/opl/getOplList','分页获取操作记录','opl','POST'),
	(29,'2023-07-13 02:33:32','2024-01-20 04:54:16',NULL,'/opl/deleteOpl','删除操作记录','opl','POST'),
	(30,'2023-07-13 06:48:47','2024-01-20 04:54:23',NULL,'/opl/deleteOplByIds','批量删除操作记录','opl','POST'),
	(31,'2023-08-27 06:05:00','2023-08-27 06:05:00',NULL,'/file/upload','文件上传','file','POST'),
	(32,'2023-08-27 06:06:43','2023-08-27 06:06:43',NULL,'/file/getFileList','分页获取文件信息','file','POST'),
	(33,'2024-01-04 03:10:15','2024-01-04 03:10:41',NULL,'/file/download','下载文件','file','GET'),
	(34,'2024-01-04 03:16:04','2024-01-04 03:16:04',NULL,'/file/delete','删除文件','file','GET'),
	(35,'2024-02-23 08:31:57','2024-02-23 08:31:57',NULL,'/cron/getCronList','分页获取cron','cron','POST'),
	(36,'2024-02-23 08:33:56','2024-02-23 08:33:56',NULL,'/cron/addCron','添加cron','cron','POST'),
	(37,'2024-02-23 08:34:25','2024-02-23 08:34:25',NULL,'/cron/deleteCron','删除cron','cron','POST'),
	(38,'2024-02-23 08:34:50','2024-02-23 08:34:50',NULL,'/cron/editCron','编辑cron','cron','POST'),
	(39,'2024-02-23 08:35:21','2024-02-23 08:35:21',NULL,'/cron/switchOpen','cron开关','cron','POST');

/*!40000 ALTER TABLE `authority_api` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table authority_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `authority_menu`;

CREATE TABLE `authority_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `pid` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `redirect` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `meta` json DEFAULT NULL,
  `sort` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`),
  KEY `idx_authority_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `authority_menu` WRITE;
/*!40000 ALTER TABLE `authority_menu` DISABLE KEYS */;

INSERT INTO `authority_menu` (`id`, `created_at`, `updated_at`, `deleted_at`, `pid`, `name`, `path`, `redirect`, `component`, `meta`, `sort`)
VALUES
	(1,NULL,'2024-01-22 07:32:30',NULL,0,'Authority','/authority','/authority/user','Layout','{\"title\": \"权限管理\", \"svgIcon\": \"lock\"}',1),
	(2,NULL,'2024-02-05 09:38:41',NULL,1,'User','user','','authority/user/index.vue','{\"title\": \"用户管理\"}',1),
	(3,NULL,'2023-06-28 08:12:06',NULL,1,'Role','role','','authority/role/index.vue','{\"title\": \"角色管理\"}',2),
	(4,NULL,'2023-06-28 08:12:16',NULL,1,'Menu','menu','','authority/menu/index.vue','{\"title\": \"菜单管理\"}',3),
	(5,'2023-03-07 01:50:48','2023-06-28 08:11:38',NULL,1,'Api','api','','authority/api/index.vue','{\"title\": \"接口管理\"}',4),
	(6,NULL,'2023-08-25 09:55:12',NULL,0,'Cenu','/cenu','/cenu/cenu1','Layout','{\"title\": \"多级菜单\", \"svgIcon\": \"menu\", \"alwaysShow\": true}',2),
	(7,NULL,'2023-06-28 08:42:39',NULL,6,'Cenu1','cenu1','/cenu/cenu1/cenu1-1','cenu/cenu1/index.vue','{\"title\": \"cenu1\"}',1),
	(8,NULL,'2023-06-28 08:42:44',NULL,7,'Cenu1-1','cenu1-1','','cenu/cenu1/cenu1-1/index.vue','{\"title\": \"cenu1-1\"}',1),
	(9,'2023-03-13 06:14:27','2023-06-28 08:43:02',NULL,7,'Cenu1-2','cenu1-2','','cenu/cenu1/cenu1-2/index.vue','{\"title\": \"cenu1-2\"}',2),
	(10,'2023-08-26 08:57:01','2023-08-26 09:02:58',NULL,0,'FileM','/fileM','/fileM/file','Layout','{\"title\": \"文件管理\", \"svgIcon\": \"file\", \"alwaysShow\": true}',3),
	(11,'2023-08-26 08:58:51','2023-08-26 08:58:51',NULL,10,'File','/fileM/file','','fileM/file/index.vue','{\"title\": \"文件\"}',1),
	(12,'2024-01-19 07:47:49','2024-03-25 06:14:36',NULL,0,'Monitor','/monitor','/monitor/operationLog','Layout','{\"title\": \"系统监控\", \"svgIcon\": \"monitor\", \"alwaysShow\": true}',5),
	(13,'2023-03-07 01:50:48','2024-01-19 07:48:52',NULL,12,'OperationLog','operationLog','','monitor/operationLog/index.vue','{\"title\": \"操作日志\"}',1),
	(14,'2024-02-05 09:56:33','2024-02-05 09:56:33',NULL,0,'SysTool','/systool','/systool/cron','Layout','{\"title\": \"系统工具\", \"svgIcon\": \"config\", \"alwaysShow\": true}',4),
	(15,'2024-02-06 10:00:00','2024-02-06 10:00:00',NULL,14,'Cron','cron','','sysTool/cron/index.vue','{\"title\": \"定时任务\"}',1);

/*!40000 ALTER TABLE `authority_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table authority_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `authority_role`;

CREATE TABLE `authority_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_name` (`role_name`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`),
  KEY `idx_authority_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `authority_role` WRITE;
/*!40000 ALTER TABLE `authority_role` DISABLE KEYS */;

INSERT INTO `authority_role` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_name`)
VALUES
	(1,NULL,'2024-03-25 07:01:12',NULL,'root');

/*!40000 ALTER TABLE `authority_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table authority_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `authority_user`;

CREATE TABLE `authority_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) DEFAULT NULL,
  `role_model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_username` (`username`),
  KEY `idx_base_user_deleted_at` (`deleted_at`),
  KEY `idx_base_user_username` (`username`),
  KEY `idx_authority_user_deleted_at` (`deleted_at`),
  KEY `idx_authority_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `authority_user` WRITE;
/*!40000 ALTER TABLE `authority_user` DISABLE KEYS */;

INSERT INTO `authority_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `phone`, `email`, `active`, `role_model_id`)
VALUES
	(1,'2023-02-20 12:51:58','2024-03-25 09:29:30',NULL,'admin','e10adc3949ba59abbe56e057f20f883e','11111111111','pddzl5@163.com',1,1);

/*!40000 ALTER TABLE `authority_user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table casbin_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	(173,'p','1','/api/addApi','POST','','',''),
	(175,'p','1','/api/deleteApi','POST','','',''),
	(178,'p','1','/api/deleteApiById','POST','','',''),
	(176,'p','1','/api/editApi','POST','','',''),
	(174,'p','1','/api/getApis','POST','','',''),
	(177,'p','1','/api/getElTreeApis','POST','','',''),
	(155,'p','1','/casbin/editCasbin','POST','','',''),
	(187,'p','1','/cron/addCron','POST','','',''),
	(188,'p','1','/cron/deleteCron','POST','','',''),
	(189,'p','1','/cron/editCron','POST','','',''),
	(186,'p','1','/cron/getCronList','POST','','',''),
	(190,'p','1','/cron/switchOpen','POST','','',''),
	(185,'p','1','/file/delete','GET','','',''),
	(184,'p','1','/file/download','GET','','',''),
	(183,'p','1','/file/getFileList','POST','','',''),
	(182,'p','1','/file/upload','POST','','',''),
	(152,'p','1','/logReg/captcha','POST','','',''),
	(153,'p','1','/logReg/login','POST','','',''),
	(154,'p','1','/logReg/logout','POST','','',''),
	(169,'p','1','/menu/addMenu','POST','','',''),
	(171,'p','1','/menu/deleteMenu','POST','','',''),
	(170,'p','1','/menu/editMenu','POST','','',''),
	(172,'p','1','/menu/getElTreeMenus','POST','','',''),
	(168,'p','1','/menu/getMenus','GET','','',''),
	(180,'p','1','/opl/deleteOpl','POST','','',''),
	(181,'p','1','/opl/deleteOplByIds','POST','','',''),
	(179,'p','1','/opl/getOplList','POST','','',''),
	(164,'p','1','/role/addRole','POST','','',''),
	(165,'p','1','/role/deleteRole','POST','','',''),
	(166,'p','1','/role/editRole','POST','','',''),
	(167,'p','1','/role/editRoleMenu','POST','','',''),
	(163,'p','1','/role/getRoles','POST','','',''),
	(159,'p','1','/user/addUser','POST','','',''),
	(158,'p','1','/user/deleteUser','POST','','',''),
	(160,'p','1','/user/editUser','POST','','',''),
	(156,'p','1','/user/getUserInfo','GET','','',''),
	(157,'p','1','/user/getUsers','POST','','',''),
	(161,'p','1','/user/modifyPass','POST','','',''),
	(162,'p','1','/user/switchActive','POST','','','');

/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table fileM_file
# ------------------------------------------------------------

DROP TABLE IF EXISTS `fileM_file`;

CREATE TABLE `fileM_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件名',
  `full_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件完整路径',
  `mime` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件类型',
  PRIMARY KEY (`id`),
  KEY `idx_fileM_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table monitor_operationLog
# ------------------------------------------------------------

DROP TABLE IF EXISTS `monitor_operationLog`;

CREATE TABLE `monitor_operationLog` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ip` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `user_agent` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `req_param` text COLLATE utf8mb4_unicode_ci COMMENT '请求Body',
  `resp_data` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
  `resp_time` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  `user_name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名称',
  PRIMARY KEY (`id`),
  KEY `idx_monitor_operationLog_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table role_menus
# ------------------------------------------------------------

DROP TABLE IF EXISTS `role_menus`;

CREATE TABLE `role_menus` (
  `menu_model_id` bigint unsigned NOT NULL,
  `role_model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`menu_model_id`,`role_model_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `role_menus` WRITE;
/*!40000 ALTER TABLE `role_menus` DISABLE KEYS */;

INSERT INTO `role_menus` (`menu_model_id`, `role_model_id`)
VALUES
	(1,1),
	(2,1),
	(3,1),
	(4,1),
	(5,1),
	(6,1),
	(7,1),
	(8,1),
	(9,1),
	(10,1),
	(11,1),
	(12,1),
	(13,1),
	(14,1),
	(15,1);

/*!40000 ALTER TABLE `role_menus` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sysTool_cron
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sysTool_cron`;

CREATE TABLE `sysTool_cron` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '任务名称',
  `method` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务方法',
  `expression` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '表达式',
  `strategy` enum('always','once') COLLATE utf8mb4_unicode_ci DEFAULT 'always' COMMENT '执行策略',
  `open` tinyint(1) DEFAULT NULL COMMENT '活跃状态',
  `extraParams` json DEFAULT NULL COMMENT '额外参数',
  `entryId` bigint DEFAULT NULL COMMENT 'cron ID',
  `comment` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_sysTool_cron_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
