# Designing an index: equality → range → sort

The user learned a repeatable decision procedure for composite-index column order: equality columns first (most-selective-first is a tie-breaker, not a law), then the single range column, then the `ORDER BY`/`GROUP BY` columns, then stop. Cardinality (`SHOW INDEX`) measures selectivity — high cardinality means a selective filter. One index serves several queries through its shared leftmost prefix — `(a, b)` serves `a` and `a,b` but not `b` — so a redundant `(a)` index should be dropped, while a query on `b` alone needs a separate `(b)`. Always verify the design with `EXPLAIN`.

**Evidence**: completed Lesson 6 — built `(customer_id, status, created_at)` for `WHERE customer_id = ? AND status = ? ORDER BY created_at`, confirmed via `SHOW INDEX` (cardinality 500 / 3 / 10000) and `EXPLAIN` (no `Using filesort`).

**Implications**: the user now has a recipe for choosing an index, closing the mission's "choose the right index and explain the column order" criterion. See [[MISSION.md]].
