+ You’ve already practiced Extract Class, back in the Extracting BottleNumber section of chapter 5 where a new class was created to cure the Primitive Obsession on number.

前面我们已经联系过从已有代码里提出class了

这么做的理由是: 让代码不再依赖原始数据类型

+ Here’s a reminder of that recipe:
    1. Choose a class name and create the new class.
    2. Add an attr_reader and an initialize method to encapsulate primitive data. Copy the methods from the old class to the new.
    3. Forward messages from the old class to the new.
    4. One by one, remove arguments from the methods in the new class, and corresponding parameters from the message sends in the old class.

步骤如下(保证每步里, 测试都是通过的):

    1. 选择类名, 创建一个新类
    2. 为类新增初始化接口, 并增加reader方法
    3. 将旧代码copy进新类里面, 然后逐步替换掉旧代码里的调用
    4. 将新类里方法的参数逐渐去掉


