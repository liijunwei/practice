# Mission: MySQL Indexes

## Why

The user wants to be able to debug and fix slow queries in production MySQL databases. When a `SELECT` is slow, they want to read `EXPLAIN`, see *why* it's slow, and know which index to add (or remove) to make it fast — without guessing.

## Success looks like

- Given a slow query and its `EXPLAIN` output, identify the access type and explain in one sentence why it's slow.
- Choose the right index (single-column or composite) for a given `WHERE` / `JOIN` / `ORDER BY` pattern, and explain the column order.
- Predict, before running a query, whether MySQL will use a given index — and spot the common traps (function on the column, leading `%`, wrong column order) that silently disable one.
- Back every claim with a reproducible experiment on their own MySQL instance.

## Constraints

- Runs MySQL locally on a Mac mini; lessons should include real SQL they can execute.
- **Purely practical** depth: behavior (what an index does) plus structure (B+ tree shape, column order, leftmost prefix). Deep disk/page internals (page splits, buffer pool, write amplification) are out of scope for now.
- Starting level: comfortable with basic SQL (`SELECT`/`JOIN`/`GROUP BY`), but no index knowledge. Do not re-teach SQL syntax.

## Out of scope

- InnoDB physical storage internals (pages, buffer pool, redo logs, write amplification).
- Index types beyond B-tree for now (FULLTEXT, spatial, hash).
- Database administration topics (replication, backups, sharding).
