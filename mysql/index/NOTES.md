# Notes

- **Setup**: user runs MySQL locally on a Mac mini. Lessons should always include real, copy-pasteable SQL they can execute.
- **Depth preference**: "purely practical" — behavior (what indexes do) + structure (B+ tree shape, column order, leftmost prefix). Skip deep disk/page internals for now (per MISSION.md out-of-scope).
- **Starting level**: comfortable with basic SQL, no index knowledge. Don't re-teach SQL syntax; build index concepts directly on top of it.
- **Goal anchor**: everything should trace back to "debug a slow query" — reading EXPLAIN, choosing an index, proving it helps.
- **Format**: user values hands-on experiments over theory-only prose. Prefer "run this, observe this, now explain it."
