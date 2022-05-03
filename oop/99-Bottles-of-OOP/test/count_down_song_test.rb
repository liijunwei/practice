require 'minitest/autorun'
require_relative '../lib/bottles'

class VerseFake
  def self.lyrics(number)
    "This is verse #{number}.\n"
  end
end

class CountdownSongTest < Minitest::Test
  def test_a_couple_verses
    expected =
    "This is verse 99.\n" +
    "\n" +
    "This is verse 98.\n" +
    "\n" +
    "This is verse 97.\n"

    assert_equal expected, CountdownSong.new(verse_template: VerseFake).verses(99, 97)
  end

  def test_a_few_verses
    expected =
   "2 bottles of beer on the wall, " +
   "2 bottles of beer.\n" +
   "Take one down and pass it around, " +
   "1 bottle of beer on the wall.\n" +
   "\n" +
   "1 bottle of beer on the wall, " +
   "1 bottle of beer.\n" +
   "Take it down and pass it around, " +
   "no more bottles of beer on the wall.\n" +
   "\n" +
   "No more bottles of beer on the wall, " +
   "no more bottles of beer.\n" +
   "Go to the store and buy some more, " +
   "99 bottles of beer on the wall.\n"
    assert_equal expected, CountdownSong.new.verses(2, 0)
  end

  def test_the_whole_song
    filepath = File.expand_path(".", "./whole-lyric-six-pack.md")
    expected = File.read(filepath)
    assert_equal expected, CountdownSong.new.song
  end
end
