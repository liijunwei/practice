require 'pry'
require 'rspec'

class Course
  def initialize(name, advanced)
    @name = name
    @advanced = advanced
  end
end

class Person
  attr_accessor :courses
end

RSpec.describe Person do
end
