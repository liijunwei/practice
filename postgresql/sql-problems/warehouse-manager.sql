create table if not exists warehouse (name varchar(50), product_id int, units int);
create table if not exists products (product_id int, product_name varchar(50), width int,length int,height int);
truncate table warehouse;
insert into warehouse (name, product_id, units) values ('lchouse1', '1', '1');
insert into warehouse (name, product_id, units) values ('lchouse1', '2', '10');
insert into warehouse (name, product_id, units) values ('lchouse1', '3', '5');
insert into warehouse (name, product_id, units) values ('lchouse2', '1', '2');
insert into warehouse (name, product_id, units) values ('lchouse2', '2', '2');
insert into warehouse (name, product_id, units) values ('lchouse3', '4', '1');
truncate table products;
insert into products (product_id, product_name, width, length, height) values ('1', 'lc-tv', '5', '50', '40');
insert into products (product_id, product_name, width, length, height) values ('2', 'lc-keychain', '5', '5', '5');
insert into products (product_id, product_name, width, length, height) values ('3', 'lc-phone', '2', '10', '10');
insert into products (product_id, product_name, width, length, height) values ('4', 'lc-t-shirt', '4', '10', '20');

-- CTE
with t as(
    select
    product_id
    ,(width*length*height) volume_per_unit
    from
    products
)
select
w.name warehouse_name
,sum(w.units*t.volume_per_unit) volume
from warehouse w
join t on w.product_id = t.product_id
group by 1

-- more straightforward
select
w.name warehouse_name
,sum(w.units * p.width*p.length*p.height) volume
from warehouse w
join products p on w.product_id = p.product_id
group by 1
