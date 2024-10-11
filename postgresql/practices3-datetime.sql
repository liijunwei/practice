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

