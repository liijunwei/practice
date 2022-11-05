+ If you are unclear about how to make it open (which is often the case), the way forward is to start removing code smells.

如果你不确定在么把代码变得对扩展开放, 可以先从消除代码的坏味道开始

+ If the smells aren’t immediately obvious, start by making a list of the things you find objectionable(令人反感的).

如果代码的坏味道也不好找, 就从列出代码中最让你反感的列表开始

```ruby
def verse(number)
  case number
  when 0
    "No more bottles of beer on the wall, " +
    "no more bottles of beer.\n" +
    "Go to the store and buy some more, " +
    "99 bottles of beer on the wall.\n"
  when 1
    "1 bottle of beer on the wall, " +
    "1 bottle of beer.\n" +
    "Take it down and pass it around, " +
    "no more bottles of beer on the wall.\n"
  when 2
    "2 bottles of beer on the wall, " +
    "2 bottles of beer.\n" +
    "Take one down and pass it around, " +
    "1 bottle of beer on the wall.\n"
  else
    "#{number} bottles of beer on the wall, " +
    "#{number} bottles of beer.\n" +
    "Take one down and pass it around, " +
    "#{number-1} bottles of beer on the wall.\n"
  end
end
```

+ This method contains a case statement (the Switch Statements smell) whose branches contain many duplicated strings (Duplicated Code). Of these two smells, Duplicated Code is the most straightforward and so will be tackled first.

+ Therefore, the current task is to refactor the verse method to remove the duplication, in hope and expectation that the resulting code will be more open to the six-pack requirement.

先来处理冗余的代码

+ Before undertaking(进行/承诺/保证) this refactoring, it must be admitted that there is no direct connection between removing the duplication, and succeeding in making the code open to the six-pack requirement.

在进行重构前, 得先意识到(承认)"删除冗余代码"和"使得代码对应扩展开放"之间没有必然(直接)的练习

+ That, however, is the beauty of this technique. You don’t have to know how to solve the whole problem in advance. The plan is to nibble away, one code smell at a time, in faith that the path to openness will be revealed.

这也正是这项技术的美妙之处

你不需要提前知道解决问题的所有技巧

我们的计划是一次一点地蚕食(nibble away)一种代码坏味道, 并且相信通往开放的道路将被揭示





