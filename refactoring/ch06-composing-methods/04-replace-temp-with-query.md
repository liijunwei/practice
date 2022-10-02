# Replace Temp with Query

You are using a temporary variable to hold the result of an expression.

Extract the expression into a method. Replace all references to the temp with the expression. The new method can then be used in other methods.

+ The problem with temps is that they are temporary and local.
+ Because they can be seen only in the context of the method in which they are used, temps tend to encourage longer methods, because that’s the only way you can reach the temp.

+ Replace Temp with Query often is a vital step before Extract Method.
+ Local variables make it difficult to extract, so replace as many variables as you can with queries.

+ Initially mark the method as private. You may find more use for it later, but you can easily relax the protection then.

+ Ensure the extracted method is free of side effects—that is, it does not modify any object.
    + If it is not free of side effects, use Separate Query from Modifier.
    + what is "Separate Query from Modifier" ?

+ You may be concerned about performance in this case.
    + As with other performance issues, let it slide for the moment.
    + Nine times out of ten, it won’t matter.
    + When it does matter, you will fix the problem during optimization.
    + With your code better factored, you often find more powerful optimizations that you would have missed without refactoring.
    + If worse comes to worst, it’s easy to put the temp back.
    +
    + do not worry about performance too early, use profilers before optimizating

```ruby
# before
def price
  base_price = @quantity * @item_price

  if base_price > 1000
    discount_factor = 0.95
  else
    discount_factor = 0.98
  end

  base_price * discount_factor
end

# after
def price
  base_price * discount_factor
end

def base_price
  @quantity * @item_price
end

def discount_factor
  base_price > 1000 ? 0.95 : 0.98
end
```
