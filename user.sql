/*
 Navicat Premium Data Transfer

 Source Server         : 1
 Source Server Type    : MySQL
 Source Server Version : 80035
 Source Host           : localhost:3306
 Source Schema         : counseling_statistics

 Target Server Type    : MySQL
 Target Server Version : 80035
 File Encoding         : 65001

 Date: 21/03/2024 18:29:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `class` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone_number` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `activity` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `activity_number` int NULL DEFAULT NULL,
  `date` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `user` VALUES ('202205567003', '董峻宾', '阮2', '15562120176', '打麻将', NULL, NULL);
INSERT INTO `user` VALUES ('202205567004', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `user` VALUES ('202205567005', NULL, NULL, NULL, NULL, NULL, '2022-05-10 00:00:00');
INSERT INTO `user` VALUES ('202205567006', 'djb', 'r2', '15562120176', '打麻将', 4, '2006-01-02 15:04:05');
INSERT INTO `user` VALUES ('202205567007', 'djb', 'r2', '15562120176', '打麻将', 4, '2006-01-02 15:04:05');
INSERT INTO `user` VALUES ('202205567010', 'djb', 'r2', '15562120176', '打麻将', 4, '2006-01-02 15:04:05');
INSERT INTO `user` VALUES ('202205567011', 'djb', 'r2', '15562120176', '打麻将', 4, '2024-03-20 00:00:00');
INSERT INTO `user` VALUES ('202205567012', 'djb', 'r2', '15562120176', '打麻将', 4, '2007-01-02 15:04:05');

SET FOREIGN_KEY_CHECKS = 1;
