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



