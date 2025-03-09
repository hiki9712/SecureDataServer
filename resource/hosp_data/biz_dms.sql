/*
 Navicat Premium Data Transfer

 Source Server         : Trust_525
 Source Server Type    : MySQL
 Source Server Version : 80027 (8.0.27)
 Source Host           : 47.94.19.196:3306
 Source Schema         : biz_dms

 Target Server Type    : MySQL
 Target Server Version : 80027 (8.0.27)
 File Encoding         : 65001

 Date: 19/02/2025 17:23:46
*/
CREATE DATABASE biz_dms;
USE biz_dms;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for 1234_POB_POINT_1026
-- ----------------------------
DROP TABLE IF EXISTS `1234_POB_POINT_1026`;
CREATE TABLE `1234_POB_POINT_1026` (
  `POINTID` decimal(8,0) NOT NULL,
  `DISC` varchar(60) DEFAULT NULL,
  `CODE` varchar(20) DEFAULT NULL,
  `CODE_NEW` varchar(20) DEFAULT NULL,
  `RTUID` decimal(8,0) DEFAULT NULL,
  `NUMBERID` decimal(4,0) DEFAULT NULL,
  `COMMADDRESS` varchar(20) DEFAULT NULL,
  `STATUSID` decimal(2,0) DEFAULT NULL,
  `LASTSTATETIME` datetime DEFAULT NULL,
  `CTID` decimal(4,0) DEFAULT NULL,
  `PTID` decimal(4,0) DEFAULT NULL,
  `METERTYPEID` decimal(4,0) DEFAULT NULL,
  `METERDIRECTION` decimal(1,0) DEFAULT NULL,
  `LASTCTTIME` datetime DEFAULT NULL,
  `TYPE` decimal(2,0) DEFAULT NULL,
  `CUSTOMERID` decimal(8,0) DEFAULT NULL,
  `LINEID` decimal(8,0) DEFAULT NULL,
  `VOLTAGECLASSID` decimal(4,0) DEFAULT NULL,
  `METERSITEID` decimal(2,0) DEFAULT NULL,
  `VOLTMONTYPEID` decimal(4,0) DEFAULT NULL,
  `ISSTANDBY` decimal(1,0) DEFAULT NULL,
  `ACTIVEPOINTID` decimal(8,0) DEFAULT NULL,
  `ISNOLOSS` decimal(1,0) DEFAULT NULL,
  `ISBYPASS` decimal(1,0) DEFAULT NULL,
  `ISSUM` decimal(1,0) DEFAULT NULL,
  `MEASUREMODE` decimal(2,0) DEFAULT NULL,
  `GWTID` decimal(2,0) DEFAULT NULL,
  `CAPABILITY` varchar(20) DEFAULT NULL,
  `BALOAD` decimal(12,3) DEFAULT NULL,
  `ADDRESS` varchar(20) DEFAULT NULL,
  `CREATETIME` datetime DEFAULT NULL,
  `METERCODE` varchar(20) DEFAULT NULL,
  `TARRIFORDER` varchar(14) DEFAULT NULL,
  `HH` varchar(20) DEFAULT NULL,
  `LOADCLASS` decimal(2,0) DEFAULT NULL,
  `SXJL` decimal(2,0) DEFAULT NULL,
  PRIMARY KEY (`POINTID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of 1234_POB_POINT_1026
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for 1234_ZDT_BM_1306_1026
-- ----------------------------
DROP TABLE IF EXISTS `1234_ZDT_BM_1306_1026`;
CREATE TABLE `1234_ZDT_BM_1306_1026` (
  `POINTID` decimal(8,0) NOT NULL,
  `DATATIME` datetime NOT NULL,
  `PHASETYPE` decimal(2,0) NOT NULL,
  `TARRIFTYPEID` decimal(4,0) NOT NULL,
  `ZYBM` decimal(10,0) DEFAULT NULL,
  `ZYBM_NEW` decimal(10,0) DEFAULT NULL,
  `FYBM` decimal(10,0) DEFAULT NULL,
  `ZWBM` decimal(10,0) DEFAULT NULL,
  `FWBM` decimal(10,0) DEFAULT NULL,
  `XX1BM` decimal(10,0) DEFAULT NULL,
  `XX2BM` decimal(10,0) DEFAULT NULL,
  `XX3BM` decimal(10,0) DEFAULT NULL,
  `XX4BM` decimal(10,0) DEFAULT NULL,
  `ZYPROP` decimal(2,0) DEFAULT NULL,
  `FYPROP` decimal(2,0) DEFAULT NULL,
  `ZWPROP` decimal(2,0) DEFAULT NULL,
  `FWPROP` decimal(2,0) DEFAULT NULL,
  `XX1PROP` decimal(2,0) DEFAULT NULL,
  `XX2PROP` decimal(2,0) DEFAULT NULL,
  `XX3PROP` decimal(2,0) DEFAULT NULL,
  `XX4PROP` decimal(2,0) DEFAULT NULL,
  PRIMARY KEY (`POINTID`,`DATATIME`,`PHASETYPE`,`TARRIFTYPEID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of 1234_ZDT_BM_1306_1026
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for POB_NEW
-- ----------------------------
DROP TABLE IF EXISTS `POB_NEW`;
CREATE TABLE `POB_NEW` (
  `POINTID` varchar(255) NOT NULL,
  `DISC` varchar(255) DEFAULT NULL,
  `CODE` varchar(255) DEFAULT NULL,
  `RTUID` varchar(255) DEFAULT NULL,
  `NUMBERID` varchar(255) DEFAULT NULL,
  `COMMADDRESS` varchar(255) DEFAULT NULL,
  `STATUSID` varchar(255) DEFAULT NULL,
  `LASTSTATETIME` varchar(255) DEFAULT NULL,
  `CTID` varchar(255) DEFAULT NULL,
  `PTID` varchar(255) DEFAULT NULL,
  `METERTYPEID` varchar(255) DEFAULT NULL,
  `METERDIRECTION` varchar(255) DEFAULT NULL,
  `LASTCTTIME` varchar(255) DEFAULT NULL,
  `TYPE` varchar(255) DEFAULT NULL,
  `CUSTOMERID` varchar(255) DEFAULT NULL,
  `CUSTOMERID_NEW` varchar(255) DEFAULT NULL,
  `LINEID` varchar(255) DEFAULT NULL,
  `VOLTAGECLASSID` varchar(255) DEFAULT NULL,
  `METERSITEID` varchar(255) DEFAULT NULL,
  `VOLTMONTYPEID` varchar(255) DEFAULT NULL,
  `ISSTANDBY` varchar(255) DEFAULT NULL,
  `ACTIVEPOINTID` varchar(255) DEFAULT NULL,
  `ISNOLOSS` varchar(255) DEFAULT NULL,
  `ISBYPASS` varchar(255) DEFAULT NULL,
  `ISSUM` varchar(255) DEFAULT NULL,
  `MEASUREMODE` varchar(255) DEFAULT NULL,
  `GWTID` varchar(255) DEFAULT NULL,
  `CAPABILITY` varchar(255) DEFAULT NULL,
  `BALOAD` varchar(255) DEFAULT NULL,
  `ADDRESS` varchar(255) DEFAULT NULL,
  `CREATETIME` varchar(255) DEFAULT NULL,
  `METERCODE` varchar(255) DEFAULT NULL,
  `TARRIFORDER` varchar(255) DEFAULT NULL,
  `HH` varchar(255) DEFAULT NULL,
  `LOADCLASS` varchar(255) DEFAULT NULL,
  `SXJL` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`POINTID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of POB_NEW
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for logs
-- ----------------------------
DROP TABLE IF EXISTS `logs`;
CREATE TABLE `logs` (
  `log_id` bigint NOT NULL COMMENT '脱敏日志ID',
  `task_id` bigint DEFAULT NULL COMMENT '脱敏任务ID',
  `task_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '脱敏任务名称',
  `securetable_field` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '脱敏数据表结构',
  `securetable_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '脱敏数据表名称',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '脱敏信息',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '脱敏任务状态',
  `exception_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '脱敏异常信息',
  `del_flag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '删除符号（0存在1删除）',
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '最后更新者',
  `update_time` datetime DEFAULT NULL COMMENT '最后更新时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for 用电统计_test_1_1024_POB_CUSTOMER
-- ----------------------------
DROP TABLE IF EXISTS `用电统计_test_1_1024_POB_CUSTOMER`;
CREATE TABLE `用电统计_test_1_1024_POB_CUSTOMER` (
  `CUSTOMERID` decimal(8,0) NOT NULL,
  `DISC` varchar(60) DEFAULT NULL,
  `DISC_NEW` varchar(60) DEFAULT NULL,
  `HH` varchar(20) DEFAULT NULL,
  `CODE` varchar(20) DEFAULT NULL,
  `TYPE` decimal(1,0) DEFAULT NULL,
  `POWERSUPPLYID` decimal(6,0) DEFAULT NULL,
  `LINEID` decimal(8,0) DEFAULT NULL,
  `LINEID2` decimal(8,0) DEFAULT NULL,
  `CUSTOMERSTATEID` decimal(2,0) DEFAULT NULL,
  `MEASUREMODE` decimal(4,0) DEFAULT NULL,
  `ECONOMICPROTID` decimal(4,0) DEFAULT NULL,
  `DEMANDPROTID` decimal(4,0) DEFAULT NULL,
  `BUSINESSCLASSID` decimal(4,0) DEFAULT NULL,
  `CAPABILITY` varchar(20) DEFAULT NULL,
  `CREATETIME` datetime DEFAULT NULL,
  `ADDRESS` varchar(50) DEFAULT NULL,
  `LINKMAN` varchar(30) DEFAULT NULL,
  `TELEPHONE` varchar(30) DEFAULT NULL,
  `MOBILE` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`CUSTOMERID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of 用电统计_test_1_1024_POB_CUSTOMER
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
