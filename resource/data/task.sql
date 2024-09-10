SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
    `service_id` bigint(64) UNSIGNED NOT NULL ,
    `service_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '服务名称',
    `task_id` bigint(64) UNSIGNED NOT NULL ,
    `handle_id` bigint(64) UNSIGNED NOT NULL,
    `db_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '库名',
    `table_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表名',
    `status` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '状态'
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'task记录表' ROW_FORMAT = COMPACT;