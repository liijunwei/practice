You have a class with many equal instances that you want to replace with a single object.

*Turn the object into a reference object.*

如果每个order里包含一个customer, 我们把customer重构成了Customer的实例

当有多个order的时候, 即使这多个order对应的customer都是同一个人, 在我们的代码中, customer还是会不停的被重新初始化(多个customer实例, 即使他们表示的是同一个人)

因此需要使用这个重构技巧, 当我们识别出这些order里对应的customer是同一个人的时候, 我们用一个customer的引用来表示这个客户, 不需要创建那么多customer的实例

> At the moment Customer is a value object.
> Each order has its own customer object even if they are for the same conceptual customer.
> I want to change this so that if we have several orders for the same conceptual customer, they share a single customer object.
> For this case this means that there should be only one customer object for each customer name.

+ TODO question: what is "Replace Constructor with Factory Method." and why we ues it?

+ reference objects
    + things like customer or account
    + stands for one object in the real world
    + use the object iden- tity to test whether they are equal

+ value objects
    + things like date or money.(also can be reference object???)
    + They are defined entirely through their data values(unclear)
    + You don't mind that copies exist
