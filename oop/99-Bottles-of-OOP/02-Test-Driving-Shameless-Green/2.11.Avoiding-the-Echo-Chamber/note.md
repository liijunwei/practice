+ remedy 补救

```ruby
def test_the_whole_song
  bottles = Bottles.new
  assert_equal bottles.verses(99, 0), bottles.song
end
```

+ This flaw lies dormant until something changes, so the benefits of writing tests like this accrue to the writer today, while the costs are paid by an unfortunate maintainer in the future.

这种测试用例有效果, 但是有个很严重的问题: 代码一旦需要修改, 这个测试用例也必须跟着修改, 并且还得了解修改的细节, 这个问题会一直处于"未被发现"的状态, 直到代码发生变化

写出这种测试用例的人, 写的时候很轻松, 但会让维护的人很痛苦

+ Understanding this flaw requires being clear about song 's responsibilities.

为了更清晰的了解这种用例的问题, 需要先明确`song`方法的职责

+ From the message sender’s point of view, song is responsible for returning the lyrics for all 100 verses.

从消息发送者的角度看, `song`方法负责返回整篇歌词

+ Imagine that you were tasked to test this method but knew nothing about how Bottles was implemented. You would be unaware of the existence of the verses method, and would have no choice other than to test song by asserting that its output matched those lyrics.

想象一下你要测试这个方法, 但是你对`Bottles`的实现一无所知

这时候你甚至不知道`verses`方法的存在

这时候你只能用整篇歌词来做测试

+ Asserting that song returns the expected lyrics is very different from asserting that song returns the same thing as verses .

断定song的返回值是整篇歌词 和 断定song的返回值和verses的一样 之间有巨大的差异

+ In the first case, the song test is independent of implementation details and so tolerates changes to other parts of the class without breaking.

前者认为song的测试用例和实现细节无关

因此能够在实现细节改变时, 不需要改动测试用例

+ In the second case, the song test is coupled to the current Bottles implementation such that it will break if the signature or behavior of verses changes, even if song continues to return the correct lyrics.

后者认为song的测试用例和具体实现有了耦合

这种情况下, 即使song的行为没有改变, 但是verses的实现改变了, 也可能导致测试用例的崩溃

+ There’s nothing more frustrating than making a change that preserves the behavior of an application but breaks apparently unrelated tests.

对代码做出改动时, 虽然没有改变整体行为, 但是还是让不相关的测试用例失败了, 这种事儿让人很崩溃

+  If you change an implementation detail while retaining existing behavior and are then confronted with **a sea of red**(生动形象), you are right to be exasperated(被激怒了).
+ This is completely avoidable, and a sign of tests that are too tightly coupled to code. Such tests impede change and increase costs.

+ Not only is the above song test too tightly-coupled to the current Bottles implementation, it doesn’t even force you to write the right code.

上面的那种测试用例的问题 不仅在于测试和代码耦合过深, 而且会迫使你写出错误的代码(page 80): 错误的verses的实现


```ruby
def test_the_whole_song
  bottles = Bottles.new
  expected = 99.downto(0).map {|i| bottles.verse(i)}.join("\n")
  assert_equal expected, bottles.song
end
```
+ Above test case doesn’t improve the test, but just tightly couples the test to code that’s one step farther back in the stack. If that more-distant code changes, this test might break.

+ The song test should know nothing about how the Bottles class produces the song.

song的测试用例应该不对Bottle的实现做任何假设

+ The clear and unambiguous expectation here is that song return the complete set of lyrics, and the best and easiest way to test song is to explicitly assert that it does.

直接测试整首歌的歌词, 就是最简单和直接的方法

```ruby
def test_the_whole_song
  filepath = File.expand_path(".", '../whole-lyric.md')
  expected = File.read(filepath)
  assert_equal expected, Bottles.new.song
end
```

+ The text needed for 100 verses is fairly lengthy, and you may resist writing out the full string because of concerns about duplication.





