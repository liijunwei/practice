.PHONY: help confirm run-api db-console db-seed db-reset db-migrate db-dump-example-data show-pg-config-file install-tools golang-migrate-example test-rate-limiter load-test sqlc sse-example

include .envrc

help: ## show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

confirm: ## confirm before proceed
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

run-api: ## start api process
	source .envrc && go run cmd/api/main.go -limiter-enabled=true -limiter-burst=30 -debug-enabled=true -cors-trusted-origins="http://localhost:9000 http://localhost:9001"

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
db-console: ## start database console
	PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight

db-seed: ## seed database
	ls ./database/seeds/*.sql | xargs -I {} sh -c 'PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight --quiet -f {} && echo "seed {}"'
	@echo "done seed"
	@echo

db-reset: ## reset database
	migrate -path=./database/migrations -database=${GREENLIGHT_DB_DSN} down -all
	@echo "done reset"
	@echo

db-migrate: ## run database migration
	migrate -path=./database/migrations -database=${GREENLIGHT_DB_DSN} up
	@echo "done migrate"
	@echo
	rm ./database/structure.sql
	pg_dump --schema-only ${GREENLIGHT_DB_DSN} --file ./database/structure.sql
	@echo "done pgdump"
	@echo

db-dump-example-data: ## dump users table data
	pg_dump --data-only --inserts --table users --no-comments ${GREENLIGHT_DB_DSN}

db-reset-add-seed: db-reset db-migrate db-seed ## reset and seed database(WARN)

show-pg-config-file: ## check postgres configuration file
	sudo -u postgres psql -c 'show config_file;'
	# /Library/PostgreSQL/17/data/postgresql.conf

install-tools: ## install 3rd tools
	brew install postgres
	brew install golang-migrate
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

golang-migrate-example: ## show example of using golang-migrate
	@echo "migrate create -seq -ext=.sql -dir=./database/migrations create_movies_table"

# test-optimistic-locking
# xargs -I % -P8 curl --location --request PUT 'http://localhost:4000/v1/movies/9' --header 'Content-Type: application/json' --data '{"title":"movie4","year":2007,"runtime":1500,"genres":["crime","war","sci-fi"]}' < <(printf '%s\n' {1..8})

# test timeout with `select pg_sleep(3)`
# curl --write-out '\nTime: %{time_total}s \n' http://localhost:4000/v1/movies/9

test-rate-limiter: ## test rate limiter
	for i in {1..10}; do curl --location 'http://localhost:4000/v1/healthcheck'; done

BODY='{"email":"Emiliano_Cruickshank@example.com","password":"pa55word"}'
load-test: ## dummy load test
	hey -d ${BODY} -m "POST" -z 5s http://localhost:4000/v1/tokens/authentication

sqlc: ## generate go files from sql
	sqlc generate -f database/sqlc.yaml

# first wire_gen.go needs to cd to specific director and run wire
generate: sqlc ## regenerate wire/sqlc...
	go generate ./...

sse-example: ## run sse api example
	go run internal/sse/example/main.go
