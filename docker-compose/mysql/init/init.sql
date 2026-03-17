-- MySQL dump 10.13  Distrib 9.6.0, for macos26.2 (arm64)
--
-- Host: 127.0.0.1    Database: td27
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=526 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (517,'p','1','/api/create','POST','','',''),(519,'p','1','/api/delete','POST','','',''),(522,'p','1','/api/deleteByIds','POST','','',''),(521,'p','1','/api/getElTree','POST','','',''),(518,'p','1','/api/list','POST','','',''),(520,'p','1','/api/update','POST','','',''),(509,'p','1','/captcha','POST','','',''),(478,'p','1','/casbin/update','POST','','',''),(496,'p','1','/cron/create','POST','','',''),(497,'p','1','/cron/delete','POST','','',''),(495,'p','1','/cron/list','POST','','',''),(499,'p','1','/cron/switchOpen','POST','','',''),(498,'p','1','/cron/update','POST','','',''),(501,'p','1','/dict/create','POST','','',''),(503,'p','1','/dict/delete','POST','','',''),(500,'p','1','/dict/list','POST','','',''),(502,'p','1','/dict/update','POST','','',''),(505,'p','1','/dictDetail/create','POST','','',''),(507,'p','1','/dictDetail/delete','POST','','',''),(508,'p','1','/dictDetail/flat','POST','','',''),(504,'p','1','/dictDetail/list','POST','','',''),(506,'p','1','/dictDetail/update','POST','','',''),(494,'p','1','/file/delete','GET','','',''),(493,'p','1','/file/download','GET','','',''),(492,'p','1','/file/list','POST','','',''),(491,'p','1','/file/upload','POST','','',''),(510,'p','1','/login','POST','','',''),(511,'p','1','/logout','POST','','',''),(513,'p','1','/menu/create','POST','','',''),(515,'p','1','/menu/delete','POST','','',''),(516,'p','1','/menu/getElTreeMenus','POST','','',''),(512,'p','1','/menu/list','GET','','',''),(514,'p','1','/menu/update','POST','','',''),(524,'p','1','/opl/delete','POST','','',''),(525,'p','1','/opl/deleteByIds','POST','','',''),(523,'p','1','/opl/list','POST','','',''),(487,'p','1','/role/create','POST','','',''),(488,'p','1','/role/delete','POST','','',''),(486,'p','1','/role/list','POST','','',''),(489,'p','1','/role/update','POST','','',''),(490,'p','1','/role/updateRoleMenu','POST','','',''),(482,'p','1','/user/create','POST','','',''),(481,'p','1','/user/delete','POST','','',''),(479,'p','1','/user/getUserInfo','GET','','',''),(480,'p','1','/user/list','POST','','',''),(484,'p','1','/user/modifyPasswd','POST','','',''),(485,'p','1','/user/switchActive','POST','','',''),(483,'p','1','/user/update','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_api`
--

DROP TABLE IF EXISTS `sys_management_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'POST' COMMENT '方法',
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`),
  KEY `idx_authority_api_deleted_at` (`deleted_at`),
  KEY `idx_sys_management_api_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_api`
--

LOCK TABLES `sys_management_api` WRITE;
/*!40000 ALTER TABLE `sys_management_api` DISABLE KEYS */;
INSERT INTO `sys_management_api` VALUES (1,'2023-03-10 06:24:36','2025-11-23 15:33:45',NULL,'/captcha','获取验证码（必选）','base','POST',NULL),(2,'2023-03-08 06:36:24','2024-01-24 08:59:21',NULL,'/login','登录（必选）','base','POST',NULL),(3,'2024-06-09 15:32:30','2024-06-09 15:32:30',NULL,'/logout','登出（必选）','base','POST',NULL),(4,'2023-03-10 07:21:37','2026-03-06 09:28:12',NULL,'/casbin/update','编辑casbin规则','casbin','POST',NULL),(5,'2023-03-08 08:56:13','2023-03-10 07:11:53',NULL,'/user/getUserInfo','获取用户信息（必选）','user','GET',NULL),(6,'2023-03-08 08:56:54','2025-12-04 21:26:40',NULL,'/user/list','获取所有用户','user','POST',NULL),(7,'2023-03-10 06:41:32','2025-12-04 21:26:51',NULL,'/user/delete','删除用户','user','POST',NULL),(8,'2023-03-10 06:42:24','2025-12-04 21:26:58',NULL,'/user/create','添加用户','user','POST',NULL),(9,'2023-03-10 06:47:18','2025-12-04 21:27:06',NULL,'/user/update','编辑用户','user','POST',NULL),(10,'2023-03-10 06:47:59','2023-03-10 06:47:59',NULL,'/user/modifyPasswd','修改用户密码','user','POST',NULL),(11,'2023-03-10 06:48:43','2023-03-10 06:48:43',NULL,'/user/switchActive','切换用户状态','user','POST',NULL),(12,'2023-03-10 06:58:30','2025-12-04 21:23:27',NULL,'/role/list','获取所有角色','role','POST',NULL),(13,'2023-03-10 06:59:08','2025-12-04 21:25:17',NULL,'/role/create','添加角色','role','POST',NULL),(14,'2023-03-10 06:59:54','2025-12-04 21:25:24',NULL,'/role/delete','删除角色','role','POST',NULL),(15,'2023-03-10 07:00:14','2025-12-04 21:25:33',NULL,'/role/update','编辑角色','role','POST',NULL),(16,'2023-03-10 07:01:44','2023-03-10 07:01:44',NULL,'/role/updateRoleMenu','编辑角色菜单','role','POST',NULL),(17,'2023-03-10 07:14:44','2025-12-04 21:17:10',NULL,'/menu/list','获取所有菜单','menu','GET',NULL),(18,'2023-03-10 07:15:25','2025-12-04 21:17:02',NULL,'/menu/create','添加菜单','menu','POST',NULL),(19,'2023-03-10 07:15:50','2025-12-04 21:16:51',NULL,'/menu/update','编辑菜单','menu','POST',NULL),(20,'2023-03-10 07:16:18','2025-12-04 21:14:39',NULL,'/menu/delete','删除菜单','menu','POST',NULL),(21,'2023-03-10 07:17:13','2023-03-10 07:17:13',NULL,'/menu/getElTreeMenus','获取所有菜单（el-tree结构）','menu','POST',NULL),(22,'2023-03-10 07:23:21','2025-12-04 17:50:26',NULL,'/api/create','Create','api','POST',NULL),(23,'2023-03-10 07:24:00','2025-12-04 21:21:02',NULL,'/api/list','List','api','POST',NULL),(24,'2023-03-10 07:24:33','2025-12-04 17:50:58',NULL,'/api/delete','Delete','api','POST',NULL),(25,'2023-03-10 07:26:15','2025-12-04 17:51:09',NULL,'/api/update','Update','api','POST',NULL),(26,'2023-03-10 07:34:08','2025-12-04 21:32:14',NULL,'/api/getElTree','获取所有api（el-tree结构）','api','POST',NULL),(27,'2024-01-03 06:20:38','2024-01-03 06:20:38',NULL,'/api/deleteByIds','批量删除API','api','POST',NULL),(28,'2023-07-13 02:32:16','2024-01-20 04:50:50',NULL,'/opl/list','分页获取操作记录','opl','POST',NULL),(29,'2023-07-13 02:33:32','2024-01-20 04:54:16',NULL,'/opl/delete','删除操作记录','opl','POST',NULL),(30,'2023-07-13 06:48:47','2024-01-20 04:54:23',NULL,'/opl/deleteByIds','批量删除操作记录','opl','POST',NULL),(31,'2023-08-27 06:05:00','2023-08-27 06:05:00',NULL,'/file/upload','文件上传','file','POST',NULL),(32,'2023-08-27 06:06:43','2023-08-27 06:06:43',NULL,'/file/list','分页获取文件信息','file','POST',NULL),(33,'2024-01-04 03:10:15','2024-01-04 03:10:41',NULL,'/file/download','下载文件','file','GET',NULL),(34,'2024-01-04 03:16:04','2024-01-04 03:16:04',NULL,'/file/delete','删除文件','file','GET',NULL),(35,'2024-02-23 08:31:57','2024-02-23 08:31:57',NULL,'/cron/list','分页获取cron','cron','POST',NULL),(36,'2024-02-23 08:33:56','2024-02-23 08:33:56',NULL,'/cron/create','添加cron','cron','POST',NULL),(37,'2024-02-23 08:34:25','2024-02-23 08:34:25',NULL,'/cron/delete','删除cron','cron','POST',NULL),(38,'2024-02-23 08:34:50','2024-02-23 08:34:50',NULL,'/cron/update','编辑cron','cron','POST',NULL),(39,'2024-02-23 08:35:21','2024-02-23 08:35:21',NULL,'/cron/switchOpen','cron开关','cron','POST',NULL),(40,'2025-09-12 16:08:40','2026-03-17 10:01:34',NULL,'/dict/list','List','dict','POST',NULL),(41,'2025-09-12 16:09:21','2025-09-12 16:09:21',NULL,'/dict/create','Add','dict','POST',NULL),(42,'2025-09-12 16:09:46','2025-09-12 16:09:46',NULL,'/dict/update','Edit','dict','POST',NULL),(43,'2025-09-12 16:10:13','2025-09-12 16:10:13',NULL,'/dict/delete','Delete','dict','POST',NULL),(44,'2025-09-12 16:11:17','2025-09-12 16:11:17',NULL,'/dictDetail/list','Get','dictDetail','POST',NULL),(45,'2025-09-12 16:11:58','2025-09-12 16:11:58',NULL,'/dictDetail/create','Add','dictDetail','POST',NULL),(46,'2025-09-12 16:12:36','2025-09-12 16:12:36',NULL,'/dictDetail/update','Edit','dictDetail','POST',NULL),(47,'2025-09-12 16:13:15','2025-09-12 16:14:13',NULL,'/dictDetail/delete','Delete','dictDetail','POST',NULL),(48,'2025-09-29 15:53:05','2025-09-29 15:53:05',NULL,'/dictDetail/flat','Get Flat','dictDetail','POST',NULL);
/*!40000 ALTER TABLE `sys_management_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_dict`
--

DROP TABLE IF EXISTS `sys_management_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `cn_name` varchar(191) DEFAULT NULL,
  `en_name` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_management_dict_cn_name` (`cn_name`),
  UNIQUE KEY `uni_sys_management_dict_en_name` (`en_name`),
  KEY `idx_sys_management_dict_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_dict`
--

LOCK TABLES `sys_management_dict` WRITE;
/*!40000 ALTER TABLE `sys_management_dict` DISABLE KEYS */;
INSERT INTO `sys_management_dict` VALUES (1,NULL,'2026-03-17 15:30:27',NULL,'主机状态','HostStatus');
/*!40000 ALTER TABLE `sys_management_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_dictDetail`
--

DROP TABLE IF EXISTS `sys_management_dictDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_dictDetail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `label` varchar(191) DEFAULT NULL,
  `value` varchar(191) DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  `dict_id` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL,
  `description` varchar(191) DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sysSet_dictDetail_deleted_at` (`deleted_at`),
  KEY `idx_sys_management_dictDetail_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_dictDetail`
--

LOCK TABLES `sys_management_dictDetail` WRITE;
/*!40000 ALTER TABLE `sys_management_dictDetail` DISABLE KEYS */;
INSERT INTO `sys_management_dictDetail` VALUES (1,'2025-09-13 08:51:09','2025-09-13 08:51:09',NULL,'运行中','running',0,1,NULL,NULL,NULL),(2,'2025-09-13 08:51:36','2025-09-13 08:51:36',NULL,'不可达','unreached',1,1,NULL,NULL,NULL),(4,'2025-09-28 11:19:51','2025-09-29 06:33:03',NULL,'t1','t1',0,1,2,'test',NULL),(10,'2025-09-29 03:01:18','2025-09-29 06:26:57',NULL,'t10','t10',2,1,1,NULL,NULL);
/*!40000 ALTER TABLE `sys_management_dictDetail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_menu`
--

DROP TABLE IF EXISTS `sys_management_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_menu` (
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
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_management_menu_path` (`path`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`),
  KEY `idx_authority_menu_deleted_at` (`deleted_at`),
  KEY `idx_sys_management_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_menu`
--

LOCK TABLES `sys_management_menu` WRITE;
/*!40000 ALTER TABLE `sys_management_menu` DISABLE KEYS */;
INSERT INTO `sys_management_menu` VALUES (1,'2026-03-05 11:03:52','2025-12-04 21:33:00',NULL,0,'SysManagement','/sysManagement','/sysManagement/user','Layout','{\"title\": \"系统管理\", \"svgIcon\": \"lock\"}',1,NULL),(2,'2026-03-05 11:03:50','2024-02-05 09:38:41',NULL,1,'User','user','','sysManagement/user/index.vue','{\"title\": \"用户管理\"}',1,NULL),(3,'2026-03-05 11:03:48','2023-06-28 08:12:06',NULL,1,'Role','role','','sysManagement/role/index.vue','{\"title\": \"角色管理\"}',2,NULL),(4,'2026-03-05 11:03:45','2023-06-28 08:12:16',NULL,1,'Menu','menu','','sysManagement/menu/index.vue','{\"title\": \"菜单管理\"}',3,NULL),(5,'2023-03-07 01:50:48','2023-06-28 08:11:38',NULL,1,'Api','api','','sysManagement/api/index.vue','{\"title\": \"接口管理\"}',4,NULL),(6,'2024-02-06 10:00:00','2025-09-29 06:18:23',NULL,1,'Dict','dict','','sysManagement/dict/index.vue','{\"title\": \"字典管理\"}',1,NULL),(20,'2024-02-05 09:56:33','2024-02-05 09:56:33',NULL,0,'SysTool','/systool','/systool/cron','Layout','{\"title\": \"系统工具\", \"svgIcon\": \"config\", \"alwaysShow\": true}',4,NULL),(21,'2024-02-06 10:00:00','2024-02-06 10:00:00',NULL,20,'Cron','cron','','sysTool/cron/index.vue','{\"title\": \"定时任务\"}',1,NULL),(22,'2023-08-26 08:58:51','2023-08-26 08:58:51',NULL,20,'File','file','','sysTool/file/index.vue','{\"title\": \"文件管理\"}',1,NULL),(40,'2024-01-19 07:47:49','2024-03-25 06:14:36',NULL,0,'SysMonitor','/sysMonitor','/sysMonitor/operationLog','Layout','{\"title\": \"系统监控\", \"svgIcon\": \"monitor\", \"alwaysShow\": true}',5,NULL),(41,'2023-03-07 01:50:48','2024-01-19 07:48:52',NULL,40,'OperationLog','operationLog','','sysMonitor/operationLog/index.vue','{\"title\": \"操作日志\"}',1,NULL),(100,'2026-03-05 11:03:58','2023-08-25 09:55:12',NULL,0,'Cenu','/cenu','/cenu/cenu1','Layout','{\"title\": \"多级菜单\", \"svgIcon\": \"menu\", \"alwaysShow\": true}',2,NULL),(101,'2026-03-05 11:03:55','2023-06-28 08:42:39',NULL,100,'Cenu1','cenu1','/cenu/cenu1/cenu1-1','cenu/cenu1/index.vue','{\"title\": \"cenu1\"}',1,NULL),(102,'2026-03-05 11:03:57','2023-06-28 08:42:44',NULL,101,'Cenu1-1','cenu1-1','','cenu/cenu1/cenu1-1/index.vue','{\"title\": \"cenu1-1\"}',1,NULL),(103,'2023-03-13 06:14:27','2023-06-28 08:43:02',NULL,101,'Cenu1-2','cenu1-2','','cenu/cenu1/cenu1-2/index.vue','{\"title\": \"cenu1-2\"}',2,NULL);
/*!40000 ALTER TABLE `sys_management_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_role`
--

DROP TABLE IF EXISTS `sys_management_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_management_role_role_name` (`role_name`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`),
  KEY `idx_authority_role_deleted_at` (`deleted_at`),
  KEY `idx_sys_management_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_role`
--

LOCK TABLES `sys_management_role` WRITE;
/*!40000 ALTER TABLE `sys_management_role` DISABLE KEYS */;
INSERT INTO `sys_management_role` VALUES (1,NULL,'2026-03-06 10:23:09',NULL,'root',NULL);
/*!40000 ALTER TABLE `sys_management_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_role_menus`
--

DROP TABLE IF EXISTS `sys_management_role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_role_menus` (
  `menu_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`menu_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_role_menus`
--

LOCK TABLES `sys_management_role_menus` WRITE;
/*!40000 ALTER TABLE `sys_management_role_menus` DISABLE KEYS */;
INSERT INTO `sys_management_role_menus` VALUES (1,1),(2,1),(3,1),(4,1),(5,1),(6,1),(20,1),(21,1),(22,1),(40,1),(41,1),(100,1),(101,1),(102,1),(103,1);
/*!40000 ALTER TABLE `sys_management_role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_management_user`
--

DROP TABLE IF EXISTS `sys_management_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_management_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL COMMENT '用户名',
  `password` varchar(191) NOT NULL COMMENT '密码',
  `phone` varchar(191) DEFAULT NULL COMMENT '手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) DEFAULT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_management_user_username` (`username`),
  KEY `idx_authority_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_management_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_management_user`
--

LOCK TABLES `sys_management_user` WRITE;
/*!40000 ALTER TABLE `sys_management_user` DISABLE KEYS */;
INSERT INTO `sys_management_user` VALUES (1,'2025-11-22 03:30:11','2025-12-04 21:28:18',NULL,'admin','e10adc3949ba59abbe56e057f20f883e','','',1,1);
/*!40000 ALTER TABLE `sys_management_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_monitor_operationLog`
--

DROP TABLE IF EXISTS `sys_monitor_operationLog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_monitor_operationLog` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `user_agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `req_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '请求Body',
  `resp_data` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
  `resp_time` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  `user_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名称',
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_monitor_operationLog_deleted_at` (`deleted_at`),
  KEY `idx_sys_monitor_operationLog_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_monitor_operationLog`
--

LOCK TABLES `sys_monitor_operationLog` WRITE;
/*!40000 ALTER TABLE `sys_monitor_operationLog` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_monitor_operationLog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_tool_cron`
--

DROP TABLE IF EXISTS `sys_tool_cron`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_tool_cron` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '任务名称',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务方法',
  `expression` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '表达式',
  `strategy` enum('always','once') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'always' COMMENT '执行策略',
  `open` tinyint(1) DEFAULT NULL COMMENT '活跃状态',
  `extraParams` json DEFAULT NULL COMMENT '额外参数',
  `entryId` bigint DEFAULT NULL COMMENT 'cron ID',
  `comment` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_tool_cron_name` (`name`),
  KEY `idx_sysTool_cron_deleted_at` (`deleted_at`),
  KEY `idx_sys_tool_cron_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_tool_cron`
--

LOCK TABLES `sys_tool_cron` WRITE;
/*!40000 ALTER TABLE `sys_tool_cron` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_tool_cron` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_tool_file`
--

DROP TABLE IF EXISTS `sys_tool_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_tool_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件名',
  `full_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件完整路径',
  `mime` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件类型',
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_fileM_file_deleted_at` (`deleted_at`),
  KEY `idx_sys_tool_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_tool_file`
--

LOCK TABLES `sys_tool_file` WRITE;
/*!40000 ALTER TABLE `sys_tool_file` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_tool_file` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-03-17 16:03:25
