# 9.5. Summary

+ Good OO is built upon small, interchangeable objects that interact via abstractions.

+ The behavior of each individual object is often quite obvious, but the same cannot be said for the operation of the whole. Tests fill this breach.

+ TODO 还是区分不清楚 `message sender` 和 `message receiver` x.x
    + Senders   are responsible for knowing what they want(调用方法的实例, 在A类的方法里, 调用一个实例方法或类方法)
    + Receivers are responsible for knowing how to do it(具体的实现???)

+ Separating intention from implementation in this way allows you to introduce new variations without altering existing code
    + simply create a new object that responds to the original message with a different implementation.

+ Listen. **Fixing problems now is not only cheaper than fixing them later, but will improve your code, clarify your tests, and make glad your work.**
