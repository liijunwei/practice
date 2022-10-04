class Person
  attr_accessor :department
end
class Department
  attr_reader :manager

  def initialize(manager)
    @manager = manager
  end
end

# If a client wants to know a person’s manager, it needs to get the department first:

john = Person.new
manager = john.department.manager

