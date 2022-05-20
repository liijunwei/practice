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

p obj.eigenclass.superclass # D

```
         Object
         |
         C
         |
         D
         |
obj ---> #obj(obj的eigenclass)
```

