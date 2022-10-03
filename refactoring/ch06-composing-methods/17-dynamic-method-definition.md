dont agree(bad for text editor method navigation and ast manipulation)

```ruby
# before
def failure
  self.state = :failure
end

def error
  self.state = :error
end

# after
def_each :failure, :error do |method_name|
  self.state = method_name
end
```

+ I default to defining methods explicitly, but at the point when duplication begins to appear I quickly move to the dynamic definitions.

+ **The primary goal for Dynamic Method Definition is to more concisely express the method definition in a readable and maintainable format.**(agree)

+ `def_each` (it's all about readability and maintainability)
```ruby
class Class
  def def_each(*method_names, &block)
    method_names.each do |method_name|
      define_method method_name do
        instance_exec method_name, &block
      end
    end
  end
end
```

+ or use dsl(class macro)

+ too much magic :(
    + library can do this
    + application better not
