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
