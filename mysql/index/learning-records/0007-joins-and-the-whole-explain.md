# Joins are nested loops — index the join key

The user learned that MySQL joins via a nested loop: read the driving table, and for each row look up the next table by the join key. So a join key needs an index exactly like a `WHERE` column does. `eq_ref` = a `PRIMARY KEY`/`UNIQUE` lookup (one row per lookup, best); `ref` = a non-unique lookup; `ALL` = a full scan of the inner table for every outer row, multiplying to N×M work. `EXPLAIN` rows read top-to-bottom as the loop order, and the `rows` column multiplies across tables.

**Evidence**: completed Lesson 7 — a `customers`/`orders` join where `orders` flipped from `ALL` to `ref` after indexing `customer_id`, plus an `eq_ref` observation when joining onto the primary key.

**Implications**: the core mechanics are now complete. The natural next step is an applied session: take one real slow query and walk its plan end-to-end, exercising all seven lessons. See [[MISSION.md]].
