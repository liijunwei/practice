A method has conditional behavior that does not make clear the normal path of execution.

*Use guard clauses for all the special cases.*

```ruby
# before
def pay_amount
  if @dead
    result = dead_amount
  else
    if @separated
      result = separated_amount
    else
      if @retired
        result = retired_amount
      else
        result = normal_pay_amount
      end
    end
  end
  result
end

# after
def pay_amount
  # the order matters
  return dead_amount if @dead
  return separated_amount if @separated
  return retired_amount if @retired
  normal_pay_amount
end
```

+ If both are part of normal behavior, use a condition with an if and an else leg.
+ If the condition is an unusual condition, check the condition and return if the condition is true. This kind of check is often called a guard clause [Beck].
