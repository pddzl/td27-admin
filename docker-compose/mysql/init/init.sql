# ************************************************************
# Sequel Ace SQL dump
# Version 20046
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.28)
# Database: td27
# Generation Time: 2023-03-13 08:43:51 +0000
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
	(1,'p','1','/api/addApi','POST','','',''),
	(2,'p','1','/api/deleteApi','POST','','',''),
	(3,'p','1','/api/editApi','POST','','',''),
	(4,'p','1','/api/getApis','POST','','',''),
	(5,'p','1','/api/getElTreeApis','POST','','',''),
	(6,'p','1','/base/captcha','POST','','',''),
	(7,'p','1','/base/login','POST','','',''),
	(8,'p','1','/casbin/editCasbin','POST','','',''),
	(9,'p','1','/jwt/joinInBlacklist','POST','','',''),
	(10,'p','1','/menu/addMenu','POST','','',''),
	(11,'p','1','/menu/deleteMenu','POST','','',''),
	(12,'p','1','/menu/editMenu','POST','','',''),
	(13,'p','1','/menu/getElTreeMenus','POST','','',''),
	(14,'p','1','/menu/getMenus','GET','','',''),
	(15,'p','1','/role/addRole','POST','','',''),
	(16,'p','1','/role/deleteRole','POST','','',''),
	(17,'p','1','/role/editRole','POST','','',''),
	(18,'p','1','/role/editRoleMenu','POST','','',''),
	(19,'p','1','/role/getRoles','POST','','',''),
	(20,'p','1','/user/addUser','POST','','',''),
	(21,'p','1','/user/deleteUser','POST','','',''),
	(22,'p','1','/user/editUser','POST','','',''),
	(23,'p','1','/user/getUserInfo','GET','','',''),
	(24,'p','1','/user/getUsers','POST','','',''),
	(25,'p','1','/user/modifyPass','POST','','',''),
	(26,'p','1','/user/switchActive','POST','','','');

/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;


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

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;

INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`)
VALUES
	(1,'2023-03-11 10:11:53','2023-03-11 10:11:53',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjc4NjE1ODU1LCJuYmYiOjE2Nzg1Mjk0NTV9.-Ww9hn9RFISVSFpRidVSUvkIjsW2FZGJum710fhnMRA'),
	(2,'2023-03-13 08:19:48','2023-03-13 08:19:48',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjc4NzYyMDgzLCJuYmYiOjE2Nzg2NzU2ODN9.0xZoylhrvlMh7Sn-KTpNKci2pE3DNlfHt1upB5-_jyQ');

/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;


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
	(9,1);

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
	(1,'2023-03-10 06:24:36','2023-03-10 06:24:36',NULL,'/base/captcha','获取验证码（必选）','base','POST'),
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
	(27,'2023-03-11 13:05:40','2023-03-11 13:05:40',NULL,'/jwt/joinInBlacklist','拉黑token','jwt','POST');

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
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`id`, `created_at`, `updated_at`, `deleted_at`, `pid`, `name`, `path`, `redirect`, `component`, `meta`)
VALUES
	(1,NULL,NULL,NULL,0,'Setting','/setting','/setting/user','Layout','{\"title\": \"系统管理\", \"elIcon\": \"Setting\"}'),
	(2,NULL,NULL,NULL,1,'User','user',NULL,'setting/user/index.vue','{\"title\": \"用户管理\"}'),
	(3,NULL,NULL,NULL,1,'Role','role',NULL,'setting/role/index.vue','{\"title\": \"角色管理\"}'),
	(4,NULL,NULL,NULL,1,'Menu','menu',NULL,'setting/menu/index.vue','{\"title\": \"菜单管理\"}'),
	(5,'2023-03-07 01:50:48','2023-03-11 15:05:15',NULL,1,'Api','api','','setting/api/index.vue','{\"title\": \"接口管理\", \"keepAlive\": true}'),
	(6,NULL,NULL,NULL,0,'Cenu','/cenu','/cenu/cenu1','Layout','{\"title\": \"多级菜单\", \"svgIcon\": \"menu\"}'),
	(7,NULL,NULL,NULL,6,'Cenu1','cenu1','/cenu/cenu1/cenu1-1','cenu/cenu1/index.vue','{\"title\": \"cenu1\"}'),
	(8,NULL,'2023-03-11 14:06:08',NULL,7,'Cenu1-1','cenu1-1','','cenu/cenu1/cenu1-1/index.vue','{\"title\": \"cenu1-1\"}'),
	(9,'2023-03-13 06:14:27','2023-03-13 06:14:27',NULL,7,'Cenu1-2','cenu1-2','','cenu/cenu1/cenu1-2/index.vue','{\"title\": \"cenu1-2\"}');

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


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
	(1,NULL,'2023-03-13 07:24:04',NULL,'root');

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
	(1,'2023-02-20 12:51:58','2023-03-10 09:59:49',NULL,'admin','e10adc3949ba59abbe56e057f20f883e','11111111111','pddzl@163.com',1,1);

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
