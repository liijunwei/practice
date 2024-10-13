-- chapter 15 view, function and triggers

-- A view is a virtual table you can create dynamically using a saved query.
-- https://www.postgresql.org/docs/current/sql-creatematerializedview.html

-- can re-run to refresh the view
CREATE OR REPLACE VIEW nevada_counties_pop_2010 AS
    SELECT geo_name,
           state_fips,
           county_fips,
           p0010001 AS pop_2010
    FROM us_counties_2010
    WHERE state_us_abbreviation = 'NV'
    ORDER BY county_fips;

SELECT *
FROM nevada_counties_pop_2010
LIMIT 5;

--  \d+ nevada_counties_pop_2010
drop view nevada_counties_pop_2010

CREATE OR REPLACE VIEW county_pop_change_2010_2000 AS
    SELECT c2010.geo_name,
           c2010.state_us_abbreviation AS st,
           c2010.state_fips,
           c2010.county_fips,
           c2010.p0010001 AS pop_2010,
           c2000.p0010001 AS pop_2000,
           round( (CAST(c2010.p0010001 AS numeric(8,1)) - c2000.p0010001)
               / c2000.p0010001 * 100, 1 ) AS pct_change_2010_2000
    FROM us_counties_2010 c2010 INNER JOIN us_counties_2000 c2000
    ON c2010.state_fips = c2000.state_fips
       AND c2010.county_fips = c2000.county_fips
    ORDER BY c2010.state_fips, c2010.county_fips;

SELECT geo_name,
       st,
       pop_2010,
       pct_change_2010_2000
FROM county_pop_change_2010_2000
WHERE st = 'NV'
LIMIT 5;

CREATE OR REPLACE VIEW employees_tax_dept AS
     SELECT emp_id,
            first_name,
            last_name,
            dept_id
     FROM employees
     -- WHERE dept_id IN (1,2)
	 WHERE dept_id = 1
     ORDER BY emp_id
     WITH LOCAL CHECK OPTION;

SELECT * FROM employees_tax_dept;
SELECT * FROM employees;

INSERT INTO employees_tax_dept (first_name, last_name, dept_id)
VALUES ('Suzanne', 'Legere', 1);

-- fail because the view only shows dept#1, this sql is trying to update dept#2
INSERT INTO employees_tax_dept (first_name, last_name, dept_id)
VALUES ('Jamil', 'White', 2);

UPDATE employees_tax_dept
SET last_name = 'Le Gere'
WHERE emp_id = 5;

UPDATE employees_tax_dept
SET salary = 100000 -- column not in view
WHERE emp_id = 5;

DELETE FROM employees_tax_dept
WHERE emp_id = 5;


-- functions
-- percent_change = (new_value / old_value) / old_value
create or replace function
percent_change(new_value numeric, old_value numeric, decimal_places integer default 1)
returns numeric as
'select round(
	((new_value-old_value)/old_value) * 100, decimal_places
);'
language SQL
immutable
returns null on null input;

SELECT percent_change(null,null);
SELECT percent_change(0, 1, 2);
SELECT percent_change(111, 100); -- use default decimal places
SELECT percent_change(111, 100, 2);
SELECT percent_change(111, 100, 3);

SELECT c2010.geo_name,
       c2010.state_us_abbreviation AS st,
       c2010.p0010001 AS pop_2010,
       percent_change(c2010.p0010001, c2000.p0010001) AS pct_chg_func,
       round( (CAST(c2010.p0010001 AS numeric(8,1)) - c2000.p0010001)
           / c2000.p0010001 * 100, 1 ) AS pct_chg_formula
FROM us_counties_2010 c2010 INNER JOIN us_counties_2000 c2000
ON c2010.state_fips = c2000.state_fips
   AND c2010.county_fips = c2000.county_fips
ORDER BY pct_chg_func DESC
LIMIT 5;


-- We can also use a function to simplify routine updates to data. 
alter table teachers add column personal_days integer;
select first_name,last_name,hire_date,personal_days from teachers;

-- create function with PL/pgSQL(instead of plain SQL)
CREATE OR REPLACE FUNCTION update_personal_days()
RETURNS void AS $$
BEGIN
    UPDATE teachers
    SET personal_days =
        CASE WHEN (now() - hire_date) BETWEEN '5 years'::interval
                                      AND '10 years'::interval THEN 4
             WHEN (now() - hire_date) > '10 years'::interval THEN 5
             ELSE 3
        END;
    RAISE NOTICE 'personal_days updated!';
END;
$$ LANGUAGE plpgsql;

select update_personal_days()

CREATE EXTENSION plpython3u; -- not working, skip

-- A database trigger executes a function whenever a specified event, such as an INSERT, UPDATE, or DELETE, occurs on a table or a view.
-- You can set a trigger to fire before, after, or instead of the event, and you can also set it to fire once for each row affected by the event or just once per operation.

CREATE TABLE grades (
    student_id bigint,
    course_id bigint,
    course varchar(30) NOT NULL,
    grade varchar(5) NOT NULL,
    PRIMARY KEY (student_id, course_id)
);

select * from grades

INSERT INTO grades
VALUES
    (1, 1, 'Biology 2', 'F'),
    (1, 2, 'English 11B', 'D'),
    (1, 3, 'World History 11B', 'C'),
    (1, 4, 'Trig 2', 'B');

CREATE TABLE grades_history (
    student_id bigint NOT NULL,
    course_id bigint NOT NULL,
    change_time timestamp with time zone NOT NULL,
    course varchar(30) NOT NULL,
    old_grade varchar(5) NOT NULL,
    new_grade varchar(5) NOT NULL,
	PRIMARY KEY (student_id, course_id, change_time)
);

CREATE OR REPLACE FUNCTION record_if_grade_changed()
    RETURNS trigger AS
$$
BEGIN
    IF NEW.grade <> OLD.grade THEN
    INSERT INTO grades_history (
        student_id,
        course_id,
        change_time,
        course,
        old_grade,
        new_grade)
    VALUES
        (OLD.student_id,
         OLD.course_id,
         now(),
         OLD.course,
         OLD.grade,
         NEW.grade);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER grades_update
  AFTER UPDATE
  ON grades
  FOR EACH ROW
  EXECUTE PROCEDURE record_if_grade_changed();

select * from grades;
SELECT * FROM grades_history;

UPDATE grades
SET grade = 'C'
WHERE student_id = 1 AND course_id = 1;


CREATE TABLE temperature_test (
    station_name varchar(50),
    observation_date date,
    max_temp integer,
    min_temp integer,
    max_temp_group varchar(40),
	PRIMARY KEY (station_name, observation_date)
);

CREATE OR REPLACE FUNCTION classify_max_temp()
    RETURNS trigger AS
$$
BEGIN
    CASE
       WHEN NEW.max_temp >= 90 THEN
           NEW.max_temp_group := 'Hot';
       WHEN NEW.max_temp BETWEEN 70 AND 89 THEN
           NEW.max_temp_group := 'Warm';
       WHEN NEW.max_temp BETWEEN 50 AND 69 THEN
           NEW.max_temp_group := 'Pleasant';
       WHEN NEW.max_temp BETWEEN 33 AND 49 THEN
           NEW.max_temp_group :=  'Cold';
       WHEN NEW.max_temp BETWEEN 20 AND 32 THEN
           NEW.max_temp_group :=  'Freezing';
       ELSE NEW.max_temp_group :=  'Inhumane';
    END CASE;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER temperature_insert
    BEFORE INSERT
    ON temperature_test
    FOR EACH ROW
    EXECUTE PROCEDURE classify_max_temp();

select * from temperature_test

INSERT INTO temperature_test (station_name, observation_date, max_temp, min_temp)
VALUES
    ('North Station', '1/19/2019', 10, -3),
    ('North Station', '3/20/2019', 28, 19),
    ('North Station', '5/2/2019', 65, 42),
    ('North Station', '8/9/2019', 93, 74);
