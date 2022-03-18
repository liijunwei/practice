gem 'minitest', '~>5'
require 'minitest/autorun'
require 'minitest/pride'
require_relative '../lib/bottles'

# rvm use 2.5.1
# ruby test/bottles_test.rb
class BottlesTest < Minitest::Test
  def test_the_whole_song
    filepath = File.expand_path(".", '../../whole-lyric.md')
    expected = File.read(filepath)
    assert_equal expected, Bottles.new.song
  end
end
