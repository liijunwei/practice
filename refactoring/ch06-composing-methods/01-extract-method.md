Turn the fragment into a method whose name explains the purpose of the method.

e.g.
```ruby
# before
def print_owing(amount)
  print_banner
  puts "name: #{@name}"
  puts "amount: #{amount}"
end

# after
def print_owing(amount)
  print_banner
  print_details amount
end

def print_details(amount)
  puts "name: #{@name}"
  puts "amount: #{amount}"
end
```

+ I prefer short, well-named methods for several reasons:
    1. improve chance of being resused
    2. more readable
