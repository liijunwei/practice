A method returns a collection.

*Make it return a copy of the collection and provide add/remove methods.*

+ The attribute reader should not return the collection object itself, because that allows clients to manipulate the contents of the collection without the owning class knowing what is going on.

类暴露的属性读取方法不应该把 集合数据本身 暴露出去(应该暴露一个副本) (unclear)

因为如果把 集合数据本身 暴露出去, 那么客户端可能会修改这个数据, 数据的所属方可能感知不到这个修改

+ It also reveals too much to clients about the object's internal data structures.

+ An attribute reader for a multivalued attribute should return something that prevents manipulation of the collection and hides unnecessary details about its structure.

`attribute reader` 应该返回能防止客户端修改数据的 集合数据, 并且隐藏掉不必要的细节

+ In addition there should not be an attribute writer for the collection: rather, there should be operations to add and remove elements.

除此以外, 类还不该暴露属性的 writer 方法, 应该暴露用有一定限制的接口去让客户端修改数据(更加可控)

+ A few years ago I was concerned that moving this kind of behavior over to Person would lead to a bloated Person class. In practice, I’ve found that this usually isn’t a problem.

