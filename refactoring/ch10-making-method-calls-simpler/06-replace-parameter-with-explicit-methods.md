You have a method that runs different code depending on the values of an enu- merated parameter.

*Create a separate method for each value of the parameter.*

没见过这种写法...
```ruby
# before
def set_value(name, value)
  if name == "height"
    @height = value
  elsif name == "width"
    @width = value
  else
    raise "Should never reach here"
  end
end

# after
def height=(value)
  @height = value
end

def width=(value)
  @width = value
end
```
