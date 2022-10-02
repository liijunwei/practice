Turn the fragment into a method whose **name** explains the purpose of the method.

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

+ And small methods really work only when you have **good names**, so you need to pay attention to naming.

+ To me method length is not the issue.
    + The key is the semantic distance between the method name and the method body.
    + If extracting improves clarity, do it, even if the name is longer than the code you have extracted.

+ **name it by what it does, not by how it does it**

+ If you can’t come up with a more meaningful name, don’t extract the code.

+ If you see an assignment to a parameter, you should immediately use **Remove Assignments to Parameters**.

+ Though parallel assignment can be used to return multiple values, I prefer to use single return values as much as possible.

+ **Temporary variables** often are so plentiful that they make extraction very awkward.
