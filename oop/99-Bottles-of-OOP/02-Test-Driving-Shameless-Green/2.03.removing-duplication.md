+ This next test should do the simplest, most useful thing that proves your existing code to be incorrect.

写完第一个测试后, 接下来要写更多的测试用例

这个测试用例也得很简单, 但是这个用例得能打破第一个用例带来的"正确"的感觉

+ this next test is often easier because it checks something relative to the first.

新的测试用例需要检查一些和第一个用例不太一样的地方, 这种检查一般能让第一个用例失效


+ TDD requires that you pass tests by writing simple code.

使用TDD时, 需要能写一些简单的代码, 让测试能够通过

+ However, most programming problems have many solutions, and it’s not always clear which one is simplest.

很多的问题会有很多的解决方案, 找出最简单的方案并不容易

+ This conditional adds a new execution path through the code, and additional execution paths increase complexity.

如果仅仅通过加一个判断, 使得测试通过了, 乍看起来好像很简单, 实际上代码的复杂度提升了, 并且它能处理的情况也很有限

conflates 合并

