my_var = "Success"

MyClass = Class.new do
  puts my_var

  def my_method
    puts my_var
  end
end

MyClass
MyClass.new.my_method # undefined local variable or method `my_var' for #<MyClass:0x0000000105465bd0> (NameError)
