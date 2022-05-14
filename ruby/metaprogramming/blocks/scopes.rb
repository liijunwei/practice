v1 = 1
puts "line:#{__LINE__} #{local_variables}"

class MyClass
  v2 = 2
  puts "line:#{__LINE__} #{local_variables}"

  def my_method
    v3 = 3
    puts "line:#{__LINE__} #{local_variables}"
  end

  puts "line:#{__LINE__} #{local_variables}"
end

puts "================"
obj = MyClass.new
obj.my_method
obj.my_method
puts "line:#{__LINE__} #{local_variables}"

# Ruby中的作用域是截然分开的:
# 一旦进入一个新的作用域, 原先的绑定就会被替换为一组新的绑定
# 这意味着在程序进入MyClass后, v1就超出了"作用域范围", 从而不可见了

# MyClass类定义结束前, 程序不会再打开一个新的作用域
# 在定义完成后, 用`class`关键字打开的作用域会永远关闭, 同事程序回到顶级作用域

