DROP TABLE IF EXISTS `handle_reg`;
CREATE TABLE `handle_reg` (
                              `handle_id` bigint NOT NULL COMMENT '句柄id',
                              `handle_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '句柄类型',
                              `handle_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '句柄名',
                              `service_id` bigint NOT NULL COMMENT '服务id',
                              `service_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '服务名',
                              `provider_id` bigint NOT NULL COMMENT '提供方id',
                              `keyValueCount` int DEFAULT NULL COMMENT '乱码键值数量',
                              `keyValueContent` LONGTEXT DEFAULT NULL COMMENT '乱码键值内容',
                              `del_flag` int DEFAULT NULL COMMENT '删除标志',
                              `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '创建者',
                              `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                              `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '更新者',
                              `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                              `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
                              PRIMARY KEY (`handle_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `negotiation`;
CREATE TABLE `negotiation` (
                               `service_id` bigint NOT NULL COMMENT '服务的唯一ID',
                               `service_owner_id` bigint NOT NULL COMMENT '服务注册者ID',
                               `service_owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '服务注册者',
                               `service_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '注册的服务名',
                               `provider_id` bigint DEFAULT NULL COMMENT '数据提供方的ID',
                               `provider_name_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '数据提供方名称列表',
                               `provide_table_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '数据提供方的具体数据表',
                               `privide_db_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '数据提供方的数据库列表',
                               `securetable_field` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '脱敏表的表结构（json文件）',
                               `securetable_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '脱敏表名称',
                               `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '数据协商的状态：success表示成功；fail表示失败；ing表示正在进行',
                               `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '协商信息',
                               `del_flag` int DEFAULT NULL COMMENT '删除标志（0存在1删除）',
                               `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '创建者',
                               `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                               `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '更新者',
                               `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                               `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;