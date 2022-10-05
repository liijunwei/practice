require 'pry'
require 'rspec'

class Course
  def initialize(name, advanced)
    @name = name
    @advanced = advanced
  end

  def advanced?
    @advanced
  end
end

class Person
  attr_accessor :courses

  def add_course(course)
    @courses << course
  end

  def remove_course(course)
    @courses.delete(course)
  end
end

RSpec.describe Person do
  specify do
    kent = Person.new

    courses = []
    courses << Course.new("Smalltalk Programming", false)
    courses << Course.new("Appreciating Single Malts", true)

    kent.courses = courses

    expect(kent.courses.size).to eq(2)

    refactoring = Course.new("Refactoring", true)
    kent.courses << refactoring
    kent.courses << Course.new("Brutal Sarcasm", false)

    expect(kent.courses.size).to eq(4)
    expect(kent.courses.select {|c| c.advanced?}.size).to eq(2)

    kent.courses.delete(refactoring)
    expect(kent.courses.size).to eq(3)

    expect(kent.courses.select {|c| c.advanced?}.size).to eq(1)
  end
end
