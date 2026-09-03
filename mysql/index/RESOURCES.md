# MySQL Indexes Resources

## Knowledge

- [Book: _High Performance MySQL_, 4th Edition — Silvia Botros & Jeremy Tinley (O'Reilly, 2021)](https://www.oreilly.com/library/view/high-performance-mysql/9781492080503/)
  The canonical MySQL performance book. Use for: query optimization, index selection, schema design, and understanding the optimizer's decisions.

- [Site: _Use The Index, Luke!_ — Markus Winand](https://use-the-index-luke.com/)
  Free web book on SQL indexing aimed at developers (not DBAs). Vendor-neutral with MySQL notes. Use for: index anatomy, WHERE/JOIN/ORDER BY optimization, and the *why* behind column order and leftmost-prefix.

- [Doc: MySQL 8.0 Reference Manual — "How MySQL Uses Indexes"](https://dev.mysql.com/doc/refman/8.0/en/mysql-indexes.html)
  The authoritative list of when MySQL can and cannot use an index. Use for: precise, version-correct claims about index behavior.

- [Doc: MySQL 8.0 Reference Manual — "Multiple-Column Indexes"](https://dev.mysql.com/doc/refman/8.0/en/multiple-column-indexes.html)
  Leftmost-prefix rules and column ordering for composite indexes. Use for: composite index design.

- [Use The Index, Luke! — "Concatenated keys"](https://use-the-index-luke.com/sql/where-clause/the-equals-operator/concatenated-keys)
  The classic phone-book intuition for composite indexes and column order. Use for: teaching the *why* behind leftmost prefix.

- [Doc: MySQL 8.0 Reference Manual — "Range Optimization"](https://dev.mysql.com/doc/refman/8.0/en/range-optimization.html)
  How MySQL evaluates range conditions and when a range uses the index. Use for: the range-stops-the-prefix rule (Lesson 3).

- [Doc: MySQL 8.0 Reference Manual — "Covering Indexes"](https://dev.mysql.com/doc/refman/8.0/en/covering-indexes.html)
  When an index can answer a query without reading the table. Use for: covering index / `Using index` (Lesson 4).

- [Doc: MySQL 8.0 Reference Manual — "Type Conversion in Expression Evaluation"](https://dev.mysql.com/doc/refman/8.0/en/type-conversion.html)
  The cast rules behind the string-column-vs-number trap. Use for: implicit type conversion disabling an index (Lesson 5).

- [Doc: MySQL 8.0 Reference Manual — "SHOW INDEX Statement"](https://dev.mysql.com/doc/refman/8.0/en/show-index.html)
  The `Cardinality` column for judging selectivity. Use for: index design (Lesson 6).

- [Doc: MySQL 8.0 Reference Manual — "Nested-Loop Join Algorithms"](https://dev.mysql.com/doc/refman/8.0/en/nested-loop-joins.html)
  How MySQL walks a join as a nested loop. Use for: join-key indexing and reading multi-table plans (Lesson 7).

- [Doc: MySQL 8.0 Reference Manual — "Understanding the Query Execution Plan"](https://dev.mysql.com/doc/refman/8.0/en/execution-plan-information.html)
  How to read `EXPLAIN`. Use for: interpreting plans end to end.

- [Doc: MySQL 8.0 Reference Manual — "EXPLAIN Output Format"](https://dev.mysql.com/doc/refman/8.0/en/explain-output.html)
  Field-by-field reference for `EXPLAIN`, including the `type` (access type) values. Use for: decoding any `EXPLAIN` column.

## Wisdom (Communities)

- [Database Administrators Stack Exchange](https://dba.stackexchange.com/)
  High-signal Q&A; use the `mysql` and `query-performance` tags. Use for: query-plan critique and hard optimization problems.

- [MySQL Forums](https://forums.mysql.com/)
  Official forums. Use for: version-specific behavior and bugs.

- [Percona Blog](https://www.percona.com/blog/)
  Deep engineering posts on MySQL/InnoDB performance. Use for: when you want to go deeper than purely-practical (optimizer + InnoDB internals).

## Gaps

- None yet — will add as the mission reveals missing areas.
