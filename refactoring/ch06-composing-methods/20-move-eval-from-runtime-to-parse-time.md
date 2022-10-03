# Move Eval from `Runtime` to `Parse Time`

Move the use of eval from within the method definition to defining the method
itself.

```ruby
# before
class Person
  def self.attr_with_default(options)
    options.each_pair do |attribute, default_value|
      define_method attribute do
        eval "@#{attribute} ||= #{default_value}"
      end
    end
  end

  attr_with_default :emails => "[]", :employee_number =>"EmployeeNumberGenerator.next"
end

# after
class Person
  def self.attr_with_default(options)
    options.each_pair do |attribute, default_value|
      eval define_method attribute do
        "@#{attribute} ||= #{default_value}"
      end
    end
  end

  attr_with_default :emails => "[]", :employee_number =>"EmployeeNumberGenerator.next"
end
```
