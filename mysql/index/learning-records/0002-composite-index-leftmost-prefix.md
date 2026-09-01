# Composite indexes and the leftmost-prefix rule

The user learned that a composite index is sorted by its columns in order — like a phone book (last name, then first name) — and therefore obeys the **leftmost-prefix rule**: an index `(a, b)` serves lookups on `a` and `(a, b)`, never `b` alone. They also saw the production payoff: "filter on `a`, sort by `b`" is served for free, because the rows come out of the index already in `b` order — eliminating the `Using filesort`.

**Evidence**: completed Lesson 2. Built an `orders` table (100k rows, 50 customers), watched `type=ALL` + `Using filesort` flip to `type=ref` with no filesort after adding `(customer_id, created_at)`, and confirmed `WHERE created_at` alone falls back to `ALL`.

**Implications**: the user can now choose the column order of a composite index and predict (via `EXPLAIN`) whether a `WHERE` clause can use it. Still in the ZPD and worth covering next: (1) a **range condition in the middle** — `a = ? AND b > ?` uses both columns, but `a > ? AND b = ?` uses only `a`; (2) **covering indexes** — answering the query from the index alone so the table row is never read. See [[MISSION.md]].
