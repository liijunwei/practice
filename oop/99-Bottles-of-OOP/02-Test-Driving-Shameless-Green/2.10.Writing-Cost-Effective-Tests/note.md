+ TDD does not claim to be free, merely that its benefits outweigh its costs.

TDD的开发方式是有代价的, 只是这种开发方式的 好处 和 代价权衡的结果是好处更多

+ Belief in the value of TDD has become mainstream, and the pressure to follow this practice approaches an unspoken mandate(授权).

+ Even those who don’t test seem to believe they ought to.

+ Despite this general agreement, the sad truth is that the promise of TDD has not been universally fulfilled.

虽然大家(多数人)都认可TDD的好处和价值, 但是事实是TDD的优势还没有被发挥出来, 或者说没有没正确是用

+ Many applications have tests that are difficult to understand, challenging to change, and prohibitively(令人望而却步) time-consuming to run.

很多使用TDD的人写出的测试用例很难理解, 很难修改, 运行起来非常耗时

+ Instead of enabling change, these tests actively impede(阻碍/妨碍/阻止) it.

+ A great deal of this pain originates with tests that are tied too closely to code. When this is true, every improvement to the code breaks the tests, forcing them to change in turn.

来自于测试用例的痛苦体验, 有一大部分来自于测试用例和代码的耦合太强

结果就是代码的一点点改动, 都会让测试用例失败, 导致用例需要修改

+ Therefore, the first step in learning the art of testing is to understand how to write tests that confirm what your code does without any knowledge of how your code does it.

学习测试的艺术的第一步是: 要理解怎么才能写出能验证代码功能的用例, 同时不需要了解代码是如何实现功能的


