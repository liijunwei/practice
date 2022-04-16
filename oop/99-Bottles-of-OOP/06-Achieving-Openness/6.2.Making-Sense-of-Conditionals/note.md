+ Switch your attention to the BottleNumber class. It’s full of conditionals, all of which have the same shape.(Code Smell)

+ vociferously 大声地(voice)

```ruby
# depend on the number argument
class Bottles
  def container(number)
    if number == 1
      "bottle"
    else
      "bottles"
    end
  end
end

# depend on the number message
class BottleNumber
  attr_reader :number

  def initialize(number)
    @number = number
  end

  def container
    if number == 1
      "bottle"
    else
      "bottles"
    end
  end
end

```