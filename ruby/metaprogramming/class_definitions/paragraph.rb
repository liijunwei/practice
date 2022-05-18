# page 111

class Paragraph
  def initialize(text)
    @text = text
  end

  def title?
    @text.upcase == @text
  end

  def reverse
    @text.reverse
  end

  def upcase
    @text.upcase
  end
end

require 'test/unit'

class TestParagraph < Test::Unit::TestCase
  def test_upcase
    Loan.instance_eval { @time_class = FakeTime }
    loan = Loan.new("War and Peace")

    assert_equal "WAR AND PEACE loaned on 2022-05-18 09:02:28 +0800", loan.to_s
  end
end

