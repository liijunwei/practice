include .envrc

help:
	go run cmd/api/* -h

run/api:
	source .envrc && go run cmd/api/main.go -limiter-enabled=false -limiter-burst=30 -debug-enabled=true -cors-trusted-origins="http://localhost:9000 http://localhost:9001"

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
	ls db/seeds/*.sql | xargs -I {} sh -c 'PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight --quiet -f {} && echo "seed {}"'
	@echo "done seed"
	@echo

db-reset:
	migrate -path=./db/migrations -database=${GREENLIGHT_DB_DSN} down -all
	@echo "done reset"
	@echo

db-migrate:
	migrate -path=./db/migrations -database=${GREENLIGHT_DB_DSN} up
	@echo "done migrate"
	@echo
	rm ./db/structure.sql
	pg_dump --schema-only ${GREENLIGHT_DB_DSN} --file ./db/structure.sql
	@echo "done pgdump"
	@echo

db-dump-example-data:
	pg_dump --data-only --inserts --table users --no-comments ${GREENLIGHT_DB_DSN}

db-reset-add-seed: db-reset db-migrate db-seed

show-pg-config-file:
	sudo -u postgres psql -c 'show config_file;'
	# /Library/PostgreSQL/17/data/postgresql.conf

install-tools:
	brew install postgres
	brew install golang-migrate
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

create-migration-file-example:
	@echo "migrate create -seq -ext=.sql -dir=./db/migrations create_movies_table"

# test-optimistic-locking
# xargs -I % -P8 curl --location --request PUT 'http://localhost:4000/v1/movies/9' --header 'Content-Type: application/json' --data '{"title":"movie4","year":2007,"runtime":1500,"genres":["crime","war","sci-fi"]}' < <(printf '%s\n' {1..8})

# test timeout with `select pg_sleep(3)`
# curl --write-out '\nTime: %{time_total}s \n' http://localhost:4000/v1/movies/9

test-rate-limiter:
	for i in {1..10}; do curl --location 'http://localhost:4000/v1/healthcheck'; done

BODY='{"email":"Emiliano_Cruickshank@example.com","password":"pa55word"}'
load-test:
	hey -d ${BODY} -m "POST" -z 5s http://localhost:4000/v1/tokens/authentication

sqlc:
	sqlc generate -f db/sqlc.yaml
