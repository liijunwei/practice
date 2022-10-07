You have a variable that is *acting as a control flag* for a series of boolean expressions.

*Use a break or return instead.*

+ This is why languages have break and next (or continue) statements to get out of a complex conditional.

```ruby
# before
def check_security(people)
  found = false

  people.each do |person|
    unless found
      if person == "Don"
        send_alert
        found = true
      end

      if person == "John"
        send_alert
        found = true
      end
    end
  end
end

# after
def check_security(people)
  people.each do |person|
    if person == "Don"
      send_alert
      break
    end

    if person == "John"
      send_alert
      break
    end
  end
end
```

+ `Separate Query from Modifier`
