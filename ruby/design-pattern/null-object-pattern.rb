# https://en.wikipedia.org/wiki/Null_object_pattern#:~:text=In%20object%2Doriented%20computer%20programming,of%20Program%20Design%20book%20series.
class Dog
  def sound
    "bark"
  end
end
 
class NilAnimal
  def sound(*); end
end

def get_animal(animal=NilAnimal.new)
  animal
end

p get_animal(Dog.new).sound
p get_animal.sound

