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
psql:
	@PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight

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
