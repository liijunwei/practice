# Three traps that disable an index

The user learned three ways a `WHERE` clause hides a column from its index, each producing `type=ALL` despite a valid index existing: (1) a function on the column — `YEAR(created_at) = 2020` — fixed by rewriting as a range over the raw column; (2) a leading wildcard — `LIKE '%foo'` — fixed by matching a prefix (`LIKE 'foo%'`); (3) an implicit type conversion — a `VARCHAR` column compared to a number — fixed by comparing to a literal of the same type.

**Evidence**: completed Lesson 5 — a `customers` table with six `EXPLAIN` comparisons (one broken + one fixed per trap).

**Implications**: "does this `WHERE` let the index see the raw column?" is now the second half of index debugging — the first being "does the right index exist?" See [[MISSION.md]].
