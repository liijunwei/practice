require 'minitest/autorun'
require_relative '../lib/bottles'

class VerseFake
  def self.lyrics(number)
    "This is verse #{number}.\n"
  end
end

class CountdownSongTest < Minitest::Test
  def test_verse
    expected = "This is verse 500.\n"

    assert_equal expected, CountdownSong.new(verse_template: VerseFake).verse(500)
  end

  def test_verses
    expected =
    "This is verse 99.\n" +
    "\n" +
    "This is verse 98.\n" +
    "\n" +
    "This is verse 97.\n"

    assert_equal expected, CountdownSong.new(verse_template: VerseFake).verses(99, 97)
  end

  def test_song
    expected =
    "This is verse 47.\n" +
    "\n" +
    "This is verse 46.\n" +
    "\n" +
    "This is verse 45.\n" +
    "\n" +
    "This is verse 44.\n" +
    "\n" +
    "This is verse 43.\n"

    assert_equal expected, CountdownSong.new(verse_template: VerseFake, max: 47, min: 43).song
  end
end
