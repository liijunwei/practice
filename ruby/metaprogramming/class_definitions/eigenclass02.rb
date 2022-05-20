# page 118

obj = Object.new

eigenclass =
  class << obj
    self
  end

# 每个eigenclass只有一个实例, 这就是他们被称为单件类的原因
# eigenclass不能被继承
p eigenclass # #<Class:#<Object:0x000000010a71e610>>
p eigenclass.class # Class

def obj.my_singleton_method
end

p eigenclass.instance_methods.grep(/my_/) # [:my_singleton_method]

demo = eigenclass.new
p demo