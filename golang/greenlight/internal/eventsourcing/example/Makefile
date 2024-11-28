help: ## show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## start api
	go run .

sqlc: ## regenerate sqlc changes
	sqlc generate -f sqlcquery/sqlc.yaml

create-db-user:
	PGPASSWORD=${PGPASSWORD} createuser --username=postgres --pwprompt es-testuser

create-db:
	PGPASSWORD=${PGPASSWORD} createuser --username=postgres es-testuser
	PGPASSWORD=${PGPASSWORD} createdb --username=es-testuser --owner=es-test-user es-test-db

# manually copy migration in psql console to do migration(for now)
# TODO use dedicated example database
# prepare-db-dsn:
# export GREENLIGHT_DB_DSN="postgres://greenlight:${PGPASSWORD}@localhost/greenlight?sslmode=disable&application_name=greenlight-api"
