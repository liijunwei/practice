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

+ If you have to add some tests before you can clean up with confidence, add them. You’ll be glad you did.
如果你觉得在清理代码前有必要先增加一点测试来提升信心, 尽管加吧, 好处多多

+ Touching the code will remind you how it works.

+ When you decide to undertake a large refactoring, try to pick off pieces that can be integrated back into the main development branch as quickly as pos- sible.
