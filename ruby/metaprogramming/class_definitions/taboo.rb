# page 110

# class MyClass < Array
#   def my_method
#     "Hello"
#   end
# end

# 目标: 在不使用关键字`class`的情况下写成和上面等价的ruby代码

my_class = Class.new(Array) do
  def my_method
    "Hello"
  end
end

p my_class.new.my_method
