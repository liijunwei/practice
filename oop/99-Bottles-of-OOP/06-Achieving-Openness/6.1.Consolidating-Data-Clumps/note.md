+ Despite much refactoring, the code is still not open to the six-pack requirement.

尽管前面做了很多重构工作, 但是现有代码还是没有能够对新需求的扩展"开放"

+ Once again, you must decide whether to continue forward with the existing code, or to retreat and strike out in a different direction.

此时需要再次确认, 我们想要继续重构, 直至它对新需求开放, 还是在现有代码基础上实现新需求

+ Consolidating Data Clumps

整合数据束(???)

+ Before focusing on that problem, however, there’s a simpler code smell that can be addressed.

处理问题前, 先把code smell消除掉

+ Data Clump is officially about data, and is defined as the situation in which several (three or more) data fields routinely occur together.

"Data Clump"是关于数据的官方术语, 意思是: 多个(3+)数据字段通常同时出现的情况

+ Having a clump of data usually means you are missing a concept.

代码出现了"data clump"通常意味着你忽略了一个概念(这个概念没被识别出来)


