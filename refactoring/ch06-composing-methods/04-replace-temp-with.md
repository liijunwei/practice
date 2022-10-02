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
