SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
    `task_id` bigint(64) UNSIGNED NOT NULL ,
    `service_id` bigint(64) UNSIGNED NOT NULL ,
    `service_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '服务名称',
    `service_owner_id` bigint(64) UNSIGNED NOT NULL,
    `provider_id` bigint(64) UNSIGNED NOT NULL,
    `handle_id` bigint(64) UNSIGNED NOT NULL,
    `handle_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'handle名称',
    `db_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '库名',
    `table_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表名',
    `status` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '状态',
    PRIMARY KEY (`task_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'task记录表' ROW_FORMAT = COMPACT;