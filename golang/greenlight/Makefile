.PHONY: help run-api db-console db-seed db-reset db-migrate db-dump-example-data show-pg-config-file install-tools golang-migrate-example test-rate-limiter load-test sqlc sse-example

include .envrc

GREENLIGHT_DB_DSN="postgres://greenlight:123@localhost:55432/greenlight?sslmode=disable&application_name=greenlight-api"

help: ## show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# steps:
# make docker-compose
# make run-api
docker-compose: ## start docker compose
	docker-compose down -v && docker-compose up -d && sleep 3 && make db-reset-add-seed && psql ${GREENLIGHT_DB_DSN} -c '\d+'

run-api: ## start api process
	source .envrc && GREENLIGHT_DB_DSN=${GREENLIGHT_DB_DSN} go run cmd/api/main.go -limiter-enabled=true -limiter-burst=30 -debug-enabled=true -cors-trusted-origins="http://localhost:9000 http://localhost:9001"

db-console: ## start database console
	PGPASSWORD=${PGPASSWORD} psql ${GREENLIGHT_DB_DSN}

db-seed: ## seed database
	ls ./database/seeds/*.sql | xargs -I {} sh -c 'PGPASSWORD=${PGPASSWORD} psql ${GREENLIGHT_DB_DSN} --quiet -f {} && echo "seed {}"'
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

install-tools: ## install 3rd party tools
	brew install postgres
	brew install golang-migrate
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0

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
