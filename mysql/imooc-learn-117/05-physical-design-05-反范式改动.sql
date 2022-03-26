-- setup tables

/*
 Navicat MySQL Data Transfer

 Source Server         : mac
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : 练习-慕课网

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 26/03/2022 22:48:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_products
-- ----------------------------
DROP TABLE IF EXISTS `order_products`;
CREATE TABLE `order_products` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) DEFAULT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `product_count` int(11) DEFAULT NULL,
  `product_price` decimal(10,2) DEFAULT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `product_expired_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of order_products
-- ----------------------------
BEGIN;
INSERT INTO `order_products` VALUES (1, 1, 1, 58, 69.00, 'Irish Cream - Butterscotch', '2022-01-03 19:28:29');
INSERT INTO `order_products` VALUES (2, 1, 1, 93, 28.00, 'Irish Cream - Butterscotch', '2022-01-03 19:28:29');
INSERT INTO `order_products` VALUES (3, 3, 1, 2, 66.00, 'Irish Cream - Butterscotch', '2022-01-03 19:28:29');
INSERT INTO `order_products` VALUES (4, 3, 2, 86, 37.00, 'Nut - Peanut, Roasted', '2021-05-12 15:54:45');
INSERT INTO `order_products` VALUES (5, 4, 6, 46, 47.00, 'Sausage - Blood Pudding', '2021-04-05 10:08:40');
INSERT INTO `order_products` VALUES (6, 5, 7, 60, 76.00, 'Bread - Focaccia Quarter', '2022-02-11 01:07:18');
INSERT INTO `order_products` VALUES (7, 5, 3, 48, 12.00, 'Pears - Bartlett', '2021-03-28 04:13:28');
INSERT INTO `order_products` VALUES (8, 7, 4, 75, 43.00, 'Foam Espresso Cup Plain White', '2022-01-02 16:20:38');
INSERT INTO `order_products` VALUES (9, 2, 4, 63, 69.00, 'Foam Espresso Cup Plain White', '2022-01-02 16:20:38');
INSERT INTO `order_products` VALUES (10, 1, 1, 33, 65.00, 'Irish Cream - Butterscotch', '2022-01-03 19:28:29');
COMMIT;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `ordered_at` datetime NOT NULL,
  `pay_type` varchar(255) NOT NULL,
  `order_type` varchar(255) NOT NULL,
  `total_price` decimal(10,2) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `user_address` varchar(255) DEFAULT NULL,
  `user_cellphone` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
INSERT INTO `orders` VALUES (1, 1, '2022-02-08 22:12:11', 'cash', 'online', 162.00, 'mlathey0', '18720 Arrowood Trail', '591-133-0440');
INSERT INTO `orders` VALUES (2, 2, '2021-05-26 23:05:48', 'cash', 'offline', 69.00, 'fbatistelli1', '7 Novick Alley', '320-109-2445');
INSERT INTO `orders` VALUES (3, 2, '2021-07-11 02:23:41', 'alipay', 'offline', 103.00, 'fbatistelli1', '7 Novick Alley', '320-109-2445');
INSERT INTO `orders` VALUES (4, 2, '2021-11-15 21:35:46', 'wechatpay', 'offline', 47.00, 'fbatistelli1', '7 Novick Alley', '320-109-2445');
INSERT INTO `orders` VALUES (5, 4, '2022-02-09 04:33:19', 'wechatpay', 'online', 88.00, 'wmollitt3', '3388 Kinsman Crossing', '313-796-5179');
INSERT INTO `orders` VALUES (7, 5, '2021-09-01 20:16:31', 'wechatpay', 'offline', 43.00, 'jkhristoforov4', '16 Stone Corner Junction', '401-558-9526');
COMMIT;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `expired_at` datetime DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of products
-- ----------------------------
BEGIN;
INSERT INTO `products` VALUES (1, 'Irish Cream - Butterscotch', 'Lucía, Lucía (Hija del caníbal, La)', '2022-01-03 19:28:29');
INSERT INTO `products` VALUES (2, 'Nut - Peanut, Roasted', 'Love Crime (Crime d\'amour)', '2021-05-12 15:54:45');
INSERT INTO `products` VALUES (3, 'Pears - Bartlett', 'Late Phases', '2021-03-28 04:13:28');
INSERT INTO `products` VALUES (4, 'Foam Espresso Cup Plain White', 'Play Time (a.k.a. Playtime)', '2022-01-02 16:20:38');
INSERT INTO `products` VALUES (5, 'Lotus Rootlets - Canned', 'Frankenweenie', '2021-09-03 04:53:22');
INSERT INTO `products` VALUES (6, 'Sausage - Blood Pudding', 'Rounders', '2021-04-05 10:08:40');
INSERT INTO `products` VALUES (7, 'Bread - Focaccia Quarter', 'Griffin & Phoenix', '2022-02-11 01:07:18');
INSERT INTO `products` VALUES (8, 'Water - Green Tea Refresher', 'High School Musical', '2021-05-09 19:32:36');
INSERT INTO `products` VALUES (9, 'Shrimp - Tiger 21/25', 'Hunger, The', '2021-09-10 17:22:32');
INSERT INTO `products` VALUES (10, 'Sweet Pea Sprouts', 'The Vixen', '2021-04-17 19:10:18');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `cellphone` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `zipcode` varchar(255) DEFAULT NULL COMMENT '邮编',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'mlathey0', '591-133-0440', '18720 Arrowood Trail', NULL);
INSERT INTO `users` VALUES (2, 'fbatistelli1', '320-109-2445', '7 Novick Alley', '216533');
INSERT INTO `users` VALUES (3, 'hkirwood2', '171-745-7729', '5 Boyd Alley', '6326');
INSERT INTO `users` VALUES (4, 'wmollitt3', '313-796-5179', '3388 Kinsman Crossing', NULL);
INSERT INTO `users` VALUES (5, 'jkhristoforov4', '401-558-9526', '16 Stone Corner Junction', NULL);
INSERT INTO `users` VALUES (6, 'fstone5', '582-204-7069', '65 Petterle Road', '3240');
INSERT INTO `users` VALUES (7, 'lhortop6', '734-273-7142', '37 Hermina Drive', '10802');
INSERT INTO `users` VALUES (8, 'ispargo7', '851-990-0422', '4 Michigan Drive', 'L6S');
INSERT INTO `users` VALUES (9, 'peastgate8', '455-174-9480', '590 Darwin Way', '71703');
INSERT INTO `users` VALUES (10, 'cdevaney9', '351-307-3874', '2 Loftsgordon Road', '97435');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;


-- 问题: 如何查询订单信息?

SELECT
  user_name,
  user_cellphone,
  user_address,
  orders.id AS order_id,
  SUM( order_products.product_price * order_products.product_count ) AS order_total_price
FROM
  orders
  INNER JOIN order_products ON order_products.order_id = orders.id
GROUP BY
  user_name,
  user_cellphone,
  user_address,
  orders.id
ORDER BY
  order_total_price DESC

-- 问题: 如何查询订单详情?

SELECT
  user_name,
  user_cellphone,
  user_address,
  orders.id AS order_id,
  SUM( order_products.product_price * order_products.product_count ) AS order_total_price
FROM
  orders
  INNER JOIN order_products ON order_products.order_id = orders.id
GROUP BY
  user_name,
  user_cellphone,
  user_address,
  orders.id
HAVING
  order_id = 1

