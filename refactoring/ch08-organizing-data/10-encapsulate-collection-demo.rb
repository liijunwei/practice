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
  def initialize
    @courses = []
  end

  def initialize_courses(courses)
    raise "courses should be empty" unless @courses.empty?
    @courses += courses
  end

  def courses
    @courses.dup
  end

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

    kent.add_course(Course.new("Smalltalk Programming", false))
    kent.add_course(Course.new("Appreciating Single Malts", true))

    refactoring = Course.new("Refactoring", true)
    kent.add_course(refactoring)
    kent.add_course(Course.new("Brutal Sarcasm", false))

    expect(kent.courses.size).to eq(4)
    expect(kent.courses.select {|c| c.advanced?}.size).to eq(2)

    kent.courses.delete(refactoring)
    expect(kent.courses.size).to eq(4)

    expect(kent.courses.select {|c| c.advanced?}.size).to eq(2)
  end
end
