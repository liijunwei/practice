class MyClass
  def initialize(value)
    @value = value
  end

  def my_method
    @value
  end
end

object = MyClass.new(1)
m = object.method :my_method
p m.call


