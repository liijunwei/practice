Initialize an attribute on access instead of at construction time.

```ruby
# before
class Employee
  def initialize
    @emails = []
  end
end

# after
class Employee
  def emails
    @emails ||= []
  end
end
```

+ The motivation for converting attributes to be lazily initialized is for code readability purposes(TODO unclear)

+ The procedural behavior of initializing each attribute in a constructor is sometimes unnecessary and less maintainable than a class that deals exclusively with attributes.

+ Lazily Initialized Attributes can encapsulate all their initialization logic within the methods themselves.

+ Using `||=` for `Lazily Initialized Attributes` is a common idiom; however, this idiom falls down when nil or false are valid values for the attribute.


