+ [`Self Encapsulate Field`](https://martinfowler.com/bliki/SelfEncapsulation.html) explained
    + "Self encapsulaton means that getAmount needs to access both fields through getters."

+ `Replace Data Value with Object` allows you to turn dumb data into articulate objects.

+ If you see an Array or Hash acting as a data structure, you can make the data structure clearer with `Replace Array with Object` or `Replace Hash with Object`.

+ In all these cases the object is but the first step. The **real advantage** comes when you use `Move Method` to add behavior to the new objects.

+ "Links between objects can be one-way or two-way"
    + TODO question: what is link?
    + `Change Unidirectional Association to Bidirectional`
    + `Change Bidirectional Association to Unidirectional`
