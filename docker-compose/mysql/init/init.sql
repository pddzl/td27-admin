-- MySQL dump 10.13  Distrib 9.5.0, for macos26.1 (arm64)
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
-- Table structure for table `authority_api`
--

DROP TABLE IF EXISTS `authority_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authority_api` (
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
  KEY `idx_authority_api_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_api`
--

LOCK TABLES `authority_api` WRITE;
/*!40000 ALTER TABLE `authority_api` DISABLE KEYS */;
INSERT INTO `authority_api` VALUES (1,'2023-03-10 06:24:36','2025-11-23 15:33:45',NULL,'/logReg/captcha','获取验证码（必选）','logReg','POST',NULL),(2,'2023-03-08 06:36:24','2024-01-24 08:59:21',NULL,'/logReg/login','登录（必选）','logReg','POST',NULL),(3,'2024-06-09 15:32:30','2024-06-09 15:32:30',NULL,'/logReg/logout','登出（必选）','logReg','POST',NULL),(4,'2023-03-10 07:21:37','2023-03-10 07:21:37',NULL,'/casbin/editCasbin','编辑casbin规则','casbin','POST',NULL),(5,'2023-03-08 08:56:13','2023-03-10 07:11:53',NULL,'/user/getUserInfo','获取用户信息（必选）','user','GET',NULL),(6,'2023-03-08 08:56:54','2023-03-08 08:56:54',NULL,'/user/getUsers','获取所有用户','user','POST',NULL),(7,'2023-03-10 06:41:32','2023-03-10 06:41:32',NULL,'/user/deleteUser','删除用户','user','POST',NULL),(8,'2023-03-10 06:42:24','2023-03-10 06:42:24',NULL,'/user/addUser','添加用户','user','POST',NULL),(9,'2023-03-10 06:47:18','2023-03-10 06:47:18',NULL,'/user/editUser','编辑用户','user','POST',NULL),(10,'2023-03-10 06:47:59','2023-03-10 06:47:59',NULL,'/user/modifyPass','修改用户密码','user','POST',NULL),(11,'2023-03-10 06:48:43','2023-03-10 06:48:43',NULL,'/user/switchActive','切换用户状态','user','POST',NULL),(12,'2023-03-10 06:58:30','2023-03-10 06:58:30',NULL,'/role/getRoles','获取所有角色','role','POST',NULL),(13,'2023-03-10 06:59:08','2023-03-10 06:59:08',NULL,'/role/addRole','添加角色','role','POST',NULL),(14,'2023-03-10 06:59:54','2023-03-10 06:59:54',NULL,'/role/deleteRole','删除角色','role','POST',NULL),(15,'2023-03-10 07:00:14','2023-03-10 07:00:53',NULL,'/role/editRole','编辑角色','role','POST',NULL),(16,'2023-03-10 07:01:44','2023-03-10 07:01:44',NULL,'/role/editRoleMenu','编辑角色菜单','role','POST',NULL),(17,'2023-03-10 07:14:44','2023-03-10 07:14:44',NULL,'/menu/getMenus','获取所有菜单','menu','GET',NULL),(18,'2023-03-10 07:15:25','2023-03-10 07:15:25',NULL,'/menu/addMenu','添加菜单','menu','POST',NULL),(19,'2023-03-10 07:15:50','2023-03-10 07:15:50',NULL,'/menu/editMenu','编辑菜单','menu','POST',NULL),(20,'2023-03-10 07:16:18','2023-03-10 07:16:18',NULL,'/menu/deleteMenu','删除菜单','menu','POST',NULL),(21,'2023-03-10 07:17:13','2023-03-10 07:17:13',NULL,'/menu/getElTreeMenus','获取所有菜单（el-tree结构）','menu','POST',NULL),(22,'2023-03-10 07:23:21','2023-03-10 07:33:01',NULL,'/api/addApi','添加api','api','POST',NULL),(23,'2023-03-10 07:24:00','2023-03-10 07:24:00',NULL,'/api/getApis','获取所有api','api','POST',NULL),(24,'2023-03-10 07:24:33','2023-03-10 07:24:33',NULL,'/api/deleteApi','删除api','api','POST',NULL),(25,'2023-03-10 07:26:15','2023-03-10 07:26:15',NULL,'/api/editApi','编辑api','api','POST',NULL),(26,'2023-03-10 07:34:08','2023-03-10 07:35:04',NULL,'/api/getElTreeApis','获取所有api（el-tree结构）','api','POST',NULL),(27,'2024-01-03 06:20:38','2024-01-03 06:20:38',NULL,'/api/deleteApiById','批量删除API','api','POST',NULL),(28,'2023-07-13 02:32:16','2024-01-20 04:50:50',NULL,'/opl/getOplList','分页获取操作记录','opl','POST',NULL),(29,'2023-07-13 02:33:32','2024-01-20 04:54:16',NULL,'/opl/deleteOpl','删除操作记录','opl','POST',NULL),(30,'2023-07-13 06:48:47','2024-01-20 04:54:23',NULL,'/opl/deleteOplByIds','批量删除操作记录','opl','POST',NULL),(31,'2023-08-27 06:05:00','2023-08-27 06:05:00',NULL,'/file/upload','文件上传','file','POST',NULL),(32,'2023-08-27 06:06:43','2023-08-27 06:06:43',NULL,'/file/getFileList','分页获取文件信息','file','POST',NULL),(33,'2024-01-04 03:10:15','2024-01-04 03:10:41',NULL,'/file/download','下载文件','file','GET',NULL),(34,'2024-01-04 03:16:04','2024-01-04 03:16:04',NULL,'/file/delete','删除文件','file','GET',NULL),(35,'2024-02-23 08:31:57','2024-02-23 08:31:57',NULL,'/cron/getCronList','分页获取cron','cron','POST',NULL),(36,'2024-02-23 08:33:56','2024-02-23 08:33:56',NULL,'/cron/addCron','添加cron','cron','POST',NULL),(37,'2024-02-23 08:34:25','2024-02-23 08:34:25',NULL,'/cron/deleteCron','删除cron','cron','POST',NULL),(38,'2024-02-23 08:34:50','2024-02-23 08:34:50',NULL,'/cron/editCron','编辑cron','cron','POST',NULL),(39,'2024-02-23 08:35:21','2024-02-23 08:35:21',NULL,'/cron/switchOpen','cron开关','cron','POST',NULL),(40,'2025-09-12 16:08:40','2025-09-12 16:08:40',NULL,'/dict/getDict','Get','dict','GET',NULL),(41,'2025-09-12 16:09:21','2025-09-12 16:09:21',NULL,'/dict/addDict','Add','dict','POST',NULL),(42,'2025-09-12 16:09:46','2025-09-12 16:09:46',NULL,'/dict/editDict','Edit','dict','POST',NULL),(43,'2025-09-12 16:10:13','2025-09-12 16:10:13',NULL,'/dict/delDict','Delete','dict','POST',NULL),(44,'2025-09-12 16:11:17','2025-09-12 16:11:17',NULL,'/dictDetail/getDictDetail','Get','dictDetail','POST',NULL),(45,'2025-09-12 16:11:58','2025-09-12 16:11:58',NULL,'/dictDetail/addDictDetail','Add','dictDetail','POST',NULL),(46,'2025-09-12 16:12:36','2025-09-12 16:12:36',NULL,'/dictDetail/editDictDetail','Edit','dictDetail','POST',NULL),(47,'2025-09-12 16:13:15','2025-09-12 16:14:13',NULL,'/dictDetail/delDictDetail','Delete','dictDetail','POST',NULL),(48,'2025-09-29 15:53:05','2025-09-29 15:53:05',NULL,'/dictDetail/getDictDetailFlat','Get Flat','dictDetail','POST',NULL);
/*!40000 ALTER TABLE `authority_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authority_menu`
--

DROP TABLE IF EXISTS `authority_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_authority_menu_path` (`path`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`),
  KEY `idx_authority_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_menu`
--

LOCK TABLES `authority_menu` WRITE;
/*!40000 ALTER TABLE `authority_menu` DISABLE KEYS */;
INSERT INTO `authority_menu` VALUES (1,NULL,'2024-01-22 07:32:30',NULL,0,'Authority','/authority','/authority/user','Layout','{\"title\": \"权限管理\", \"svgIcon\": \"lock\"}',1,NULL),(2,NULL,'2024-02-05 09:38:41',NULL,1,'User','user','','authority/user/index.vue','{\"title\": \"用户管理\"}',1,NULL),(3,NULL,'2023-06-28 08:12:06',NULL,1,'Role','role','','authority/role/index.vue','{\"title\": \"角色管理\"}',2,NULL),(4,NULL,'2023-06-28 08:12:16',NULL,1,'Menu','menu','','authority/menu/index.vue','{\"title\": \"菜单管理\"}',3,NULL),(5,'2023-03-07 01:50:48','2023-06-28 08:11:38',NULL,1,'Api','api','','authority/api/index.vue','{\"title\": \"接口管理\"}',4,NULL),(6,NULL,'2023-08-25 09:55:12',NULL,0,'Cenu','/cenu','/cenu/cenu1','Layout','{\"title\": \"多级菜单\", \"svgIcon\": \"menu\", \"alwaysShow\": true}',2,NULL),(7,NULL,'2023-06-28 08:42:39',NULL,6,'Cenu1','cenu1','/cenu/cenu1/cenu1-1','cenu/cenu1/index.vue','{\"title\": \"cenu1\"}',1,NULL),(8,NULL,'2023-06-28 08:42:44',NULL,7,'Cenu1-1','cenu1-1','','cenu/cenu1/cenu1-1/index.vue','{\"title\": \"cenu1-1\"}',1,NULL),(9,'2023-03-13 06:14:27','2023-06-28 08:43:02',NULL,7,'Cenu1-2','cenu1-2','','cenu/cenu1/cenu1-2/index.vue','{\"title\": \"cenu1-2\"}',2,NULL),(10,'2023-08-26 08:57:01','2023-08-26 09:02:58',NULL,0,'FileM','/fileM','/fileM/file','Layout','{\"title\": \"文件管理\", \"svgIcon\": \"file\", \"alwaysShow\": true}',3,NULL),(11,'2023-08-26 08:58:51','2023-08-26 08:58:51',NULL,10,'File','/fileM/file','','fileM/file/index.vue','{\"title\": \"文件\"}',1,NULL),(12,'2024-01-19 07:47:49','2024-03-25 06:14:36',NULL,0,'Monitor','/monitor','/monitor/operationLog','Layout','{\"title\": \"系统监控\", \"svgIcon\": \"monitor\", \"alwaysShow\": true}',5,NULL),(13,'2023-03-07 01:50:48','2024-01-19 07:48:52',NULL,12,'OperationLog','operationLog','','monitor/operationLog/index.vue','{\"title\": \"操作日志\"}',1,NULL),(14,'2024-02-05 09:56:33','2024-02-05 09:56:33',NULL,0,'SysTool','/systool','/systool/cron','Layout','{\"title\": \"系统工具\", \"svgIcon\": \"config\", \"alwaysShow\": true}',4,NULL),(15,'2024-02-06 10:00:00','2024-02-06 10:00:00',NULL,14,'Cron','cron','','sysTool/cron/index.vue','{\"title\": \"定时任务\"}',1,NULL),(16,'2024-02-05 09:56:33','2025-09-12 16:05:36',NULL,0,'SysSet','/sysSet','/sysSet/dict','Layout','{\"title\": \"系统设置\", \"svgIcon\": \"config\", \"alwaysShow\": true}',6,NULL),(17,'2024-02-06 10:00:00','2025-09-29 06:18:23',NULL,16,'Dict','dict','','sysSet/dict/index.vue','{\"title\": \"字典管理\"}',1,NULL);
/*!40000 ALTER TABLE `authority_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authority_role`
--

DROP TABLE IF EXISTS `authority_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authority_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_authority_role_role_name` (`role_name`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`),
  KEY `idx_authority_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_role`
--

LOCK TABLES `authority_role` WRITE;
/*!40000 ALTER TABLE `authority_role` DISABLE KEYS */;
INSERT INTO `authority_role` VALUES (1,NULL,'2025-09-12 16:04:46',NULL,'root',NULL);
/*!40000 ALTER TABLE `authority_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authority_user`
--

DROP TABLE IF EXISTS `authority_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authority_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL COMMENT '用户名',
  `password` varchar(191) NOT NULL COMMENT '密码',
  `phone` varchar(191) DEFAULT NULL COMMENT '手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) DEFAULT NULL,
  `role_model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_authority_user_username` (`username`),
  KEY `idx_authority_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_user`
--

LOCK TABLES `authority_user` WRITE;
/*!40000 ALTER TABLE `authority_user` DISABLE KEYS */;
INSERT INTO `authority_user` VALUES (1,'2025-11-22 03:30:11',NULL,NULL,'admin','e10adc3949ba59abbe56e057f20f883e',NULL,NULL,1,1);
/*!40000 ALTER TABLE `authority_user` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB AUTO_INCREMENT=380 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (353,'p','1','/api/addApi','POST','','',''),(355,'p','1','/api/deleteApi','POST','','',''),(358,'p','1','/api/deleteApiById','POST','','',''),(356,'p','1','/api/editApi','POST','','',''),(354,'p','1','/api/getApis','POST','','',''),(357,'p','1','/api/getElTreeApis','POST','','',''),(335,'p','1','/casbin/editCasbin','POST','','',''),(367,'p','1','/cron/addCron','POST','','',''),(368,'p','1','/cron/deleteCron','POST','','',''),(369,'p','1','/cron/editCron','POST','','',''),(366,'p','1','/cron/getCronList','POST','','',''),(370,'p','1','/cron/switchOpen','POST','','',''),(372,'p','1','/dict/addDict','POST','','',''),(374,'p','1','/dict/delDict','POST','','',''),(373,'p','1','/dict/editDict','POST','','',''),(371,'p','1','/dict/getDict','GET','','',''),(376,'p','1','/dictDetail/addDictDetail','POST','','',''),(378,'p','1','/dictDetail/delDictDetail','POST','','',''),(377,'p','1','/dictDetail/editDictDetail','POST','','',''),(375,'p','1','/dictDetail/getDictDetail','POST','','',''),(379,'p','1','/dictDetail/getDictDetailFlat','POST','','',''),(365,'p','1','/file/delete','GET','','',''),(364,'p','1','/file/download','GET','','',''),(363,'p','1','/file/getFileList','POST','','',''),(362,'p','1','/file/upload','POST','','',''),(332,'p','1','/logReg/captcha','POST','','',''),(333,'p','1','/logReg/login','POST','','',''),(334,'p','1','/logReg/logout','POST','','',''),(349,'p','1','/menu/addMenu','POST','','',''),(351,'p','1','/menu/deleteMenu','POST','','',''),(350,'p','1','/menu/editMenu','POST','','',''),(352,'p','1','/menu/getElTreeMenus','POST','','',''),(348,'p','1','/menu/getMenus','GET','','',''),(360,'p','1','/opl/deleteOpl','POST','','',''),(361,'p','1','/opl/deleteOplByIds','POST','','',''),(359,'p','1','/opl/getOplList','POST','','',''),(344,'p','1','/role/addRole','POST','','',''),(345,'p','1','/role/deleteRole','POST','','',''),(346,'p','1','/role/editRole','POST','','',''),(347,'p','1','/role/editRoleMenu','POST','','',''),(343,'p','1','/role/getRoles','POST','','',''),(339,'p','1','/user/addUser','POST','','',''),(338,'p','1','/user/deleteUser','POST','','',''),(340,'p','1','/user/editUser','POST','','',''),(336,'p','1','/user/getUserInfo','GET','','',''),(337,'p','1','/user/getUsers','POST','','',''),(341,'p','1','/user/modifyPass','POST','','',''),(342,'p','1','/user/switchActive','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fileM_file`
--

DROP TABLE IF EXISTS `fileM_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fileM_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件名',
  `full_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件完整路径',
  `mime` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件类型',
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_fileM_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fileM_file`
--

LOCK TABLES `fileM_file` WRITE;
/*!40000 ALTER TABLE `fileM_file` DISABLE KEYS */;
/*!40000 ALTER TABLE `fileM_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_operationLog`
--

DROP TABLE IF EXISTS `monitor_operationLog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `monitor_operationLog` (
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
  KEY `idx_monitor_operationLog_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_operationLog`
--

LOCK TABLES `monitor_operationLog` WRITE;
/*!40000 ALTER TABLE `monitor_operationLog` DISABLE KEYS */;
INSERT INTO `monitor_operationLog` VALUES (1,NULL,NULL,NULL,'127.0.0.1','GET','/user/getUserInfo',200,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','{}','{\"code\":0,\"data\":{\"id\":1,\"createdAt\":\"2025-11-22T03:30:11+08:00\",\"updatedAt\":\"0001-01-01T00:00:00Z\",\"username\":\"admin\",\"phone\":\"\",\"email\":\"\",\"active\":true,\"roleId\":1,\"roleName\":\"root\"},\"msg\":\"获取成功\"}',1,1,'admin',NULL),(2,NULL,NULL,NULL,'127.0.0.1','GET','/menu/getMenus',200,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','{}','{\"code\":0,\"data\":[{\"id\":1,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2024-01-22T07:32:30+08:00\",\"pid\":0,\"name\":\"Authority\",\"path\":\"/authority\",\"redirect\":\"/authority/user\",\"component\":\"Layout\",\"sort\":1,\"meta\":{\"title\":\"权限管理\",\"svgIcon\":\"lock\"},\"children\":[{\"id\":2,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2024-02-05T09:38:41+08:00\",\"pid\":1,\"name\":\"User\",\"path\":\"user\",\"component\":\"authority/user/index.vue\",\"sort\":1,\"meta\":{\"title\":\"用户管理\"}},{\"id\":3,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2023-06-28T08:12:06+08:00\",\"pid\":1,\"name\":\"Role\",\"path\":\"role\",\"component\":\"authority/role/index.vue\",\"sort\":2,\"meta\":{\"title\":\"角色管理\"}},{\"id\":4,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2023-06-28T08:12:16+08:00\",\"pid\":1,\"name\":\"Menu\",\"path\":\"menu\",\"component\":\"authority/menu/index.vue\",\"sort\":3,\"meta\":{\"title\":\"菜单管理\"}},{\"id\":5,\"createdAt\":\"2023-03-07T01:50:48+08:00\",\"updatedAt\":\"2023-06-28T08:11:38+08:00\",\"pid\":1,\"name\":\"Api\",\"path\":\"api\",\"component\":\"authority/api/index.vue\",\"sort\":4,\"meta\":{\"title\":\"接口管理\"}}]},{\"id\":6,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2023-08-25T09:55:12+08:00\",\"pid\":0,\"name\":\"Cenu\",\"path\":\"/cenu\",\"redirect\":\"/cenu/cenu1\",\"component\":\"Layout\",\"sort\":2,\"meta\":{\"title\":\"多级菜单\",\"svgIcon\":\"menu\",\"alwaysShow\":true},\"children\":[{\"id\":7,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2023-06-28T08:42:39+08:00\",\"pid\":6,\"name\":\"Cenu1\",\"path\":\"cenu1\",\"redirect\":\"/cenu/cenu1/cenu1-1\",\"component\":\"cenu/cenu1/index.vue\",\"sort\":1,\"meta\":{\"title\":\"cenu1\"},\"children\":[{\"id\":8,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"2023-06-28T08:42:44+08:00\",\"pid\":7,\"name\":\"Cenu1-1\",\"path\":\"cenu1-1\",\"component\":\"cenu/cenu1/cenu1-1/index.vue\",\"sort\":1,\"meta\":{\"title\":\"cenu1-1\"}},{\"id\":9,\"createdAt\":\"2023-03-13T06:14:27+08:00\",\"updatedAt\":\"2023-06-28T08:43:02+08:00\",\"pid\":7,\"name\":\"Cenu1-2\",\"path\":\"cenu1-2\",\"component\":\"cenu/cenu1/cenu1-2/index.vue\",\"sort\":2,\"meta\":{\"title\":\"cenu1-2\"}}]}]},{\"id\":10,\"createdAt\":\"2023-08-26T08:57:01+08:00\",\"updatedAt\":\"2023-08-26T09:02:58+08:00\",\"pid\":0,\"name\":\"FileM\",\"path\":\"/fileM\",\"redirect\":\"/fileM/file\",\"component\":\"Layout\",\"sort\":3,\"meta\":{\"title\":\"文件管理\",\"svgIcon\":\"file\",\"alwaysShow\":true},\"children\":[{\"id\":11,\"createdAt\":\"2023-08-26T08:58:51+08:00\",\"updatedAt\":\"2023-08-26T08:58:51+08:00\",\"pid\":10,\"name\":\"File\",\"path\":\"/fileM/file\",\"component\":\"fileM/file/index.vue\",\"sort\":1,\"meta\":{\"title\":\"文件\"}}]},{\"id\":14,\"createdAt\":\"2024-02-05T09:56:33+08:00\",\"updatedAt\":\"2024-02-05T09:56:33+08:00\",\"pid\":0,\"name\":\"SysTool\",\"path\":\"/systool\",\"redirect\":\"/systool/cron\",\"component\":\"Layout\",\"sort\":4,\"meta\":{\"title\":\"系统工具\",\"svgIcon\":\"config\",\"alwaysShow\":true},\"children\":[{\"id\":15,\"createdAt\":\"2024-02-06T10:00:00+08:00\",\"updatedAt\":\"2024-02-06T10:00:00+08:00\",\"pid\":14,\"name\":\"Cron\",\"path\":\"cron\",\"component\":\"sysTool/cron/index.vue\",\"sort\":1,\"meta\":{\"title\":\"定时任务\"}}]},{\"id\":12,\"createdAt\":\"2024-01-19T07:47:49+08:00\",\"updatedAt\":\"2024-03-25T06:14:36+08:00\",\"pid\":0,\"name\":\"Monitor\",\"path\":\"/monitor\",\"redirect\":\"/monitor/operationLog\",\"component\":\"Layout\",\"sort\":5,\"meta\":{\"title\":\"系统监控\",\"svgIcon\":\"monitor\",\"alwaysShow\":true},\"children\":[{\"id\":13,\"createdAt\":\"2023-03-07T01:50:48+08:00\",\"updatedAt\":\"2024-01-19T07:48:52+08:00\",\"pid\":12,\"name\":\"OperationLog\",\"path\":\"operationLog\",\"component\":\"monitor/operationLog/index.vue\",\"sort\":1,\"meta\":{\"title\":\"操作日志\"}}]},{\"id\":16,\"createdAt\":\"2024-02-05T09:56:33+08:00\",\"updatedAt\":\"2025-09-12T16:05:36+08:00\",\"pid\":0,\"name\":\"SysSet\",\"path\":\"/sysSet\",\"redirect\":\"/sysSet/dict\",\"component\":\"Layout\",\"sort\":6,\"meta\":{\"title\":\"系统设置\",\"svgIcon\":\"config\",\"alwaysShow\":true},\"children\":[{\"id\":17,\"createdAt\":\"2024-02-06T10:00:00+08:00\",\"updatedAt\":\"2025-09-29T06:18:23+08:00\",\"pid\":16,\"name\":\"Dict\",\"path\":\"dict\",\"component\":\"sysSet/dict/index.vue\",\"sort\":1,\"meta\":{\"title\":\"字典管理\"}}]}],\"msg\":\"获取成功\"}',7,1,'admin',NULL),(3,NULL,NULL,NULL,'127.0.0.1','POST','/opl/getOplList',200,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','{\"page\":1,\"pageSize\":10}','{\"code\":0,\"data\":{\"list\":[{\"id\":2,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"0001-01-01T00:00:00Z\",\"ip\":\"127.0.0.1\",\"method\":\"GET\",\"path\":\"/menu/getMenus\",\"status\":200,\"userAgent\":\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36\",\"reqParam\":\"{}\",\"respData\":\"{\\\"code\\\":0,\\\"data\\\":[{\\\"id\\\":1,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2024-01-22T07:32:30+08:00\\\",\\\"pid\\\":0,\\\"name\\\":\\\"Authority\\\",\\\"path\\\":\\\"/authority\\\",\\\"redirect\\\":\\\"/authority/user\\\",\\\"component\\\":\\\"Layout\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"权限管理\\\",\\\"svgIcon\\\":\\\"lock\\\"},\\\"children\\\":[{\\\"id\\\":2,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2024-02-05T09:38:41+08:00\\\",\\\"pid\\\":1,\\\"name\\\":\\\"User\\\",\\\"path\\\":\\\"user\\\",\\\"component\\\":\\\"authority/user/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"用户管理\\\"}},{\\\"id\\\":3,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2023-06-28T08:12:06+08:00\\\",\\\"pid\\\":1,\\\"name\\\":\\\"Role\\\",\\\"path\\\":\\\"role\\\",\\\"component\\\":\\\"authority/role/index.vue\\\",\\\"sort\\\":2,\\\"meta\\\":{\\\"title\\\":\\\"角色管理\\\"}},{\\\"id\\\":4,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2023-06-28T08:12:16+08:00\\\",\\\"pid\\\":1,\\\"name\\\":\\\"Menu\\\",\\\"path\\\":\\\"menu\\\",\\\"component\\\":\\\"authority/menu/index.vue\\\",\\\"sort\\\":3,\\\"meta\\\":{\\\"title\\\":\\\"菜单管理\\\"}},{\\\"id\\\":5,\\\"createdAt\\\":\\\"2023-03-07T01:50:48+08:00\\\",\\\"updatedAt\\\":\\\"2023-06-28T08:11:38+08:00\\\",\\\"pid\\\":1,\\\"name\\\":\\\"Api\\\",\\\"path\\\":\\\"api\\\",\\\"component\\\":\\\"authority/api/index.vue\\\",\\\"sort\\\":4,\\\"meta\\\":{\\\"title\\\":\\\"接口管理\\\"}}]},{\\\"id\\\":6,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2023-08-25T09:55:12+08:00\\\",\\\"pid\\\":0,\\\"name\\\":\\\"Cenu\\\",\\\"path\\\":\\\"/cenu\\\",\\\"redirect\\\":\\\"/cenu/cenu1\\\",\\\"component\\\":\\\"Layout\\\",\\\"sort\\\":2,\\\"meta\\\":{\\\"title\\\":\\\"多级菜单\\\",\\\"svgIcon\\\":\\\"menu\\\",\\\"alwaysShow\\\":true},\\\"children\\\":[{\\\"id\\\":7,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2023-06-28T08:42:39+08:00\\\",\\\"pid\\\":6,\\\"name\\\":\\\"Cenu1\\\",\\\"path\\\":\\\"cenu1\\\",\\\"redirect\\\":\\\"/cenu/cenu1/cenu1-1\\\",\\\"component\\\":\\\"cenu/cenu1/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"cenu1\\\"},\\\"children\\\":[{\\\"id\\\":8,\\\"createdAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updatedAt\\\":\\\"2023-06-28T08:42:44+08:00\\\",\\\"pid\\\":7,\\\"name\\\":\\\"Cenu1-1\\\",\\\"path\\\":\\\"cenu1-1\\\",\\\"component\\\":\\\"cenu/cenu1/cenu1-1/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"cenu1-1\\\"}},{\\\"id\\\":9,\\\"createdAt\\\":\\\"2023-03-13T06:14:27+08:00\\\",\\\"updatedAt\\\":\\\"2023-06-28T08:43:02+08:00\\\",\\\"pid\\\":7,\\\"name\\\":\\\"Cenu1-2\\\",\\\"path\\\":\\\"cenu1-2\\\",\\\"component\\\":\\\"cenu/cenu1/cenu1-2/index.vue\\\",\\\"sort\\\":2,\\\"meta\\\":{\\\"title\\\":\\\"cenu1-2\\\"}}]}]},{\\\"id\\\":10,\\\"createdAt\\\":\\\"2023-08-26T08:57:01+08:00\\\",\\\"updatedAt\\\":\\\"2023-08-26T09:02:58+08:00\\\",\\\"pid\\\":0,\\\"name\\\":\\\"FileM\\\",\\\"path\\\":\\\"/fileM\\\",\\\"redirect\\\":\\\"/fileM/file\\\",\\\"component\\\":\\\"Layout\\\",\\\"sort\\\":3,\\\"meta\\\":{\\\"title\\\":\\\"文件管理\\\",\\\"svgIcon\\\":\\\"file\\\",\\\"alwaysShow\\\":true},\\\"children\\\":[{\\\"id\\\":11,\\\"createdAt\\\":\\\"2023-08-26T08:58:51+08:00\\\",\\\"updatedAt\\\":\\\"2023-08-26T08:58:51+08:00\\\",\\\"pid\\\":10,\\\"name\\\":\\\"File\\\",\\\"path\\\":\\\"/fileM/file\\\",\\\"component\\\":\\\"fileM/file/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"文件\\\"}}]},{\\\"id\\\":14,\\\"createdAt\\\":\\\"2024-02-05T09:56:33+08:00\\\",\\\"updatedAt\\\":\\\"2024-02-05T09:56:33+08:00\\\",\\\"pid\\\":0,\\\"name\\\":\\\"SysTool\\\",\\\"path\\\":\\\"/systool\\\",\\\"redirect\\\":\\\"/systool/cron\\\",\\\"component\\\":\\\"Layout\\\",\\\"sort\\\":4,\\\"meta\\\":{\\\"title\\\":\\\"系统工具\\\",\\\"svgIcon\\\":\\\"config\\\",\\\"alwaysShow\\\":true},\\\"children\\\":[{\\\"id\\\":15,\\\"createdAt\\\":\\\"2024-02-06T10:00:00+08:00\\\",\\\"updatedAt\\\":\\\"2024-02-06T10:00:00+08:00\\\",\\\"pid\\\":14,\\\"name\\\":\\\"Cron\\\",\\\"path\\\":\\\"cron\\\",\\\"component\\\":\\\"sysTool/cron/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"定时任务\\\"}}]},{\\\"id\\\":12,\\\"createdAt\\\":\\\"2024-01-19T07:47:49+08:00\\\",\\\"updatedAt\\\":\\\"2024-03-25T06:14:36+08:00\\\",\\\"pid\\\":0,\\\"name\\\":\\\"Monitor\\\",\\\"path\\\":\\\"/monitor\\\",\\\"redirect\\\":\\\"/monitor/operationLog\\\",\\\"component\\\":\\\"Layout\\\",\\\"sort\\\":5,\\\"meta\\\":{\\\"title\\\":\\\"系统监控\\\",\\\"svgIcon\\\":\\\"monitor\\\",\\\"alwaysShow\\\":true},\\\"children\\\":[{\\\"id\\\":13,\\\"createdAt\\\":\\\"2023-03-07T01:50:48+08:00\\\",\\\"updatedAt\\\":\\\"2024-01-19T07:48:52+08:00\\\",\\\"pid\\\":12,\\\"name\\\":\\\"OperationLog\\\",\\\"path\\\":\\\"operationLog\\\",\\\"component\\\":\\\"monitor/operationLog/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"操作日志\\\"}}]},{\\\"id\\\":16,\\\"createdAt\\\":\\\"2024-02-05T09:56:33+08:00\\\",\\\"updatedAt\\\":\\\"2025-09-12T16:05:36+08:00\\\",\\\"pid\\\":0,\\\"name\\\":\\\"SysSet\\\",\\\"path\\\":\\\"/sysSet\\\",\\\"redirect\\\":\\\"/sysSet/dict\\\",\\\"component\\\":\\\"Layout\\\",\\\"sort\\\":6,\\\"meta\\\":{\\\"title\\\":\\\"系统设置\\\",\\\"svgIcon\\\":\\\"config\\\",\\\"alwaysShow\\\":true},\\\"children\\\":[{\\\"id\\\":17,\\\"createdAt\\\":\\\"2024-02-06T10:00:00+08:00\\\",\\\"updatedAt\\\":\\\"2025-09-29T06:18:23+08:00\\\",\\\"pid\\\":16,\\\"name\\\":\\\"Dict\\\",\\\"path\\\":\\\"dict\\\",\\\"component\\\":\\\"sysSet/dict/index.vue\\\",\\\"sort\\\":1,\\\"meta\\\":{\\\"title\\\":\\\"字典管理\\\"}}]}],\\\"msg\\\":\\\"获取成功\\\"}\",\"respTime\":7,\"userID\":1,\"userName\":\"admin\"},{\"id\":1,\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"0001-01-01T00:00:00Z\",\"ip\":\"127.0.0.1\",\"method\":\"GET\",\"path\":\"/user/getUserInfo\",\"status\":200,\"userAgent\":\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36\",\"reqParam\":\"{}\",\"respData\":\"{\\\"code\\\":0,\\\"data\\\":{\\\"id\\\":1,\\\"createdAt\\\":\\\"2025-11-22T03:30:11+08:00\\\",\\\"updatedAt\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"username\\\":\\\"admin\\\",\\\"phone\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"active\\\":true,\\\"roleId\\\":1,\\\"roleName\\\":\\\"root\\\"},\\\"msg\\\":\\\"获取成功\\\"}\",\"respTime\":1,\"userID\":1,\"userName\":\"admin\"}],\"total\":2,\"page\":1,\"pageSize\":10},\"msg\":\"获取成功\"}',2,1,'admin',NULL);
/*!40000 ALTER TABLE `monitor_operationLog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_menus`
--

DROP TABLE IF EXISTS `role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_menus` (
  `menu_model_id` bigint unsigned NOT NULL,
  `role_model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`menu_model_id`,`role_model_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_menus`
--

LOCK TABLES `role_menus` WRITE;
/*!40000 ALTER TABLE `role_menus` DISABLE KEYS */;
INSERT INTO `role_menus` VALUES (1,1),(2,1),(3,1),(4,1),(5,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1),(14,1),(15,1),(16,1),(17,1);
/*!40000 ALTER TABLE `role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sysSet_dict`
--

DROP TABLE IF EXISTS `sysSet_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sysSet_dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ch_name` varchar(191) DEFAULT NULL,
  `en_name` varchar(191) DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sysSet_dict_ch_name` (`ch_name`),
  UNIQUE KEY `uni_sysSet_dict_en_name` (`en_name`),
  KEY `idx_sysSet_dict_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sysSet_dict`
--

LOCK TABLES `sysSet_dict` WRITE;
/*!40000 ALTER TABLE `sysSet_dict` DISABLE KEYS */;
INSERT INTO `sysSet_dict` VALUES (1,'2025-09-12 16:20:56','2025-09-12 16:20:56',NULL,'主机状态','HostStatus',NULL);
/*!40000 ALTER TABLE `sysSet_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sysSet_dictDetail`
--

DROP TABLE IF EXISTS `sysSet_dictDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sysSet_dictDetail` (
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
  KEY `idx_sysSet_dictDetail_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sysSet_dictDetail`
--

LOCK TABLES `sysSet_dictDetail` WRITE;
/*!40000 ALTER TABLE `sysSet_dictDetail` DISABLE KEYS */;
INSERT INTO `sysSet_dictDetail` VALUES (1,'2025-09-13 08:51:09','2025-09-13 08:51:09',NULL,'运行中','running',0,1,NULL,NULL,NULL),(2,'2025-09-13 08:51:36','2025-09-13 08:51:36',NULL,'不可达','unreached',1,1,NULL,NULL,NULL),(4,'2025-09-28 11:19:51','2025-09-29 06:33:03',NULL,'t1','t1',0,1,2,'test',NULL),(10,'2025-09-29 03:01:18','2025-09-29 06:26:57',NULL,'t10','t10',2,1,1,NULL,NULL);
/*!40000 ALTER TABLE `sysSet_dictDetail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sysTool_cron`
--

DROP TABLE IF EXISTS `sysTool_cron`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sysTool_cron` (
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
  UNIQUE KEY `uni_sysTool_cron_name` (`name`),
  KEY `idx_sysTool_cron_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sysTool_cron`
--

LOCK TABLES `sysTool_cron` WRITE;
/*!40000 ALTER TABLE `sysTool_cron` DISABLE KEYS */;
/*!40000 ALTER TABLE `sysTool_cron` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-11-23 16:10:11
