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

def not_summer(date)
  date < SUMMER_START || date > SUMMER_END
end

def winter_charge(quantity)
  quantity * @winter_rate + @winter_service_charge
end

def summer_charge(quantity)
  quantity * @summer_rate
end
```

+ One of the most common areas of complexity in a program lies in complex conditional logic.

+ If I find a nested conditional, I usually first look to see whether I should use `Replace Nested Conditional with Guard Clauses`
+ If that does not make sense, I decompose each of the conditionals


