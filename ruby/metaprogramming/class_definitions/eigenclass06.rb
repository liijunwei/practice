# page 121

# 为了实验时方便, 写一个辅助方法返回对象的eigenclass
class Object
  def eigenclass
    class << self
      self
    end
  end
end

p "abc".eigenclass # #<Class:#<String:0x000000010be1e6f8>>

obj = Object.new

class << obj
  def a_singleton_method
    "obj#a_singleton_method()"
  end
end

p obj # #<Object:0x0000000103129068>
p obj.eigenclass # #<Class:#<Object:0x0000000103129068>>
p obj.eigenclass.superclass

