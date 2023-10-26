SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `create_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
    `name` varchar(32) NOT NULL DEFAULT '' COMMENT '书名',
    `price` int unsigned NOT NULL DEFAULT 0 COMMENT '价格',
    `desc` varchar(255) DEFAULT '' COMMENT '描述',
    PRIMARY KEY(`id`),
    KEY `index_name` (`name`),
    KEY `index_price` (`price`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='书籍表';

SET FOREIGN_KEY_CHECKS = 1;
