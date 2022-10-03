# Replace Method with Method Object

+ question: what is "method object"?

background: You have a long method that uses local variables in such a way that you cannot apply Extract Method.

solution: **Turn the method into its own object so that all the local variables become instance variables on that object.** You can then decompose the method into other methods on the same object.


+ Using Replace Temp with Query helps to reduce this burden, but occasionally you may find you cannot break down a method that needs breaking
+ In this case you reach deep into the tool bag and get out your Method Object

## Mechanics

将这个方法提取成一个类的实例方法, 然后把这些局部变量转化为实例变量, 然后重构这个类
