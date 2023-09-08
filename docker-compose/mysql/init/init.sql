# ************************************************************
# Sequel Ace SQL dump
# Version 20050
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.28)
# Database: td27
# Generation Time: 2023-08-30 11:19:15 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table casbin_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	(309,'p','1','/api/addApi','POST','','',''),
	(311,'p','1','/api/deleteApi','POST','','',''),
	(312,'p','1','/api/editApi','POST','','',''),
	(310,'p','1','/api/getApis','POST','','',''),
	(313,'p','1','/api/getElTreeApis','POST','','',''),
	(289,'p','1','/base/captcha','POST','','',''),
	(290,'p','1','/base/login','POST','','',''),
	(308,'p','1','/casbin/editCasbin','POST','','',''),
	(321,'p','1','/file/delete','GET','','',''),
	(320,'p','1','/file/download','GET','','',''),
	(319,'p','1','/file/getFileList','POST','','',''),
	(318,'p','1','/file/upload','POST','','',''),
	(314,'p','1','/jwt/joinInBlacklist','POST','','',''),
	(304,'p','1','/menu/addMenu','POST','','',''),
	(306,'p','1','/menu/deleteMenu','POST','','',''),
	(305,'p','1','/menu/editMenu','POST','','',''),
	(307,'p','1','/menu/getElTreeMenus','POST','','',''),
	(303,'p','1','/menu/getMenus','GET','','',''),
	(316,'p','1','/or/deleteOr','POST','','',''),
	(317,'p','1','/or/deleteOrByIds','POST','','',''),
	(315,'p','1','/or/getOrList','POST','','',''),
	(299,'p','1','/role/addRole','POST','','',''),
	(300,'p','1','/role/deleteRole','POST','','',''),
	(301,'p','1','/role/editRole','POST','','',''),
	(302,'p','1','/role/editRoleMenu','POST','','',''),
	(298,'p','1','/role/getRoles','POST','','',''),
	(294,'p','1','/user/addUser','POST','','',''),
	(293,'p','1','/user/deleteUser','POST','','',''),
	(295,'p','1','/user/editUser','POST','','',''),
	(291,'p','1','/user/getUserInfo','GET','','',''),
	(292,'p','1','/user/getUsers','POST','','',''),
	(296,'p','1','/user/modifyPass','POST','','',''),
	(297,'p','1','/user/switchActive','POST','','','');

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
  `file_name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件名',
  `full_path` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件完整路径',
  `mime` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件类型',
  PRIMARY KEY (`id`),
  KEY `idx_fileM_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table jwt_blacklists
# ------------------------------------------------------------

DROP TABLE IF EXISTS `jwt_blacklists`;

CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `jwt` text COLLATE utf8mb4_unicode_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
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
	(12,1),
	(13,1);

/*!40000 ALTER TABLE `role_menus` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_api`;

CREATE TABLE `sys_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api路径',
  `description` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api组',
  `method` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_api` WRITE;
/*!40000 ALTER TABLE `sys_api` DISABLE KEYS */;

INSERT INTO `sys_api` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`)
VALUES
	(1,'2023-03-10 06:24:36','2023-07-13 05:42:10',NULL,'/base/captcha','获取验证码（必选）','base','POST'),
	(2,'2023-03-08 06:36:24','2023-03-09 08:50:20',NULL,'/base/login','登录（必选）','base','POST'),
	(3,'2023-03-08 08:56:13','2023-03-10 07:11:53',NULL,'/user/getUserInfo','获取用户信息（必选）','user','GET'),
	(4,'2023-03-08 08:56:54','2023-03-08 08:56:54',NULL,'/user/getUsers','获取所有用户','user','POST'),
	(5,'2023-03-10 06:41:32','2023-03-10 06:41:32',NULL,'/user/deleteUser','删除用户','user','POST'),
	(6,'2023-03-10 06:42:24','2023-03-10 06:42:24',NULL,'/user/addUser','添加用户','user','POST'),
	(7,'2023-03-10 06:47:18','2023-03-10 06:47:18',NULL,'/user/editUser','编辑用户','user','POST'),
	(8,'2023-03-10 06:47:59','2023-03-10 06:47:59',NULL,'/user/modifyPass','修改用户密码','user','POST'),
	(9,'2023-03-10 06:48:43','2023-03-10 06:48:43',NULL,'/user/switchActive','切换用户状态','user','POST'),
	(10,'2023-03-10 06:58:30','2023-03-10 06:58:30',NULL,'/role/getRoles','获取所有角色','role','POST'),
	(11,'2023-03-10 06:59:08','2023-03-10 06:59:08',NULL,'/role/addRole','添加角色','role','POST'),
	(12,'2023-03-10 06:59:54','2023-03-10 06:59:54',NULL,'/role/deleteRole','删除角色','role','POST'),
	(13,'2023-03-10 07:00:14','2023-03-10 07:00:53',NULL,'/role/editRole','编辑角色','role','POST'),
	(14,'2023-03-10 07:01:44','2023-03-10 07:01:44',NULL,'/role/editRoleMenu','编辑角色菜单','role','POST'),
	(15,'2023-03-10 07:14:44','2023-03-10 07:14:44',NULL,'/menu/getMenus','获取所有菜单','menu','GET'),
	(16,'2023-03-10 07:15:25','2023-03-10 07:15:25',NULL,'/menu/addMenu','添加菜单','menu','POST'),
	(17,'2023-03-10 07:15:50','2023-03-10 07:15:50',NULL,'/menu/editMenu','编辑菜单','menu','POST'),
	(18,'2023-03-10 07:16:18','2023-03-10 07:16:18',NULL,'/menu/deleteMenu','删除菜单','menu','POST'),
	(19,'2023-03-10 07:17:13','2023-03-10 07:17:13',NULL,'/menu/getElTreeMenus','获取所有菜单（el-tree结构）','menu','POST'),
	(20,'2023-03-10 07:21:37','2023-03-10 07:21:37',NULL,'/casbin/editCasbin','编辑casbin规则','casbin','POST'),
	(21,'2023-03-10 07:23:21','2023-03-10 07:33:01',NULL,'/api/addApi','添加api','api','POST'),
	(22,'2023-03-10 07:24:00','2023-03-10 07:24:00',NULL,'/api/getApis','获取所有api','api','POST'),
	(23,'2023-03-10 07:24:33','2023-03-10 07:24:33',NULL,'/api/deleteApi','删除api','api','POST'),
	(24,'2023-03-10 07:26:15','2023-03-10 07:26:15',NULL,'/api/editApi','编辑api','api','POST'),
	(25,'2023-03-10 07:34:08','2023-03-10 07:35:04',NULL,'/api/getElTreeApis','获取所有api（el-tree结构）','api','POST'),
	(27,'2023-03-11 13:05:40','2023-03-11 13:05:40',NULL,'/jwt/joinInBlacklist','拉黑token','jwt','POST'),
	(28,'2023-07-13 02:32:16','2023-07-13 02:35:41',NULL,'/or/getOrList','分页获取操作记录','operationRecord','POST'),
	(29,'2023-07-13 02:33:32','2023-07-13 02:35:50',NULL,'/or/deleteOr','删除操作记录','operationRecord','POST'),
	(30,'2023-07-13 06:48:47','2023-07-13 06:48:47',NULL,'/or/deleteOrByIds','批量删除操作记录','operationRecord','POST'),
	(31,'2023-08-27 06:05:00','2023-08-27 06:05:00',NULL,'/file/upload','文件上传','file','POST'),
	(32,'2023-08-27 06:06:43','2023-08-27 06:06:43',NULL,'/file/getFileList','分页获取文件信息','file','POST'),
	(33,'2023-08-28 15:38:40','2023-08-28 15:38:40',NULL,'/file/download','下载文件','file','GET'),
	(34,'2023-08-29 15:54:07','2023-08-29 15:54:07',NULL,'/file/delete','删除文件','file','GET');

/*!40000 ALTER TABLE `sys_api` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `pid` bigint unsigned DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `redirect` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `component` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `meta` json DEFAULT NULL,
  `sort` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`id`, `created_at`, `updated_at`, `deleted_at`, `pid`, `name`, `path`, `redirect`, `component`, `meta`, `sort`)
VALUES
	(1,NULL,'2023-07-11 09:26:52',NULL,0,'Setting','/setting','/setting/user','Layout','{\"title\": \"系统管理\", \"svgIcon\": \"setting\"}',1),
	(2,NULL,'2023-06-28 08:11:56',NULL,1,'User','user','','setting/user/index.vue','{\"title\": \"用户管理\"}',1),
	(3,NULL,'2023-06-28 08:12:06',NULL,1,'Role','role','','setting/role/index.vue','{\"title\": \"角色管理\"}',2),
	(4,NULL,'2023-06-28 08:12:16',NULL,1,'Menu','menu','','setting/menu/index.vue','{\"title\": \"菜单管理\"}',3),
	(5,'2023-03-07 01:50:48','2023-06-28 08:11:38',NULL,1,'Api','api','','setting/api/index.vue','{\"title\": \"接口管理\"}',4),
	(6,'2023-03-07 01:50:48','2023-06-28 08:11:38',NULL,1,'OperationRecord','operationRecord','','setting/operationRecord/index.vue','{\"title\": \"操作记录\"}',5),
	(7,NULL,'2023-08-25 09:55:12',NULL,0,'Cenu','/cenu','/cenu/cenu1','Layout','{\"title\": \"多级菜单\", \"svgIcon\": \"menu\", \"alwaysShow\": true}',2),
	(8,NULL,'2023-06-28 08:42:39',NULL,7,'Cenu1','cenu1','/cenu/cenu1/cenu1-1','cenu/cenu1/index.vue','{\"title\": \"cenu1\"}',1),
	(9,NULL,'2023-06-28 08:42:44',NULL,8,'Cenu1-1','cenu1-1','','cenu/cenu1/cenu1-1/index.vue','{\"title\": \"cenu1-1\"}',1),
	(10,'2023-03-13 06:14:27','2023-06-28 08:43:02',NULL,8,'Cenu1-2','cenu1-2','','cenu/cenu1/cenu1-2/index.vue','{\"title\": \"cenu1-2\"}',2),
	(12,'2023-08-26 08:57:01','2023-08-26 09:02:58',NULL,0,'FileM','/fileM','/fileM/file','Layout','{\"title\": \"文件管理\", \"svgIcon\": \"file\", \"alwaysShow\": true}',3),
	(13,'2023-08-26 08:58:51','2023-08-26 08:58:51',NULL,12,'File','/fileM/file','','fileM/file/index.vue','{\"title\": \"文件\"}',1);

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_operation_record
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_operation_record`;

CREATE TABLE `sys_operation_record` (
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
  KEY `idx_sys_operation_record_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_name` (`role_name`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_name`)
VALUES
	(1,NULL,'2023-08-26 08:59:36',NULL,'root');

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
  `phone` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) DEFAULT NULL,
  `role_model_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `phone`, `email`, `active`, `role_model_id`)
VALUES
	(1,'2023-02-20 12:51:58','2023-03-10 09:59:49',NULL,'admin','e10adc3949ba59abbe56e057f20f883e','11111111111','pddzl5@163.com',1,1);

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
