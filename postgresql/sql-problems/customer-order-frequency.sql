-- \i ~/practice/postgresql/sql-problems/customer-order-frequency.sql
-- 消费者下单频率

drop table if exists customers;
drop table if exists product;
drop table if exists orders;
create table if not exists customers (customer_id int, name varchar(30), country varchar(30));
create table if not exists product (product_id int, description varchar(30), price int);
create table if not exists orders (order_id int, customer_id int, product_id int, order_date date, quantity int);
truncate table customers;
insert into customers (customer_id, name, country) values ('1', 'winston', 'usa');
insert into customers (customer_id, name, country) values ('2', 'jonathan', 'peru');
insert into customers (customer_id, name, country) values ('3', 'moustafa', 'egypt');
truncate table product;
insert into product (product_id, description, price) values ('10', 'lc phone', '300');
insert into product (product_id, description, price) values ('20', 'lc t-shirt', '10');
insert into product (product_id, description, price) values ('30', 'lc book', '45');
insert into product (product_id, description, price) values ('40', 'lc keychain', '2');
truncate table orders;
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('1', '1', '10', '2020-06-10', '1');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('2', '1', '20', '2020-07-01', '1');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('3', '1', '30', '2020-07-08', '2');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('4', '2', '10', '2020-06-15', '2');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('5', '2', '40', '2020-07-01', '10');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('6', '3', '20', '2020-06-24', '2');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('7', '3', '30', '2020-06-25', '2');
insert into orders (order_id, customer_id, product_id, order_date, quantity) values ('9', '3', '30', '2020-05-08', '3');

-- solution1
with t as(
    select
    c.customer_id
    ,c.name
    ,date_trunc('month', o.order_date) date
    ,sum(p.price * o.quantity)
    from
    orders o
    join product p on o.product_id = p.product_id
    join customers c on o.customer_id = c.customer_id
    group by 1,2,3
    order by c.customer_id
)
select customer_id, name from t
group by 1,2
having sum(case when date = '2020-06-01' then sum else 0 end) >= 100
and sum(case when date = '2020-07-01' then sum else 0 end) >= 100

-- solution2
select
c.customer_id
,c.name
-- ,sum(case when date_trunc('month', o.order_date)='2020-06-01' then p.price*o.quantity else 0 end) june
-- ,sum(case when date_trunc('month', o.order_date)='2020-07-01' then p.price*o.quantity else 0 end) july
from customers c, orders o, product p
where c.customer_id = o.customer_id and p.product_id = o.product_id
group by 1,2
having sum(case when date_trunc('month', o.order_date)='2020-06-01' then p.price*o.quantity else 0 end)>=100
and sum(case when date_trunc('month', o.order_date)='2020-07-01' then p.price*o.quantity else 0 end)>=100

