Replace all references to that temp with the expression.

```ruby
# before
base_price = an_order.base_price
return (base_price > 1000)

# after
return (an_order.base_price > 1000)
```
