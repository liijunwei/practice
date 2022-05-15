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
p obj.method_two

