# page 117

class MyClass
  def my_method
  end
end

obj = MyClass.new

p obj.my_method        # nil
p obj.class            # MyClass
p obj.class.superclass # Object

# 问题: 这个obj的方法, 放在类体系结构的哪个位置?
def obj.my_singleton_method
end

p obj.my_singleton_method # nil

# 问题: 这个MyClass的方法, 放在类体系结构的哪个位置?
def MyClass.my_class_method
end

p MyClass.my_class_method
