# Account & DebitHold example

how to run it?
```sh
make example-db-start
make db-migrate
make run

curl -X POST 'http://127.0.0.1:3000/api/account'

open http://localhost:8080/?pgsql=db
username: greenlight
password: 123
database: greenlight
```
