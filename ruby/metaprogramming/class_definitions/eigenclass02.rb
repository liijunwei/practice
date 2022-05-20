# page 118

class MyClass
  def my_method
  end
end

obj = MyClass.new

class << obj
  # 这里是eigenclass的作用域
  puts "in eigenclass scope"
end

obj1 = Object.new
eigenclass = class << obj
  self
end


