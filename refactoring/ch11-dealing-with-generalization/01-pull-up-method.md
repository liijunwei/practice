You have methods with identical results on subclasses.

*Move them to the superclass.*

+ Eliminating duplicate behavior is important.
+ Whenever there is duplication, you face the risk that an alteration to one will not be made to the other.

+ If the methods look like they do the same thing but are not identical, use `Substitute Algorithm` on one of them to **make them identical**.
