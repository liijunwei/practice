# Covering indexes (Using index)

The user learned that a covering index contains every column the query needs (across `SELECT`/`WHERE`/`ORDER BY`/`GROUP BY`), so MySQL answers from the index alone — `Extra: Using index`, no table read. InnoDB secondary indexes secretly append the primary key, so an index covers one more column than declared. A non-covering query shows `Using index condition` (ICP) and still reads the table row for the missing columns. `SELECT *` defeats coverage.

**Evidence**: completed Lesson 4 — compared `SELECT name` / `SELECT name, age` / `SELECT id, name` against the same indexed `WHERE` and watched `Extra` flip between `Using index` and `Using index condition`.

**Implications**: a narrower `SELECT` list is a real optimization lever, and column selection now factors into index design alongside `WHERE` columns. See [[MISSION.md]].
