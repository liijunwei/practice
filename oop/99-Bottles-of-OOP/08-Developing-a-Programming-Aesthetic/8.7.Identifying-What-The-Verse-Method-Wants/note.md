```ruby
class Bottles
  # ...
  def verse(number)
    verse_template.new(number).lyrics
  end
end
```

+ line 5 above violates `the law of Demeter`

+ From the classes-are-just-objects perspective, this line of code is very much a LoD violation.

+ Sending new to verse_template returns a object that conforms to a different API than the original verse_template receiver.

+ Put another way:
    + The API of the object that responds to new is different from the API of the object that responds to lyrics.
    + While a class and an instance of that class are certainly both objects, they are very different kinds of things.

+ The rule for injecting dependencies is that you should inject the thing you want to talk to.



