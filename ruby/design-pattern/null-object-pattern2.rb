# example from "Agile Software Development, Principles, Patterns, and Practices"
# chapter 17: null object pattern

class Employee
  attr_reader :name

  def initialize(name)
    @name = name
  end

  def is_time_to_pay(time)
    time == "2022-05-19 20:51:23 +0800"
  end

  def pay
    "#{name} was paied..."
  end
end

class NullEmployee
  attr_reader :name

  def initialize(name)
    @name = name
  end

  def is_time_to_pay(time)
    false
  end

  def pay
    raise "can not pay to a non-exist employee [#{name}]"
  end
end

class DB
  def self.get_employee(name)
    if name == 'Bob'
      return Employee.new(name)
    else
      return NullEmployee.new(name)
    end
  end
end

require 'test/unit'

class TestEmployee < Test::Unit::TestCase
  def setup
    @employee      = DB.get_employee("Bob")
    @null_employee = DB.get_employee("David")
  end

  def test_is_time_to_pay
    refute @employee.is_time_to_pay("2023-05-19 20:51:23 +0800")
    assert @employee.is_time_to_pay("2022-05-19 20:51:23 +0800")
    refute @null_employee.is_time_to_pay(Time.now.to_s)
  end

  def test_pay
    assert_equal "#{@employee.name} was paied...", @employee.pay
    assert_raise_with_message(RuntimeError, "can not pay to a non-exist employee [#{@null_employee.name}]") { @null_employee.pay }
  end
end

