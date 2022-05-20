# page 119

class C
  def a_method
    "C#a_method()"
  end
end

class D < C

end

obj = D.new

p obj.a_method


