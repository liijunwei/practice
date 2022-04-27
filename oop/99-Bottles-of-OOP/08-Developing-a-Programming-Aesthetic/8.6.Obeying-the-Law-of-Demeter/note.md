# 8.6. Obeying the Law of Demeter

+ Lines that contain many dots might violate the Law of Demeter (LoD).
+ This section defines that law, determines where it applies, explores the consequences of ignoring it, and explains how to fix violations.

+ Listing 8.19: Verse Method Contains Many Dependencies
```ruby
class Bottles
  # ...
  def verse(number)
    verse_template.new(number).lyrics
  end
end
```
On line 4 above the verse method knows:

