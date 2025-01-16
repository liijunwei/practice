
-- https://leetcode.cn/problems/highest-grade-for-each-student/?envType=study-plan-v2&envId=sql-premium-50

Create table If Not Exists Enrollments (student_id int, course_id int, grade int);
Truncate table Enrollments;
insert into Enrollments (student_id, course_id, grade) values ('2', '2', '95');
insert into Enrollments (student_id, course_id, grade) values ('2', '3', '95');
insert into Enrollments (student_id, course_id, grade) values ('1', '1', '90');
insert into Enrollments (student_id, course_id, grade) values ('1', '2', '99');
insert into Enrollments (student_id, course_id, grade) values ('3', '1', '80');
insert into Enrollments (student_id, course_id, grade) values ('3', '2', '75');
insert into Enrollments (student_id, course_id, grade) values ('3', '3', '82');



-- solution1
select DISTINCT ON (pri.student_id) -- this is new to me
pri.student_id
,pri.course_id
,pri.grade
from
enrollments pri
where
pri.grade = (select max(grade) from enrollments sub where sub.student_id = pri.student_id)
order by
pri.student_id asc
,pri.course_id asc;


-- solution2, better
with t as(
    select *, rank() over(partition by student_id order by grade desc, course_id asc) rk from enrollments
)
select
student_id
,course_id
,grade
from
t
where rk=1
order by student_id asc
