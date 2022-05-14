def my_method
  x = "Goodbye"
  yield("crule")
end

x = "Hello"

p my_method { |y| "#{x}, #{y} world" }

# my_method 中定义的 x 对块来说是不可见的
#     块 看不见 my_method 里定义的x
#     块 能看见 和它平级定义的 x

