# 9.3. Seeking Context Independence

+ An object’s **context** is its surrounding environment, or the interrelated(相互关联) conditions under which it can exist.

+ Objects are more usable when they know less.
    + Objects that expect little of their surroundings can be more easily applied to novel situations, and so provide greater utility.
    + The most useful are those that are entirely independent of context; they have no expectations about, and make no demands upon the external world.

+ As far as context goes, ignorance truly is bliss.

就上下文而言, 无知真的很幸福

+ This section explores these ideas by examining the Bottles class and its remaining tests with an eye towards reducing their context.

本节通过检查`Bottles`类和它剩下的测试用例来探索这些想法, 希望能减少`Bottles`的上下文

## 9.3.1. Examining Bottles' Responsibilities

```ruby
class Bottles
  attr_reader :verse_template

  def initialize(verse_template: BottleVerse)
    @verse_template = verse_template
  end

  def song
    verses(99, 0)
  end

  def verses(starting, ending)
    starting.downto(ending).map {|num| verse(num)}.join("\n")
  end

  def verse(number)
    verse_template.lyrics(number)
  end
end
```

+ Take a fresh look at this code and describe what it does. Classes should be named after what they are: what is this?
    + **This is a CountdownSong.**

+ parochial [pəˈroʊkiəl] 只关心本地区的;教区的;堂区的;地方观念的

+ Bottles is parochial and useful in one specific case.
+ CountdownSong is generic and useful in many cases.

+ Bottles can be renamed to CountdownSong with a straightforward refactoring.

## 9.3.2. Purifying Tests With Fakes

+ It’s a given that tests should prove that code works. But they have other purposes too, one of which is to explain the domain to the reader.

+ Tests should explain the essence of a class.

+ design patterns: named, re-usable solutions to common software problems

+ Since pattern names are so meaningful, it can be tempting to stick them in class names.

+ Class names should reflect concepts in your domain, not the patterns used to create them.

## 9.3.3. Purging Redundant Tests

+  As long as you have tests to prove that every verse is correct, you don’t need to test verses against more than one range. If verses works with one, it will work with all.








