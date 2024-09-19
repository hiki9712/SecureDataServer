SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `compute_reg`;

CREATE TABLE `compute_reg` (
  `compute_task_id` bigint NOT NULL COMMENT '计算任务的唯一ID',
  `compute_type` int NOT NULL COMMENT '计算类型',
  `service_id` bigint NOT NULL COMMENT '服务的唯一ID',
  `service_owner_id` bigint NOT NULL COMMENT '服务注册者ID',
  `service_owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '服务注册者',
  `service_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '注册的服务名',
  `query_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '查询的联系人姓名',
  `query_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '查询的手机号',
  `query_start_time` datetime DEFAULT NULL COMMENT '查询的开始时间',
  `query_end_time` datetime DEFAULT NULL COMMENT '查询的结束时间',
  `provider_id_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '数据提供者的ID列表',
  `handle_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '该服务已经注册完成的handle列表',
  `del_flag` int DEFAULT NULL COMMENT '删除标志（0存在1删除）',
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `query_hh` varchar(255) DEFAULT NULL COMMENT '查询的户号',
  PRIMARY KEY (`compute_task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;