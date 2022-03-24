class Bottles
  def verses(_, _)
   "99 bottles of beer on the wall, " +
   "99 bottles of beer.\n" +
   "Take one down and pass it around, " +
   "98 bottles of beer on the wall.\n" +
   "\n" +
   "98 bottles of beer on the wall, " +
   "98 bottles of beer.\n" +
   "Take one down and pass it around, " +
   "97 bottles of beer on the wall.\n"
  end

  def verse(number)
    case number
    when 0
      "No more bottles of beer on the wall, " +
      "No more bottles of beer.\n" +
      "Go to the store and bug some more, " +
      "99 bottles of beer on the wall.\n"
    when 1
      "1 bottle of beer on the wall, " +
      "1 bottle of beer.\n" +
      "Take one down and pass it around, " +
      "no more bottles of beer on the wall.\n"
    when 2
      "2 bottles of beer on the wall, " +
      "2 bottles of beer.\n" +
      "Take one down and pass it around, " +
      "1 bottle of beer on the wall.\n"
    else
      "#{number} bottles of beer on the wall, " +
      "#{number} bottles of beer.\n" +
      "Take one down and pass it around, " +
      "#{number - 1} bottles of beer on the wall.\n"
    end
  end
end
