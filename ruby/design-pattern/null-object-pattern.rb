# https://en.wikipedia.org/wiki/Null_object_pattern#:~:text=In%20object%2Doriented%20computer%20programming,of%20Program%20Design%20book%20series.
class Dog
  # return object if found
  # return null object if not found, do not retur nil or raise exception
  def self.find(id)
    id == 1 ? new : NilAnimal.new
  end

  def sound
    "bark"
  end
end

class NilAnimal
  def sound
  end
end

dog = Dog.find(1)
p dog.sound
dog = Dog.find(2)
p dog.sound

