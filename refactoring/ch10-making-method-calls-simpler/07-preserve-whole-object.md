You are getting several values from an object and passing these values as param- eters in a method call.

*Send the whole object instead.*

```ruby
# before
low = days_temperature_range.low
high = days_temperature_range.high
plan.within_range?(low, high)

# after
plan.within_range?(days_temperature_range)
```

+ When you are considering Preserve Whole Object, consider Move Method as an alternative.
