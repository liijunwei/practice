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

  def number_of_advanced_courses
    courses.select { |course| course.advanced? }.size
  end

  def number_of_courses
    courses.size
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


    expect(kent.number_of_courses).to eq(4)
    expect(kent.number_of_advanced_courses).to eq(2)

    kent.remove_course(refactoring)
    expect(kent.number_of_courses).to eq(3)

    expect(kent.number_of_advanced_courses).to eq(1)
  end
end
