+ **New requirement: Your customer wants other songs that are similar to "99 Bottles" but contain different lyrics.**

+ When your customer asks you to produce other songs like the "99 Bottles" song, they’re asking you to write code to produce other songs that also count down.

+ Follow the Open-Closed Flowchart:

1. Is the code open to the vary the verses requirement?

NO

2. Do you know how to make it open?

have no idea

3. What’s the best-understood code smell?

have no idea

+ Paradoxically, when faced with uncertainty about what to do next, it can sometimes help to sigh deeply, ignore everything you’ve learned, and just write a new conditional.

自相矛盾的时, 当面对下一步该做什么的不确定性时, 有时深叹一口气, 忽略你所学的一切, 加一个条件, 把需求实现了往往是有助于进展的

+ This conditional’s purpose is to supply more information about the problem, and writing it can clarify what needs to change.

先写下来, 先实现了功能, 能够为理解问题提供更多的信息

当了解到的信息增多以后, 那里需要做改变就变得清晰了

+ Once you understand what should change, you can discard the conditional and write better code.

当你理解了问题后, 你就可以逐渐把条件判断去掉了

+ the conditional’s mere presence helps clarify what needs to be done.
