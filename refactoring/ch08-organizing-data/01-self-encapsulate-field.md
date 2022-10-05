You are accessing a field directly, but the coupling to the field is becoming awk- ward.

*Create getting and setting methods for the field and use only those to access the field.*

+ Essentially the advantages of *indirect variable access* are that it allows a subclass to override how to get that information with a method and that it supports more flexibility in managing the data, such as lazy initialization, which initializes the value only when you need to use it.

+ I'm always of two minds with this choice. I'm usually happy to do what the rest of the team wants to do.
    + and I prefer using `attr_reader`
