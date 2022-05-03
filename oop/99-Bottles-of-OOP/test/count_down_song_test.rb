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

  # OK 没看懂 page 265 讲的 为什么不需要测试这 2,1,0 这几个特殊的
  # 因为对CountdownSong来说, 他不知道, 或者说他不关心有没有特殊的情况吗?
  # 对, CountdownSong 里面, 没有特殊情况, 不一样的地方交给 BottleNumber.for 工厂处理了

  def test_the_whole_song
    filepath = File.expand_path(".", "./whole-lyric-six-pack.md")
    expected = File.read(filepath)
    assert_equal expected, CountdownSong.new.song
  end
end
