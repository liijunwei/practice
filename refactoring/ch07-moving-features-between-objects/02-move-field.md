A field is, or will be, used by another class more than the class on which it is defined.

*Create a new attribute reader (and if necessary, a writer) in the target class, and change all its users.*

+ **Moving state and behavior between classes is the essence of refactoring**
    + 随着系统的演进, 你需要新的类去实现新的功能, 不停的理清他们的职责
    + 也许某个设计之前看起来是合理的/正确的, 但是随着系统发生变化, 他可能不再合适
    + 这不是问题, 但是发现了这种问题缺什么也不做是个大问题

+ I consider moving a field if I see more methods on another class using the information in the field than the class itself.
    + I may choose to move the methods; this decision is based on **interface**.
    + But if the methods seem sensible where they are, I move the field.

+ `Self Encapsulate Field`
