# page 118

obj = Object.new

eigenclass =
  class << obj
    self
  end

# eigenclass是一个很特殊的类
# 每个eigenclass只有一个实例(这就是他们被称为单件类的原因)
# eigenclass不能被继承
# eigenclass是一个对象(class or instance)的单件方法的存活之所
p eigenclass # #<Class:#<Object:0x000000010a71e610>>
p eigenclass.class # Class

def obj.my_singleton_method
end

p eigenclass.instance_methods.grep(/my_/) # [:my_singleton_method]

begin
  eigenclass.new
rescue => e
  p [e.class, e.message] # [TypeError, "can't create instance of singleton class"]
end

begin
  class Demo < eigenclass
  end
rescue => e
  p [e.class, e.message] # [TypeError, "can't make subclass of singleton class"]
end

