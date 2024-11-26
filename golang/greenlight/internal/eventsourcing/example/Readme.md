# Account & DebitHold example

## Run

run migrations in [sqlcquery/migrations](./sqlcquery/migrations).

``` shell
go run .
```

## Navigating the code

Aggregates:


- [Account](./account.go): account aggregate
- [DebitHold](./debit_hold.go): debit hold aggregate

Events:

- [DebitHold events](./debit_hold_events.go): event definition for DebitHold
- [Account events](./account_events.go): event definition for Account

Repository:

- [DebitHoldRepository](./debit_hold_repository.go): a repository without projected view, demonstrate basic AggregateRepository usage
- [AccountRepository](./account_repository.go): demonstrates AggregateLoader, AggregateSaver, sqlc integration, currency handling with locks.

Handles:

- [Account handles](./account_handlers.go): demonstrates basic usage of repository
- [DebitHold handles](./debit_hold_handlers.go): demonstrates transactions across multiple repository
