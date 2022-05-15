class MyClass
  define_method :my_method do |my_arg|
    my_arg * 2
  end
end

obj = MyClass.new

p obj.my_method(3)
puts "=" * 30

p obj.send(:my_method, 3)
