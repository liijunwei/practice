You have an inheritance hierarchy, but never intend to explicitly instantiate an instance of the superclass.

*Replace the superclass with a module to better communicate your intention.*

+ **Writing intentional code is important**, and it would be nice if we could communicate that instances of our abstract superclass are not meant to be instantiated directly.

+ the `inherited` hook
