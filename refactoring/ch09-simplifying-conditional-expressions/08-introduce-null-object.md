You have repeated checks for a nil value.

*Replace the nil value with a null object.*

+ null object pattern

+ The essence of polymorphism is that **instead of asking an object what type it is and then invoking some behavior based on the answer, you just invoke the behavior**.

+ An interesting characteristic of using null objects is that **things almost never blow up**. Because the null object responds to all the same messages as a real one, the system generally behaves normally.
+ This can sometimes make it difficult to detect or find a problem, because nothing ever breaks.
+ Of course, as soon as you begin inspecting the objects, you'll find the null object somewhere where it shouldn't be.(unclear)

+ **Remember, null objects are always constant; nothing about them ever changes.**
    + Apply `Singleton pattern`

+ Ruby allows us **two main options** for implementing the null object.
    + 1. subclass
    + 2. duck type
    + add methods to `NilClass` itself(dangerous)

+ consider whether to implement a `message-eating null object`
    + A message-eating null will accept any message sent to it and return another message-eating null object.
