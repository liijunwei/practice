+ The final unresolved issue is that the lyrics instance method in BottleVerse directly references the concrete BottleNumber class.

```ruby
class BottleVerse
  # ...
  def lyrics
    bottle_number = BottleNumber.for(number)

    "#{bottle_number} of beer on the wall, ".capitalize +
    "#{bottle_number} of beer.\n" +
    "#{bottle_number.action}, " +
    "#{bottle_number.successor} of beer on the wall.\n"
  end
end
```


