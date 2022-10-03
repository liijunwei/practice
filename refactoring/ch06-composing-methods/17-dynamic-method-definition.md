dont agree(bad for text editor method navigation and ast manipulation)

```ruby
# before
def failure
  self.state = :failure
end

def error
  self.state = :error
end

# after
def_each :failure, :error do |method_name|
  self.state = method_name
end
```

+ I default to defining methods explicitly, but at the point when duplication begins to appear I quickly move to the dynamic definitions.

+ **The primary goal for Dynamic Method Definition is to more concisely express the method definition in a readable and maintainable format.**(agree)


