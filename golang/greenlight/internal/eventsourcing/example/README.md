# Account & DebitHold example

how to run it?
```sh
make example-db-start && export GREENLIGHT_DB_DSN='postgres://greenlight:123@localhost:5439/greenlight?sslmode=disable&application_name=greenlight-api' && echo "waiting for db is ready..." && sleep 5 && make db-migrate && make run-api

account_id=$(curl -sX POST 'http://127.0.0.1:3000/api/account' | jq --raw-output .ID)

curl -s --location 'http://127.0.0.1:3000/api/debit-hold' \
--header 'Content-Type: application/json' \
--data '{
    "AccountID": "'$account_id'",
    "Amount": "100.0"
}' | jq

open http://localhost:8080/?pgsql=db&username=greenlight&db=greenlight
# enter password: 123
```
