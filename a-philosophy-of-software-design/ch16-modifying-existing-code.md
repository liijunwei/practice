+ It isn’t possible to conceive the right design for a system at the outset; the design of a mature system is determined more by changes made during the system’s evolution than by any initial conception.

系统一般是慢慢向更好的方向演进的，不太可能一上来就是尽善尽美的

+ Previous chapters described how to squeeze out complexity during the initial design and implementation; this chapter discusses how to keep complexity from creeping in as the system evolves.

前面的章节介绍了怎么在设计和实现阶段降低系统的复杂性

这一章讲会讨论怎么在系统演进的工程中降低或者排除复杂性

+ If you want to have a system that is easy to maintain and enhance, then “working” isn’t a high enough standard;
+ you have to *prioritize design and think strategically*.
+ This idea also applies when you are modifying existing code.

+ *Ideally, when you have finished with each change, the system will have the structure it would have had if you had designed it from the start with that change in mind.* 不好懂... 最理想的情况下，如果做完修改，系统还是一个很理想的样子，而不需要额外的修正或者重构是最好的，说明在最初设计的时候就考虑到了一些可能发生变化的地方，并且做了准备
+ To achieve this goal, you must resist the temptation to make a quick fix.
    + Instead, think about whether the current system design is still the best one, in light of the desired change.
    + If not, refactor the system so that you end up with the best possible design.
    + With this approach, the system design improves with every modification.
    + 积少成多，持续的小的重构是必要的，这样才能保证系统的整体设计逐渐变的更好
    + **if you invest a little extra time to refactor and improve the system design, you’ll end up with a cleaner system.**
    + **If you’re not making the design better, you are probably making it worse.**

+ an investment mindset sometimes conflicts with the realities of commercial software development.
    + 理想和现实可能会有矛盾
+ Nonetheless, you should resist these compromises as much as possible. Ask yourself “Is this the best I can possibly do to create a clean system design, given my current constraints?”
    + 但是无论如何，应该尽量避免轻易的妥协

+ Comments belong in the code, not the commit log
    + 注释应该写在代码里，而不是commit信息里

+ maintaining comments: avoid duplication
+ try to document each design decision exactly once.
+ find the most **obvious** single place to put the documentation.

+ If information is already documented someplace outside your program, don’t repeat the documentation inside the program; just reference the external documentation.

+ One final thought on maintaining documentation: comments are easier to maintain if they are higher-level and more abstract than the code.
    + describe the interface, instead of the implementation
