class C
  def initialize
    @x = 1
    @y = 2
  end
end

# ruby 1.9 中引入了一个类似instance_eval的方法: instance_exec
# 他们功能相似, 但是instance_exec可以传入参数
p C.new.instance_exec(3) { |arg| (@x + @y) * arg }
p C.new.instance_exec(4) { |arg| (@x + @y) * arg }
