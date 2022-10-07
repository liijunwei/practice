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
+ If the condition is an unusual condition, check the condition and return if the condition is true.
    + This kind of check is often called a **guard clause** [Beck].

+ The key point about `Replace Nested Conditional with Guard Clauses` is one of emphasis.
    + 使用 guard clause 的重点是 强调某个条件(条件之一)
    + "This is rare, and if it happens, do something and get out"

+ If you are using an if-then-else construct **you are giving equal weight** to the if leg and the else leg
    + "this two are equally likely and important"

+ "Clarity is the key principle"

+ Trick: "Reversing the Conditions"
    + `Replace Nested Conditional with Guard Clauses` by **reversing the conditional expressions**

