# Replace Temp with Chain

Change the methods to support chaining, thus removing the need for a temp.

```ruby
# before
mock = Mock.new
expectation = mock.expects(:a_method_name)
expectation.with("arguments")
expectation.returns([1, :array])

# after
mock = Mock.new
mock.expects(:a_method_name).with("arguments").returns([1, :array])
```

+ The solution utilizing Method Chaining **removes the need for the local variable**.
+ Method Chaining can also improve maintainability by providing an interface that allows you to compose code that **reads naturally**.

+ Replace Temp With Chain involves only one object.
+ It’s about improving the fluency of one object by allowing chaining of its method calls.
