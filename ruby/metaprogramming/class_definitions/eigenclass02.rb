# page 118

obj = Object.new

eigenclass =
  class << obj
    self
  end

p eigenclass
p eigenclass.class

def obj.my_singleton_method
end

p eigenclass.instance_methods.grep(/my_/)

