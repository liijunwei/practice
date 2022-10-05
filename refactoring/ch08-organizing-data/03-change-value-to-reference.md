You have a class with many equal instances that you want to replace with a single object.

*Turn the object into a reference object.*

如果每个order里包含一个customer, 我们把customer重构成了Customer的实例

当有多个order的时候, 即使这多个order对应的customer都是同一个人, 在我们的代码中, customer还是会不停的被重新初始化(多个customer实例, 即使他们表示的是同一个人)

+ TODO question: what is "Replace Constructor with Factory Method." and why we ues it?

+ reference objects
    + things like customer or account
    + stands for one object in the real world
    + use the object iden- tity to test whether they are equal

+ value objects
    + things like date or money.(also can be reference object???)
    + They are defined entirely through their data values(unclear)
    + You don't mind that copies exist
