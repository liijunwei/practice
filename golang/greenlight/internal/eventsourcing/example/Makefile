help: ## show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## start api
	go run .

sqlc: ## regenerate sqlc changes
	sqlc generate -f sqlcquery/sqlc.yaml

example-db-start: example-db-stop ## start example db and admin web
	docker system prune -f --volumes && docker-compose up -d
	@echo "export GREENLIGHT_DB_DSN='postgres://greenlight:123@localhost:5439/greenlight?sslmode=disable&application_name=greenlight-api'"
	@echo 'open http://localhost:8080'

example-db-stop: ## stop example db and admin web
	docker compose down
