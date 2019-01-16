CREATE TABLE `go_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT  '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8  COMMENT='用户管理';

INSERT INTO `go_user` (`id`, `username`, `password`) VALUES ('1', 'admin', '123456');
INSERT INTO `go_user` (`id`, `username`, `password`) VALUES ('2', 'hequan', '123456');

CREATE TABLE `go_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '名字',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `go_role` (`id`,`name`) VALUES ('1', 'test');

CREATE TABLE `go_user_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned DEFAULT NULL COMMENT '角色ID',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户_角色ID_管理';

INSERT INTO `go_user_role` (`id`,`user_id`,`role_id`) VALUES ('1', '2','1');

CREATE TABLE `go_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `path` varchar(50) DEFAULT '' COMMENT '访问路径',
  `method` varchar(50) DEFAULT '' COMMENT '资源请求方式',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `go_menu` (`id`,`path`,`method`) VALUES ('1', '/api/v1/users','GET');

CREATE TABLE `go_role_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned DEFAULT NULL COMMENT '角色ID',
  `menu_id` int(11) unsigned DEFAULT NULL COMMENT '菜单ID',
  `deleted_on` int(11) unsigned DEFAULT  '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户_角色ID_管理';

INSERT INTO `go_role_menu` (`id`,`role_id`,`menu_id`) VALUES ('1', '1','1');