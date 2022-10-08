You have a conditional that chooses different behavior depending on the type of an object.

*Move each leg of the conditional to a method in an object that can be called polymorphically.*

+ similiar?
    + ../ch08-organizing-data/12-replace-type-code-with-polymorphism.md

+ Polymorphism gives you many advantages.
+ The biggest gain occurs when this same set of conditions appears in many places in the program.

+ TODO unclear(p280)

> If you want to add a new type, you have to find and update all the conditionals.
> But with polymorphism you just create a new class and provide the appropriate methods.
> Clients of the class don’t need to know about the polymorphism, which reduces the dependencies in your system and makes it easier to update.

+ **We already have our clients creating specific mountain bike objects**, and calling them polymorphically
    + don't worry about the object creation
