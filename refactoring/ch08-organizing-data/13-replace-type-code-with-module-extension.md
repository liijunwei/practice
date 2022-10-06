You have a type code that affects the behavior of a class.

*Replace the type code with dynamic module extension.*

+ The one catch with module extension is that **modules cannot be unmixed easily**. Once they are mixed into an object, their behavior is hard to remove.
    + So use Replace Type Code with Module Extension when you don’t care about removing behavior. If you do care, use Replace Type Code with State/Strategy instead.
