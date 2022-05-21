# page 136

class Integer
  alias :old_plus :+

  def +(value)
    old_plus(value).old_plus(1)
  end
end

p  1 + 1
p -1 + 1
p  100 + 10

