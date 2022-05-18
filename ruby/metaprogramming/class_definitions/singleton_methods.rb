# page 111

require 'test/unit'

class TestParagraph < Test::Unit::TestCase
  def setup
    @str = "just a regular string"

    def @str.title?
      self.upcase == self
    end
  end

  def test_title_for_str
    refute @str.title?
  end

  def test_methods_for_str
    assert_equal [:"title?"], @str.methods.grep(/title?/)
  end

  def test_singleton_methods_for_str
    assert_equal [:"title?"], @str.singleton_methods
  end
end

