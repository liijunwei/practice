# page 119

class Object
  def eigenclass
    class << self
      self
    end
  end
end

class C
  def a_method
    "C#a_method()"
  end
end

class D < C

end

obj = D.new

p obj.a_method

p "abc".eigenclass # #<Class:#<String:0x000000010be1e6f8>>
