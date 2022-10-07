The same fragment of code is in all branches of a conditional expression.

*Move it outside the expression.*

```ruby
if special_deal?
  total = price * 0.95
  send_order
else
  total = price * 0.98
  send_order
end
```
