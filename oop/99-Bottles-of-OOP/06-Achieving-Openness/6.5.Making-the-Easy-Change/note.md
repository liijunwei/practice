+ The previous horizontal refactoring is complete, and it is again time to ask if the code is open to the six-pack requirement.

现在代码是否对扩展开放了呢?

+ You can now meet the six-pack requirement by adding a new class that stands in for bottle number 6. This new class will report its quantity as "1" and its container, "six-pack."

已经满足开闭原则里的 "开放" 了

你只需要新加一个类"BottleNumber6"

里面覆写"quantity"和"container"方法, 然后在工厂里增加一个分支就好了


