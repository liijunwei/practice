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

+ the verse method knows:
    + that verse_template responds to new
    + that new expects an argument
    + that the argument to new must be a number
    + that the object returned from new(number) responds to the message lyrics that lyrics returns the actual lyrics of interest

+ **This list enumerates many things that Bottles knows about but doesn’t control, which means they’re dependencies**
+ Dependencies are vulnerabilities

依赖 是 脆弱性的来源

+ Dependencies can’t be avoided but should certainly be minimized.

应用程序无法避免依赖, 但是 应该尽量减少依赖

+ Be alert for superfluous dependencies and remove them with extreme prejudice.

警惕多余的依赖, 并以极端偏见消除它们

+ [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter)
    + In its general form, the LoD is a specific case of loose coupling.
    +
    + Each unit should have only limited knowledge about other units: only units "closely" related to the current unit.
    + Each unit should only talk to its friends; don't talk to strangers.
    + Only talk to your immediate friends.

迪米特法则, 是松耦合的一种特例



