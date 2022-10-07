You have conditional code that is unnecessarily verbose and does not use the most readable Ruby construct.

*Replace the conditional code with the more idiomatic Ruby construct.*

```ruby
# before
parameters = params ? params : []

# after
parameters = params || []
```

+ improve the expressiveness of our code
