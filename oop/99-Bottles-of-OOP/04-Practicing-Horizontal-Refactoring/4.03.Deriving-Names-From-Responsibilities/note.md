page 128

```bash
tig oop/99-Bottles-of-OOP/04-Practicing-Horizontal-Refactoring/lib/bottles.rb
```

```ruby
"no more bottles"

# vs

"#{number-1} #{container(number-1)}"
```

+ If it’s not clear how to proceed, look for a way to make the lines more alike (even if not yet identical), using code you’ve already written.

如果不清楚如何继续, 就使用已有的代码, 寻找一种使行更相似(即使尚未完全相同)的方法

+ Remember that the goal is to locate the next small difference, not the next clump of differences.

记住, 目标是找到一点区别, 不是找到所有的区别

+ If each verse variant reflects a more general verse abstraction, then the differences between the variants must represent smaller concepts within that larger abstraction.

如果每一个诗歌变体都反映了一个更一般的诗歌抽象, 那么变体之间的差异必须代表更大抽象中的较小概念

+ Again, you can resolve this difference by following the pattern you learned from container and pronoun.

可以通过仿照`container`和`pronoun`的思路来处理这个差异

+ Name the concept, create the method, and replace the difference with a common message send.

为概念命名, 创建一个方法, 然后把不同的地方替换为同一方法

+ To help you name the new concept, remember the "what would the column header be?" technique. The following table shows a sampling of numbers and associated values:

Number | XXX?
-------|--------
99     | '99'
50     | '50'
1      | '1'
0      | 'no more'

(Number to XXX Column Header)

+ When creating an abstraction, first describe its responsibility as you understand it at this moment, then choose a name which reflects that responsibility.

+ The effort you put into selecting good names right now pays off by making it easier to recognize perfect names later.

