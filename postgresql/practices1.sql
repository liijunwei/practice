-- login with psql
-- export PGPASSWORD=
-- psql -U postgres -d analysis

select * from teachers;

select id from teachers;


-- create database foo;

SELECT DISTINCT school FROM teachers
-- distinct keyword works for more than 1 column at a time, it returns each unique pair of values
SELECT DISTINCT school,salary FROM teachers ORDER BY school DESC


select first_name,last_name,salary from teachers ORDER BY salary DESC
show all

select first_name,last_name,hire_date from teachers ORDER BY school ASC, hire_date DESC


SELECT first_name
FROM teachers
WHERE first_name LIKE 'sam%';

SELECT first_name
FROM teachers
WHERE first_name ILIKE 'sam%'; -- ignore case like

SELECT first_name, last_name, school, hire_date, salary
FROM teachers
WHERE school LIKE '%Roos%'
ORDER BY hire_date DESC;

SELECT first_name, last_name, school, hire_date, salary
FROM teachers
WHERE school ILIKE '%roos%'
ORDER BY hire_date DESC;

SELECT first_name, last_name, school, hire_date, salary
FROM teachers
WHERE school ILIKE '%roos'
ORDER BY hire_date DESC;

SELECT DISTINCT school, last_name FROM teachers ORDER BY school ASC, last_name ASC

select * from teachers where first_name like 'S%' and salary > 40000

select * from teachers where hire_date > '2010-01-01' order by salary desc

create table char_data_types (
	varchar_column varchar(10),
	char_column char(10),
	text_column text
)

INSERT INTO char_data_types
VALUES
    ('abc', 'abc', 'abc'),
    ('defghi', 'defghi', 'defghi');

select * from char_data_types

copy char_data_types to '/tmp/typetest.txt' with (format CSV,HEADER,DELIMITER '|')
copy char_data_types to '/tmp/typetest.txt' with (format CSV,HEADER,DELIMITER ',')

select * from teachers_id_seq
select * from teachers

create table people(
	id serial,
	person_name varchar(100)
)

insert into people (person_name) 
values ('name1'),
	   ('name2'),
	   ('name3')

CREATE TABLE number_data_types (
    numeric_column numeric(20,5),
    real_column real,
    double_column double precision
);


INSERT INTO number_data_types
VALUES
    (.7, .7, .7),
    (2.13579, 2.13579, 2.13579),
    (2.1357987654, 2.1357987654, 2.1357987654);

SELECT * FROM number_data_types;

SELECT
    numeric_column * 10000000 AS "Fixed",
    real_column * 10000000 AS "Float"
FROM number_data_types
WHERE numeric_column = .7;

create table date_time_types (
	timestamp_column timestamp with time zone,
	interval_column interval
)

delete from date_time_types
select * from date_time_types;
insert into date_time_types values ('2024-10-06 10:51:00 UTC+8', '2 days');
insert into date_time_types values ('2024-10-06 10:51:00 UTC+8', '2 day');
insert into date_time_types values ('2024-10-06 10:51:00 UTC+8', '2 month');
insert into date_time_types values (now(), '2 week');

INSERT INTO date_time_types
VALUES
    ('2018-12-31 01:00 EST','2 days'),
    ('2018-12-31 01:00 PST','1 month'),
    ('2018-12-31 01:00 Australia/Melbourne','1 century'),
    (now(),'1 week');

select timestamp_column, interval_column,
	timestamp_column-interval_column as new_date
from date_time_types;

SELECT *, CAST(timestamp_column AS varchar(10)) as casted
FROM date_time_types;

select * from number_data_types;
SELECT numeric_column,
       CAST(numeric_column AS integer),
       CAST(numeric_column AS varchar(6))
FROM number_data_types;

SELECT CAST(char_column AS integer) FROM char_data_types;

-- obvious
SELECT CAST(timestamp_column AS varchar(10)) FROM date_time_types;
-- less obvious
SELECT timestamp_column::varchar(10) FROM date_time_types;

CREATE TABLE us_counties_2010 (
    geo_name varchar(90),                    -- Name of the geography
    state_us_abbreviation varchar(2),        -- State/U.S. abbreviation
    summary_level varchar(3),                -- Summary Level
    region smallint,                         -- Region
    division smallint,                       -- Division
    state_fips varchar(2),                   -- State FIPS code
    county_fips varchar(3),                  -- County code
    area_land bigint,                        -- Area (Land) in square meters
    area_water bigint,                        -- Area (Water) in square meters
    population_count_100_percent integer,    -- Population count (100%)
    housing_unit_count_100_percent integer,  -- Housing Unit count (100%)
    internal_point_lat numeric(10,7),        -- Internal point (latitude)
    internal_point_lon numeric(10,7),        -- Internal point (longitude)

    -- This section is referred to as P1. Race:
    p0010001 integer,   -- Total population
    p0010002 integer,   -- Population of one race:
    p0010003 integer,       -- White Alone
    p0010004 integer,       -- Black or African American alone
    p0010005 integer,       -- American Indian and Alaska Native alone
    p0010006 integer,       -- Asian alone
    p0010007 integer,       -- Native Hawaiian and Other Pacific Islander alone
    p0010008 integer,       -- Some Other Race alone
    p0010009 integer,   -- Population of two or more races
    p0010010 integer,   -- Population of two races:
    p0010011 integer,       -- White; Black or African American
    p0010012 integer,       -- White; American Indian and Alaska Native
    p0010013 integer,       -- White; Asian
    p0010014 integer,       -- White; Native Hawaiian and Other Pacific Islander
    p0010015 integer,       -- White; Some Other Race
    p0010016 integer,       -- Black or African American; American Indian and Alaska Native
    p0010017 integer,       -- Black or African American; Asian
    p0010018 integer,       -- Black or African American; Native Hawaiian and Other Pacific Islander
    p0010019 integer,       -- Black or African American; Some Other Race
    p0010020 integer,       -- American Indian and Alaska Native; Asian
    p0010021 integer,       -- American Indian and Alaska Native; Native Hawaiian and Other Pacific Islander
    p0010022 integer,       -- American Indian and Alaska Native; Some Other Race
    p0010023 integer,       -- Asian; Native Hawaiian and Other Pacific Islander
    p0010024 integer,       -- Asian; Some Other Race
    p0010025 integer,       -- Native Hawaiian and Other Pacific Islander; Some Other Race
    p0010026 integer,   -- Population of three races
    p0010047 integer,   -- Population of four races
    p0010063 integer,   -- Population of five races
    p0010070 integer,   -- Population of six races

    -- This section is referred to as P2. HISPANIC OR LATINO, AND NOT HISPANIC OR LATINO BY RACE
    p0020001 integer,   -- Total
    p0020002 integer,   -- Hispanic or Latino
    p0020003 integer,   -- Not Hispanic or Latino:
    p0020004 integer,   -- Population of one race:
    p0020005 integer,       -- White Alone
    p0020006 integer,       -- Black or African American alone
    p0020007 integer,       -- American Indian and Alaska Native alone
    p0020008 integer,       -- Asian alone
    p0020009 integer,       -- Native Hawaiian and Other Pacific Islander alone
    p0020010 integer,       -- Some Other Race alone
    p0020011 integer,   -- Two or More Races
    p0020012 integer,   -- Population of two races
    p0020028 integer,   -- Population of three races
    p0020049 integer,   -- Population of four races
    p0020065 integer,   -- Population of five races
    p0020072 integer,   -- Population of six races

    -- This section is referred to as P3. RACE FOR THE POPULATION 18 YEARS AND OVER
    p0030001 integer,   -- Total
    p0030002 integer,   -- Population of one race:
    p0030003 integer,       -- White alone
    p0030004 integer,       -- Black or African American alone
    p0030005 integer,       -- American Indian and Alaska Native alone
    p0030006 integer,       -- Asian alone
    p0030007 integer,       -- Native Hawaiian and Other Pacific Islander alone
    p0030008 integer,       -- Some Other Race alone
    p0030009 integer,   -- Two or More Races
    p0030010 integer,   -- Population of two races
    p0030026 integer,   -- Population of three races
    p0030047 integer,   -- Population of four races
    p0030063 integer,   -- Population of five races
    p0030070 integer,   -- Population of six races

    -- This section is referred to as P4. HISPANIC OR LATINO, AND NOT HISPANIC OR LATINO BY RACE
    -- FOR THE POPULATION 18 YEARS AND OVER
    p0040001 integer,   -- Total
    p0040002 integer,   -- Hispanic or Latino
    p0040003 integer,   -- Not Hispanic or Latino:
    p0040004 integer,   -- Population of one race:
    p0040005 integer,   -- White alone
    p0040006 integer,   -- Black or African American alone
    p0040007 integer,   -- American Indian and Alaska Native alone
    p0040008 integer,   -- Asian alone
    p0040009 integer,   -- Native Hawaiian and Other Pacific Islander alone
    p0040010 integer,   -- Some Other Race alone
    p0040011 integer,   -- Two or More Races
    p0040012 integer,   -- Population of two races
    p0040028 integer,   -- Population of three races
    p0040049 integer,   -- Population of four races
    p0040065 integer,   -- Population of five races
    p0040072 integer,   -- Population of six races

    -- This section is referred to as H1. OCCUPANCY STATUS
    h0010001 integer,   -- Total housing units
    h0010002 integer,   -- Occupied
    h0010003 integer    -- Vacant
);

SELECT * FROM us_counties_2010;

COPY us_counties_2010
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_04/us_counties_2010.csv'
WITH (FORMAT CSV, HEADER);

SELECT geo_name, state_us_abbreviation, area_land
FROM us_counties_2010
ORDER BY area_land DESC
LIMIT 5;

SELECT geo_name, state_us_abbreviation, internal_point_lon
FROM us_counties_2010
ORDER BY internal_point_lon DESC
LIMIT 5;

select geo_name,state_us_abbreviation,summary_level,region,division,state_fips,county_fips,area_land,area_water,population_count_100_percent,housing_unit_count_100_percent,internal_point_lat,internal_point_lon from us_counties_2010

CREATE TABLE supervisor_salaries (
    town varchar(30),
    county varchar(30),
    supervisor varchar(30),
    start_date date,
    salary money,
    benefits money
);

SELECT * FROM supervisor_salaries;

COPY supervisor_salaries (town,supervisor,salary)
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_04/supervisor_salaries.csv'
WITH (FORMAT CSV, HEADER);

DELETE FROM supervisor_salaries;

CREATE TEMPORARY TABLE supervisor_salaries_temp (LIKE supervisor_salaries);
SELECT * FROM supervisor_salaries_temp;
COPY supervisor_salaries_temp (town,supervisor,salary)
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_04/supervisor_salaries.csv'
WITH (FORMAT CSV, HEADER);

INSERT INTO supervisor_salaries (town, county, supervisor, salary)
SELECT town, 'Some County', supervisor, salary
FROM supervisor_salaries_temp;

DROP TABLE supervisor_salaries_temp;
SELECT * FROM supervisor_salaries LIMIT 2;


COPY us_counties_2010
TO '/tmp/us_counties_export.txt'
WITH (FORMAT CSV, HEADER, DELIMITER '|');

COPY us_counties_2010 (geo_name, internal_point_lat, internal_point_lon)
TO '/tmp/us_counties_latlon_export.txt'
WITH (FORMAT CSV, HEADER, DELIMITER '|');

COPY (
    SELECT geo_name, state_us_abbreviation
    FROM us_counties_2010
    WHERE geo_name ILIKE '%mill%'
     )
TO '/tmp/us_counties_mill_export.txt'
WITH (FORMAT CSV, HEADER, DELIMITER '|');


SELECT 2 + 2;
SELECT 11 / 6;
SELECT 11 / 6.0;
SELECT CAST(11 AS numeric(3,1)) / 6;

SELECT 3 ^ 4;
SELECT 2 ^ 15;
select |/16 -- square root of 16 is 4
select |/4  -- square root of 16 is 2

SELECT factorial(40);
SELECT 3 ^ 3 - 1;

SELECT geo_name,
       state_us_abbreviation AS "st",
       p0010001 AS "Total Population",
       p0010003 AS "White Alone",
       p0010004 AS "Black or African American Alone",
       p0010005 AS "Am Indian/Alaska Native Alone",
       p0010006 AS "Asian Alone",
       p0010007 AS "Native Hawaiian and Other Pacific Islander Alone",
       p0010008 AS "Some Other Race Alone",
       p0010009 AS "Two or More Races"
FROM us_counties_2010;

SELECT geo_name,
       state_us_abbreviation AS "st",
       p0010003 AS "White Alone",
       p0010004 AS "Black Alone",
       p0010003 + p0010004 AS "Total White and Black"
FROM us_counties_2010;

SELECT geo_name,
       state_us_abbreviation AS "st",
       p0010001 AS "Total",
       p0010003 + p0010004 + p0010005 + p0010006 + p0010007
           + p0010008 + p0010009 AS "All Races",
       (p0010003 + p0010004 + p0010005 + p0010006 + p0010007
           + p0010008 + p0010009) - p0010001 AS "Difference"
FROM us_counties_2010
ORDER BY "Difference" DESC;

-- the question to ask: "how to reference column aliases defined in a CTE directly in the outer query"?
with t as (
	SELECT geo_name,
	       state_us_abbreviation AS "st",
	       p0010001 AS "Total",
	       p0010003 + p0010004 + p0010005 + p0010006 + p0010007
	           + p0010008 + p0010009 AS "All Races",
	       (p0010003 + p0010004 + p0010005 + p0010006 + p0010007
	           + p0010008 + p0010009) - p0010001 AS "Difference"
	FROM us_counties_2010
)
select distinct "Difference" from t

SELECT geo_name,
       state_us_abbreviation AS "st",
       (CAST(p0010006 AS numeric(8,1)) / p0010001) * 100 AS "pct_asian"
FROM us_counties_2010
ORDER BY "pct_asian" DESC;

CREATE TABLE percent_change (
    department varchar(20),
    spend_2014 numeric(10,2),
    spend_2017 numeric(10,2)
);


INSERT INTO percent_change
VALUES
    ('Building', 250000, 289000),
    ('Assessor', 178556, 179500),
    ('Library', 87777, 90001),
    ('Clerk', 451980, 650000),
    ('Police', 250000, 223000),
    ('Recreation', 199000, 195000);

SELECT department,
       spend_2014,
       spend_2017,
       round( (spend_2017 - spend_2014) /
                    spend_2014 * 100, 1 ) AS "pct_change"
FROM percent_change;

SELECT sum(p0010001) AS "County Sum",
       round(avg(p0010001), 0) AS "County Average"
FROM us_counties_2010;

-- Finding the Median with Percentile Functions
-- In statistics, percentiles indicate the point in an ordered set of data below which a certain percentage of the data is found.

CREATE TABLE percentile_test (
    numbers integer
);

INSERT INTO percentile_test (numbers) VALUES
    (1), (2), (3), (4), (5), (6);

select * from percentile_test;

SELECT
    percentile_cont(.5) -- continuous
    WITHIN GROUP (ORDER BY numbers),
    percentile_disc(.5) -- discrete
    WITHIN GROUP (ORDER BY numbers)
FROM percentile_test;

-- The median and average are far apart, which shows that averages can mislead. 
SELECT sum(p0010001) AS "County Sum",
       round(avg(p0010001), 0) AS "County Average",
       percentile_cont(.5)
       WITHIN GROUP (ORDER BY p0010001) AS "County Median"
FROM us_counties_2010;

-- However, entering values one at a time is laborious if you want to generate multiple cut points. Instead, you can pass values into percentile_cont() using an array, a SQL data type that contains a list of items
SELECT percentile_cont(array[.25,.5,.75])
       WITHIN GROUP (ORDER BY p0010001) AS "quartiles"
FROM us_counties_2010;

SELECT unnest (percentile_cont(array[.25,.5,.75])
       WITHIN GROUP (ORDER BY p0010001)) AS "quartiles"
FROM us_counties_2010;

-- https://www.postgresql.org/docs/current/functions-array.html
select ARRAY[4,5,6];
select ARRAY[1,2,3] || ARRAY[4,5,6,7];
select ARRAY[1,2,3] || ARRAY[[4,5,6],[7,8,9.9]];
select ARRAY[1,4,3] && ARRAY[2,1]  -- true
select ARRAY[1,4,3] && ARRAY[2,10] -- false

SELECT percentile_cont(array[.2,.4,.6,.8])
       WITHIN GROUP (ORDER BY p0010001) AS "quintiles"
FROM us_counties_2010;

SELECT unnest (percentile_cont(array[.1,.2,.3,.4,.5,.6,.7,.8,.9])
       WITHIN GROUP (ORDER BY p0010001)) AS "deciles"
FROM us_counties_2010;

SELECT unnest(
	percentile_cont(array[.25,.5,.75])
	WITHIN GROUP (ORDER BY p0010001)
	) AS "quartiles"
FROM us_counties_2010;

-- define sql function
-- https://wiki.postgresql.org/wiki/Aggregate_Median
CREATE OR REPLACE FUNCTION _final_median(numeric[])
   RETURNS numeric AS
$$
   SELECT AVG(val)
   FROM (
     SELECT val
     FROM unnest($1) val
     ORDER BY 1
     LIMIT  2 - MOD(array_upper($1, 1), 2)
     OFFSET CEIL(array_upper($1, 1) / 2.0) - 1
   ) sub;
$$
LANGUAGE 'sql' IMMUTABLE;

CREATE AGGREGATE median(numeric) (
  SFUNC=array_append,
  STYPE=numeric[],
  FINALFUNC=_final_median,
  INITCOND='{}'
);

SELECT sum(p0010001) AS "County Sum",
       round(avg(p0010001), 0) AS "County Average",
       median(p0010001) AS "County Median",
       percentile_cont(.5)
       WITHIN GROUP (ORDER BY P0010001) AS "50th Percentile"
FROM us_counties_2010;

-- In PostgreSQL, the mode() function is an aggregate function that returns the most frequent value (mode) in a set of values.
SELECT mode() WITHIN GROUP (ORDER BY p0010001)
FROM us_counties_2010;

select pi() * (5 ^ 2);

SELECT median(p0010001) AS "County Median"
FROM us_counties_2010 where state_us_abbreviation = 'CA'
UNION
SELECT median(p0010001) AS "County Median"
FROM us_counties_2010 where state_us_abbreviation = 'NY'
