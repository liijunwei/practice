You have a two-way association but one class no longer needs features from the other.

*Drop the unneeded end of the association.*

+ Bidirectional associations are useful, but they carry a price.
    + The price is the added complexity of maintaining the two-way links and ensuring that objects are properly created and removed.
    + Bidirectional associations are not natural for many programmers, so they often are a source of errors.
    + Having many two-way links also makes it easy for mistakes to lead to zombies: objects that should be dead but still hang around because of a reference that was not cleared.
    + Bidirectional associations force an interdependency between the two classes.

