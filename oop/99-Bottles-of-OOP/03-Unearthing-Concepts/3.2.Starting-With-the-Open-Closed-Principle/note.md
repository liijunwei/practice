+ The decision about whether to refactor in the first place should be determined by whether your code is already "open" to the new requirement.

决定是否需要重构代码前, 需要先确定你的代码是否支持 对新需求进行扩展

+ "Open" is short for "Open/Closed," which in turn is short for "open for extension and closed for modification."

Open 是 开闭原则的简写, 表示代码要对扩展开放, 对修改关闭

+ Code is open to a new requirement when you can meet that new requirement without changing existing code.

即: 代码写好以后, 为适应新的需求, 应该尽量少去考虑修改原有代码, 而应考虑去扩展原有代码

