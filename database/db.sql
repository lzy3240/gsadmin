-- name: create-sys-user
INSERT INTO `sys_user` (`id`, `login_name`, `real_name`, `password`, `level`, `role_ids`, `phone`, `email`, `avatar`, `remark`, `salt`, `last_ip`, `last_login`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
    (1, 'admin', 'admin', '63b11d6a04a9d76c6c8134bf99a31306', '99', '', '13000000000', '2222@qq.com', '', '超级管理员', 'JsnOHR5tZk', '127.0.0.1', '2024-11-04 10:03:06', '1', '0', '0', '2020-10-14 17:04:30', '2024-11-04 10:03:07'),
    (2, 'user', 'user', '3fae7af9815b0886a2a95dba8356f589', '1', '', '18912345678', 'ceshi@qq.com', '', '普通用户', 'vbKG0Djx0U', '127.0.0.1', '2024-10-31 20:59:15', '1', '1', '0', '2020-10-14 17:04:30', '2024-11-04 11:08:49');


-- name: create-sys-auth
INSERT INTO `sys_auth` (`id`, `auth_name`, `auth_url`, `user_id`, `pid`, `sort`, `icon`, `is_show`, `status`, `power_type`, `is_default`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
    (1, '首页', '/', '1', '0', '1', 'layui-icon layui-icon-console', '1', '1', '0', '1', '0', '0', '2021-05-21 15:16:44', '2021-05-21 15:16:44'),
    (2, '系统管理', '/', '1', '0', '97', 'layui-icon layui-icon-app', '1', '1', '0', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (3, '用户管理', '/system/user/list', '1', '2', '1', 'layui-icon fa-user-o', '1', '1', '1', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (4, '角色管理', '/system/role/list', '1', '2', '2', 'fa-user-circle-o', '1', '1', '1', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (5, '新增', '/system/user/add', '1', '3', '1', '', '0', '1', '2', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (6, '修改', '/system/user/edit', '1', '3', '2', '', '0', '1', '2', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (7, '删除', '/system/user/delete', '1', '3', '3', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (8, '新增', '/system/role/add', '1', '4', '1', '', '1', '1', '2', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (9, '修改', '/system/role/edit', '1', '4', '2', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (10, '删除', '/system/role/delete', '1', '4', '3', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (11, '权限管理', '/system/auth/list', '1', '2', '3', 'layui-icon fa-list', '1', '1', '1', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (12, '新增', '/system/auth/edit', '1', '11', '1', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (13, '修改', '/system/auth/edit', '1', '11', '2', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (14, '删除', '/system/auth/delete', '1', '11', '3', '', '0', '1', '2', '1', '1', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (15, '权限树', '/system/auth/nodes', '0', '11', '4', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (16, '单个权限获取', '/system/auth/node', '0', '11', '5', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (17, '用户列表接口', '/system/user/json', '0', '3', '4', 'layui-icon ', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (18, '角色列表', '/system/role/json', '0', '4', '4', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (19, '仪表盘', '/system/main', '1', '1', '1', 'layui-icon layui-icon-rate', '1', '1', '1', '1', '0', '0', '2021-05-21 15:22:36', '2021-05-21 15:22:36'),
    (20, '个人信息', 'profile/edit', '1', '0', '98', 'layui-icon layui-icon-user', '1', '1', '0', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (21, '资料修改', '/system/user/profile', '1', '20', '1', 'fa-edit', '1', '1', '1', '1', '0', '1', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (22, '更新头像', '/system/user/avatar', '0', '21', '2', '', '0', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (23, '修改密码', '/system/user/pwd', '0', '21', '3', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
    (24, '头像上传', '/system/user/uploadPage', '0', '21', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 12:03:00', '2024-10-17 12:03:00'),
    (25, '文件上传', '/system/upload', '0', '21', '4', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 12:11:55', '2024-10-17 12:11:55'),
    (26, '网站设置', '/system/site/edit', '0', '2', '5', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2023-07-11 13:37:09', '2023-07-11 13:37:09'),
    (27, '服务监控', '/monitor/server', '0', '2', '6', 'layui-icon ', '1', '1', '1', '1', '1', '1', '2025-02-27 12:32:52', '2025-02-27 12:33:03'),
    (30, '日志管理', '/', '0', '0', '99', 'layui-icon layui-icon-form', '1', '1', '0', '1', '0', '1', '2021-10-09 18:48:25', '2024-11-03 18:00:24'),
    (31, '日志列表', '/system/log/list', '0', '30', '1', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2021-10-09 18:49:08', '2021-10-09 18:49:08'),
    (32, '操作日志', '/system/log/operate', '0', '31', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2021-10-09 18:50:07', '2021-10-09 18:50:07'),
    (33, '登陆日志', '/system/log/login', '0', '31', '3', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2021-10-09 18:50:44', '2021-10-09 18:50:44'),
    (40, '字典管理', '/system/dict/type/list', '0', '2', '4', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2024-10-17 10:45:39', '2024-10-17 10:45:39'),
    (41, '新增', '/system/dict/type/add', '0', '40', '1', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 10:47:38', '2024-10-17 10:47:38'),
    (42, '修改', '/system/dict/type/edit', '0', '40', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 10:48:14', '2024-10-17 10:48:14'),
    (43, '删除', '/system/dict/type/delete', '0', '40', '3', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-17 10:48:39', '2024-10-17 10:48:39'),
    (44, '字典类型接口', '/system/dict/type/json', '0', '40', '4', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-21 10:02:06', '2024-10-21 10:02:06'),
    (45, '字典值接口', '/system/dict/data/json', '0', '40', '5', 'layui-icon ', '1', '1', '2', '1', '0', '1', '2024-10-17 10:48:39', '2024-10-30 21:39:18'),
    (50, '消息中心', '/', '0', '0', '2', 'layui-icon layui-icon-note', '1', '1', '0', '1', '0', '0', '2024-10-21 15:28:00', '2024-10-21 15:28:00'),
    (51, '消息列表', '/system/notice/list', '0', '50', '1', 'layui-icon ', '1', '1', '1', '1', '0', '0', '2024-10-21 15:29:29', '2024-10-21 15:29:29'),
    (52, '消息确认', '/system/notice/status', '0', '51', '2', 'layui-icon ', '1', '1', '2', '1', '0', '1', '2024-10-21 15:30:33', '2024-10-29 21:22:34'),
    (53, '消息列表接口', '/system/notice/json', '0', '51', '2', 'layui-icon ', '1', '1', '2', '1', '0', '0', '2024-10-23 11:57:51', '2024-10-23 11:57:51'),
    (54, '未读消息接口', '/system/notice', '0', '51', '4', 'layui-icon ', '1', '1', '2', '1', '0', '1', '2024-10-23 11:57:51', '2024-10-29 21:22:54'),
    (55, '消息查看', '/system/notice/edit', '0', '51', '1', 'layui-icon ', '1', '1', '2', '0', '1', '1', '2024-10-29 21:22:23', '2024-10-29 21:22:23'),
    (112, '文章管理', '/', '0', '0', '3', 'layui-icon layui-icon-tabs', '1', '1', '0', '0', '1', '1', '2025-03-13 10:35:37', '2025-03-13 10:36:11'),
    (113, '书籍管理', '/', '0', '0', '4', 'layui-icon layui-icon-read', '1', '1', '0', '0', '1', '1', '2025-03-13 10:36:45', '2025-03-13 10:36:45'),
    (114, '文章列表', '/cms/article/list', '0', '112', '1', 'layui-icon ', '1', '1', '1', '0', '1', '1', '2025-03-13 10:37:30', '2025-03-13 10:37:30'),
    (115, '书籍列表', '/cms/book/list', '0', '113', '1', 'layui-icon ', '1', '1', '1', '0', '1', '1', '2025-03-13 10:38:02', '2025-03-13 10:38:02');


-- name: create-sys-role
INSERT INTO `sys_role` (`id`, `role_name`, `detail`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
    (1, '管理员', '拥有超级管理权限', '1', '0', '0', '2020-09-28 14:00:10', '2020-09-28 14:00:06'),
    (2, '普通用户', '拥有监控操作权限', '1', '0', '0', '2020-09-28 13:59:53', '2020-09-28 14:00:03');


-- name: create-sys-role-auth
INSERT INTO `sys_role_auth` (`role_id`, `auth_id`)
VALUES
    (2,1),
    (2,2);


-- name: create-sys-conf
INSERT INTO `sys_conf` (`id`, `type`, `info`, `status`, `created_at`, `updated_at`, `create_id`, `update_id`)
VALUES
    (1, 'base_conf', '{\r\n	\"colors\": [\r\n {\r\n \"color\": \"#FF0000\",\r\n \"id\": \"1\"\r\n },\r\n {\r\n \"color\": \"#43B899\",\r\n \"id\": \"2\"\r\n },\r\n {\r\n \"color\": \"#1E9FFF\",\r\n \"id\": \"3\"\r\n },\r\n {\r\n \"color\": \"#FFB800\",\r\n \"id\": \"4\"\r\n },\r\n {\r\n \"color\": \"darkgray\",\r\n \"id\": \"5\"\r\n }\r\n ],\r\n	\"header\": {\r\n \"message\": \"/system/notice\"\r\n },\r\n \"links\": [\r\n {\r\n \"href\": \"http://www.anylink.vip\",\r\n \"icon\": \"layui-icon layui-icon-auz\",\r\n \"title\": \"官方网站\"\r\n },\r\n {\r\n \"href\": \"https://github.com/lzy3240/gsadmin\",\r\n \"icon\": \"layui-icon layui-icon-auz\",\r\n \"title\": \"开发文档\"\r\n },\r\n {\r\n \"href\": \"https://github.com/lzy3240/gsadmin\",\r\n \"icon\": \"layui-icon layui-icon-auz\",\r\n \"title\": \"开源地址\"\r\n }\r\n ],\r\n	\"logo\": {\r\n \"image\": \"/static/admin/images/logo.png\",\r\n \"title\": \"gsadmin\"\r\n },\r\n	\"menu\": {\r\n \"accordion\": true,\r\n \"control\": false,\r\n \"data\": \"/system/menu\",\r\n \"method\": \"GET\",\r\n \"select\": \"60\"\r\n },\r\n	\"other\": {\r\n \"autoHead\": false,\r\n \"keepLoad\": 100\r\n	},\r\n	\"tab\": {\r\n \"index\": {\r\n \"href\": \"/system/main\",\r\n \"id\": \"60\",\r\n \"title\": \"首页\"\r\n },\r\n \"keepState\": true,\r\n \"muiltTab\": true,\r\n \"tabMax\": 30\r\n	},\r\n	\"theme\": {\r\n \"allowCustom\": true,\r\n \"defaultColor\": \"2\",\r\n \"defaultMenu\": \"dark-theme\"\r\n }\r\n}', '1', '2021-05-28 10:48:35', '2021-05-28 10:48:35', NULL, NULL),
    (2, 'site_conf', '{\"web_name\":\"gsdoc\",\"web_url\":\"http://123/\",\"logo_url\":\"http://123/\",\"key_words\":\"admin\",\"description\":\"GsAdmin是基于Golang + Layui + MySql的轻量级极速后台开发框架，干净不臃肿、操作简单、开箱即用；通用型的后台权限管理机制，容易功能定制和二次开发，帮助开发者简单高效降低二次开发成本。\",\"copyright\":\"Copyright © 2025 gsadmin.com MIT license\",\"icp\":\"皖ICP备 1234567890号\",\"site_status\":1}', '1', '2021-05-28 10:48:35', '2025-04-22 13:31:00', NULL, NULL);


-- name: create-sys-dict-type
INSERT INTO `sys_dict_type` (`id`, `dict_type`, `dict_name`, `status`, `is_default`, `create_id`, `update_id`, `created_at`, `updated_at`, `remark`)
VALUES
    (1, 'sys_user_sex', '用户性别', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '用户性别列表'),
    (2, 'sys_show_hide', '菜单状态', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '菜单状态列表'),
    (3, 'sys_normal_disable', '系统开关', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '系统开关列表'),
    (4, 'sys_job_status', '任务状态', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '任务状态列表'),
    (5, 'sys_job_group', '任务分组', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '任务分组列表'),
    (6, 'sys_yes_no', '系统是否', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '系统是否列表'),
    (7, 'sys_notice_type', '通知类型', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '通知类型列表'),
    (8, 'sys_notice_status', '通知状态', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '通知状态列表'),
    (9, 'sys_oper_type', '操作类型', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '操作类型列表'),
    (10, 'sys_common_status', '系统状态', '1', '1', '0', '0', '2023-05-08 15:56:38', NULL, '登录状态列表'),
    (11, 'sys_audit_status', '审核状态', '1', '1', '0', '1', '2023-03-23 09:46:42', NULL, '审核状态列表'),
    (12, 'sys_notice_read', '通知已读', '1', '1', '0', '0', '2023-03-23 09:46:42', NULL, '通知已读列表'),
    (13, 'cms_show_status', '公开状态', '1', '0', '1', '1', '2025-03-26 09:21:12', '2025-03-26 09:30:55', '主页是否公开状态');


-- name: create-sys-dict-data
INSERT INTO `sys_dict_data` (`id`, `dict_type`, `dict_label`, `dict_value`, `dict_sort`, `status`, `remark`, `is_default`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
    (1, 'sys_user_sex', '男', '0', '1', '1', '性别男', '0', '0', '0', '2023-05-08 15:56:39', NULL),
	(2, 'sys_user_sex', '女', '1', '2', '1', '性别女', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (3, 'sys_user_sex', '未知', '2', '3', '1', '性别未知', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (4, 'sys_show_hide', '显示', '0', '1', '1', '显示菜单', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (5, 'sys_show_hide', '隐藏', '1', '2', '1', '隐藏菜单', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (6, 'sys_normal_disable', '正常', '1', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (7, 'sys_normal_disable', '停用', '0', '2', '1', '停用状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (8, 'sys_job_status', '正常', '0', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (9, 'sys_job_status', '暂停', '1', '2', '1', '停用状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (10, 'sys_job_group', '默认', 'DEFAULT', '1', '1', '默认分组', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (11, 'sys_job_group', '系统', 'SYSTEM', '2', '1', '系统分组', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (12, 'sys_yes_no', '是', 'Y', '1', '1', '系统默认是', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (13, 'sys_yes_no', '否', 'N', '2', '1', '系统默认否', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (14, 'sys_notice_type', '通知', '1', '1', '1', '通知', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (15, 'sys_notice_type', '公告', '2', '2', '1', '公告', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (16, 'sys_notice_status', '正常', '0', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (17, 'sys_notice_status', '关闭', '1', '2', '1', '关闭状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (18, 'sys_oper_type', '其他', '0', '99', '1', '其他操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (19, 'sys_oper_type', '新增', '1', '1', '1', '新增操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (20, 'sys_oper_type', '修改', '2', '2', '1', '修改操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (21, 'sys_oper_type', '删除', '3', '3', '1', '删除操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (22, 'sys_oper_type', '授权', '4', '4', '1', '授权操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (23, 'sys_oper_type', '导出', '5', '5', '1', '导出操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (24, 'sys_oper_type', '导入', '6', '6', '1', '导入操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (25, 'sys_oper_type', '强退', '7', '7', '1', '强退操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (26, 'sys_oper_type', '生成代码', '8', '8', '1', '生成操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (27, 'sys_oper_type', '清空数据', '9', '9', '1', '清空操作', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (28, 'sys_common_status', '成功', '0', '1', '1', '正常状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (29, 'sys_common_status', '失败', '1', '2', '1', '停用状态', '0', '0', '0', '2023-05-08 15:56:39', NULL),
    (30, 'sys_notice_read', '已读', '1', '1', '1', '已读状态', '0', '0', '0', '2024-10-22 22:09:59', NULL),
    (31, 'sys_notice_read', '未读', '0', '2', '1', '未读状态', '0', '0', '0', '2024-10-22 22:10:02', NULL),
    (32, 'sys_audit_status', '待审核11', '0', '1', '1', '待审核', '0', '0', '1', '2024-11-04 10:09:54', NULL),
    (33, 'sys_audit_status', '通过', '1', '2', '1', '审核通过', '0', '0', '0', '2024-11-04 10:11:16', NULL),
    (34, 'sys_audit_status', '不通过', '2', '3', '1', '审核不通过', '0', '0', '0', '2024-11-04 10:17:07', NULL),
    (35, 'cms_show_status', '隐藏', '0', '0', '1', '主页隐藏', '0', '1', '1', '2025-03-26 09:29:31', '2025-03-26 09:29:31'),
    (36, 'cms_show_status', '展示', '1', '1', '1', '主页展示', '0', '1', '1', '2025-03-26 09:29:44', '2025-03-26 09:29:44');
