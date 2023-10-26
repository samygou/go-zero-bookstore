/*
 Navicat MySQL Data Transfer

 Source Server         : bookstore
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:3306
 Source Schema         : bookstore_usercenter

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 21/10/2023 17:14:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `create_time` bigint unsigned NOT NULL COMMENT '创建时间',
    `update_time` bigint unsigned NOT NULL COMMENT '更新时间',
    `del_status` tinyint(1) unsigned NOT NULL COMMENT '状态, 0正常, 1删除',
    `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `sex` tinyint(1) unsigned NOT NULL COMMENT '性别, 0未知, 1男, 2女',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `remark` varchar(255) DEFAULT '' COMMENT '备注信息',
    PRIMARY KEY(`id`),
    UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='账户表';

SET FOREIGN_KEY_CHECKS = 1;
