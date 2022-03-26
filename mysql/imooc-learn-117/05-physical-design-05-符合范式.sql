-- setup tables

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `cellphone` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `zipcode` varchar(255) DEFAULT NULL COMMENT '邮编',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

insert into users (name, cellphone, address, zipcode) values ('mlathey0', '591-133-0440', '18720 Arrowood Trail', null);
insert into users (name, cellphone, address, zipcode) values ('fbatistelli1', '320-109-2445', '7 Novick Alley', '216533');
insert into users (name, cellphone, address, zipcode) values ('hkirwood2', '171-745-7729', '5 Boyd Alley', '6326');
insert into users (name, cellphone, address, zipcode) values ('wmollitt3', '313-796-5179', '3388 Kinsman Crossing', null);
insert into users (name, cellphone, address, zipcode) values ('jkhristoforov4', '401-558-9526', '16 Stone Corner Junction', null);
insert into users (name, cellphone, address, zipcode) values ('fstone5', '582-204-7069', '65 Petterle Road', '3240');
insert into users (name, cellphone, address, zipcode) values ('lhortop6', '734-273-7142', '37 Hermina Drive', '10802');
insert into users (name, cellphone, address, zipcode) values ('ispargo7', '851-990-0422', '4 Michigan Drive', 'L6S');
insert into users (name, cellphone, address, zipcode) values ('peastgate8', '455-174-9480', '590 Darwin Way', '71703');
insert into users (name, cellphone, address, zipcode) values ('cdevaney9', '351-307-3874', '2 Loftsgordon Road', '97435');

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `expired_at` datetime DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

insert into products (name, description, expired_at) values ('Irish Cream - Butterscotch', 'Lucía, Lucía (Hija del caníbal, La)', '2022-01-03 19:28:29');
insert into products (name, description, expired_at) values ('Nut - Peanut, Roasted', 'Love Crime (Crime d''amour)', '2021-05-12 15:54:45');
insert into products (name, description, expired_at) values ('Pears - Bartlett', 'Late Phases', '2021-03-28 04:13:28');
insert into products (name, description, expired_at) values ('Foam Espresso Cup Plain White', 'Play Time (a.k.a. Playtime)', '2022-01-02 16:20:38');
insert into products (name, description, expired_at) values ('Lotus Rootlets - Canned', 'Frankenweenie', '2021-09-03 04:53:22');
insert into products (name, description, expired_at) values ('Sausage - Blood Pudding', 'Rounders', '2021-04-05 10:08:40');
insert into products (name, description, expired_at) values ('Bread - Focaccia Quarter', 'Griffin & Phoenix', '2022-02-11 01:07:18');
insert into products (name, description, expired_at) values ('Water - Green Tea Refresher', 'High School Musical', '2021-05-09 19:32:36');
insert into products (name, description, expired_at) values ('Shrimp - Tiger 21/25', 'Hunger, The', '2021-09-10 17:22:32');
insert into products (name, description, expired_at) values ('Sweet Pea Sprouts', 'The Vixen', '2021-04-17 19:10:18');

-- 问题: 如何查询订单信息?

SELECT users.name, users.cellphone, users.address, orders.id AS 订单ID, SUM(order_products.product_price * order_products.product_count) AS 订单价格
FROM orders
INNER JOIN users ON users.id = orders.user_id
INNER JOIN order_products ON order_products.order_id = orders.id
GROUP BY users.name, users.cellphone, users.address, orders.id


-- 问题: 如何查询订单详情? (????)

SELECT users.name, users.cellphone, users.address, orders.id, SUM(order_products.product_price * order_products.product_count) AS 订单价格
FROM orders
INNER JOIN users ON orders.user_id = users.id
INNER JOIN order_products ON order_products.order_id = orders.id
INNER JOIN products ON products.id = order_products.product_id
GROUP BY users.name, users.cellphone, users.address, orders.id

