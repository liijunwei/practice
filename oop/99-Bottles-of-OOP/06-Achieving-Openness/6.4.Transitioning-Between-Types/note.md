+ The code now consists of a pleasing set of small objects with clear-cut responsibilities.

当前版本的代码已经由一系列小的对象组成了, 各个小模块职责清晰

+ However, there’s one persistent problem that can no longer be ignored: the successor methods violate the generalized Liskov Substitution Principle.

但是还有个问题需要处理: `successor`方法违反了里氏替换原则

+ You have every right to expect the successor of a bottle number to act like a bottle number

我们有理由认为 一个bottle_numer的实例的successor也具有bottle_number的行为

但是当前的`successor`方法返回了一个integer类型的数字, 显然它并不具有bottle_number的行为
