+ error handling is important, but if it obscures logic, it's wrong

+ Use exceptions rather than return codes

+ don't return `null`
    + when we return null, we are essentially creating work for ourselves and foisting problems upon our callers
    + if you're tempted to return null from a method, consider throwing an exception or returning a SPECIAL CASE object instead
    + if you're calling a null-returning method from a third-party API, consider wrapping that method with a method that either throws an exception or returns a special case object

+ don't pass `null`
    + returning null from method is bad, but passing null into methods is worse
    + unless you're working with an API which expects you to pass null, you should avoid passing null in your code whenever possible
