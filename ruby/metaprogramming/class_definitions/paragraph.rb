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
  def setup
    @paragraph = Paragraph.new("Hello")
  end

  def test_title?
    refute @paragraph.title?
    assert Paragraph.new("WORLD").title?
  end

  def test_reverse
    assert_equal "olleH", @paragraph.reverse
  end

  def test_upcase
    assert_equal "HELLO", @paragraph.upcase
  end
end

