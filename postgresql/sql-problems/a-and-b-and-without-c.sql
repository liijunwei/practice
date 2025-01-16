Create table If Not Exists Customers (customer_id int, customer_name varchar(30));
Create table If Not Exists Orders (order_id int, customer_id int, product_name varchar(30));
Truncate table Customers;
insert into Customers (customer_id, customer_name) values ('1', 'Daniel');
insert into Customers (customer_id, customer_name) values ('2', 'Diana');
insert into Customers (customer_id, customer_name) values ('3', 'Elizabeth');
insert into Customers (customer_id, customer_name) values ('4', 'Jhon');
Truncate table Orders;
insert into Orders (order_id, customer_id, product_name) values ('10', '1', 'A');
insert into Orders (order_id, customer_id, product_name) values ('11', '1', 'A');
insert into Orders (order_id, customer_id, product_name) values ('20', '1', 'B');
insert into Orders (order_id, customer_id, product_name) values ('30', '1', 'D');
insert into Orders (order_id, customer_id, product_name) values ('40', '1', 'C');
insert into Orders (order_id, customer_id, product_name) values ('50', '2', 'A');
insert into Orders (order_id, customer_id, product_name) values ('60', '3', 'A');
insert into Orders (order_id, customer_id, product_name) values ('70', '3', 'B');
insert into Orders (order_id, customer_id, product_name) values ('80', '3', 'D');
insert into Orders (order_id, customer_id, product_name) values ('90', '4', 'C');

-- https://leetcode.cn/problems/customers-who-bought-products-a-and-b-but-not-c/?envType=study-plan-v2&envId=sql-premium-50

-- 请你编写解决方案，报告购买了产品 "A"，"B" 但没有购买产品 "C" 的客户的 customer_id 和 customer_name，因为我们想推荐他们购买这样的产品。
-- 返回按 customer_id 排序 的结果表。

-- solution1
select * from customers where customer_id in(
    select customer_id from orders where product_name = 'A'
)
and customer_id in(
    select customer_id from orders where product_name = 'B'
)
and customer_id NOT in(
    select customer_id from orders where product_name = 'C'
)

-- solution2
SELECT
c.customer_id
,c.customer_name
,SUM(CASE WHEN product_name = 'A' THEN 1 ELSE 0 END) aa
,SUM(CASE WHEN product_name = 'B' THEN 1 ELSE 0 END) bb
,SUM(CASE WHEN product_name = 'C' THEN 1 ELSE 0 END) cc
FROM orders o
LEFT JOIN customers c ON o.customer_id = c.customer_id
GROUP BY c.customer_id, c.customer_name
HAVING
    SUM(CASE WHEN product_name = 'A' THEN 1 ELSE 0 END) > 0 AND
    SUM(CASE WHEN product_name = 'B' THEN 1 ELSE 0 END) > 0 AND
    SUM(CASE WHEN product_name = 'C' THEN 1 ELSE 0 END) = 0;

SELECT
o.customer_id
,SUM(CASE WHEN product_name = 'A' THEN 1 ELSE 0 END) aa
,SUM(CASE WHEN product_name = 'B' THEN 1 ELSE 0 END) bb
,SUM(CASE WHEN product_name = 'C' THEN 1 ELSE 0 END) cc
,SUM(CASE WHEN product_name = 'D' THEN 1 ELSE 0 END) dd
FROM orders o
GROUP BY 1



