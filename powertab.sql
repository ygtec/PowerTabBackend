/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80012 (8.0.12)
 Source Host           : localhost:3306
 Source Schema         : powertab

 Target Server Type    : MySQL
 Target Server Version : 80012 (8.0.12)
 File Encoding         : 65001

 Date: 08/02/2026 23:16:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for calendar_events
-- ----------------------------
DROP TABLE IF EXISTS `calendar_events`;
CREATE TABLE `calendar_events`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `time` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'todo',
  `completed` tinyint(1) NULL DEFAULT 0,
  `recurring` tinyint(1) NULL DEFAULT 0,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_calendar_events_user_id`(`user_id`) USING BTREE,
  INDEX `idx_user_date`(`date`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of calendar_events
-- ----------------------------

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `label` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `order` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_categories_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES ('81012108-791d-454e-b480-9052316049b2', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '主页', 'Home', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('29cc59a4-4718-4f23-95b3-2f2dc04a2de8', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'AI 工具箱', 'Cpu', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('af1a6196-bd6d-45a5-8390-69b7e2217d2b', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'PowerTab AI', 'Sparkles', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('eeaed401-fb00-454a-8b8a-733956ac6339', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '工具箱', 'Grid', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('3f38e418-d780-4013-baf5-c08cdb037833', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'PDF 工具', 'FileText', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('51ca1eea-b36f-4f5f-9b56-c490b9b054be', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '壁纸', 'Image', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('fa91f01c-1ec9-4654-aae6-99e4f0f18043', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '日历', 'Calendar', 0, '2025-12-14 20:47:27.088');
INSERT INTO `categories` VALUES ('a53bebf9-49d5-4441-8929-f3308e4dd310', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '壁纸', 'Image', 5, '2025-12-15 09:18:12.393');
INSERT INTO `categories` VALUES ('014177c9-4b7d-4392-a8d5-cb8f347826b9', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'PDF 工具', 'FileText', 4, '2025-12-15 09:18:12.393');
INSERT INTO `categories` VALUES ('00e4e820-1d47-4796-84bf-386ee5b22df1', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '工具箱', 'Grid', 3, '2025-12-15 09:18:12.393');
INSERT INTO `categories` VALUES ('746dee8e-cecc-48e3-a747-ca66b6d7b972', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'PowerTab AI', 'Sparkles', 2, '2025-12-15 09:18:12.393');
INSERT INTO `categories` VALUES ('60bbf32d-bdad-471d-a554-49722be025a9', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'AI 工具箱', 'Cpu', 1, '2025-12-15 09:18:12.393');
INSERT INTO `categories` VALUES ('4c348ddf-47c1-4491-b813-e86931431c80', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '主页', 'Home', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('e79fd98b-a0c9-48d3-95bf-be20c995a7f1', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'AI 工具箱', 'Cpu', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('5e3aa09a-e53d-428c-8967-a8f9fe6e3619', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'PowerTab AI', 'Sparkles', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('5b3b0eea-f1d5-467d-b1c8-433aa74698ed', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '工具箱', 'Grid', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('e3042235-d9ca-4bd0-9850-48287dcf97f9', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'PDF 工具', 'FileText', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('93c07879-4aad-4cbc-a846-7f4f6a8baeaa', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '壁纸', 'Image', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('942c7dd4-4b03-4d3f-afc5-863ecb76030d', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '日历', 'Calendar', 0, '2025-12-15 09:23:39.501');
INSERT INTO `categories` VALUES ('c1b0a7cd-a7ba-49d8-810f-a9ba63624ade', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '主页', 'Home', 0, '2025-12-15 09:18:12.393');
INSERT INTO `categories` VALUES ('dashboard', 'user_169d4019-e63d-477d-9706-0213b012f16e', '主页', 'Home', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('ai-tools', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'AI 工具箱', 'Cpu', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('powertab-ai', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'PowerTab AI', 'Sparkles', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('productivity', 'user_169d4019-e63d-477d-9706-0213b012f16e', '工具箱', 'Grid', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('pdf-tools', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'PDF 工具', 'FileText', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('wallpapers', 'user_169d4019-e63d-477d-9706-0213b012f16e', '壁纸', 'Image', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('calendar', 'user_169d4019-e63d-477d-9706-0213b012f16e', '日历', 'Calendar', 0, '2025-12-15 09:33:54.057');
INSERT INTO `categories` VALUES ('7e98232c-1b2e-44a2-bb01-a7a17177a63a', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '日历', 'Calendar', 6, '2025-12-15 09:18:12.393');

-- ----------------------------
-- Table structure for chat_messages
-- ----------------------------
DROP TABLE IF EXISTS `chat_messages`;
CREATE TABLE `chat_messages`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `session_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'text',
  `media_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_chat_messages_session_id`(`session_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat_messages
-- ----------------------------
INSERT INTO `chat_messages` VALUES ('welcome', '1765763354830', 'model', '你好！我是 PowerTab AI 助手。我可以帮你解答问题、生成图片或创作视频。', 'text', '', '2025-12-15 10:57:40.659');
INSERT INTO `chat_messages` VALUES ('1765767460646', '1765763354830', 'user', '1', 'text', '', '2025-12-15 10:57:40.659');

-- ----------------------------
-- Table structure for chat_sessions
-- ----------------------------
DROP TABLE IF EXISTS `chat_sessions`;
CREATE TABLE `chat_sessions`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_chat_sessions_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat_sessions
-- ----------------------------
INSERT INTO `chat_sessions` VALUES ('1765763354830', 'user_38bf58dc-111c-4a89-95ac-6a48c2958024', '1', '2025-12-15 10:57:40.658');

-- ----------------------------
-- Table structure for point_consumptions
-- ----------------------------
DROP TABLE IF EXISTS `point_consumptions`;
CREATE TABLE `point_consumptions`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `action` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `points` bigint(20) NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_point_consumptions_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of point_consumptions
-- ----------------------------

-- ----------------------------
-- Table structure for recommended_tools
-- ----------------------------
DROP TABLE IF EXISTS `recommended_tools`;
CREATE TABLE `recommended_tools`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tool_id` bigint(20) NOT NULL,
  `order` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_recommended_tools_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of recommended_tools
-- ----------------------------

-- ----------------------------
-- Table structure for saved_wallpapers
-- ----------------------------
DROP TABLE IF EXISTS `saved_wallpapers`;
CREATE TABLE `saved_wallpapers`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_saved_wallpapers_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of saved_wallpapers
-- ----------------------------

-- ----------------------------
-- Table structure for search_engines
-- ----------------------------
DROP TABLE IF EXISTS `search_engines`;
CREATE TABLE `search_engines`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_search_engines_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of search_engines
-- ----------------------------

-- ----------------------------
-- Table structure for speed_dial_items
-- ----------------------------
DROP TABLE IF EXISTS `speed_dial_items`;
CREATE TABLE `speed_dial_items`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `color` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `order` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_speed_dial_items_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of speed_dial_items
-- ----------------------------
INSERT INTO `speed_dial_items` VALUES ('3af052bc-5d91-487d-a2ce-cb37303756b5', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'Bilibili', 'https://www.bilibili.com', '#00AEEC', '📺', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('dc8e7dbf-d446-4929-b65d-609e94201bf6', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'GitHub', 'https://github.com', '#181717', '💻', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('441cdddb-8572-42f9-9d18-9e4853c0cde6', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '知乎', 'https://www.zhihu.com', '#0084FF', '🧠', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('78819cad-f398-45c3-beca-090462b29b86', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'Youtube', 'https://www.youtube.com', '#FF0000', '▶️', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('7a91dc6a-9e59-4b0b-8bdf-f6933217b2cc', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '微博', 'https://weibo.com', '#E6162D', '👁️', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('9e505df7-30a2-43c0-a9e1-a5652a016e9e', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'ChatGPT', 'https://chat.openai.com', '#10A37F', '🤖', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('3f3ac224-8687-4ed7-a336-0250418c7cdb', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '小红书', 'https://www.xiaohongshu.com', '#FF2442', '📕', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('86785074-af94-430d-bbbb-eb444270f13c', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', '淘宝', 'https://taobao.com', '#FF5000', '🛍️', 0, '2025-12-14 20:47:27.093');
INSERT INTO `speed_dial_items` VALUES ('288bcc46-1439-4922-ace6-03b0bde75749', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'Bilibili', 'https://www.bilibili.com', '#00AEEC', '📺', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('94ca7cbf-aab2-4195-887b-d6878c5c673a', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'GitHub', 'https://github.com', '#181717', '💻', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('df8bd3b4-3255-4923-bc86-9e65bcfd286e', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '知乎', 'https://www.zhihu.com', '#0084FF', '🧠', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('d5cb89d7-5372-4f43-acc9-0e450cc2308e', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'Youtube', 'https://www.youtube.com', '#FF0000', '▶️', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('a0ba6465-3050-45c1-bb2d-05d5014999bd', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '微博', 'https://weibo.com', '#E6162D', '👁️', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('73399049-6c85-4232-9cfa-d0cc75b5844c', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'ChatGPT', 'https://chat.openai.com', '#10A37F', '🤖', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('04d9093a-3685-4b9d-8a0c-51f996e9f793', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '小红书', 'https://www.xiaohongshu.com', '#FF2442', '📕', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('6ba563f6-85b5-4c5a-8e97-490bee4a1563', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', '淘宝', 'https://taobao.com', '#FF5000', '🛍️', 0, '2025-12-15 09:18:12.401');
INSERT INTO `speed_dial_items` VALUES ('7f49d317-897e-4672-a4d4-e26af20c880d', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'Bilibili', 'https://www.bilibili.com', '#00AEEC', '📺', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('85f76779-eacf-43e9-a6d6-243b856d6cd4', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'GitHub', 'https://github.com', '#181717', '💻', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('019605d5-997f-4b7d-a4b7-aef13e5aaf4d', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '知乎', 'https://www.zhihu.com', '#0084FF', '🧠', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('3b0dd67f-9ddf-4c5c-a04d-0cc745e6fcd4', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'Youtube', 'https://www.youtube.com', '#FF0000', '▶️', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('d920aa67-e99d-468e-88c0-66803dfd6fb1', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '微博', 'https://weibo.com', '#E6162D', '👁️', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('eb937fdc-c545-4c9e-a341-b89c607b8c33', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'ChatGPT', 'https://chat.openai.com', '#10A37F', '🤖', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('0be210ac-33c6-45f0-9c70-b5f962922715', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '小红书', 'https://www.xiaohongshu.com', '#FF2442', '📕', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('012b6a22-ce3c-4b20-b2c5-a0d698fc8863', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', '淘宝', 'https://taobao.com', '#FF5000', '🛍️', 0, '2025-12-15 09:23:39.508');
INSERT INTO `speed_dial_items` VALUES ('dd784284-ea88-4046-b581-ab031158032b', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'Bilibili', 'https://www.bilibili.com', '#00AEEC', '📺', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('24f2264f-bc86-4567-85f0-abb42f9dfa79', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'GitHub', 'https://github.com', '#181717', '💻', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('63bb9987-c4b8-48f6-85d6-abc765a319c6', 'user_169d4019-e63d-477d-9706-0213b012f16e', '知乎', 'https://www.zhihu.com', '#0084FF', '🧠', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('5c0ce29c-d912-4cb1-8c63-4e64615401e5', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'Youtube', 'https://www.youtube.com', '#FF0000', '▶️', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('7ea6e977-22ae-454d-bad2-4f93314f6bc2', 'user_169d4019-e63d-477d-9706-0213b012f16e', '微博', 'https://weibo.com', '#E6162D', '👁️', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('9c207ddb-08c9-448f-b55d-918b03b036a0', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'ChatGPT', 'https://chat.openai.com', '#10A37F', '🤖', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('f458c73f-2d73-4931-8fbe-e486fda4b1e8', 'user_169d4019-e63d-477d-9706-0213b012f16e', '小红书', 'https://www.xiaohongshu.com', '#FF2442', '📕', 0, '2025-12-15 09:33:54.062');
INSERT INTO `speed_dial_items` VALUES ('df7fb0c5-837f-48ab-b74b-e0eb8858c084', 'user_169d4019-e63d-477d-9706-0213b012f16e', '淘宝', 'https://taobao.com', '#FF5000', '🛍️', 0, '2025-12-15 09:33:54.062');

-- ----------------------------
-- Table structure for tool_orders
-- ----------------------------
DROP TABLE IF EXISTS `tool_orders`;
CREATE TABLE `tool_orders`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `category` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `orders` longblob NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_tool_orders_user_id`(`user_id`) USING BTREE,
  INDEX `idx_user_category`(`category`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_orders
-- ----------------------------

-- ----------------------------
-- Table structure for tool_usage_logs
-- ----------------------------
DROP TABLE IF EXISTS `tool_usage_logs`;
CREATE TABLE `tool_usage_logs`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tool_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `action` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_tool_usage_logs_user_id`(`user_id`) USING BTREE,
  INDEX `idx_tool_usage_logs_tool_id`(`tool_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_usage_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_preferences
-- ----------------------------
DROP TABLE IF EXISTS `user_preferences`;
CREATE TABLE `user_preferences`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `theme` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'auto',
  `show_seconds` tinyint(1) NULL DEFAULT 1,
  `notifications` tinyint(1) NULL DEFAULT 1,
  `language` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'zh',
  `default_engine` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'baidu',
  `wallpaper_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `auto_wallpaper` tinyint(1) NULL DEFAULT 0,
  `icon_size` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'medium',
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_preferences_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_preferences
-- ----------------------------
INSERT INTO `user_preferences` VALUES ('2f0cc34d-200d-4753-847e-9734a4664a15', 'user_1851da51-aa86-475a-aeed-cb56f2532110', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-14 19:31:34.434');
INSERT INTO `user_preferences` VALUES ('f9b21d8b-29a9-4b5d-8d5c-a1efea4a4cf2', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-14 19:37:17.273');
INSERT INTO `user_preferences` VALUES ('f9509466-1a57-4958-996e-f88effc9baf3', 'user_0d98271a-474c-4c55-9e36-f602226ae1e9', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-14 19:37:17.327');
INSERT INTO `user_preferences` VALUES ('8fa0bee5-09c9-401e-893a-319e81ec8b1f', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-14 20:47:26.725');
INSERT INTO `user_preferences` VALUES ('c2df1e9d-83c9-454d-919d-fa1616db1592', 'user_6655e2c9-5508-4545-ab14-c7e229b53a04', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-15 09:23:31.187');
INSERT INTO `user_preferences` VALUES ('73695da7-a2fe-4839-8e8b-21cf0fe96acd', 'user_4d6591c0-9065-4a62-9b66-8079d60a7e56', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-15 09:33:53.957');
INSERT INTO `user_preferences` VALUES ('2f73acd2-b152-476e-a3cc-5eea456a274f', 'user_38bf58dc-111c-4a89-95ac-6a48c2958024', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-15 09:49:14.741');
INSERT INTO `user_preferences` VALUES ('ac83748e-9912-4fbc-be60-b2c78b6a94b0', 'user_c5d7e776-22a0-45ac-b839-77201820e42a', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-15 09:49:14.766');
INSERT INTO `user_preferences` VALUES ('120be1c6-d003-4043-89bf-7fa94ab275ef', 'user_1c890f41-06d0-44de-90f2-3a6977c7eec7', 'auto', 1, 1, 'zh', 'baidu', '', 0, 'medium', '2025-12-15 22:03:54.645');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `plan_level` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'free',
  `plan_expiry` datetime(3) NULL DEFAULT NULL,
  `points` bigint(20) NULL DEFAULT 0,
  `token` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_users_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_users_email`(`email`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('user_1851da51-aa86-475a-aeed-cb56f2532110', 'test100', 'test100@example.com', '$2a$10$gbUQJol.8DYxihK25Rds5emfyy3ZNWALkqZdt2UzZsX8kaaW2KOpu', '', 'free', NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl80NTVjNDNkOC04NDJjLTQ2YTItYTdjMS1jN2E2MWUwYTE2NjEiLCJleHAiOjE3NjU3OTgyOTQsIm5iZiI6MTc2NTcxMTg5NCwiaWF0IjoxNzY1NzExODk0fQ.Pyqaugww9MJW5Ol266Oj_qS6bhTTlbJ6YfOBDw6HpaQ', '2025-12-14 19:31:34.431', '2025-12-14 19:31:34.431');
INSERT INTO `users` VALUES ('user_13183d50-1563-419d-87e6-a6ee9cefc664', '1', '1@qq.com', '$2a$10$qwcS1aFP617u9k2BBnGtquRCZSJ7NgLCBrsCLvxY3hreLECr0RBbC', '', 'free', NULL, 0, '', '2025-12-14 19:37:17.271', '2025-12-20 17:10:30.764');
INSERT INTO `users` VALUES ('user_38bf58dc-111c-4a89-95ac-6a48c2958024', '2', '2@qq.com', '$2a$10$or2kcgwwoXyF0u8/ocxkRe7jfol10BFUmJcaMb42VD.0K5pN4YPbG', '', 'free', NULL, 0, '', '2025-12-15 09:49:14.739', '2025-12-15 11:25:12.529');
INSERT INTO `users` VALUES ('user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'testuser', 'testdata_1825447545@example.com', '$2a$10$aNkuaPFwBsRYvgjAgCDfm.G/cypsJrx2hpHbuAe9qF3NZZrL6X5bS', '', 'free', NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl80MGE3Y2Y3Mi03MTY3LTQ0YjEtOWI3NC1iZjhkN2EyN2Y0NGMiLCJleHAiOjE3NjU4MDI4NDYsIm5iZiI6MTc2NTcxNjQ0NiwiaWF0IjoxNzY1NzE2NDQ2fQ.ybeHIWiLTOp2S8FnmBQda-D6I29lHM0k9832QBF37j0', '2025-12-14 20:47:26.723', '2025-12-14 20:47:26.892');
INSERT INTO `users` VALUES ('user_6655e2c9-5508-4545-ab14-c7e229b53a04', 'newtestuser20251215092331', 'newtest20251215092331@example.com', '$2a$10$8dAC.X/Q4BcxiUbgi3wvZulM.W/afUnyrIxNOr6dkg9w9DkqRSbjy', '', 'free', NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl8wODdkZDQwYS1kYWU2LTQzNDgtYTQ1OS02ZjI3YmVlY2FhMjYiLCJleHAiOjE3NjU4NDgyMTEsIm5iZiI6MTc2NTc2MTgxMSwiaWF0IjoxNzY1NzYxODExfQ.qATluZK4jGYVDp-pP1Ms9b_7-4APxH4fWcDBAsYEQt4', '2025-12-15 09:23:31.185', '2025-12-15 09:23:31.185');
INSERT INTO `users` VALUES ('user_4d6591c0-9065-4a62-9b66-8079d60a7e56', 'finaluser20251215093353', 'final_test_20251215093353@example.com', '$2a$10$NYXZDy/gpwQ3WaT.zyHcO.1wsYXqpupzqtzlTIux4jEJu8Qh425/K', '', 'free', NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl8xNjlkNDAxOS1lNjNkLTQ3N2QtOTcwNi0wMjEzYjAxMmYxNmUiLCJleHAiOjE3NjU4NDg4MzMsIm5iZiI6MTc2NTc2MjQzMywiaWF0IjoxNzY1NzYyNDMzfQ.Vk5N2-0jxymppbdhHw2DJR94LB3RTblVRm3EgX4i1Ks', '2025-12-15 09:33:53.954', '2025-12-15 09:33:53.954');
INSERT INTO `users` VALUES ('user_1c890f41-06d0-44de-90f2-3a6977c7eec7', '3', '3@qq.com', '$2a$10$c29x9/Ed4zQLUELpHr.DDuTYrAlJTm/j7WmJg7FryhRgZLzsS9zBu', '', 'free', NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl8xYzg5MGY0MS0wNmQwLTQ0ZGUtOTBmMi0zYTY5NzdjN2VlYzciLCJleHAiOjE3NjYzMDgyNDYsIm5iZiI6MTc2NjIyMTg0NiwiaWF0IjoxNzY2MjIxODQ2fQ.N8RJb24lNcVC5lBL_xjaQ1lXSVroaUUXIVhY8ugdrGE', '2025-12-15 22:03:54.642', '2025-12-20 17:10:46.027');

-- ----------------------------
-- Table structure for widgets
-- ----------------------------
DROP TABLE IF EXISTS `widgets`;
CREATE TABLE `widgets`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `widget_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `page_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `size` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `settings` longblob NULL,
  `order` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_widgets_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of widgets
-- ----------------------------
INSERT INTO `widgets` VALUES ('462cfa1b-ad86-4467-b19f-99ebf193cabe', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'weather', 'dashboard', 'medium', 0x7B226C6F636174696F6E223A224265696A696E67227D, 0, '2025-12-14 20:47:27.091');
INSERT INTO `widgets` VALUES ('640dbea5-cea8-48d1-8560-95ae068f13a7', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'clock', 'dashboard', 'small', 0x7B22666F726D6174223A22323468227D, 0, '2025-12-14 20:47:27.091');
INSERT INTO `widgets` VALUES ('5f5a7b3e-b626-4156-9be4-9a13d44c710b', 'user_40a7cf72-7167-44b1-9b74-bf8d7a27f44c', 'todo', 'dashboard', 'medium', 0x7B7D, 0, '2025-12-14 20:47:27.091');
INSERT INTO `widgets` VALUES ('fb0488c7-3dc0-449d-ac82-2a7013816572', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'weather', 'dashboard', 'medium', 0x7B226C6F636174696F6E223A224265696A696E67227D, 0, '2025-12-15 09:18:12.397');
INSERT INTO `widgets` VALUES ('48de7798-e33e-4303-be27-c0ba0e59af48', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'clock', 'dashboard', 'small', 0x7B22666F726D6174223A22323468227D, 0, '2025-12-15 09:18:12.397');
INSERT INTO `widgets` VALUES ('cf9d53e6-810a-469b-ba73-17ce90b110a2', 'user_13183d50-1563-419d-87e6-a6ee9cefc664', 'todo', 'dashboard', 'medium', 0x7B7D, 0, '2025-12-15 09:18:12.397');
INSERT INTO `widgets` VALUES ('65e56598-0419-488e-8b59-33b68d4b2995', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'weather', 'dashboard', 'medium', 0x7B226C6F636174696F6E223A224265696A696E67227D, 0, '2025-12-15 09:23:39.504');
INSERT INTO `widgets` VALUES ('44b06de2-b788-48fa-a860-1c9d9e709288', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'clock', 'dashboard', 'small', 0x7B22666F726D6174223A22323468227D, 0, '2025-12-15 09:23:39.504');
INSERT INTO `widgets` VALUES ('eb5e26b1-d6c8-48c9-9deb-6d044f9841d9', 'user_087dd40a-dae6-4348-a459-6f27beecaa26', 'todo', 'dashboard', 'medium', 0x7B7D, 0, '2025-12-15 09:23:39.504');
INSERT INTO `widgets` VALUES ('44b5e201-40ae-4796-b8bb-1e957d6b67ef', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'weather', 'dashboard', 'medium', 0x7B226C6F636174696F6E223A224265696A696E67227D, 0, '2025-12-15 09:33:54.059');
INSERT INTO `widgets` VALUES ('e8774f4e-d4ae-4e5c-9369-af78b2eb502f', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'clock', 'dashboard', 'small', 0x7B22666F726D6174223A22323468227D, 0, '2025-12-15 09:33:54.059');
INSERT INTO `widgets` VALUES ('5dcdc117-a0ab-41ac-884b-d662c2fa138d', 'user_169d4019-e63d-477d-9706-0213b012f16e', 'todo', 'dashboard', 'medium', 0x7B7D, 0, '2025-12-15 09:33:54.059');

SET FOREIGN_KEY_CHECKS = 1;
