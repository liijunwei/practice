class Person
  attr_reader :mother, :children, :name

  def initialize(name, date_of_birth, date_of_death=nil, mother=nil)
    @name = name
    @mother = mother
    @date_of_birth = date_of_birth
    @date_of_death = date_of_death
    @children = []
    @mother.add_child(self) if @mother
  end

  def add_child(child)
    @children << child
  end

  def number_of_living_descendants
    children.inject(0) do |count, child|
      count += 1 if child.alive?
      count + child.number_of_living_descendants
    end
  end

  def number_of_descendants_named(name)
    children.inject(0) do |count, child|
      count += 1 if child.name == name
      count + child.number_of_descendants_named(name)
    end
  end

  def alive?
    @date_of_death.nil?
  end
end

mother = Person.new("mia", "1963-01-01", nil, nil)
Person.new("tiny-mia1", "1993-01-01", nil, mother)
Person.new("tiny-mia2", "1994-01-01", nil, mother)
p mother
