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


