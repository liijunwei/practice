### Q: Why are all these wonderful techniques really only the beginning?

+ Because you don't yet know when to use them and when not to, when to start and when to stop, when to go and when to wait.
+ It is the rhythm that makes for refactorsing, not the individual notes.

### Q: How will you know when you are really getting it?

+ You'll know when you start to calm down.
+ **When you feel absolute confidence that no matter how screwed up someone left it, you can make the code better, enough better to keep making progress.**

+ Mostly, though, you'll know you're getting it when you can stop with confidence.
+ **Stopping is the strongest move in the refactorer's repertoire.**

这种心态真的太好了! 我已经有一点点这种感觉了(虽然重构技术用得还很不熟练)

+ If the code is already better, integrate and release what you've done.
+ If it isn't better, walk away. Flush it. Glad to have learned a lesson, pity it didn't work out. What's on for tomorrow?

+ You bed down for the night, sure the sun will rise again in the morning.

> When you really understand refactoring, the design of the system is as fluid and plastic and moldable to you as the individual characters in a source code file.
> You can feel the whole design at once.
> You can see how it might flex and change a little this way and this is possible; a little that way and that is possible.

### I said this was a learnable skill. How do you learn it?

+ Get used to picking a goal. Somewhere your code smells bad. Resolve to get rid of the problem.
    + You aren't refactoring to pursue truth and beauty
    + **You are trying to make your world easier to understand, to regain control of a program that is flapping loose.**

+ Stop when you are unsure.
    + preserve the semantics of your program
    + **If the code is already better, go ahead and release your progress. If it isn't, throw away your changes.**

+ Backtrack(回溯)
    + take small steps
    + run your tests frequently
    + **Go back to your last known good configuration, Replay your changes one by one. Run the tests after each one.**
    + "I think my personal record for failing to follow my own advice is four hours and three separate tries."

+ Duets(二重奏?)
    + For goodness' sake, refactor with someone. **There are many advantages to working in pairs for all kinds of development.**
    + 结对编程 来做重构

+ As ugly as the mess looks now, discipline yourself to nibble away at the problem.
即使代码看起来一团糟, 有时也要抑制住一下做出很多改动的冲动, 要慢慢训练自己解决问题的能力

+ When you are going to add some new functionality to an area, take a few minutes to clean it up first.
如果你需要为应用添加新的功能, 先花点时间来清理代码;

+ If you have to add some tests before you can clean up with confidence, add them. You'll be glad you did.
如果你觉得在清理代码前有必要先增加一点测试来提升信心, 尽管加吧, 好处多多

+ Touching the code will remind you how it works.
在为新功能修改代码前先这样做一些清理/增加测试的事情, 能提示我们对代码的熟悉程度

+ When you decide to undertake a large refactoring, try to pick off pieces that can be integrated back into the main development branch as quickly as possible.
当你决定承担起一个较大的重构任务的时候, 先选出一些小的部分做修改, 然后确保他们能尽快的集成进已有系统里(small steps with confidence; keep close with the main branch)

+ When contemplating a large refactoring it's tempting to say, "I can't improve that code without adopting a big-bang approach that will take 3 days."
    + 做较大的重构时, 有一种很具诱惑的想法是: 我很难在不做出很大的breaking changes的情况下, 完成这项重构工作;
    + **实际上不是这样, 一般都会有兼容的方案 让重构尽可能平滑的完成**
    + But the extra time spent writing this throw-away adaptive code is **worth the benefit of continual integration with the development branch**.

+ Never forget the two hats:
    + **The refactoring hat**
    + **The new functionality hat**
    +
    + Only wear one hat at a time

+ When you refactor, you will inevitably discover code that doesn't work correctly.
+ You'll find bugs, test cases to add or change, and other unrelated refactorings.
+ Some of these might even be more important than the refactoring you're currently working on.
+ Resist temptation to mix an unfinished refactoring with one of these newfound tasks.

重构的时候, 不可避免的会发现一些当前目标(重构某部分代码, 添加新功能)以外的 问题/bug/不合理之处 等等(深有体会...)
有些问题可能比你当前的任务更加重要, 但是 要抵制住 将当前正在做的`重构+添加新功能`的事和 那些`新发现的问题`混在一起的诱惑

+ If the newfound task truly is an immediate priority, abandon your refactoring.
+ Revert the code and start a fresh.
如果新的问题真的很重要, 优先级非常高, 那么停止手头的工作, 把当前做的事暂时搁置(commit/stash/abort), 然后基于最新的代码开出新的分支, 专门处理那个问题

+ But if you decide to wear the refactoring hat, your goal is to leave the code computing exactly the same answers that it was when you found it; nothing more, nothing less.
+ Once you develop the discipline and rhythm to juggle the two hats, you'll find refactoring to be a rewarding and productive experience. Happy coding!
但是如果他们没那么重要, 而且你现在要专注于重构一个部分, 你应该忽略掉那些些你发现的问题(可以先记录下来, 以后再处理), 不要被他们分心, 不要修改他们
当我们适应了这种 重构-添加新功能 的节奏后, 你会真正从重构之中获益

happy coding!


