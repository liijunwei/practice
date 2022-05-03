# 9.4. Communicating With the Future

## 9.4.1. Enriching Code with Signals

+ Arrange-Act-Assert (AAA) pattern
    1. defines an expectation,
    2. executes some code, and
    3. asserts that the result matches the expectation

+ Committing to a common shape lowers the cognitive load imposed upon future readers.

+ When every test is structured in the same way, readers don’t have to waste time determining if spurious differences matter.

+ Choosing to arrange tests in a similar way is a matter of programming style.

+ You’re always on the lookout for code arrangements that send signals to these hapless readers.

+ Style is one kind of signal, but there are others.
    + This test uses prime numbers to signal arbitrariness.
    + The 500 used in test_verse is so ridiculously large that it already signals that any number, even a big one, will do.

+ **Once you start thinking about sending signals, opportunities abound.**

+ Weirich used do...end to warn that the enclosed block had side-effects, and {...} to assure that it did not. Readers of his code remain grateful for these signals.

+ Signals offer a cheap way to add valuable information to code.
    + Look for opportunities to develop and use them.
    + Even if your team consists only of you, the information imparted by a signal will be useful to your future self.

## 9.4.2. Verifying Roles

+ Ruby is dynamically typed and doesn’t have syntax to support interfaces.
    + opportunities: easy to create and use polymorphic duck types
    + You may do anything without constraint, including sending messages to objects that don’t understand them and suffering run-time failures.

ruby作为动态语言的优势和问题

优势: 灵活, 很容易使用duck type实现多态, 代码好写好读

劣势: 过于灵活, 容易出错


