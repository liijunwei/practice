class Bottles
  def song
    lyric_p1 = song_first_part
    lyric_p2 = song_second_part

    return lyric_p1 + lyric_p2
  end

  def song_first_part
    lyric = ""

    99.downto(3) do |number|
      lyric += <<~EOF
#{number} bottles of beer on the wall, #{number} bottles of beer.
Take one down and pass it around, #{number - 1} bottles of beer on the wall.

      EOF
    end

    return lyric
  end

  def song_second_part
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

