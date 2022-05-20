# page 127

# 如果想给某个特定的类增加属性, 需要 在这个类的eigenclass中定义这个属性
class MyClass
  class << self
    attr_accessor :c
  end
end

