+ Objects hide their data behind abstractions and expose functions that operate on that data
+ Data structure expose theie data and have no meaningful functions(???)

+ TODO read page97 again("Data/Object Anti-Symmetry")

+ The Law of Demeter
    + train wrecks

+ Objects expose behavior and hide data. this makes it easy to add new kinds of objects without changing existing behaviors. it also makes it hard to add new behaviors to existing objects.
+ Data Structures expose data and have no significant behavior. this makes it easy to add new hebaviors to existing data structures but makes it hard to add new data strucures to existing functions

+ In any given system we will sometimes want the flexibility to add new data types, and so we prefer objects for that part of the system.
+ Other times we will want the flexibility to add new behaviors, and so in that part of the system we prefer data types and procedures.
+ Good software developers understand these issues without prejudice and choose the approach that is best for the job at hand.
