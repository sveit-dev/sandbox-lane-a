# Known bug

`pager.Count` (pager/pager.go) has no test and returns the wrong number of
pages whenever `total` is not a multiple of `size`: it uses integer division
and drops the last partial page.

| Call | Returns | Expected |
|---|---|---|
| `Count(45, 20)` | 2 | 3 |
| `Count(1, 20)` | 0 | 1 |
| `Count(40, 20)` | 2 | 2 |

The fix rounds up, `(total + size - 1) / size`, and adds a table-driven test.
This bug is deliberate: the qa end-to-end case expects an agent to find it,
add the failing test, and fix it. Restore it after each run.
