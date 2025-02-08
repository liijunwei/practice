-- \i ~/practice/postgresql/sql-problems/countries-you-can-safely-invest-in.sql

-- 可以放心投资的国家

drop table if exists person;
drop table if exists country;
drop table if exists calls;
create table if not exists person (id int, name varchar(15), phone_number varchar(11));
create table if not exists country (name varchar(15), country_code varchar(3));
create table if not exists calls (caller_id int, callee_id int, duration int);
truncate table person;
insert into person (id, name, phone_number) values ('3', 'jonathan', '051-1234567');
insert into person (id, name, phone_number) values ('12', 'elvis', '051-7654321');
insert into person (id, name, phone_number) values ('1', 'moncef', '212-1234567');
insert into person (id, name, phone_number) values ('2', 'maroua', '212-6523651');
insert into person (id, name, phone_number) values ('7', 'meir', '972-1234567');
insert into person (id, name, phone_number) values ('9', 'rachel', '972-0011100');
truncate table country;
insert into country (name, country_code) values ('peru', '051');
insert into country (name, country_code) values ('israel', '972');
insert into country (name, country_code) values ('morocco', '212');
insert into country (name, country_code) values ('germany', '049');
insert into country (name, country_code) values ('ethiopia', '251');
truncate table calls;
insert into calls (caller_id, callee_id, duration) values ('1', '9', '33');
insert into calls (caller_id, callee_id, duration) values ('2', '9', '4');
insert into calls (caller_id, callee_id, duration) values ('1', '2', '59');
insert into calls (caller_id, callee_id, duration) values ('3', '12', '102');
insert into calls (caller_id, callee_id, duration) values ('3', '12', '330');
insert into calls (caller_id, callee_id, duration) values ('12', '3', '5');
insert into calls (caller_id, callee_id, duration) values ('7', '9', '13');
insert into calls (caller_id, callee_id, duration) values ('7', '1', '3');
insert into calls (caller_id, callee_id, duration) values ('9', '7', '1');
insert into calls (caller_id, callee_id, duration) values ('1', '7', '7');


with t as(
  select caller_id caller,duration from calls
  union all
  select callee_id caller,duration from calls
)
select
c.name country
-- ,avg(t.duration)
from t left join person p on p.id=t.caller
left join country c on left(p.phone_number,3)=c.country_code
group by c.name
having avg(t.duration) > (select avg(duration) from t)
