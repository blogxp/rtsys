/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : b5gocmf

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/11/2022 15:33:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for b5net_admin
-- ----------------------------
DROP TABLE IF EXISTS `b5net_admin`;
CREATE TABLE `b5net_admin`  (
  `id` bigint UNSIGNED NOT NULL,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `realname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '人员姓名',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '状态',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_admin
-- ----------------------------
INSERT INTO `b5net_admin` VALUES (10000, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超管', '1', '超级管理员', '2020-12-24 10:50:56', '2022-04-13 05:49:00');
INSERT INTO `b5net_admin` VALUES (104683566088065024, 'test', 'e10adc3949ba59abbe56e057f20f883e', 'test', '1', '', '2022-10-16 20:55:07', '2022-11-10 15:32:10');
INSERT INTO `b5net_admin` VALUES (113601612861149184, 'test1', 'e10adc3949ba59abbe56e057f20f883e', 'test1', '1', '', '2022-11-10 11:32:15', '2022-11-10 15:32:14');
INSERT INTO `b5net_admin` VALUES (113602333207695360, 'test2', 'e10adc3949ba59abbe56e057f20f883e', 'test2', '1', '', '2022-11-10 11:35:06', '2022-11-10 15:32:24');
INSERT INTO `b5net_admin` VALUES (113602602515566592, 'test3', 'e10adc3949ba59abbe56e057f20f883e', 'test3', '1', '', '2022-11-10 11:36:11', '2022-11-10 15:32:30');

-- ----------------------------
-- Table structure for b5net_admin_pos
-- ----------------------------
DROP TABLE IF EXISTS `b5net_admin_pos`;
CREATE TABLE `b5net_admin_pos`  (
  `admin_id` bigint NOT NULL COMMENT '用户ID',
  `pos_id` bigint NOT NULL COMMENT '职位ID',
  UNIQUE INDEX `admin_id`(`admin_id`, `pos_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和职位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_admin_pos
-- ----------------------------
INSERT INTO `b5net_admin_pos` VALUES (10000, 1);

-- ----------------------------
-- Table structure for b5net_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `b5net_admin_role`;
CREATE TABLE `b5net_admin_role`  (
  `admin_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  UNIQUE INDEX `admin_id`(`admin_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_admin_role
-- ----------------------------
INSERT INTO `b5net_admin_role` VALUES (10000, 1);
INSERT INTO `b5net_admin_role` VALUES (104683566088065024, 113603187126046720);
INSERT INTO `b5net_admin_role` VALUES (113601612861149184, 104682922295955456);
INSERT INTO `b5net_admin_role` VALUES (113602333207695360, 104682922295955456);
INSERT INTO `b5net_admin_role` VALUES (113602602515566592, 104682922295955456);

-- ----------------------------
-- Table structure for b5net_admin_struct
-- ----------------------------
DROP TABLE IF EXISTS `b5net_admin_struct`;
CREATE TABLE `b5net_admin_struct`  (
  `admin_id` bigint NOT NULL COMMENT '用户ID',
  `struct_id` bigint NOT NULL COMMENT '组织ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与组织架构关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_admin_struct
-- ----------------------------
INSERT INTO `b5net_admin_struct` VALUES (10000, 100);
INSERT INTO `b5net_admin_struct` VALUES (104683566088065024, 104677601469009920);
INSERT INTO `b5net_admin_struct` VALUES (113601612861149184, 104677839734837248);
INSERT INTO `b5net_admin_struct` VALUES (113602333207695360, 104677894931877888);
INSERT INTO `b5net_admin_struct` VALUES (113602602515566592, 104677679562756096);

-- ----------------------------
-- Table structure for b5net_app_token
-- ----------------------------
DROP TABLE IF EXISTS `b5net_app_token`;
CREATE TABLE `b5net_app_token`  (
  `token` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户类型',
  `plat` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台',
  `extend` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '额外信息',
  `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户ID',
  `exp_time` datetime NULL DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`token`) USING BTREE,
  UNIQUE INDEX `token`(`token`) USING BTREE,
  UNIQUE INDEX `type`(`type`, `user_id`, `plat`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_app_token
-- ----------------------------
INSERT INTO `b5net_app_token` VALUES ('34e86997846936d288e00ccb32aa9297', 'store', 'app', '', 1001, '2022-11-03 23:49:16');

-- ----------------------------
-- Table structure for b5net_config
-- ----------------------------
DROP TABLE IF EXISTS `b5net_config`;
CREATE TABLE `b5net_config`  (
  `id` bigint UNSIGNED NOT NULL COMMENT '配置ID',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置名称',
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置标识',
  `style` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置类型',
  `is_sys` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '是否系统内置 0否 1是',
  `groups` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '配置分组',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '配置值',
  `extra` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '配置项',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '配置说明',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `type`(`type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_config
-- ----------------------------
INSERT INTO `b5net_config` VALUES (1, '配置分组', 'sys_config_group', 'array', '1', '', 'site:基本设置\r\nwx:微信设置\r\nsys:系统配置', '', '', '2020-12-31 14:01:18', '2022-10-16 19:41:08');
INSERT INTO `b5net_config` VALUES (2, '系统名称', 'sys_config_sysname', 'text', '1', 'site', 'RTSys', '', '系统后台显示的名称', '2020-12-31 14:01:18', '2022-10-16 19:41:14');
INSERT INTO `b5net_config` VALUES (10, '公众号appid', 'wechat_appid', 'text', '0', 'wx', '', '', '微信公众号的AppId', '2021-01-12 11:05:50', '2022-10-31 23:21:40');
INSERT INTO `b5net_config` VALUES (11, '公众号secret', 'wechat_appsecret', 'text', '0', 'wx', '', '', '微信公众号-AppSecret', '2021-01-12 11:06:24', '2022-11-10 11:13:01');
INSERT INTO `b5net_config` VALUES (104665361156149248, '组织类型', 'sys_struct_type', 'array', '1', 'sys', 'group:集团\r\ncom:公司\r\ndep:部门\r\nteam:小组', '', '', '2022-10-16 19:42:46', '2022-10-16 19:42:46');

-- ----------------------------
-- Table structure for b5net_login_log
-- ----------------------------
DROP TABLE IF EXISTS `b5net_login_log`;
CREATE TABLE `b5net_login_log`  (
  `id` bigint NOT NULL COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ip_addr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `net` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `create_time` datetime NULL DEFAULT NULL COMMENT '访问时间',
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_login_log
-- ----------------------------
INSERT INTO `b5net_login_log` VALUES (110669113432477696, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-02 09:19:32', '2022-11-02 09:19:32');
INSERT INTO `b5net_login_log` VALUES (110756760209330176, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-02 15:07:49', '2022-11-02 15:07:49');
INSERT INTO `b5net_login_log` VALUES (111100041241825280, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-03 13:51:53', '2022-11-03 13:51:53');
INSERT INTO `b5net_login_log` VALUES (111259127178596352, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-04 00:24:03', '2022-11-04 00:24:03');
INSERT INTO `b5net_login_log` VALUES (111396667311263744, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-04 09:30:35', '2022-11-04 09:30:35');
INSERT INTO `b5net_login_log` VALUES (111810395072630784, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-05 12:54:35', '2022-11-05 12:54:35');
INSERT INTO `b5net_login_log` VALUES (112145177912545280, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-06 11:04:53', '2022-11-06 11:04:53');
INSERT INTO `b5net_login_log` VALUES (112512221086486528, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-07 11:23:23', '2022-11-07 11:23:23');
INSERT INTO `b5net_login_log` VALUES (113596695165538304, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '账号或密码错误', '0', '2022-11-10 11:12:42', '2022-11-10 11:12:42');
INSERT INTO `b5net_login_log` VALUES (113596724919930880, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 11:12:49', '2022-11-10 11:12:49');
INSERT INTO `b5net_login_log` VALUES (113603705848205312, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 11:40:34', '2022-11-10 11:40:34');
INSERT INTO `b5net_login_log` VALUES (113627209993818112, 'admin', '127.0.0.1', '  ', ' 11.2.5170.400', ' ', '', '登录成功', '1', '2022-11-10 13:13:57', '2022-11-10 13:13:57');
INSERT INTO `b5net_login_log` VALUES (113627828334891008, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 13:16:25', '2022-11-10 13:16:25');
INSERT INTO `b5net_login_log` VALUES (113628685348638720, 'test', '127.0.0.1', '  ', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 13:19:49', '2022-11-10 13:19:49');
INSERT INTO `b5net_login_log` VALUES (113633480675430400, 'test1', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '账号或密码错误', '0', '2022-11-10 13:38:53', '2022-11-10 13:38:53');
INSERT INTO `b5net_login_log` VALUES (113633533322334208, 'test1', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '账号或密码错误', '0', '2022-11-10 13:39:05', '2022-11-10 13:39:05');
INSERT INTO `b5net_login_log` VALUES (113633577601601536, 'test1', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 13:39:16', '2022-11-10 13:39:16');
INSERT INTO `b5net_login_log` VALUES (113637260527669248, 'test1', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 13:53:54', '2022-11-10 13:53:54');
INSERT INTO `b5net_login_log` VALUES (113649724162052096, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '账号或密码错误', '0', '2022-11-10 14:43:25', '2022-11-10 14:43:25');
INSERT INTO `b5net_login_log` VALUES (113649743510376448, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '账号或密码错误', '0', '2022-11-10 14:43:30', '2022-11-10 14:43:30');
INSERT INTO `b5net_login_log` VALUES (113649778805444608, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '账号或密码错误', '0', '2022-11-10 14:43:38', '2022-11-10 14:43:38');
INSERT INTO `b5net_login_log` VALUES (113649806412353536, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 14:43:45', '2022-11-10 14:43:45');
INSERT INTO `b5net_login_log` VALUES (113657503861968896, 'test', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 15:14:20', '2022-11-10 15:14:20');
INSERT INTO `b5net_login_log` VALUES (113661923475591168, 'admin', '127.0.0.1', '本机地址', 'Chrome 105.0.0.0', 'Windows Windows 10', '', '登录成功', '1', '2022-11-10 15:31:54', '2022-11-10 15:31:54');
INSERT INTO `b5net_login_log` VALUES (113662168980787200, 'test3', '127.0.0.1', '  ', ' 11.2.5170.400', ' ', '', '登录成功', '1', '2022-11-10 15:32:52', '2022-11-10 15:32:52');

-- ----------------------------
-- Table structure for b5net_menu
-- ----------------------------
DROP TABLE IF EXISTS `b5net_menu`;
CREATE TABLE `b5net_menu`  (
  `id` bigint NOT NULL COMMENT '菜单ID',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父菜单ID',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求地址',
  `target` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '打开方式（0页签 1新窗口）',
  `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `note` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '菜单状态（1显示 0隐藏）',
  `is_refresh` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '是否刷新（0不刷新 1刷新）',
  `list_sort` int NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `perms`(`perms`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_menu
-- ----------------------------
INSERT INTO `b5net_menu` VALUES (1, '系统管理', 0, '', '0', 'M', '', 'fa fa-cog', '系统管理', '1', '0', 10, '2021-01-03 07:25:11', '2022-10-12 16:25:29');
INSERT INTO `b5net_menu` VALUES (2, '权限管理', 0, '', '0', 'M', '', 'fa fa-id-card-o', '权限管理', '1', '0', 20, '2021-01-03 07:25:11', '2022-03-20 16:00:10');
INSERT INTO `b5net_menu` VALUES (3, '系统工具', 0, '', '0', 'M', '', 'fa fa-cloud', '', '1', '0', 30, '2021-07-29 20:28:41', '2022-03-20 15:59:55');
INSERT INTO `b5net_menu` VALUES (90, '官方网站', 0, 'http://www.b5net.com', '1', 'C', '', 'fa fa-send', '官方网站', '1', '0', 99, '2021-01-05 12:05:30', '2022-10-13 15:28:23');
INSERT INTO `b5net_menu` VALUES (100, '人员管理', 2, 'system/admin/index', '0', 'C', 'system:admin:index', 'fa fa-user-o', '人员管理', '1', '0', 1, '2021-01-03 07:25:11', '2022-03-20 16:02:24');
INSERT INTO `b5net_menu` VALUES (101, '角色管理', 2, 'system/role/index', '0', 'C', 'system:role:index', 'fa fa-address-book-o', '角色管理', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (102, '组织架构', 2, 'system/struct/index', '0', 'C', 'system:struct:index', 'fa fa-sitemap', '组织架构', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (103, '菜单管理', 2, 'system/menu/index', '0', 'C', 'system:menu:index', 'fa fa-server', '菜单管理', '1', '0', 4, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (104, '登录日志', 2, 'system/login_log/index', '0', 'C', 'system:login_log:index', 'fa fa-paw', '登录日志', '1', '0', 5, '2021-01-03 07:25:11', '2021-01-07 12:54:43');
INSERT INTO `b5net_menu` VALUES (105, '参数配置', 1, 'system/config/index', '0', 'C', 'system:config:index', 'fa fa-sliders', '参数配置', '1', '0', 1, '2021-01-03 07:25:11', '2021-01-05 12:20:56');
INSERT INTO `b5net_menu` VALUES (106, '网站设置', 1, 'system/config/site', '0', 'C', 'system:config:site', 'fa fa-object-group', '网站设置', '1', '0', 0, '2021-01-11 22:17:31', '2021-01-11 22:39:46');
INSERT INTO `b5net_menu` VALUES (107, '通知公告', 1, 'system/notice/index', '0', 'C', 'system:notice:index', 'fa fa-bullhorn', '通知公告', '1', '0', 10, '2021-01-03 07:25:11', '2021-03-17 14:05:34');
INSERT INTO `b5net_menu` VALUES (108, '岗位管理', 1, 'system/position/index', '0', 'C', 'system:position:index', '', '', '1', '0', 2, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (140, '实例演示', 3, '', '0', 'M', '', '', '', '1', '0', 0, '2022-10-29 18:22:20', '2022-10-29 18:22:20');
INSERT INTO `b5net_menu` VALUES (150, '代码生成', 3, 'demo/gen/index', '0', 'C', 'demo:gen:index', '', '', '1', '0', 3, '2021-07-29 20:29:15', '2022-04-12 13:40:34');
INSERT INTO `b5net_menu` VALUES (151, '表单构建', 3, 'demo/build/index', '0', 'C', 'demo:build:index', '', '', '1', '0', 2, '2021-07-29 20:29:15', '2022-10-23 14:33:40');
INSERT INTO `b5net_menu` VALUES (152, '图片操作', 140, 'demo/media/index', '0', 'C', 'demo:media:index', '', '', '1', '0', 1, '2021-07-29 20:29:15', '2022-10-29 18:23:43');
INSERT INTO `b5net_menu` VALUES (153, '测试数据权限', 3, 'demo/test_info/index', '0', 'C', 'demo:test_info:index', '', '', '1', '0', 4, '2022-11-01 20:05:52', '2022-11-01 20:05:57');
INSERT INTO `b5net_menu` VALUES (10001, '用户修改', 100, '', '0', 'F', 'system:admin:edit', '', '用户修改', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10002, '用户删除', 100, '', '0', 'F', 'system:admin:drop', '', '用户删除', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10004, '用户状态', 100, '', '0', 'F', 'system:admin:set_status', '', '用户状态', '1', '0', 4, '2021-01-03 07:25:11', '2021-01-08 10:47:09');
INSERT INTO `b5net_menu` VALUES (10100, '角色新增', 101, '', '0', 'F', 'system:role:add', '', '角色新增', '1', '0', 1, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10101, '角色修改', 101, '', '0', 'F', 'system:role:edit', '', '角色修改', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10102, '角色删除', 101, '', '0', 'F', 'system:role:drop', '', '角色删除', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10104, '角色状态', 101, '', '0', 'F', 'system:role:set_status', '', '角色状态', '1', '0', 4, '2021-01-03 07:25:11', '2021-01-08 10:47:31');
INSERT INTO `b5net_menu` VALUES (10105, '菜单授权', 101, '', '0', 'F', 'system:role:auth', '', '菜单授权', '1', '0', 10, '2021-01-03 07:25:11', '2021-01-07 13:32:41');
INSERT INTO `b5net_menu` VALUES (10106, '数据权限', 101, '', '0', 'F', 'system:role:data_scope', '', '数据权限', '1', '0', 11, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10200, '组织新增', 102, '', '0', 'F', 'system:struct:add', '', '组织新增', '1', '0', 1, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10201, '组织修改', 102, '', '0', 'F', 'system:struct:edit', '', '组织修改', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10202, '组织删除', 102, '', '0', 'F', 'system:struct:drop', '', '组织删除', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10300, '菜单新增', 103, '', '0', 'F', 'system:menu:add', '', '菜单新增', '1', '0', 1, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10301, '菜单修改', 103, '', '0', 'F', 'system:menu:edit', '', '菜单修改', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10302, '菜单删除', 103, '', '0', 'F', 'system:menu:drop', '', '菜单删除', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10400, '日志删除', 104, '', '0', 'F', 'system:loginlog:drop', '', '日志删除', '1', '0', 0, '2021-01-07 13:03:15', '2021-01-07 13:03:15');
INSERT INTO `b5net_menu` VALUES (10401, '日志清空', 104, '', '0', 'F', 'system:loginlog:trash', '', '日志清空', '1', '0', 0, '2021-01-07 13:04:06', '2021-01-07 13:04:06');
INSERT INTO `b5net_menu` VALUES (10500, '参数新增', 105, '', '0', 'F', 'system:config:add', '', '参数新增', '1', '0', 1, '2021-01-03 07:25:11', '2021-01-05 06:00:02');
INSERT INTO `b5net_menu` VALUES (10501, '参数修改', 105, '', '0', 'F', 'system:config:edit', '', '参数修改', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-05 06:00:25');
INSERT INTO `b5net_menu` VALUES (10502, '参数删除', 105, '', '0', 'F', 'system:config:drop', '', '参数删除', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-05 06:00:59');
INSERT INTO `b5net_menu` VALUES (10503, '参数批量删除', 105, '', '0', 'F', 'system:config:drop_all', '', '参数批量删除', '1', '0', 4, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10504, '清除缓存', 105, '', '0', 'F', 'system:config:del_cache', '', '清除缓存', '1', '0', 5, '2021-01-03 07:25:11', '2021-01-08 10:46:47');
INSERT INTO `b5net_menu` VALUES (10700, '公告新增', 107, '', '0', 'F', 'system:notice:add', '', '公告新增', '1', '0', 1, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10701, '公告修改', 107, '', '0', 'F', 'system:notice:edit', '', '公告修改', '1', '0', 2, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10702, '公告删除', 107, '', '0', 'F', 'system:notice:drop', '', '公告删除', '1', '0', 3, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10703, '公告批量删除', 107, '', '0', 'F', 'system:notice:drop_all', '', '公告批量删除', '1', '0', 4, '2021-01-03 07:25:11', '2021-01-03 07:25:11');
INSERT INTO `b5net_menu` VALUES (10801, '添加岗位', 108, '', '0', 'F', 'system:position:index', '', '', '1', '0', 1, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (10802, '编辑岗位', 108, '', '0', 'F', 'system:position:add', '', '', '1', '0', 2, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (10803, '删除岗位', 108, '', '0', 'F', 'system:position:drop_all', '', '', '1', '0', 3, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15201, '图片添加', 152, '', '0', 'F', 'demo:media:add', '', '', '1', '0', 1, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15202, '图片编辑', 152, '', '0', 'F', 'demo:media:edit', '', '', '1', '0', 2, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15203, '图片删除', 152, '', '0', 'F', 'demo:mediadrop', '', '', '1', '0', 3, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15204, '图片批量删除', 152, '', '0', 'F', 'demo:media:drop_all', '', '', '1', '0', 4, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15301, '新增', 153, '', '0', 'F', 'demo:test_info:add', '', '', '1', '0', 0, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15302, '编辑', 153, '', '0', 'F', 'demo:test_info:edit', '', '', '1', '0', 0, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15303, '删除', 153, '', '0', 'F', 'demo:test_info:drop', '', '', '1', '0', 0, NULL, NULL);
INSERT INTO `b5net_menu` VALUES (15304, '批量删除', 153, '', '0', 'F', 'demo:test_info:drop_all', '', '', '1', '0', 0, NULL, NULL);

-- ----------------------------
-- Table structure for b5net_notice
-- ----------------------------
DROP TABLE IF EXISTS `b5net_notice`;
CREATE TABLE `b5net_notice`  (
  `id` bigint NOT NULL COMMENT '公告ID',
  `title` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公告标题',
  `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '公告类型（1通知 2公告）',
  `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公告内容',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '公告状态（1正常 0关闭）',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '通知公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_notice
-- ----------------------------
INSERT INTO `b5net_notice` VALUES (1, '【公告】： B5LaravelCMF新版本发布啦', '2', NULL, '<p>新版本内容</p><p><img src=\"http://127.0.0.1:8080/uploads/editor/2022/10/31/a91878e315bcc0748a3516a51769fcf9.jpg\" style=\"width: 500px;\" data-filename=\"u=3671441873,259090506&amp;fm=26&amp;gp=0.jpg\"><br></p><p>新版本内容</p><p>新版本内容</p><p>新版本内容</p><p><br></p>', '1', '2022-03-12 11:33:42', '2022-10-31 19:14:28');
INSERT INTO `b5net_notice` VALUES (2, '【通知】：B5LaravelCMF系统凌晨维护', '1', NULL, '<p><font color=\"#0000ff\">维护内容</font></p><p><img src=\"http://127.0.0.1:8080/uploads/editor/2022/10/31/14ab608de35646b5ca1967ca9a680ceb.jpg\" style=\"width: 500px;\" data-filename=\"下载.jpg\"><font color=\"#0000ff\"><br></font></p>', '1', '2022-03-20 11:33:42', '2022-10-31 19:14:39');

-- ----------------------------
-- Table structure for b5net_position
-- ----------------------------
DROP TABLE IF EXISTS `b5net_position`;
CREATE TABLE `b5net_position`  (
  `id` bigint UNSIGNED NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
  `pos_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位标识',
  `list_sort` int NOT NULL DEFAULT 100 COMMENT '排序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '状态',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_position
-- ----------------------------
INSERT INTO `b5net_position` VALUES (1, '总经理', 'ceo', 1, '1', '', '2022-04-04 23:04:49', '2022-10-13 15:27:40');
INSERT INTO `b5net_position` VALUES (2, '部门经理', 'cpo', 2, '1', '', '2022-04-04 23:25:34', '2022-10-28 15:00:07');
INSERT INTO `b5net_position` VALUES (3, '组长', 'cgo', 3, '1', '', '2022-04-04 23:26:08', '2022-04-08 12:53:33');
INSERT INTO `b5net_position` VALUES (4, '员工', 'user', 4, '1', '12', '2022-04-04 23:26:50', '2022-10-19 18:14:03');

-- ----------------------------
-- Table structure for b5net_role
-- ----------------------------
DROP TABLE IF EXISTS `b5net_role`;
CREATE TABLE `b5net_role`  (
  `id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `role_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色权限字符串',
  `data_scope` int NOT NULL DEFAULT 0 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `list_sort` int NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '角色状态（1正常 0停用）',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `rolekey`(`role_key`) USING BTREE,
  INDEX `listsort`(`list_sort`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_role
-- ----------------------------
INSERT INTO `b5net_role` VALUES (1, '超级管理员', 'administrator', 0, 1, '1', '超级管理员', '2020-12-28 07:42:31', '2022-10-14 11:18:52');
INSERT INTO `b5net_role` VALUES (104682922295955456, '员工角色', 'test', 16, 2, '1', '只能看自己信息', '2022-10-16 20:52:33', '2022-11-10 11:39:04');
INSERT INTO `b5net_role` VALUES (113603187126046720, '部门领导', 'dept_leader', 2, 1, '1', '本部门及以下部门', '2022-11-10 11:38:30', '2022-11-10 11:38:59');

-- ----------------------------
-- Table structure for b5net_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `b5net_role_menu`;
CREATE TABLE `b5net_role_menu`  (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `menu_id` bigint NOT NULL COMMENT '菜单ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_role_menu
-- ----------------------------
INSERT INTO `b5net_role_menu` VALUES (104682922295955456, 3);
INSERT INTO `b5net_role_menu` VALUES (104682922295955456, 153);
INSERT INTO `b5net_role_menu` VALUES (104682922295955456, 15301);
INSERT INTO `b5net_role_menu` VALUES (104682922295955456, 15302);
INSERT INTO `b5net_role_menu` VALUES (104682922295955456, 15303);
INSERT INTO `b5net_role_menu` VALUES (104682922295955456, 15304);
INSERT INTO `b5net_role_menu` VALUES (113603187126046720, 3);
INSERT INTO `b5net_role_menu` VALUES (113603187126046720, 153);
INSERT INTO `b5net_role_menu` VALUES (113603187126046720, 15301);
INSERT INTO `b5net_role_menu` VALUES (113603187126046720, 15302);
INSERT INTO `b5net_role_menu` VALUES (113603187126046720, 15303);
INSERT INTO `b5net_role_menu` VALUES (113603187126046720, 15304);

-- ----------------------------
-- Table structure for b5net_role_struct
-- ----------------------------
DROP TABLE IF EXISTS `b5net_role_struct`;
CREATE TABLE `b5net_role_struct`  (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `struct_id` bigint NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `struct_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for b5net_struct
-- ----------------------------
DROP TABLE IF EXISTS `b5net_struct`;
CREATE TABLE `b5net_struct`  (
  `id` bigint UNSIGNED NOT NULL COMMENT '部门id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组织类型',
  `parent_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父部门id',
  `levels` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '祖级列表',
  `list_sort` int NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '部门状态（1正常 0停用）',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织架构' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of b5net_struct
-- ----------------------------
INSERT INTO `b5net_struct` VALUES (100, '冰舞科技', 'group', '', 0, '0', 0, '冰舞', '15888888888', '', '1', '2020-12-24 11:33:42', '2022-10-16 20:30:40');
INSERT INTO `b5net_struct` VALUES (104677500600193024, '山东总公司', 'com', '冰舞科技', 100, '0,100', 1, '哈哈', '13333333333', '', '1', '2022-10-16 20:31:00', '2022-10-16 20:31:00');
INSERT INTO `b5net_struct` VALUES (104677561774116864, '江苏分公司', 'com', '冰舞科技', 100, '0,100', 2, '', '', '', '1', '2022-10-16 20:31:15', '2022-10-16 20:32:49');
INSERT INTO `b5net_struct` VALUES (104677601469009920, '研发部', 'dep', '冰舞科技,山东总公司', 104677500600193024, '0,100,104677500600193024', 10, '', '', '', '1', '2022-10-16 20:31:25', '2022-10-16 20:31:30');
INSERT INTO `b5net_struct` VALUES (104677679562756096, '财务部', 'dep', '冰舞科技,山东总公司', 104677500600193024, '0,100,104677500600193024', 20, '', '', '', '1', '2022-10-16 20:31:43', '2022-10-16 20:31:48');
INSERT INTO `b5net_struct` VALUES (104677742527647744, '财务部', 'com', '冰舞科技,山东总公司', 104677500600193024, '0,100,104677500600193024', 30, '', '', '', '1', '2022-10-16 20:31:58', '2022-10-16 20:31:58');
INSERT INTO `b5net_struct` VALUES (104677839734837248, '前端开发组', 'team', '冰舞科技,山东总公司,研发部', 104677601469009920, '0,100,104677500600193024,104677601469009920', 10, '', '', '', '1', '2022-10-16 20:32:21', '2022-10-16 20:32:21');
INSERT INTO `b5net_struct` VALUES (104677894931877888, '后端开发组', 'team', '冰舞科技,山东总公司,研发部', 104677601469009920, '0,100,104677500600193024,104677601469009920', 10, '', '', '', '1', '2022-10-16 20:32:34', '2022-10-16 20:32:34');

-- ----------------------------
-- Table structure for demo_media
-- ----------------------------
DROP TABLE IF EXISTS `demo_media`;
CREATE TABLE `demo_media`  (
  `id` bigint UNSIGNED NOT NULL,
  `img` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单图',
  `imgs` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '多图',
  `crop` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '裁剪图片',
  `video` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '视频',
  `file` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单文件',
  `files` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '多文件',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_media
-- ----------------------------
INSERT INTO `demo_media` VALUES (2, '/uploads/demo/2022/10/31/614b057dd1aac2a0f49154b8881e6c88.jpeg', '/uploads/demo/2022/10/31/641b86c83cce91834a3f53292e41b0a9.jpeg,/uploads/demo/2022/10/31/fa0ca4f53ebe257d7267c5a37b8aff9f.jpeg,/uploads/demo/2022/10/31/bea312b6fb3eef15ef6aa6bed8fe9d2c.jpeg,/uploads/demo/2022/10/31/89fb936dea6eddb1e02abbf960ccfcf6.jpg,/uploads/demo/2022/10/31/01fd4f2bb8b59ca925c838afeb7a3092.jpeg', '/uploads/demo/2022/10/31/a2910bad9e5cfa1be90f97a307515394.jpg', '/uploads/demo/2022/10/31/c09283b2e06fa7103b07bc8076487e0a.mp4', '/uploads/demo/2022/10/31/08a4e005e5854351f7f59cd2e97bfdb7.txt', '/uploads/demo/2022/10/31/3de0bab44d59737430c3cea85b9a457e.txt', '2022-04-13 07:18:02', '2022-10-31 19:28:00');
INSERT INTO `demo_media` VALUES (110097978957500416, '/uploads/demo/2022/10/31/62089a568d79f346f4a52bd06de918a1.jpeg', '/uploads/demo/2022/10/31/b2302f2f617603c7c1cf1e759dde5b69.jpeg,/uploads/demo/2022/10/31/e8af31c3221dd17ac04f22fcd4024ffd.jpeg,/uploads/demo/2022/10/31/acbf99d09620334b9da429b5b8db49bc.jpeg', '/uploads/demo/2022/10/31/e6ba26b8b848ee5b2017fa02b08d220f.jpg,/uploads/demo/2022/10/31/770d205aea973678c865da1f15682b56.jpg', '/uploads/demo/2022/10/31/df9fb3174a68b1364fb682ccc430618c.mp4', '/uploads/demo/2022/10/31/38b816236f1358e5319691eef239b982.png', '/uploads/demo/2022/10/31/c2d3d615526cc12b4d1075b63c2dca76.png', '2022-10-31 19:30:03', '2022-11-01 12:36:23');

-- ----------------------------
-- Table structure for test_info
-- ----------------------------
DROP TABLE IF EXISTS `test_info`;
CREATE TABLE `test_info`  (
  `id` bigint NOT NULL,
  `struct_id` bigint NOT NULL DEFAULT 0 COMMENT '组织ID',
  `struct_levels` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组织levels',
  `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '状态',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '介绍',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试数据权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of test_info
-- ----------------------------
INSERT INTO `test_info` VALUES (1, 104677601469009920, '0,100,104677500600193024', 104683566088065024, '研发部数据1', '1', '阿萨大1', '2022-11-10 11:33:43', '2022-11-10 15:14:31');
INSERT INTO `test_info` VALUES (2, 104677601469009920, '0,100,104677500600193024', 104683566088065024, '前端数据1', '1', 'aaaa', '2022-11-10 11:34:33', '2022-11-10 15:14:30');
INSERT INTO `test_info` VALUES (3, 104677601469009920, '0,100,104677500600193024', 104683566088065024, '后端数据11', '1', '222', '2022-11-10 15:14:28', '2022-11-10 15:14:28');
INSERT INTO `test_info` VALUES (4, 104677679562756096, '0,100,104677500600193024', 113602602515566592, '财务数据1', '1', '333', '2022-11-10 15:14:28', '2022-11-10 15:32:59');
INSERT INTO `test_info` VALUES (5, 100, '0', 10000, '根组织数据', '1', '44', '2022-11-10 15:14:28', '2022-11-10 15:14:28');
INSERT INTO `test_info` VALUES (113640979554111488, 104677601469009920, '0,100,104677500600193024', 104683566088065024, 'test1前端数据222', '1', 's大苏打撒旦', '2022-11-10 14:08:40', '2022-11-10 15:14:26');

SET FOREIGN_KEY_CHECKS = 1;
