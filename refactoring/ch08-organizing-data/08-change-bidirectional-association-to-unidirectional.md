You have a two-way association but one class no longer needs features from the other.

*Drop the unneeded end of the association.*

+ Bidirectional associations are useful, but they carry a price.
    + The price is the added complexity of maintaining the two-way links and ensuring that objects are properly created and removed.
    + Bidirectional associations are not natural for many programmers, so they often are a source of errors.(确实有点不好理解, 可能是之前没处理过这种)
    + Having many two-way links also makes it easy for mistakes to lead to zombies: objects that should be dead but still hang around because of a reference that was not cleared.
    + Bidirectional associations force an interdependency between the two classes.

+ **You should use bidirectional associations when you need to but avoid them when you don’t.**
    + As soon as you see a bidirectional association is no longer pull- ing its weight, drop the unnecessary end.

+ "Can I find another way to provide the customer(dependent) object?"

```ruby
# before
class Order
  def discounted_price
    gross_price * (1 - @customer.discount)
  end
end

# after
class Order
  def discounted_price(customer)
    gross_price * (1 - customer.discount)
  end
end
```

+ TODO question: what is `Introduce Assertion` ?

+ there are many ways "remove the back-pointer but retain the interdependencies between the two classes"
