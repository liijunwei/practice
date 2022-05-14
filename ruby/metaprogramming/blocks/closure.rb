def my_method
  x = "Goodbye"
  yield("crule")
end

x = "Hello"

p my_method { |y| "#{x}, #{y} world" }


