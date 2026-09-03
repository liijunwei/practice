# A range stops the prefix (key_len as proof)

The user learned that in a composite index, MySQL uses equality columns up to the first range; the range column itself is used, and anything after it can no longer narrow the search. So `a = ? AND b > ?` uses both columns, while `a > ? AND b = ?` uses only `a`. `key_len` (in `EXPLAIN`) reports the byte length of the index key actually used — an objective count of consumed columns (e.g. 4 vs 8 for two `INT` columns). `IN (...)` behaves like a multi-value equality, not an open range, so it does not block later columns.

**Evidence**: completed Lesson 3 — built a `sales` table and observed `key_len` 4 / 8 / 4 across the equality, equality-then-range, and range-then-equality cases.

**Implications**: the user can now predict range behavior and prove it with `key_len` instead of guessing. Leads directly into covering indexes (Lesson 4) and index design (Lesson 6). See [[MISSION.md]].
