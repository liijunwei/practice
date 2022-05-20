# page 121

class Object
  def eigenclass
    class << self
      self
    end
  end
end

p "abc".eigenclass

