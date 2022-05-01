require 'minitest/autorun'
require_relative '../lib/bottles'

class CountdownSongTest < Minitest::Test
  def test_a_couple_verses
    expected =
   "99 bottles of beer on the wall, " +
   "99 bottles of beer.\n" +
   "Take one down and pass it around, " +
   "98 bottles of beer on the wall.\n" +
   "\n" +
   "98 bottles of beer on the wall, " +
   "98 bottles of beer.\n" +
   "Take one down and pass it around, " +
   "97 bottles of beer on the wall.\n"
    assert_equal expected, CountdownSong.new.verses(99, 98)
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
    filepath = File.expand_path(".", './whole-lyric-six-pack.md')
    expected = File.read(filepath)
    assert_equal expected, CountdownSong.new.song
  end
end
