Put the method's body into the body of its callers and remove the method when A method's body is just as clear as its name.

+ Indirection can be helpful, but needless indirection is irritating.

+ Another time to use Inline Method is when you have a group of methods that seem badly factored.
    + You can inline them all into one big method and then re-extract the methods.
    + Kent Beck finds it is often good to do this before using Replace Method with Method Object.

+ **I commonly use Inline Method when someone is using too much indirection, and it seems that every method does simple delegation to another method, and I get lost in all the delegation.**

