# Replace Temp with Query

You are using a temporary variable to hold the result of an expression.

Extract the expression into a method. Replace all references to the temp with the expression. The new method can then be used in other methods.

+ The problem with temps is that they are temporary and local.
+ Because they can be seen only in the context of the method in which they are used, temps tend to encourage longer methods, because that’s the only way you can reach the temp.

