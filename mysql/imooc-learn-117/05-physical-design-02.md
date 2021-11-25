### [MYSQL常用存储引擎](https://www.imooc.com/video/1909)

> 只要实现了mysql的存储协议, 各厂商可以开发自己的存储引擎

+ MyISAM
+ MRG_MYISAM
+ Innodb: 支持MVCC行级锁(多版本并发控制)
+ Archive
+ Ndb cluster


### [数据库表及字段的命名规则](https://www.imooc.com/video/1910)

所有对象命名各应该遵循下述原则:

+ 可读性原则
    + 使用大写和小写来格式化 库对象名字, 以获得良好的可读性(EmailAddress / emailaddress)
    + 要注意有些DBMS对表明的大小写敏感

+ 表意性原则
    + 对象的名字应该能够描述它所标识的对象
    + e.g. 对于表来说, 表的名称应该能体现出表中存储的数据内容
    + e.g. 对于存储过程来说, 存储过程表的名称应该能体现出存储过程的功能(???)

+ 长名原则
    + 尽可能少用或者不用缩写
    + 适用于数据库(Database)名外的任一对象
