-- 创建数据库
CREATE DATABASE IF NOT EXISTS chronos CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 切换数据库
USE chronos;

-- 用户
CREATE TABLE user
(
    `id`          bigint(20)    NOT NULL COMMENT '主键',
    `created_at`  datetime(3)   NULL COMMENT '创建时间',
    `updated_at`  datetime(3)   NULL COMMENT '更新时间',
    `deleted_at`  datetime(3)   NULL COMMENT '删除时间',
    `email`       varchar(255)  NOT NULL COMMENT '邮箱',
    `password`    varchar(255)  NOT NULL COMMENT '密码',
    `nickname`    varchar(255)  NOT NULL DEFAULT '' COMMENT '昵称',
    `gender`      tinyint       NOT NULL DEFAULT 0  COMMENT '性别',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_id` (`id`) USING BTREE,
    KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';
