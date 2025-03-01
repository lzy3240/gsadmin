/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 80027
Source Host           : localhost:3306
Source Database       : gsadmin

Target Server Type    : MYSQL
Target Server Version : 80027
File Encoding         : 65001

Date: 2024-11-05 09:47:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_auth
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth`;
CREATE TABLE `sys_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(255) DEFAULT NULL,
  `auth_url` varchar(255) DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `pid` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `is_show` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  `power_type` int DEFAULT NULL,
  `is_default` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_auth
-- ----------------------------
INSERT INTO `sys_auth` VALUES ('1', '首页', '/', '1', '0', '1', 'layui-icon layui-icon-console', '1', '1', '0', '1', '0', '0', '2021-05-21 15:16:44', '2021-05-21 15:16:44');
INSERT INTO `sys_auth` VALUES ('2', '系统管理', '/', '1', '0', '97', 'layui-icon layui-icon-app', '1', '1', '0', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('3', '用户管理', '/system/user/list', '1', '2', '1', 'layui-icon fa-user-o', '1', '1', '1', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('4', '角色管理', '/system/role/list', '1', '2', '2', 'fa-user-circle-o', '1', '1', '1', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('5', '新增', '/system/user/add', '1', '3', '1', '', '0', '1', '2', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('6', '修改', '/system/user/edit', '1', '3', '2', '', '0', '1', '2', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('7', '删除', '/system/user/delete', '1', '3', '3', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('8', '新增', '/system/role/add', '1', '4', '1', '', '1', '1', '2', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('9', '修改', '/system/role/edit', '1', '4', '2', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('10', '删除', '/system/role/delete', '1', '4', '3', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('11', '权限管理', '/system/auth/list', '1', '2', '3', 'layui-icon fa-list', '1', '1', '1', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('12', '新增', '/system/auth/edit', '1', '11', '1', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('13', '修改', '/system/auth/edit', '1', '11', '2', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('14', '删除', '/system/auth/delete', '1', '11', '3', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('15', '权限树', '/system/auth/nodes', '0', '11', '4', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('16', '单个权限获取', '/system/auth/node', '0', '11', '5', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('17', '用户列表接口', '/system/user/json', '0', '3', '4', 'layui-icon ', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('18', '角色列表', '/system/role/json', '0', '4', '4', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('19', '仪表盘', '/system/main', '1', '1', '1', 'layui-icon layui-icon-rate', '1', '1', '1', '1', '0', '0', '2021-05-21 15:22:36', '2021-05-21 15:22:36');
INSERT INTO `sys_auth` VALUES ('20', '个人信息', 'profile/edit', '1', '0', '98', 'layui-icon layui-icon-user', '1', '1', '0', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('21', '资料修改', '/system/user/profile', '1', '20', '1', 'fa-edit', '1', '1', '1', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('22', '更新头像', '/system/user/avatar', '0', '21', '2', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('23', '修改密码', '/system/user/pwd', '0', '21', '3', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30');
INSERT INTO `sys_auth` VALUES ('24', '头像上传', '/system/user/uploadPage', '0', '21', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 12:03:00', '2024-10-17 12:03:00');
INSERT INTO `sys_auth` VALUES ('25', '文件上传', '/system/upload', '0', '21', '4', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 12:11:55', '2024-10-17 12:11:55');
INSERT INTO `sys_auth` VALUES ('26', '网站设置', '/system/site/edit', '0', '2', '5', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2023-07-11 13:37:09', '2023-07-11 13:37:09');
INSERT INTO `sys_auth` VALUES ('27', '服务监控', '/monitor/server', '0', '2', '6', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2025-02-27 12:32:52', '2025-02-27 12:33:03');
INSERT INTO `sys_auth` VALUES ('30', '日志管理', '/', '0', '0', '99', 'layui-icon layui-icon-form', '1', '1', '0', '1', '0', '1', '2021-10-09 18:48:25', '2024-11-03 18:00:24');
INSERT INTO `sys_auth` VALUES ('31', '日志列表', '/system/log/list', '0', '30', '1', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2021-10-09 18:49:08', '2021-10-09 18:49:08');
INSERT INTO `sys_auth` VALUES ('32', '操作日志', '/system/log/operate', '0', '31', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2021-10-09 18:50:07', '2021-10-09 18:50:07');
INSERT INTO `sys_auth` VALUES ('33', '登陆日志', '/system/log/login', '0', '31', '3', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2021-10-09 18:50:44', '2021-10-09 18:50:44');
INSERT INTO `sys_auth` VALUES ('40', '字典管理', '/system/dict/type/list', '0', '2', '4', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2024-10-17 10:45:39', '2024-10-17 10:45:39');
INSERT INTO `sys_auth` VALUES ('41', '新增', '/system/dict/type/add', '0', '40', '1', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 10:47:38', '2024-10-17 10:47:38');
INSERT INTO `sys_auth` VALUES ('42', '修改', '/system/dict/type/edit', '0', '40', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 10:48:14', '2024-10-17 10:48:14');
INSERT INTO `sys_auth` VALUES ('43', '删除', '/system/dict/type/delete', '0', '40', '3', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 10:48:39', '2024-10-17 10:48:39');
INSERT INTO `sys_auth` VALUES ('44', '字典类型接口', '/system/dict/type/json', '0', '40', '4', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-21 10:02:06', '2024-10-21 10:02:06');
INSERT INTO `sys_auth` VALUES ('45', '字典值接口', '/system/dict/data/json', '0', '40', '5', 'layui-icon ', '1', '1', '2', '1', '0', '1', '2024-10-17 10:48:39', '2024-10-30 21:39:18');
INSERT INTO `sys_auth` VALUES ('50', '消息中心', '/', '0', '0', '2', 'layui-icon layui-icon-note', '1', '1', '0', '1', '0', '0', '2024-10-21 15:28:00', '2024-10-21 15:28:00');
INSERT INTO `sys_auth` VALUES ('51', '消息列表', '/system/notice/list', '0', '50', '1', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2024-10-21 15:29:29', '2024-10-21 15:29:29');
INSERT INTO `sys_auth` VALUES ('52', '消息确认', '/system/notice/status', '0', '51', '2', 'layui-icon ', '1', '1', '2', '1', '0', '1', '2024-10-21 15:30:33', '2024-10-29 21:22:34');
INSERT INTO `sys_auth` VALUES ('53', '消息列表接口', '/system/notice/json', '0', '51', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-23 11:57:51', '2024-10-23 11:57:51');
INSERT INTO `sys_auth` VALUES ('54', '未读消息接口', '/system/notice', '0', '51', '4', 'layui-icon ', '1', '1', '2', '1', '0', '1', '2024-10-23 11:57:51', '2024-10-29 21:22:54');
INSERT INTO `sys_auth` VALUES ('55', '消息查看', '/system/notice/edit', '0', '51', '1', 'layui-icon ', '1', '1', '2', '0', '1', '1', '2024-10-29 21:22:23', '2024-10-29 21:22:23');

-- ----------------------------
-- Table structure for sys_conf
-- ----------------------------
DROP TABLE IF EXISTS `sys_conf`;
CREATE TABLE `sys_conf` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) DEFAULT NULL,
  `info` varchar(3000) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_conf
-- ----------------------------
INSERT INTO `sys_conf` VALUES ('1', 'base_conf', '{\"colors\": [{\"color\": \"#FF0000\",\"id\": \"1\"},{\"color\": \"#5FB878\",\"id\": \"2\"},{\"color\": \"#1E9FFF\",\"id\": \"3\"},{\"color\": \"#FFB800\",\"id\": \"4\"},{\"color\": \"darkgray\",\"id\": \"5\"}],\"header\": {\"message\": \"/system/notice\"},\"links\": [{\"href\": \"http://www.baidu.com\",\"icon\": \"layui-icon layui-icon-auz\",\"title\": \"官方网站\"},{\"href\": \"https://github.com/lzy3240/gsadmin\",\"icon\": \"layui-icon layui-icon-auz\",\"title\": \"开发文档\"},{\"href\": \"https://github.com/lzy3240/gsadmin\",\"icon\": \"layui-icon layui-icon-auz\",\"title\": \"开源地址\"}],\"logo\": {\"image\": \"/static/admin/images/logo.png\",\"title\": \"gsadmin\"},\"menu\": {\"accordion\": true,\"control\": false,\"data\": \"/system/menu\",\"method\": \"GET\",\"select\": \"60\"},\"other\": {\"autoHead\": false,\"keepLoad\": 100},\"tab\": {\"index\": {\"href\": \"/system/main\",\"id\": \"60\",\"title\": \"首页\"},\"keepState\": true,\"muiltTab\": true,\"tabMax\": 30},\"theme\": {\"allowCustom\": true,\"defaultColor\": \"2\",\"defaultMenu\": \"dark-theme\"}}', '1', null, null, '2021-05-28 10:48:35', '2021-05-28 10:48:35');
INSERT INTO `sys_conf` VALUES ('2', 'site_conf', '{\"web_name\":\"gsadmin\",\"web_url\":\"http://123/\",\"logo_url\":\"http://123/\",\"key_words\":\"admin\",\"description\":\"11223344\",\"copyright\":\"456542\",\"icp\":\"789873\",\"site_status\":1}', '1', null, null, '2021-05-28 10:48:35', '2024-11-03 19:51:28');

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `dict_type` varchar(255) DEFAULT NULL,
  `dict_label` varchar(255) DEFAULT NULL,
  `dict_value` varchar(255) DEFAULT NULL,
  `dict_sort` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `is_default` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES ('1', 'sys_user_sex', '男', '0', '1', '1', '性别男', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('2', 'sys_user_sex', '女', '1', '2', '1', '性别女', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('3', 'sys_user_sex', '未知', '2', '3', '1', '性别未知', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('4', 'sys_show_hide', '显示', '0', '1', '1', '显示菜单', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('5', 'sys_show_hide', '隐藏', '1', '2', '1', '隐藏菜单', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('6', 'sys_normal_disable', '正常', '1', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('7', 'sys_normal_disable', '停用', '0', '2', '1', '停用状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('8', 'sys_job_status', '正常', '0', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('9', 'sys_job_status', '暂停', '1', '2', '1', '停用状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('10', 'sys_job_group', '默认', 'DEFAULT', '1', '1', '默认分组', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('11', 'sys_job_group', '系统', 'SYSTEM', '2', '1', '系统分组', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('12', 'sys_yes_no', '是', 'Y', '1', '1', '系统默认是', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('13', 'sys_yes_no', '否', 'N', '2', '1', '系统默认否', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('14', 'sys_notice_type', '通知', '1', '1', '1', '通知', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('15', 'sys_notice_type', '公告', '2', '2', '1', '公告', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('16', 'sys_notice_status', '正常', '0', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('17', 'sys_notice_status', '关闭', '1', '2', '1', '关闭状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('18', 'sys_oper_type', '其他', '0', '99', '1', '其他操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('19', 'sys_oper_type', '新增', '1', '1', '1', '新增操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('20', 'sys_oper_type', '修改', '2', '2', '1', '修改操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('21', 'sys_oper_type', '删除', '3', '3', '1', '删除操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('22', 'sys_oper_type', '授权', '4', '4', '1', '授权操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('23', 'sys_oper_type', '导出', '5', '5', '1', '导出操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('24', 'sys_oper_type', '导入', '6', '6', '1', '导入操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('25', 'sys_oper_type', '强退', '7', '7', '1', '强退操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('26', 'sys_oper_type', '生成代码', '8', '8', '1', '生成操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('27', 'sys_oper_type', '清空数据', '9', '9', '1', '清空操作', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('28', 'sys_common_status', '成功', '0', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('29', 'sys_common_status', '失败', '1', '2', '1', '停用状态', '0', '0', '0', '2023-05-08 15:56:39', null);
INSERT INTO `sys_dict_data` VALUES ('30', 'sys_notice_read', '已读', '1', '1', '1', '已读状态', '0', '0', '0', '2024-10-22 22:09:59', null);
INSERT INTO `sys_dict_data` VALUES ('31', 'sys_notice_read', '未读', '0', '2', '1', '未读状态', '0', '0', '0', '2024-10-22 22:10:02', null);
INSERT INTO `sys_dict_data` VALUES ('32', 'sys_audit_status', '待审核11', '0', '1', '1', '待审核', '0', '0', '1', '2024-11-04 10:09:54', '2024-11-04 16:22:53');
INSERT INTO `sys_dict_data` VALUES ('33', 'sys_audit_status', '通过', '1', '2', '1', '审核通过', '0', '0', '0', '2024-11-04 10:11:16', null);
INSERT INTO `sys_dict_data` VALUES ('34', 'sys_audit_status', '不通过', '2', '3', '1', '审核不通过', '0', '0', '0', '2024-11-04 10:17:07', null);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `dict_type` varchar(255) DEFAULT NULL,
  `dict_name` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `is_default` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES ('1', 'sys_user_sex', '用户性别', '1', '用户性别列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('2', 'sys_show_hide', '菜单状态', '1', '菜单状态列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('3', 'sys_normal_disable', '系统开关', '1', '系统开关列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('4', 'sys_job_status', '任务状态', '1', '任务状态列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('5', 'sys_job_group', '任务分组', '1', '任务分组列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('6', 'sys_yes_no', '系统是否', '1', '系统是否列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('7', 'sys_notice_type', '通知类型', '1', '通知类型列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('8', 'sys_notice_status', '通知状态', '1', '通知状态列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('9', 'sys_oper_type', '操作类型', '1', '操作类型列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('10', 'sys_common_status', '系统状态', '1', '登录状态列表', '1', '0', '0', '2023-05-08 15:56:38', null);
INSERT INTO `sys_dict_type` VALUES ('11', 'sys_audit_status', '审核状态', '1', '审核状态列表', '1', '0', '1', '2023-03-23 09:46:42', null);
INSERT INTO `sys_dict_type` VALUES ('12', 'sys_notice_read', '通知已读', '1', '通知已读列表', '1', '0', '0', '2023-03-23 09:46:42', null);

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `login_name` varchar(255) DEFAULT NULL,
  `ip_addr` varchar(255) DEFAULT NULL,
  `login_location` varchar(255) DEFAULT NULL,
  `browser` varchar(255) DEFAULT NULL,
  `os` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES ('1', 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '1', '登陆成功', '2024-11-04 16:23:15');
INSERT INTO `sys_login_log` VALUES ('2', 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '1', '登陆成功', '2024-11-05 09:08:04');

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `notice_type` varchar(255) DEFAULT NULL,
  `notice_title` varchar(255) DEFAULT NULL,
  `notice_context` varchar(255) DEFAULT NULL,
  `notice_from` varchar(255) DEFAULT NULL,
  `notice_to` int unsigned DEFAULT NULL,
  `status` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `business_type` int DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  `request_method` varchar(255) DEFAULT NULL,
  `operator_type` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `oper_name` varchar(255) DEFAULT NULL,
  `dept_name` varchar(255) DEFAULT NULL,
  `oper_url` varchar(255) DEFAULT NULL,
  `oper_ip` varchar(255) DEFAULT NULL,
  `oper_location` varchar(255) DEFAULT NULL,
  `oper_param` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `json_result` varchar(255) DEFAULT NULL,
  `latency_time` int DEFAULT NULL,
  `error_msg` varchar(255) DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  `oper_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES ('1', '修改用户', '2', 'POST', 'POST', '1', '1', 'admin', '', '/system/user/status', '127.0.0.1', '内网IP', '{\"form\":{\"id\":[\"2\"],\"status\":[\"0\"]}}', '1', '{\"code\":200,\"msg\":\"更新成功\",\"data\":null,\"type\":2,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:16:01');
INSERT INTO `sys_oper_log` VALUES ('2', '修改用户', '2', 'POST', 'POST', '1', '1', 'admin', '', '/system/user/status', '127.0.0.1', '内网IP', '{\"form\":{\"id\":[\"2\"],\"status\":[\"1\"]}}', '1', '{\"code\":200,\"msg\":\"更新成功\",\"data\":null,\"type\":2,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:16:02');
INSERT INTO `sys_oper_log` VALUES ('3', '修改用户', '2', 'POST', 'POST', '1', '1', 'admin', '', '/system/user/edit', '127.0.0.1', '内网IP', '{\"form\":{\"email\":[\"ceshi@qq.com\"],\"id\":[\"2\"],\"login_name\":[\"user\"],\"phone\":[\"18912345678\"],\"real_name\":[\"user\"],\"role_ids\":[\"2\"],\"status\":[\"on\"]}}', '1', '{\"code\":200,\"msg\":\"更新成功\",\"data\":null,\"type\":2,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:16:45');
INSERT INTO `sys_oper_log` VALUES ('4', '修改用户', '2', 'POST', 'POST', '1', '1', 'admin', '', '/system/user/edit', '127.0.0.1', '内网IP', '{\"form\":{\"email\":[\"2222@qq.com\"],\"id\":[\"1\"],\"login_name\":[\"admin\"],\"phone\":[\"13000000000\"],\"real_name\":[\"admin\"],\"role_ids\":[\"1\"],\"status\":[\"on\"]}}', '1', '{\"code\":200,\"msg\":\"更新成功\",\"data\":null,\"type\":2,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:16:51');
INSERT INTO `sys_oper_log` VALUES ('5', '修改用户', '2', 'POST', 'POST', '1', '1', 'admin', '', '/system/user/status', '127.0.0.1', '内网IP', '{\"form\":{\"id\":[\"2\"],\"status\":[\"0\"]}}', '1', '{\"code\":200,\"msg\":\"更新成功\",\"data\":null,\"type\":2,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:18:39');
INSERT INTO `sys_oper_log` VALUES ('6', '修改用户', '2', 'POST', 'POST', '1', '1', 'admin', '', '/system/user/status', '127.0.0.1', '内网IP', '{\"form\":{\"id\":[\"2\"],\"status\":[\"1\"]}}', '1', '{\"code\":200,\"msg\":\"更新成功\",\"data\":null,\"type\":2,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:18:46');
INSERT INTO `sys_oper_log` VALUES ('7', '用户登陆', '0', 'POST', 'POST', '1', '1', 'admin', '', '/login', '127.0.0.1', '内网IP', 'null', '1', '{\"code\":200,\"msg\":\"登陆成功\",\"data\":null,\"type\":0,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-04 16:23:14');
INSERT INTO `sys_oper_log` VALUES ('8', '用户登陆', '0', 'POST', 'POST', '1', '1', 'admin', '', '/login', '127.0.0.1', '内网IP', 'null', '1', '{\"code\":200,\"msg\":\"登陆成功\",\"data\":null,\"type\":0,\"count\":0}', '','', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', '2024-11-05 09:08:04');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) DEFAULT NULL,
  `detail` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '管理员', '拥有超级管理权限', '1', '0', '0', '2020-09-28 14:00:10', '2020-09-28 14:00:06');
INSERT INTO `sys_role` VALUES ('2', '普通用户', '拥有普通用户权限', '1', '0', '0', '2020-09-28 13:59:53', '2024-11-05 09:46:33');

-- ----------------------------
-- Table structure for sys_role_auth
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_auth`;
CREATE TABLE `sys_role_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` bigint unsigned DEFAULT NULL,
  `auth_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_role_auth
-- ----------------------------
INSERT INTO `sys_role_auth` VALUES ('1', '2', '1');
INSERT INTO `sys_role_auth` VALUES ('2', '2', '2');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(255) DEFAULT NULL,
  `real_name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `level` int DEFAULT NULL,
  `role_ids` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `last_ip` varchar(255) DEFAULT NULL,
  `last_login` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `create_id` int DEFAULT NULL,
  `update_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'admin', 'admin', '63b11d6a04a9d76c6c8134bf99a31306', '99', '1', '13000000000', '2222@qq.com', '/runtime/upload/images/20241101/1730447790000.jpg', '超级管理员', 'JsnOHR5tZk', '127.0.0.1', '2024-11-05 09:08:04', '1', '0', '0', '2020-10-14 17:04:30', '2024-11-05 09:08:04');
INSERT INTO `sys_user` VALUES ('2', 'user', 'user', '3fae7af9815b0886a2a95dba8356f589', '1', '2', '18912345678', 'ceshi@qq.com', '/runtime/upload/images/20241030/1730302348000.jpg', '普通用户', 'vbKG0Djx0U', '127.0.0.1', '2024-10-31 20:59:15', '1', '1', '0', '2020-10-14 17:04:30', '2024-11-04 16:18:47');

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online` (
  `session_id` varchar(255) DEFAULT NULL,
  `login_name` varchar(255) DEFAULT NULL,
  `dept_name` varchar(255) DEFAULT NULL,
  `ip_addr` varchar(255) DEFAULT NULL,
  `login_location` varchar(255) DEFAULT NULL,
  `browser` varchar(255) DEFAULT NULL,
  `os` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `start_timestamp` bigint DEFAULT NULL,
  `last_access_time` varchar(255) DEFAULT NULL,
  `expire_time` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES ('1', 'admin', '', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', 'on_line', '1730768884', '2024-11-05 09:08:04', '1440');
