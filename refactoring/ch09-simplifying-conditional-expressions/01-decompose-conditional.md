You have a complicated conditional (if-then-else) statement.

*Extract methods from the condition, “then” part, and “else” parts.*

```ruby
# before
if date < SUMMER_START || date > SUMMER_END
  charge = quantity * @winter_rate + @winter_service_charge
else
  charge = quantity * @summer_rate
end

# after
if not_summer(date)
  charge = winter_charge(quantity)
else
  charge = summer_charge(quantity)
end
```

+ One of the most common areas of complexity in a program lies in complex conditional logic.
