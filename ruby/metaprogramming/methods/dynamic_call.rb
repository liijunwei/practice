class MyClass
  def my_method(my_arg)
    my_arg * 2
  end
end

obj = MyClass.new

p obj.my_method(3)
puts "=" * 30

# 这种技术称为动态派发(dynamic dispatch)
# page 41
p obj.send(:my_method, 3)
