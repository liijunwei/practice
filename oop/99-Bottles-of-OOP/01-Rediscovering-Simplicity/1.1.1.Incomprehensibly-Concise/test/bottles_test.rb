gem 'minitest', '~>5'
require 'minitest/autorun'
require 'minitest/pride'
require_relative '../lib/bottles'

class BottlesTest < Minitest::Test
  def test_the_whole_song
    filepath = File.expand_path(".", '../../whole-lyric.md')
    expected = File.read(filepath)
    assert_equal expected, Bottles.new.song
  end
end
