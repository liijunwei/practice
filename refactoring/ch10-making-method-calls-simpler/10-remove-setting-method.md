A field should be set at creation time and never altered.

*Remove any setting method for that field.*

+ Providing a setting method indicates that a field may be changed.
+ If you don't want that field to change once the object is created, don't provide a setting method.

