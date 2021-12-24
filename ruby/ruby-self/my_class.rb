
# "如果你想成为Ruby大师, 你应该在任何时候都知道当前那个对象在充当self的角色"
class MyClass
  def testing_self
    @var = 10
    my_method()
    self
  end

  def my_method()
    @var += 1
  end
end

obj = MyClass.new
p obj.testing_self