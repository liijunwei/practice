class MyClass
  def my_method
  end
end

obj = MyClass.new
p obj.my_method # nil
p obj.class # MyClass
p obj.class.superclass # Object

