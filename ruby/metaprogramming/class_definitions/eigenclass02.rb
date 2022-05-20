# page 118

obj = Object.new

eigenclass =
  class << obj
    self
  end

p eigenclass # #<Class:#<Object:0x000000010a71e610>>
p eigenclass.class # Class

def obj.my_singleton_method
end

p eigenclass.instance_methods.grep(/my_/) # [:my_singleton_method]

