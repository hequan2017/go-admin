# noinspection SqlNoDataSourceInspectionForFile

/*
Navicat MySQL Data Transfer

Source Server         : 192.168.100.50
Source Server Version : 50722
Source Host           : 192.168.100.50:3306
Source Database       : go

Target Server Type    : MYSQL
Target Server Version : 50722
File Encoding         : 65001

Date: 2019-02-20 14:03:31
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for go_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_menu`;
CREATE TABLE `go_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '名字',
  `path` varchar(50) DEFAULT '' COMMENT '访问路径',
  `method` varchar(50) DEFAULT '' COMMENT '资源请求方式',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_menu
-- ----------------------------
INSERT INTO `go_menu` VALUES ('1', '查询所有菜单', '/api/v1/menus', 'GET', null, null, '0');
INSERT INTO `go_menu` VALUES ('2', '查询单个菜单', '/api/v1/menus/:id', 'GET', null, null, '0');
INSERT INTO `go_menu` VALUES ('3', '创建单个菜单', '/api/v1/menus', 'POST', null, null, '0');
INSERT INTO `go_menu` VALUES ('4', '更新单个菜单', '/api/v1/menus/:id', 'PUT', null, null, '0');
INSERT INTO `go_menu` VALUES ('5', '删除单个菜单', '/api/v1/menus/:id', 'DELETE', null, null, '0');
INSERT INTO `go_menu` VALUES ('6', '查询所有用户', '/api/v1/users', 'GET', null, null, '0');
INSERT INTO `go_menu` VALUES ('7', '查询单个用户', '/api/v1/users/:id', 'GET', null, null, '0');
INSERT INTO `go_menu` VALUES ('8', '创建单个用户', '/api/v1/users', 'POST', null, null, '0');
INSERT INTO `go_menu` VALUES ('9', '更新单个用户', '/api/v1/users/:id', 'PUT', null, null, '0');
INSERT INTO `go_menu` VALUES ('10', '删除单个用户', '/api/v1/users/:id', 'DELETE', null, null, '0');
INSERT INTO `go_menu` VALUES ('11', '查询所有角色', '/api/v1/roles', 'GET', null, null, '0');
INSERT INTO `go_menu` VALUES ('12', '查询单个角色', '/api/v1/roles/:id', 'GET', null, null, '0');
INSERT INTO `go_menu` VALUES ('13', '创建单个角色', '/api/v1/roles', 'POST', null, null, '0');
INSERT INTO `go_menu` VALUES ('14', '更新单个角色', '/api/v1/roles/:id', 'PUT', null, null, '0');
INSERT INTO `go_menu` VALUES ('15', '删除单个角色', '/api/v1/roles/:id', 'DELETE', null, null, '0');
INSERT INTO `go_menu` VALUES ('16', '登录', '/auth', 'GET', null, null, '0');

-- ----------------------------
-- Table structure for go_role
-- ----------------------------
DROP TABLE IF EXISTS `go_role`;
CREATE TABLE `go_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '名字',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_role
-- ----------------------------
INSERT INTO `go_role` VALUES ('1', 'test', null, null, '0');

-- ----------------------------
-- Table structure for go_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_role_menu`;
CREATE TABLE `go_role_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned DEFAULT NULL COMMENT '角色ID',
  `menu_id` int(11) unsigned DEFAULT NULL COMMENT '菜单ID',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='用户_角色ID_管理';

-- ----------------------------
-- Records of go_role_menu
-- ----------------------------
INSERT INTO `go_role_menu` VALUES ('1', '1', '1', '0');
INSERT INTO `go_role_menu` VALUES ('2', '1', '2', '0');
INSERT INTO `go_role_menu` VALUES ('3', '1', '3', '0');
INSERT INTO `go_role_menu` VALUES ('4', '1', '4', '0');
INSERT INTO `go_role_menu` VALUES ('5', '1', '5', '0');
INSERT INTO `go_role_menu` VALUES ('6', '1', '6', '0');
INSERT INTO `go_role_menu` VALUES ('7', '1', '7', '0');
INSERT INTO `go_role_menu` VALUES ('8', '1', '8', '0');
INSERT INTO `go_role_menu` VALUES ('9', '1', '9', '0');
INSERT INTO `go_role_menu` VALUES ('10', '1', '10', '0');
INSERT INTO `go_role_menu` VALUES ('11', '1', '11', '0');
INSERT INTO `go_role_menu` VALUES ('12', '1', '12', '0');
INSERT INTO `go_role_menu` VALUES ('13', '1', '13', '0');
INSERT INTO `go_role_menu` VALUES ('14', '1', '14', '0');
INSERT INTO `go_role_menu` VALUES ('15', '1', '15', '0');

-- ----------------------------
-- Table structure for go_user
-- ----------------------------
DROP TABLE IF EXISTS `go_user`;
CREATE TABLE `go_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COMMENT='用户管理';

-- ----------------------------
-- Records of go_user
-- ----------------------------
INSERT INTO `go_user` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', null, null, '0');
INSERT INTO `go_user` VALUES ('2', 'hequan', 'e10adc3949ba59abbe56e057f20f883e', '1550642309', '1550642309', '0');

-- ----------------------------
-- Table structure for go_user_role
-- ----------------------------
DROP TABLE IF EXISTS `go_user_role`;
CREATE TABLE `go_user_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned DEFAULT NULL COMMENT '角色ID',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='用户_角色ID_管理';

-- ----------------------------
-- Records of go_user_role
-- ----------------------------
INSERT INTO `go_user_role` VALUES ('1', '2', '1', '0');
