> Note
> The SHOW PROFILE and SHOW PROFILES statements are deprecated; expect them to be removed in a future MySQL release. Use the [Performance Schema](https://dev.mysql.com/doc/refman/8.0/en/performance-schema.html) instead; see [Section 27.19.1, “Query Profiling Using Performance Schema”](https://dev.mysql.com/doc/refman/8.0/en/performance-schema-query-profiling.html).

```sql
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `zh_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

BEGIN;
INSERT INTO `roles` VALUES (1, 'admin', '管理员', '2019-12-20 09:26:29', '2019-12-20 09:26:29');
INSERT INTO `roles` VALUES (2, 'agent', '客服', '2019-12-20 09:26:29', '2019-12-20 09:26:29');
INSERT INTO `roles` VALUES (3, 'customer', '客户', '2019-12-20 09:26:29', '2019-12-20 09:26:29');
INSERT INTO `roles` VALUES (4, 'invalid_agent', '无效客服', '2019-12-20 09:26:29', '2019-12-20 09:26:29');
INSERT INTO `roles` VALUES (5, 'disable_agent', '禁用客服', '2019-12-24 07:37:12', '2019-12-24 07:37:12');
COMMIT;

-- required to enable
-- https://dev.mysql.com/doc/refman/8.0/en/show-profile.html
set profiling=1;

select * from roles;
desc roles;
show profiles;
show profile for query 1;
show profile for query 2;
```
