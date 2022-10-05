You have a data item that needs additional data or behavior.

*Turn the data item into an object.*

把实例里的某个实例变量 由简单的值 改为 某个类的实例

```ruby
# before
class Order...
  attr_accessor :customer # the data item

  def initialize(customer)
    @customer = customer
  end
end

# after
class Order...
  attr_accessor :customer

  def initialize(customer_name)
    @customer = Customer.new(customer_name)
  end
end

class Customer
  def initialize(name)
    @name = name
  end
end
```
