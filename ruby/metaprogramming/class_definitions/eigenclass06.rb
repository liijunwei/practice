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

