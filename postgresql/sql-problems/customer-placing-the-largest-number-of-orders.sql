drop table if exists orders;
create table if not exists orders (order_number int, customer_number int);
truncate table orders;
insert into orders (order_number, customer_number) values ('1', '1');
insert into orders (order_number, customer_number) values ('2', '2');
insert into orders (order_number, customer_number) values ('3', '3');
insert into orders (order_number, customer_number) values ('4', '3');
insert into orders (order_number, customer_number) values ('5', '2');

select
customer_number
-- ,count(*)
from orders
group by 1
order by count(*) desc
limit 1

-- 支持有多位顾客订单数并列最多
with t as(
    select
    rank() over (order by count(*) desc) rk
    ,customer_number
    ,count(*) order_count
    from
    orders
    group by 2
    order by order_count desc
)
select customer_number from t where rk = 1
