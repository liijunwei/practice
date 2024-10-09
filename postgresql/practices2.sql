-- fy: fiscal year 财政年度
CREATE TABLE pls_fy2014_pupld14a (
    stabr varchar(2) NOT NULL,
    fscskey varchar(6) CONSTRAINT fscskey2014_key PRIMARY KEY,
    libid varchar(20) NOT NULL,
    libname varchar(100) NOT NULL,
    obereg varchar(2) NOT NULL,
    rstatus integer NOT NULL,
    statstru varchar(2) NOT NULL,
    statname varchar(2) NOT NULL,
    stataddr varchar(2) NOT NULL,
    longitud numeric(10,7) NOT NULL,
    latitude numeric(10,7) NOT NULL,
    fipsst varchar(2) NOT NULL,
    fipsco varchar(3) NOT NULL,
    address varchar(35) NOT NULL,
    city varchar(20) NOT NULL,
    zip varchar(5) NOT NULL,
    zip4 varchar(4) NOT NULL,
    cnty varchar(20) NOT NULL,
    phone varchar(10) NOT NULL,
    c_relatn varchar(2) NOT NULL,
    c_legbas varchar(2) NOT NULL,
    c_admin varchar(2) NOT NULL,
    geocode varchar(3) NOT NULL,
    lsabound varchar(1) NOT NULL,
    startdat varchar(10),
    enddate varchar(10),
    popu_lsa integer NOT NULL,
    centlib integer NOT NULL,
    branlib integer NOT NULL,
    bkmob integer NOT NULL,
    master numeric(8,2) NOT NULL,
    libraria numeric(8,2) NOT NULL,
    totstaff numeric(8,2) NOT NULL,
    locgvt integer NOT NULL,
    stgvt integer NOT NULL,
    fedgvt integer NOT NULL,
    totincm integer NOT NULL,
    salaries integer,
    benefit integer,
    staffexp integer,
    prmatexp integer NOT NULL,
    elmatexp integer NOT NULL,
    totexpco integer NOT NULL,
    totopexp integer NOT NULL,
    lcap_rev integer NOT NULL,
    scap_rev integer NOT NULL,
    fcap_rev integer NOT NULL,
    cap_rev integer NOT NULL,
    capital integer NOT NULL,
    bkvol integer NOT NULL,
    ebook integer NOT NULL,
    audio_ph integer NOT NULL,
    audio_dl float NOT NULL,
    video_ph integer NOT NULL,
    video_dl float NOT NULL,
    databases integer NOT NULL,
    subscrip integer NOT NULL,
    hrs_open integer NOT NULL,
    visits integer NOT NULL,
    referenc integer NOT NULL,
    regbor integer NOT NULL,
    totcir integer NOT NULL,
    kidcircl integer NOT NULL,
    elmatcir integer NOT NULL,
    loanto integer NOT NULL,
    loanfm integer NOT NULL,
    totpro integer NOT NULL,
    totatten integer NOT NULL,
    gpterms integer NOT NULL,
    pitusr integer NOT NULL,
    wifisess integer NOT NULL,
    yr_sub integer NOT NULL
);

CREATE INDEX libname2014_idx ON pls_fy2014_pupld14a (libname);
CREATE INDEX stabr2014_idx ON pls_fy2014_pupld14a (stabr);
CREATE INDEX city2014_idx ON pls_fy2014_pupld14a (city);
CREATE INDEX visits2014_idx ON pls_fy2014_pupld14a (visits);

COPY pls_fy2014_pupld14a
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_08/pls_fy2014_pupld14a.csv'
WITH (FORMAT CSV, HEADER);

select * from pls_fy2014_pupld14a

CREATE TABLE pls_fy2009_pupld09a (
    stabr varchar(2) NOT NULL,
    fscskey varchar(6) CONSTRAINT fscskey2009_key PRIMARY KEY,
    libid varchar(20) NOT NULL,
    libname varchar(100) NOT NULL,
    address varchar(35) NOT NULL,
    city varchar(20) NOT NULL,
    zip varchar(5) NOT NULL,
    zip4 varchar(4) NOT NULL,
    cnty varchar(20) NOT NULL,
    phone varchar(10) NOT NULL,
    c_relatn varchar(2) NOT NULL,
    c_legbas varchar(2) NOT NULL,
    c_admin varchar(2) NOT NULL,
    geocode varchar(3) NOT NULL,
    lsabound varchar(1) NOT NULL,
    startdat varchar(10),
    enddate varchar(10),
    popu_lsa integer NOT NULL,
    centlib integer NOT NULL,
    branlib integer NOT NULL,
    bkmob integer NOT NULL,
    master numeric(8,2) NOT NULL,
    libraria numeric(8,2) NOT NULL,
    totstaff numeric(8,2) NOT NULL,
    locgvt integer NOT NULL,
    stgvt integer NOT NULL,
    fedgvt integer NOT NULL,
    totincm integer NOT NULL,
    salaries integer,
    benefit integer,
    staffexp integer,
    prmatexp integer NOT NULL,
    elmatexp integer NOT NULL,
    totexpco integer NOT NULL,
    totopexp integer NOT NULL,
    lcap_rev integer NOT NULL,
    scap_rev integer NOT NULL,
    fcap_rev integer NOT NULL,
    cap_rev integer NOT NULL,
    capital integer NOT NULL,
    bkvol integer NOT NULL,
    ebook integer NOT NULL,
    audio integer NOT NULL,
    video integer NOT NULL,
    databases integer NOT NULL,
    subscrip integer NOT NULL,
    hrs_open integer NOT NULL,
    visits integer NOT NULL,
    referenc integer NOT NULL,
    regbor integer NOT NULL,
    totcir integer NOT NULL,
    kidcircl integer NOT NULL,
    loanto integer NOT NULL,
    loanfm integer NOT NULL,
    totpro integer NOT NULL,
    totatten integer NOT NULL,
    gpterms integer NOT NULL,
    pitusr integer NOT NULL,
    yr_sub integer NOT NULL,
    obereg varchar(2) NOT NULL,
    rstatus integer NOT NULL,
    statstru varchar(2) NOT NULL,
    statname varchar(2) NOT NULL,
    stataddr varchar(2) NOT NULL,
    longitud numeric(10,7) NOT NULL,
    latitude numeric(10,7) NOT NULL,
    fipsst varchar(2) NOT NULL,
    fipsco varchar(3) NOT NULL
);

CREATE INDEX libname2009_idx ON pls_fy2009_pupld09a (libname);
CREATE INDEX stabr2009_idx ON pls_fy2009_pupld09a (stabr);
CREATE INDEX city2009_idx ON pls_fy2009_pupld09a (city);
CREATE INDEX visits2009_idx ON pls_fy2009_pupld09a (visits);

COPY pls_fy2009_pupld09a
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_08/pls_fy2009_pupld09a.csv'
WITH (FORMAT CSV, HEADER);

select * from pls_fy2009_pupld09a

-- Aggregate functions combine values from multiple rows and return a single result based on an operation on those values.


SELECT count(*) FROM pls_fy2014_pupld14a; -- 9305
SELECT count(salaries) FROM pls_fy2014_pupld14a; -- 5983

SELECT count(*) FROM pls_fy2009_pupld09a;
SELECT count(salaries) FROM pls_fy2009_pupld09a;

SELECT count(libname) FROM pls_fy2014_pupld14a; -- 9305
SELECT count(DISTINCT libname) FROM pls_fy2014_pupld14a; -- 8515

-- Bonus: find duplicate libnames
SELECT libname, count(libname)
FROM pls_fy2014_pupld14a
GROUP BY libname
ORDER BY count(libname) DESC;

SELECT libname, count(libname) AS lib_total
FROM pls_fy2014_pupld14a
GROUP BY libname
ORDER BY lib_total DESC;

SELECT libname, city, stabr
FROM pls_fy2014_pupld14a
WHERE libname = 'OXFORD PUBLIC LIBRARY';

-- -1 indicates a “nonresponse”
-- -3 indicates “not applicable” 
-- A better alternative for this negative value scenario is to use NULL in rows in the visits column where response data is absent
-- and then create a separate isits_flag column to hold codes explaining why.
-- This technique separates number values from information about them
SELECT max(visits), min(visits)
FROM pls_fy2014_pupld14a;

-- When you use the GROUP BY clause with aggregate functions, you can group results according to the values in one or more columns. 
-- On its own, GROUP BY, which is also part of standard ANSI SQL, eliminates duplicate values from the results, similar to DISTINCT.
-- You’re not limited to grouping just one column. 
SELECT stabr FROM pls_fy2014_pupld14a -- 9305 rows
SELECT stabr FROM pls_fy2014_pupld14a GROUP BY stabr ORDER BY stabr -- 56 rows

SELECT stabr FROM pls_fy2009_pupld09a -- 9299
SELECT DISTINCT stabr FROM pls_fy2009_pupld09a -- 55
SELECT stabr FROM pls_fy2009_pupld09a GROUP BY stabr ORDER BY stabr; -- 55

SELECT distinct city, stabr FROM pls_fy2014_pupld14a ORDER BY city, stabr; -- 9088
SELECT city, stabr FROM pls_fy2014_pupld14a GROUP BY city, stabr ORDER BY city, stabr; -- 9088

SELECT city, stabr, count(*) as total
FROM pls_fy2014_pupld14a
GROUP BY city, stabr
ORDER BY total DESC

SELECT stabr, count(*)
FROM pls_fy2014_pupld14a
GROUP BY stabr
ORDER BY count(*) DESC;

SELECT stabr, stataddr, count(*)
FROM pls_fy2014_pupld14a
GROUP BY stabr, stataddr
ORDER BY stabr ASC, count(*) DESC;

SELECT stabr, stataddr, count(*)
FROM pls_fy2014_pupld14a
GROUP BY stabr -- stataddr" must appear in the GROUP BY clause
ORDER BY stabr ASC, count(*) DESC;

SELECT sum(visits) AS visits_2014
FROM pls_fy2014_pupld14a
WHERE visits >= 0
UNION 
SELECT sum(visits) AS visits_2009
FROM pls_fy2009_pupld09a
WHERE visits >= 0

-- total visits on joined 2014 and 2009 library tables
SELECT sum(pls14.visits) AS visits_2014,
       sum(pls09.visits) AS visits_2009
FROM pls_fy2014_pupld14a pls14 JOIN pls_fy2009_pupld09a pls09
ON pls14.fscskey = pls09.fscskey
WHERE pls14.visits >= 0 AND pls09.visits >= 0;

-- track percent change in library visits by state
SELECT pls14.stabr,
       sum(pls09.visits) AS visits_2009,
       sum(pls14.visits) AS visits_2014,
       round( (CAST(sum(pls14.visits) AS decimal(10,1)) - sum(pls09.visits)) /
                    sum(pls09.visits) * 100, 2 ) AS pct_change
FROM pls_fy2014_pupld14a pls14 JOIN pls_fy2009_pupld09a pls09
ON pls14.fscskey = pls09.fscskey
WHERE pls14.visits >= 0 AND pls09.visits >= 0
GROUP BY pls14.stabr
ORDER BY pct_change DESC;

-- To filter the results of aggregate functions, we need to use the HAVING clause that’s part of standard ANSI SQL.
-- where work on row level
-- aggregate functions work across rows
-- The HAVING clause places conditions on groups created by aggregating. 

-- filter the results of an aggregate query
SELECT pls14.stabr,
       sum(pls09.visits) AS visits_2009,
       sum(pls14.visits) AS visits_2014,
       round( (CAST(sum(pls14.visits) AS decimal(10,1)) - sum(pls09.visits)) /
                    sum(pls09.visits) * 100, 2 ) AS pct_change
FROM pls_fy2014_pupld14a pls14 JOIN pls_fy2009_pupld09a pls09
ON pls14.fscskey = pls09.fscskey
WHERE pls14.visits >= 0 AND pls09.visits >= 0
GROUP BY pls14.stabr
HAVING sum(pls14.visits) > 50000000
ORDER BY pct_change DESC;

-- https://www.postgresql.org/docs/current/functions-datetime.html#FUNCTIONS-DATETIME-TRUNC

SELECT DATE_TRUNC('day', timestamp) AS day,
       SUM(amount) AS total_amount
FROM sales_data
GROUP BY day
ORDER BY day;

SELECT date_trunc('hour', TIMESTAMP '2001-02-16 20:38:40');
SELECT date_trunc('year', TIMESTAMP '2001-02-16 20:38:40');
SELECT date_trunc('day', TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40+00');
SELECT date_trunc('day', TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40+00', 'Australia/Sydney');
SELECT date_trunc('hour', INTERVAL '3 days 02:47:33');

-- chapter9
CREATE TABLE meat_poultry_egg_inspect (
    est_number varchar(50) CONSTRAINT est_number_key PRIMARY KEY,
    company varchar(100),
    street varchar(100),
    city varchar(30),
    st varchar(2),
    zip varchar(5),
    phone varchar(14),
    grant_date date,
    activities text,
    dbas text
);

COPY meat_poultry_egg_inspect
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_09/MPI_Directory_by_Establishment_Name.csv'
WITH (FORMAT CSV, HEADER, DELIMITER ',');

CREATE INDEX company_idx ON meat_poultry_egg_inspect (company);

select * from meat_poultry_egg_inspect
SELECT count(*) FROM meat_poultry_egg_inspect;

SELECT company,
       street,
       city,
       st as state,
       count(*) AS address_count
FROM meat_poultry_egg_inspect
GROUP BY company, street, city, st
HAVING count(*) > 1
ORDER BY company, street, city, st;

SELECT company,
       street,
       city,
       st as state,
       count(*) AS address_count
FROM meat_poultry_egg_inspect
GROUP BY 1,2,3,4
HAVING count(*) > 1
ORDER BY 2,3,4,1

SELECT company,
       street,
       city,
       st as state
FROM meat_poultry_egg_inspect
ORDER BY street, city, st, company;


SELECT st,
       count(*) AS st_count
FROM meat_poultry_egg_inspect
GROUP BY st
ORDER BY st;

SELECT est_number,
       company,
       city,
       st,
       zip
FROM meat_poultry_egg_inspect
WHERE st IS NULL;

SELECT company,
       count(*) AS company_count
FROM meat_poultry_egg_inspect
GROUP BY company
ORDER BY company ASC;

SELECT length(zip),
       count(*) AS length_count
FROM meat_poultry_egg_inspect
GROUP BY length(zip)
ORDER BY length(zip) ASC;

SELECT st,
       count(*) AS st_count
FROM meat_poultry_egg_inspect
WHERE length(zip) < 5
GROUP BY st
ORDER BY st ASC;

-- Listing 9-8: Backing up a table

CREATE TABLE meat_poultry_egg_inspect_backup AS
SELECT * FROM meat_poultry_egg_inspect;
drop table if exists m1; create table m1 as select * from meat_poultry_egg_inspect; select count(*) from m1;
drop table if exists m2; create table m2 as select * from meat_poultry_egg_inspect; select count(*) from m2;
drop table if exists m3; create table m3 as select * from new_york_addresses; select count(*) from new_york_addresses;
drop table if exists m4; create table m4 as select * from new_york_addresses; select count(*) from new_york_addresses;

SELECT 
    (SELECT count(*) FROM meat_poultry_egg_inspect) AS original,
    (SELECT count(*) FROM meat_poultry_egg_inspect_backup) AS backup;

ALTER TABLE meat_poultry_egg_inspect ADD COLUMN st_copy varchar(2);
alter table meat_poultry_egg_inspect drop column st_copy;
alter table meat_poultry_egg_inspect add column st_copy varchar(2)

select count(st_copy) from meat_poultry_egg_inspect
UPDATE meat_poultry_egg_inspect SET st_copy = st;

SELECT st,
       st_copy
FROM meat_poultry_egg_inspect
ORDER BY st;

SELECT st,
       st_copy
FROM meat_poultry_egg_inspect
group by 1,2
ORDER BY st;

SELECT distinct st,
       st_copy
FROM meat_poultry_egg_inspect
ORDER BY st;

SELECT distinct st_copy
FROM meat_poultry_egg_inspect
ORDER BY st_copy;

select distinct st from meat_poultry_egg_inspect where est_number = 'V18677A'
select distinct st from meat_poultry_egg_inspect where est_number = 'M45319+P45319'
select distinct st from meat_poultry_egg_inspect where est_number = 'M263A+P263A+V263A'

update meat_poultry_egg_inspect set st = 'MN' where est_number = 'V18677A'
UPDATE meat_poultry_egg_inspect SET st = 'AL' WHERE est_number = 'M45319+P45319';
UPDATE meat_poultry_egg_inspect SET st = 'WI' WHERE est_number = 'M263A+P263A+V263A';

-- Restoring original st column values
UPDATE meat_poultry_egg_inspect SET st = st_copy;

-- Restoring from the column backup
UPDATE meat_poultry_egg_inspect original
SET st = backup.st
FROM meat_poultry_egg_inspect_backup backup
WHERE original.est_number = backup.est_number;

select distinct company_standard from meat_poultry_egg_inspect where company like 'Armour%'

ALTER TABLE meat_poultry_egg_inspect ADD COLUMN company_standard varchar(100);

UPDATE meat_poultry_egg_inspect SET company_standard = company;

UPDATE meat_poultry_egg_inspect
SET company_standard = 'Armour-Eckrich Meats'
WHERE company LIKE 'Armour%';

SELECT company, company_standard
FROM meat_poultry_egg_inspect
WHERE company LIKE 'Armour%';

ALTER TABLE meat_poultry_egg_inspect ADD COLUMN zip_copy varchar(5);
UPDATE meat_poultry_egg_inspect
SET zip_copy = zip;

SELECT zip, zip_copy
FROM meat_poultry_egg_inspect
WHERE st IN('PR','VI') AND length(zip) = 3;

UPDATE meat_poultry_egg_inspect
SET zip = '00' || zip
WHERE st IN('PR','VI') AND length(zip) = 3;

SELECT zip, zip_copy
FROM meat_poultry_egg_inspect
WHERE st IN('CT','MA','ME','NH','NJ','RI','VT') AND length(zip) = 4;

UPDATE meat_poultry_egg_inspect
SET zip = '0' || zip
WHERE st IN('CT','MA','ME','NH','NJ','RI','VT') AND length(zip) = 4;

------------------

CREATE TABLE state_regions (
    st varchar(2) CONSTRAINT st_key PRIMARY KEY,
    region varchar(20) NOT NULL
);

COPY state_regions
FROM '/Users/lijunwei/OuterGitRepo/practical-sql/Chapter_09/state_regions.csv'
WITH (FORMAT CSV, HEADER, DELIMITER ',');

select * from state_regions;

ALTER TABLE meat_poultry_egg_inspect ADD COLUMN inspection_date date;
ALTER TABLE meat_poultry_egg_inspect drop COLUMN inspection_date

SELECT inspection_date, count(*) from meat_poultry_egg_inspect
group by inspection_date

SELECT inspection_date
FROM meat_poultry_egg_inspect inspect
WHERE EXISTS (SELECT state_regions.region
              FROM state_regions
              WHERE inspect.st = state_regions.st 
              AND state_regions.region = 'New England');

UPDATE meat_poultry_egg_inspect inspect
SET inspection_date = '2019-12-01'
WHERE EXISTS (SELECT state_regions.region
              FROM state_regions
              WHERE inspect.st = state_regions.st 
              AND state_regions.region = 'New England');

SELECT st, inspection_date
FROM meat_poultry_egg_inspect
GROUP BY st, inspection_date
ORDER BY st;

select * FROM meat_poultry_egg_inspect
WHERE st IN('PR','VI');

DELETE FROM meat_poultry_egg_inspect
WHERE st IN('PR','VI');

ALTER TABLE meat_poultry_egg_inspect DROP COLUMN zip_copy;
DROP TABLE meat_poultry_egg_inspect_backup;

-- Demonstrating a transaction block
START TRANSACTION; COMMIT;
START TRANSACTION;

UPDATE meat_poultry_egg_inspect
SET company = 'AGRO Merchantss Oakland LLC'
WHERE company = 'AGRO Merchants Oakland, LLC';

-- view changes
SELECT company
FROM meat_poultry_egg_inspect
WHERE company LIKE 'AGRO%'
ORDER BY company;

ROLLBACK;

SELECT company
FROM meat_poultry_egg_inspect
WHERE company LIKE 'AGRO%'
ORDER BY company;

START TRANSACTION;

UPDATE meat_poultry_egg_inspect
SET company = 'AGRO Merchants Oakland LLC'
WHERE company = 'AGRO Merchants Oakland, LLC';

COMMIT

CREATE TABLE meat_poultry_egg_inspect_backup AS
SELECT *,
       '2018-02-07'::date AS reviewed_date
FROM meat_poultry_egg_inspect;

select * from meat_poultry_egg_inspect_backup;

-- Listing 9-26: Swapping table names using ALTER TABLE
ALTER TABLE meat_poultry_egg_inspect RENAME TO meat_poultry_egg_inspect_temp;
ALTER TABLE meat_poultry_egg_inspect_backup RENAME TO meat_poultry_egg_inspect;
ALTER TABLE meat_poultry_egg_inspect_temp RENAME TO meat_poultry_egg_inspect_backup;

