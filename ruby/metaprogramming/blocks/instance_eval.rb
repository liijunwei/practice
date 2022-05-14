class MyClass
  def initialize
    @v = 1
  end
end

obj = MyClass.new
obj.instance_eval do
  p self
  p @v
end

puts "=" * 30

v = 2
obj.instance_eval { @v = v }
obj.instance_eval { p @v }

# 可以把传递给 `instance_eval()` 方法的块称为一个上下文探针(Context Probe), 因为它就像一个深入到对象中的代码片段, 对其进行操作
