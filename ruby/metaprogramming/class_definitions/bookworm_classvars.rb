class Loan
  def initialize(book)
    @book = book
    @time = Loan.time_class.now
  end

  def self.time_class
    @time_class || Time
  end

  def to_s
    "#{@book.upcase} loaned on #{@time}"
  end
end

require 'test/unit'

class FakeTime
  def self.now
    "2022-05-18 09:02:28 +0800"
  end
end

class TestLoan < Test::Unit::TestCase
  def test_conversion_to_string
    Loan.instance_eval { @time_class = FakeTime }
    loan = Loan.new("War and Peace")

    assert_equal "WAR AND PEACE loaned on 2022-05-18 09:02:28 +0800", loan.to_s
  end
end
