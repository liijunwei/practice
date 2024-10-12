-- chapter 11
-- ISO 8601 is the international standard format and also the default PostgreSQL date output.

-- date 2024-10-11
-- time 09:14:00
-- timestamp 2024-10-11 09:14:00
SELECT
    date_part('year', '2019-12-01 18:37:12 EST'::timestamptz) AS "year",
    date_part('month', '2019-12-01 18:37:12 EST'::timestamptz) AS "month",
    date_part('day', '2019-12-01 18:37:12 EST'::timestamptz) AS "day",
    date_part('hour', '2019-12-01 18:37:12 EST'::timestamptz) AS "hour",
    date_part('minute', '2019-12-01 18:37:12 EST'::timestamptz) AS "minute",
    date_part('seconds', '2019-12-01 18:37:12 EST'::timestamptz) AS "seconds",
    date_part('timezone_hour', '2019-12-01 18:37:12 EST'::timestamptz) AS "tz",
    date_part('week', '2019-12-01 18:37:12 EST'::timestamptz) AS "week",
    date_part('quarter', '2019-12-01 18:37:12 EST'::timestamptz) AS "quarter",
    date_part('epoch', '2019-12-01 18:37:12 EST'::timestamptz) AS "epoch";

select date_part('week', '2024-12-29 18:37:12 EST'::timestamptz) AS "week"


-- unclear yet
-- Each UTC offset can refer to multiple named time zones plus standard and daylight saving time variants.
SELECT extract('year' from '2019-12-01 18:37:12 EST'::timestamptz) AS "year";

SELECT make_date(2018, 2, 22);
SELECT make_time(18, 4, 30.3);
SELECT make_time(18, 4, 30);
-- You can specify time zones in three different formats:
--   its UTC offset
--   an area/location designator
--   a standard abbreviation.
SELECT make_timestamptz(2018, 2, 22, 18, 4, 30.3, 'Europe/Lisbon');
SELECT make_timestamptz(2018, 2, 22, 18, 4, 30, 'Europe/Lisbon');

show timezone

SELECT
    current_date,
    current_time,
    current_timestamp,
    localtime,
    localtimestamp,
    now();

select clock_timestamp();

CREATE TABLE current_time_example (
    time_id bigserial,
    current_timestamp_col timestamp with time zone,
    clock_timestamp_col timestamp with time zone
);

INSERT INTO current_time_example
        (current_timestamp_col, clock_timestamp_col)
(SELECT current_timestamp,      clock_timestamp()
FROM generate_series(1,100_000));

truncate table current_time_example;
select * from current_time_example;
select count(distinct current_timestamp_col) from current_time_example; -- 1 row
select count(distinct clock_timestamp_col) from current_time_example; -- 100000 rows

-- https://www.postgresql.org/docs/current/functions-srf.html
SELECT * FROM generate_series(1,400000_000); -- too large

SELECT pid FROM pg_stat_activity WHERE state = 'active' and query like 'SELECT * FROM generate_series%';
SELECT pg_cancel_backend(28323);

-- When working with time zones in SQL, you first need know the time zone setting for your database server.
show timezone

-- TODO what is pg_catelog ?
-- \dp pg_catalog.*
select * from pg_indexes;
select * from pg_tables;

SELECT * FROM pg_timezone_abbrevs; -- this is a table
SELECT * FROM pg_timezone_names; -- this is a table
-- beijing +8
-- EDT -4
-- beijing time -> US time <=> 北京时间减12h
SELECT * FROM pg_timezone_names where name ilike '%Eastern%';

SELECT * FROM pg_timezone_names WHERE name LIKE 'Europe%';

SET timezone TO 'US/Pacific';
set timezone to 'US/Eastern';  select now();


CREATE TABLE time_zone_test (
    test_date timestamp with time zone
);
INSERT INTO time_zone_test VALUES ('2020-01-01 4:00');
SELECT * FROM time_zone_test;

SELECT '9/30/1929'::date - '9/27/1929'::date as diff_in_day;
SELECT '9/30/1929'::date + '5 years'::interval as some_date;


CREATE TABLE nyc_yellow_taxi_trips_2016_06_01 (
    trip_id bigserial PRIMARY KEY,
    vendor_id varchar(1) NOT NULL,
    tpep_pickup_datetime timestamp with time zone NOT NULL,
    tpep_dropoff_datetime timestamp with time zone NOT NULL,
    passenger_count integer NOT NULL,
    trip_distance numeric(8,2) NOT NULL,
    pickup_longitude numeric(18,15) NOT NULL,
    pickup_latitude numeric(18,15) NOT NULL,
    rate_code_id varchar(2) NOT NULL,
    store_and_fwd_flag varchar(1) NOT NULL,
    dropoff_longitude numeric(18,15) NOT NULL,
    dropoff_latitude numeric(18,15) NOT NULL,
    payment_type varchar(1) NOT NULL,
    fare_amount numeric(9,2) NOT NULL,
    extra numeric(9,2) NOT NULL,
    mta_tax numeric(5,2) NOT NULL,
    tip_amount numeric(9,2) NOT NULL,
    tolls_amount numeric(9,2) NOT NULL,
    improvement_surcharge numeric(9,2) NOT NULL,
    total_amount numeric(9,2) NOT NULL
);

-- In the COPY statement, we provide the names of columns because the input CSV file doesn’t include the trip_id column that exists in the target table. 
COPY nyc_yellow_taxi_trips_2016_06_01 (
    vendor_id,
    tpep_pickup_datetime,
    tpep_dropoff_datetime,
    passenger_count,
    trip_distance,
    pickup_longitude,
    pickup_latitude,
    rate_code_id,
    store_and_fwd_flag,
    dropoff_longitude,
    dropoff_latitude,
    payment_type,
    fare_amount,
    extra,
    mta_tax,
    tip_amount,
    tolls_amount,
    improvement_surcharge,
    total_amount
   )
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_11/yellow_tripdata_2016_06_01.csv'
WITH (FORMAT CSV, HEADER, DELIMITER ',');

CREATE INDEX tpep_pickup_idx
ON nyc_yellow_taxi_trips_2016_06_01 (tpep_pickup_datetime);

SELECT count(*) FROM nyc_yellow_taxi_trips_2016_06_01;
SELECT * FROM nyc_yellow_taxi_trips_2016_06_01;

-- The Busiest Time of Day(18:00-23:00)
SELECT
    date_part('hour', tpep_pickup_datetime) AS trip_hour,
    count(*)
FROM nyc_yellow_taxi_trips_2016_06_01
GROUP BY trip_hour
ORDER BY trip_hour desc;

with t as (
	SELECT
	    date_part('hour', tpep_pickup_datetime) AS trip_hour,
	    count(*) total
	FROM nyc_yellow_taxi_trips_2016_06_01
	GROUP BY trip_hour
	ORDER BY trip_hour desc
)
select unnest(
	percentile_cont(array[.1,.2,.3,.4,.5,.6,.7,.8,.9])
	within group (order by trip_hour)
)  as "deciles" from t;

-- Exporting taxi pickups per hour to a CSV file
COPY
    (SELECT
        date_part('hour', tpep_pickup_datetime) AS trip_hour,
        count(*)
    FROM nyc_yellow_taxi_trips_2016_06_01
    GROUP BY trip_hour
    ORDER BY trip_hour
    )
TO '/tmp/hourly_pickups_2016_06_01.csv'
WITH (FORMAT CSV, HEADER, DELIMITER ',');

-- When Do Trips Take the Longest?
-- 查询行程耗时的中位数
SELECT
    date_part('hour', tpep_pickup_datetime) AS trip_hour,
    percentile_cont(.5)
        WITHIN GROUP (
			ORDER BY tpep_dropoff_datetime - tpep_pickup_datetime
		) AS median_trip
FROM nyc_yellow_taxi_trips_2016_06_01
GROUP BY trip_hour
ORDER BY trip_hour;

-- 中位数
SELECT
    date_part('hour', tpep_pickup_datetime) AS trip_hour,
    percentile_cont(.5)
        WITHIN GROUP (
			ORDER BY tpep_dropoff_datetime - tpep_pickup_datetime
		) AS median_trip,
    percentile_cont(.5)
        WITHIN GROUP (
			ORDER BY (extract (epoch from tpep_dropoff_datetime - tpep_pickup_datetime) )
		) AS median_trip_in_seconds
FROM nyc_yellow_taxi_trips_2016_06_01
GROUP BY trip_hour
ORDER BY trip_hour;

-- 80th
SELECT
    date_part('hour', tpep_pickup_datetime) AS trip_hour,
    percentile_cont(.9)
        WITHIN GROUP (
			ORDER BY tpep_dropoff_datetime - tpep_pickup_datetime
		) AS median_trip,
    percentile_cont(.8)
        WITHIN GROUP (
			ORDER BY (extract (epoch from tpep_dropoff_datetime - tpep_pickup_datetime) )
		) AS median_trip_in_seconds
FROM nyc_yellow_taxi_trips_2016_06_01
GROUP BY trip_hour
ORDER BY trip_hour;

SELECT
    date_part('hour', tpep_pickup_datetime) AS trip_hour,
    percentile_cont(.9)
        WITHIN GROUP (
			ORDER BY tpep_dropoff_datetime - tpep_pickup_datetime
		) AS median_trip,
    percentile_cont(.8)
        WITHIN GROUP (
			ORDER BY (date_part('epoch', tpep_dropoff_datetime - tpep_pickup_datetime))
		) AS median_trip_in_seconds
FROM nyc_yellow_taxi_trips_2016_06_01
GROUP BY trip_hour
ORDER BY trip_hour;

SET timezone TO 'US/Central';

CREATE TABLE train_rides (
    trip_id bigserial PRIMARY KEY,
    segment varchar(50) NOT NULL,
    departure timestamp with time zone NOT NULL,
    arrival timestamp with time zone NOT NULL
);

INSERT INTO train_rides (segment, departure, arrival)
VALUES
    ('Chicago to New York', '2017-11-13 21:30 CST', '2017-11-14 18:23 EST'),
    ('New York to New Orleans', '2017-11-15 14:15 EST', '2017-11-16 19:32 CST'),
    ('New Orleans to Los Angeles', '2017-11-17 13:45 CST', '2017-11-18 9:00 PST'),
    ('Los Angeles to San Francisco', '2017-11-19 10:10 PST', '2017-11-19 21:24 PST'),
    ('San Francisco to Denver', '2017-11-20 9:10 PST', '2017-11-21 18:38 MST'),
    ('Denver to Chicago', '2017-11-22 19:10 MST', '2017-11-23 14:50 CST');

SELECT * FROM train_rides;

-- https://www.postgresql.org/docs/current/functions-formatting.html
SELECT segment,
       to_char(departure, 'YYYY-MM-DD HH12:MI a.m. TZ') AS departure,
       arrival - departure AS segment_time,
	   extract (epoch from arrival - departure) as segment_time_in_second
FROM train_rides;

SELECT segment,
       arrival - departure AS segment_time,
       sum(arrival - departure) OVER (ORDER BY trip_id) AS cume_time
FROM train_rides;

SELECT distinct segment FROM train_rides;

SELECT segment,
       arrival - departure AS segment_time,
       sum(date_part('epoch', (arrival - departure)))
           OVER (ORDER BY trip_id) * interval '1 second' AS cume_time
FROM train_rides;

-- chapter12
-- A subquery is nested inside another query.

-- Say you wanted to write a query to show which U.S. counties are at or above the 90th percentile, or top 10 percent, for population.
-- You can convert the CTE to a scalar subquery if it makes the query more straightforward
SELECT geo_name,
       state_us_abbreviation,
       p0010001
FROM us_counties_2010
WHERE p0010001 >= (
	    SELECT percentile_cont(.9) WITHIN GROUP (ORDER BY p0010001)
	    FROM us_counties_2010
    )
ORDER BY p0010001 DESC;

-- unclear about the difference yet
with t(percentile_value) as (
	    SELECT percentile_cont(.9) WITHIN GROUP (ORDER BY p0010001)
	    FROM us_counties_2010
    )
SELECT geo_name,
       state_us_abbreviation,
       p0010001
FROM us_counties_2010
WHERE p0010001 >= (select percentile_value from t)
ORDER BY p0010001 DESC;

-- join the CTE table(NICE)
WITH t AS (
    SELECT percentile_cont(.9) WITHIN GROUP (ORDER BY p0010001) AS percentile_value
    FROM us_counties_2010
)
SELECT geo_name, state_us_abbreviation, p0010001
FROM us_counties_2010
JOIN t ON p0010001 >= t.percentile_value
ORDER BY p0010001 DESC;

-- join CTE without any condition
WITH t AS (
    SELECT percentile_cont(.9) WITHIN GROUP (ORDER BY p0010001) AS percentile_value
    FROM us_counties_2010
)
SELECT geo_name, state_us_abbreviation, p0010001, t.percentile_value
FROM us_counties_2010
JOIN t ON true
ORDER BY p0010001 DESC;

-- TODO Q: when to use subquery and when to use CTE?

drop table us_counties_2010_top10;

CREATE TABLE us_counties_2010_top10 AS
SELECT * FROM us_counties_2010;


DELETE FROM us_counties_2010_top10
WHERE p0010001 < (
    SELECT percentile_cont(.9) WITHIN GROUP (ORDER BY p0010001)
    FROM us_counties_2010_top10
    );

SELECT count(*) FROM us_counties_2010_top10;

-- If your subquery returns rows and columns of data, you can convert that data to a table by placing it in a FROM clause, the result of which is known as a derived table.
-- A derived table behaves just like any other table, so you can query it or join it to other tables, even other derived tables.
SELECT round(calcs.average, 0) as average,
       calcs.median,
	   round(calcs.average - calcs.median, 2) AS average_median_diff,
	   round(calcs.average / calcs.median, 2) AS average_median_ratio
FROM (
     SELECT avg(p0010001) AS average,
            percentile_cont(.5)
                WITHIN GROUP (ORDER BY p0010001)::numeric(10,1) AS median
     FROM us_counties_2010
     )
AS calcs;

-- Joining Derived Tables
-- Because derived tables behave like regular tables, you can join them.
SELECT census.st AS st,
       census.st_population,
       plants.plant_count,
       round((plants.plant_count/census.st_population::numeric(10,1)) * 1000000, 1)
           AS plants_per_million
FROM
    (
         SELECT st,
                count(*) AS plant_count
         FROM meat_poultry_egg_inspect
         GROUP BY st
    )
    AS plants
JOIN
    (
        SELECT state_us_abbreviation st,
               sum(p0010001) AS st_population
        FROM us_counties_2010
        GROUP BY st
    )
    AS census
ON plants.st = census.st
ORDER BY plants_per_million DESC;

-- rewrite with CTE
with
plants as (
	SELECT st, count(*) AS plant_count
	FROM meat_poultry_egg_inspect
	GROUP BY st
),
census as (
	SELECT state_us_abbreviation as st, sum(p0010001) AS st_population
	FROM us_counties_2010
	GROUP BY st
)
SELECT census.st AS st,
       census.st_population,
       plants.plant_count,
       round((plants.plant_count/census.st_population::numeric(10,1)) * 1000000, 1)
           AS plants_per_million
FROM plants JOIN census
ON plants.st = census.st
ORDER BY plants_per_million DESC;

-- generating columns with subqueires
SELECT geo_name,
       state_us_abbreviation AS st,
       p0010001 AS total_pop,
       (SELECT percentile_cont(.5) WITHIN GROUP (ORDER BY p0010001)
        FROM us_counties_2010) AS us_median
FROM us_counties_2010;

with t(us_median) as (
	SELECT percentile_cont(.5) WITHIN GROUP (ORDER BY p0010001)
	FROM us_counties_2010
)
SELECT geo_name,
       state_us_abbreviation AS st,
       p0010001 AS total_pop,
       t.us_median
FROM us_counties_2010
join t on true;

 -- looks like cross join is another option, but not sure which one is easier to understand
with t(us_median) as (
	SELECT percentile_cont(.5) WITHIN GROUP (ORDER BY p0010001)
	FROM us_counties_2010
)
SELECT geo_name,
       state_us_abbreviation AS st,
       p0010001 AS total_pop,
       t.us_median
FROM us_counties_2010
cross join t;

-- calculate the diff between total and mdeian
SELECT geo_name,
       state_us_abbreviation AS st,
       p0010001 AS total_pop,
       (SELECT percentile_cont(.5) WITHIN GROUP (ORDER BY p0010001)
        FROM us_counties_2010) AS us_median,
       p0010001 - (SELECT percentile_cont(.5) WITHIN GROUP (ORDER BY p0010001)
                   FROM us_counties_2010) AS diff_from_median
FROM us_counties_2010
WHERE (p0010001 - (SELECT percentile_cont(.5) WITHIN GROUP (ORDER BY p0010001)
                   FROM us_counties_2010))
       BETWEEN -1000 AND 1000;

CREATE TABLE retirees (
    id int,
    first_name varchar(50),
    last_name varchar(50)
);

truncate table retirees;
drop table retirees;
select * from retirees;

INSERT INTO retirees 
VALUES (2, 'Lee', 'Smith'),
       (4, 'Janet', 'King');

-- Generating values for the IN operator
SELECT first_name, last_name
FROM employees
WHERE emp_id IN (
    SELECT id
    FROM retirees);

SELECT first_name, last_name
FROM employees
WHERE EXISTS ( -- EXISTS check the boolean result
    SELECT id
    FROM retirees);

-- mimic the behavior of IN query above
-- This approach is particularly helpful if you need to join on more than one column, which you can’t do with the IN expression.
SELECT first_name, last_name
FROM employees
WHERE EXISTS (
    SELECT id
    FROM retirees
    WHERE id = employees.emp_id);

SELECT first_name, last_name
FROM employees
WHERE NOT EXISTS (
    SELECT id
    FROM retirees
    WHERE id = employees.emp_id);

-- CTE: Common Table Expression
-- Also, the column list is optional if you’re not renaming columns, although including the list is still a good idea for clarity even if you don’t rename columns.
WITH large_counties (geo_name, st, p0010001) AS (
	SELECT geo_name, state_us_abbreviation, p0010001
	FROM us_counties_2010
	WHERE p0010001 >= 100000
)
SELECT st, count(*)
FROM large_counties
GROUP BY st
ORDER BY count(*) DESC;

SELECT state_us_abbreviation, count(*)
FROM us_counties_2010
WHERE p0010001 >= 100000
GROUP BY state_us_abbreviation
ORDER BY count(*) DESC;

-- why use CTE?
-- 1. pre-stage subsets of data to feed into larger query for more complex analysis, can reuse the table defined by CTE
-- 2. improve readability
-- 3. use CTE to simplify queries with redundant code
WITH
    counties (st, population) AS
    (SELECT state_us_abbreviation, sum(population_count_100_percent)
     FROM us_counties_2010
     GROUP BY state_us_abbreviation),
    plants (st, plants) AS
    (SELECT st, count(*) AS plants
     FROM meat_poultry_egg_inspect
     GROUP BY st)
SELECT counties.st,
       population,
       plants,
       round((plants/population::numeric(10,1))*1000000, 1) AS per_million
FROM counties JOIN plants
ON counties.st = plants.st
ORDER BY per_million DESC;

WITH us_median AS 
    (SELECT percentile_cont(.5)
     WITHIN GROUP (ORDER BY p0010001) AS us_median_pop
     FROM us_counties_2010)
SELECT geo_name,
       state_us_abbreviation AS st,
       p0010001 AS total_pop,
       us_median_pop,
       p0010001 - us_median_pop AS diff_from_median 
FROM us_counties_2010 CROSS JOIN us_median -- make sure "cross join" here is well understood
WHERE (p0010001 - us_median_pop)
       BETWEEN -1000 AND 1000;

-- Cross Tabulations
-- https://www.postgresql.org/docs/current/tablefunc.html
create extension tablefunc;

CREATE TABLE ice_cream_survey (
    response_id integer PRIMARY KEY,
    office varchar(20),
    flavor varchar(20)
);

copy ice_cream_survey
from '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_12/ice_cream_survey.csv'
with (FORMAT CSV, HEADER)

select * from ice_cream_survey

-- Generating the ice cream survey crosstab
-- a different/summarized view of the first query...
-- 数据透视表 in excel, steps:
-- 1. download query result
-- 2. open csv with wps/excel
-- 3. menu -> insert -> pivot table
-- 4. select office as row
-- 5. select flavor as column
-- 6. select count as value
-- 7. you'll get the same result
SELECT *
FROM crosstab('SELECT office,
                      flavor,
                      count(*)
               FROM ice_cream_survey
               GROUP BY office, flavor
               ORDER BY office',

              'SELECT distinct flavor
               FROM ice_cream_survey
               ORDER BY flavor')

AS (office varchar(20),
    chocolate bigint,
    strawberry bigint,
    vanilla bigint);

CREATE TABLE temperature_readings (
    reading_id bigserial PRIMARY KEY,
    station_name varchar(50),
    observation_date date,
    max_temp integer,
    min_temp integer
);

COPY temperature_readings 
     (station_name, observation_date, max_temp, min_temp)
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_12/temperature_readings.csv'
WITH (FORMAT CSV, HEADER);

select * from temperature_readings;

-- data source
SELECT
	station_name, date_part('month', observation_date),
	percentile_cont(.5)
	WITHIN GROUP (ORDER BY max_temp)
FROM temperature_readings
GROUP BY station_name,
date_part('month', observation_date)
ORDER BY station_name

-- Crosstabs do take time to set up, but viewing data sets in a matrix often makes comparisons easier than viewing the same data in a vertical list.
-- Keep in mind that the crosstab() function is CPU-intensive, so tread carefully when querying sets that have millions or billions of rows.
SELECT *
FROM crosstab('SELECT
                  station_name,
                  date_part(''month'', observation_date),
                  percentile_cont(.5)
                      WITHIN GROUP (ORDER BY max_temp)
               FROM temperature_readings
               GROUP BY station_name,
                        date_part(''month'', observation_date)
               ORDER BY station_name',

              'SELECT month
               FROM generate_series(1,12) month')

AS (station varchar(50),
    jan numeric(3,0),
    feb numeric(3,0),
    mar numeric(3,0),
    apr numeric(3,0),
    may numeric(3,0),
    jun numeric(3,0),
    jul numeric(3,0),
    aug numeric(3,0),
    sep numeric(3,0),
    oct numeric(3,0),
    nov numeric(3,0),
    dec numeric(3,0)
);

SELECT max_temp,
       CASE WHEN max_temp >= 90 THEN 'Hot'
            WHEN max_temp BETWEEN 70 AND 89 THEN 'Warm'
            WHEN max_temp BETWEEN 50 AND 69 THEN 'Pleasant'
            WHEN max_temp BETWEEN 33 AND 49 THEN 'Cold'
            WHEN max_temp BETWEEN 20 AND 32 THEN 'Freezing'
            ELSE 'Inhumane'
        END AS temperature_group
FROM temperature_readings;

WITH temps_collapsed (station_name, max_temperature_group) AS
    (SELECT station_name,
           CASE WHEN max_temp >= 90 THEN 'Hot'
                WHEN max_temp BETWEEN 70 AND 89 THEN 'Warm'
                WHEN max_temp BETWEEN 50 AND 69 THEN 'Pleasant'
                WHEN max_temp BETWEEN 33 AND 49 THEN 'Cold'
                WHEN max_temp BETWEEN 20 AND 32 THEN 'Freezing'
                ELSE 'Inhumane'
            END
    FROM temperature_readings)

SELECT station_name, max_temperature_group, count(*)
FROM temps_collapsed
GROUP BY station_name, max_temperature_group
ORDER BY station_name, count(*) DESC;
