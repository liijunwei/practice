You are raising an exception on a condition the caller could have checked first.

如果在本可以使用 条件判断 的地方, 使用异常处理来替代, 需要调整

*Change the caller to make the test first.*

+ 不要把异常处理机制 当做 控制流程 来使用

```ruby
# before
def execute(command)
  command.prepare rescue nil
  command.execute
end

# after
def execute(command)
  command.prepare if command.respond_to? :prepare
  command.execute
end
```

+ Like so many pleasures, exceptions can be used to excess(过度使用), and they **cease to be** pleasurable.

+ Exceptions should be used for exceptional behavior: behavior that is an unexpected error.
+ **They should not act as a substitute for conditional tests.**
+ If you can reasonably expect the caller to check the condition before calling the operation, you should provide a test, and the caller should use it.
