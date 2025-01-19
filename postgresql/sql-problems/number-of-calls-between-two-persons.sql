create table if not exists calls (from_id int, to_id int, duration int);
truncate table calls;
insert into calls (from_id, to_id, duration) values ('1', '2', '59');
insert into calls (from_id, to_id, duration) values ('2', '1', '11');
insert into calls (from_id, to_id, duration) values ('1', '3', '20');
insert into calls (from_id, to_id, duration) values ('3', '4', '100');
insert into calls (from_id, to_id, duration) values ('3', '4', '200');
insert into calls (from_id, to_id, duration) values ('3', '4', '200');
insert into calls (from_id, to_id, duration) values ('4', '3', '499');

-- TIL: least and greatest function
select
least(from_id,to_id) person1
,greatest(from_id,to_id) person2
,count(*) call_count
,sum(duration) total_duration
from
calls
group by 1,2
