class MyClass
  def initialize(value)
    @value = value
  end

  def my_method
    @value
  end
end

object = MyClass.new(1)
m1 = object.method :my_method
p m1.call

puts "=" * 30

unbound = m1.unbind

another_object = MyClass.new(2)
m2 = unbound.bind(another_object)
p m2.call


