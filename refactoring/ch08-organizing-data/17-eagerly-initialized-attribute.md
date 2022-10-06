Initialize an attribute at construction time instead of on the first access.

```ruby
# before
class Employee
  def emails
    @emails ||= []
  end
end

# after
class Employee
  def initialize
    @emails = []
  end
end
```

+ The motivation for converting attributes to be eagerly initialized is for code readability purposes.(WTF???)

+ I prefer `Lazily Initialized Attributes`, but Martin prefers `Eagerly Initialized Attributes`.
    + personally prefer `Eagerly Initialized Attributes` unless some resources are too expensive to initialize eagerly

+ Martin and I agree that this isn’t something worth being religious about.
