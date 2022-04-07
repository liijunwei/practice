```ruby
  def verse(number)
    case number
    when 0
      # ...
      "Go to the store and buy some more, " +
      # ...
    else
      # ...
      "#{quantity(number).capitalize} #{container(number)} of beer on the wall, " +
      # ...
    end
  end
```

+ The only thing the above lines have in common is the trailing ", " , which means that everything up to that point is a difference.

+ If the 0 and else verse variants reflect a common verse abstraction, this difference must represent a smaller concept within that larger abstraction.

+ **It doesn’t matter how long these strings are, their presence here in opposition means they reflect a single concept.**


+ **You must name the concept, create a method to represent it, and then replace this difference with a message send.**

+ The first step is therefore to name the category in which these two phrase are concrete examples.


