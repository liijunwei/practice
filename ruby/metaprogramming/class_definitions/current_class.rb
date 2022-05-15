# page 104

class MyClass
  def method_one
    def method_two
      "hello"
    end
  end
end

obj = MyClass.new

p obj.method_one
p obj.method_two # 会发现, method_two 虽然是在 method_one里定义的, 但它还是 MyClass 的实例方法


