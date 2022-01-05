class Bottles
  def song
    song_part1 + song_part2
  end

  def song_part1
    lyric = ""

    99.downto(3) do |number|
      lyric += <<~EOF
#{number} bottles of beer on the wall, #{number} bottles of beer.
Take one down and pass it around, #{number - 1} bottles of beer on the wall.

      EOF
    end

    return lyric
  end

  def song_part2
    lyric = <<-SONG
2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.

1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.

No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
    SONG

    return lyric
  end
end

