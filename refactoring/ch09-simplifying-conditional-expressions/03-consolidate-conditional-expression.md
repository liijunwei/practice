You have a sequence of conditional tests with the same result.

*Combine them into a single conditional expression and extract it.*

Consolidate 合并

```ruby
def disability_amount
  return 0 if @seniority < 2
  return 0 if @months_disabled > 12
  return 0 if @is_part_time
  # compute the disability amount
end
```

+ name the condition

+ The reasons in favor of consolidating conditionals also point to reasons for not doing it.
    + If you think the checks are really independent and shouldn't be thought of as a single check, don't do the refactoring.
    + Your code already communicates your intention.

有时候确定是独立的检查, 就不需要合并, 而不合并恰恰表明了你的意图

