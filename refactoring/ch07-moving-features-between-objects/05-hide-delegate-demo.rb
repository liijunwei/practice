class Person
  attr_accessor :department

  def manager
    @department.manager
  end
end

class Department
  attr_reader :manager

  def initialize(manager)
    @manager = manager
  end
end

# If a client wants to know a person’s manager, it needs to get the department first:

john = Person.new
# apply "Law of Demeter" ?
manager = john.department.manager

# This reveals to the client how the department class works and that the department is responsible for tracking the manager.

# I can reduce this coupling by hiding the department class from the client.
# I do this by creating a simple delegating method on person:
manager = john.manager
