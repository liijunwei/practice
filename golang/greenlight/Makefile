help:
	go run cmd/api/* -h

# export PGPASSWORD in shell before `make psql`
# OR
# psql $GREENLIGHT_DB_DSN
# create database greenlight
# \c greenlight;
# create role greenlight with login password '{password}';
# create extension if not exists citext;
# run with postgres user
# GRANT ALL PRIVILEGES ON DATABASE greenlight TO greenlight; # run with postgres user
# GRANT ALL ON SCHEMA public TO greenlight;
db-console:
	PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight

db-seed:
	PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight -a -f db/seeds/*.sql

show-pg-config-file:
	sudo -u postgres psql -c 'show config_file;'
	# /Library/PostgreSQL/17/data/postgresql.conf

install-tools:
	brew install postgres
	brew install golang-migrate

create-migration-file-example:
	@echo "migrate create -seq -ext=.sql -dir=./db/migrations create_movies_table"

migrate:
	migrate -path=./db/migrations -database=${GREENLIGHT_DB_DSN} up
	pg_dump --schema-only ${GREENLIGHT_DB_DSN} --file ./db/structure.sql

# test-optimistic-locking
# xargs -I % -P8 curl --location --request PUT 'http://localhost:4000/v1/movies/9' --header 'Content-Type: application/json' --data '{"title":"movie4","year":2007,"runtime":1500,"genres":["crime","war","sci-fi"]}' < <(printf '%s\n' {1..8})

# test timeout with `select pg_sleep(3)`
# curl --write-out '\nTime: %{time_total}s \n' http://localhost:4000/v1/movies/9
