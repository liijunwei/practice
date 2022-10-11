Behavior on a superclass is relevant only for some of its subclasses.

*Move it to those subclasses.*

+ Use an accessor on the superclass to access any fields on the superclass.
+ Make the accessor **protected** if public access isn’t required.
