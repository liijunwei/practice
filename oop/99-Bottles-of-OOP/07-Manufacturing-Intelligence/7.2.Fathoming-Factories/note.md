+ Fathoming(理解/彻底了解/弄清真相)

+ From the message sender’s point of view, all players of a role are exactly the same.

从消息发送者的角度看, 应用里多态相关的实例都是一样的

+ They know what their collaborators do, but refuse to be aware of how they do it.

消息发送者知道实例会做什么, 但是不需要知道他们是怎么做的, 不需要知道内部的实现

+ Message senders aren’t allowed to know the names of the concrete variant classes, nor may they know the logic needed to choose between them.

消息发送这不能知道使用多态的类的变体的名字, 他们也不需要知道怎么在这些变体里做出选择

+ Message senders can’t know these things, but of course someone must.

消息发送者不(需要/该)知道这些事, 但是代码里肯定得有地方维护这个知识

+ Knowledge of the class names of the variants, and of the logic necessary to choose the correct one, can be hidden in, you guessed it --- factories.

这些知识应该被隐藏在工厂里

+ A factory’s responsibility is to manufacture the right object for a given role.

工厂的职责就是为给定的角色创建合适的实例

他们不需要知道这些实例怎么实现功能, 他们只需要知道怎么为特定的角色创建合适的实例

这种选择往往需要条件判断, 把判断放到工厂里面能让你把这些分支判断隔离到一个地方, 统一管理

