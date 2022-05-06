### [数据库设计其他注意事项](https://www.imooc.com/video/1938)

#### 如何选择主键

+ 业务主键: 用于标识业务数据, 进行表与表之间的关联;
+ 数据库主键: 为了优化数据存储(Innodb 会生成6个字节的隐含主键)
    + [15.6.2.1 Clustered and Secondary Indexes](https://dev.mysql.com/doc/refman/8.0/en/innodb-index-types.html)

> If a table has no PRIMARY KEY or suitable UNIQUE index, InnoDB generates a hidden clustered index named GEN_CLUST_INDEX on a synthetic column that contains row ID values.
> The rows are ordered by the row ID that InnoDB assigns.
> The row ID is a 6-byte field that increases monotonically as new rows are inserted. Thus, the rows ordered by the row ID are physically in order of insertion.

+ 根据数据库类型, 考虑主键是否要顺序增长

+ 主键字段的类型所占空间要尽可能小: 对于使用聚集索引方式存储的表, 每个索引后都会附加主键信息
    + 主键小, 每页可加载更多数据, 提升查询效率
    + [极客时间: 04 | 深入浅出索引(上）](https://time.geekbang.org/column/article/69236)

#### 避免使用外键约束

+ 降低数据导入的效率

+ 增加维护成本

+ 虽然不建议使用外键约束, 但是相关联的列上一定要建立索引

#### 避免使用触发器

+ 降低数据导入的效率

+ 可能会出现意想不到的数据异常

+ 使业务逻辑变得复杂


#### 关于预留字段

+ 无法准确的知道预留字段的类型

+ 无法准确的知道预留字段中存储的内容

+ 后期维护预留字段索要的成本, 和增加一个字段所需要的成本是相同的

+ **严禁**使用预留字段
