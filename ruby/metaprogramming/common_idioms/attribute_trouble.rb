# page 225

class MyClass
  attr_accessor :my_attr

  def initialize_attributes
    my_attr = 10
  end
end

obj = MyClass.new
obj.initialize_attributes
p obj.my_attr # nil

# not expected to be nil

# since ruby can't tell
# if we're trying to assign value to a local variable
# or
# we're trying to call a "my_attr=" method

# the default behavior is "assign value to a local variable"

#### quick fix
class MyClassFix
  attr_accessor :my_attr

  def initialize_attributes
    self.my_attr = 10
  end
end

obj = MyClassFix.new
obj.initialize_attributes
p obj.my_attr # 10

