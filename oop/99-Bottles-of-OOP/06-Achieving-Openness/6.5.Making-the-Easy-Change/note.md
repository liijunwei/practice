+ The previous horizontal refactoring is complete, and it is again time to ask if the code is open to the six-pack requirement.

现在代码是否对扩展开放了呢?

+ You can now meet the six-pack requirement by adding a new class that stands in for bottle number 6. This new class will report its quantity as "1" and its container, "six-pack."

已经满足开闭原则里的 "开放" 了

你只需要新加一个类"BottleNumber6"

里面覆写"quantity"和"container"方法, 然后在工厂里增加一个分支就好了

+ At long last, it’s time to write a failing test.

前面所做的重构都是依靠着写好的测试用例完成的

现在引入了新的需求, TDD要求我们先写出会失败的测试用例, 然后开始新增代码, 让测试用例通过

+ You have been refactoring under green for many chapters, and now, suddenly, almost abruptly, the outstanding requirement can be met by two one-line methods in a class that has nine total lines of code. It took several refactorings to make the code open, but once so, the six-pack requirement was extraordinarily easy to fulfill.
    + `tig 3cddfa1ece74e8b18b278224982994ebed4b9d99`


