require 'minitest/autorun'
require_relative '../lib/bottles'

class BottleVerseTest < Minitest::Test
  def test_the_first_verse
    expected =
    "99 bottles of beer on the wall, " +
    "99 bottles of beer.\n" +
    "Take one down and pass it around, " +
    "98 bottles of beer on the wall.\n"

    assert_equal expected, BottleVerse.lyrics(99)
  end
end
