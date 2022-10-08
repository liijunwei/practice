A section of code assumes something about the state of the program.

*Make the assumption explicit with an assertion.*

+ Sometimes the assumptions are stated with a comment. A better technique is to make the assumption explicit by writing an assertion.

+ **An assertion is a conditional statement that is assumed to be always true.**

+ Indeed assertions usually are removed for production code.

+ **Assertions act as communication and debugging aids. In communication they help the reader understand the assumptions the code is making.**
+ In debugging, assertions can help catch bugs closer to their origin.
+ I've noticed the debugging help is less important when I write self-testing code, but I still appreciate the value of assertions in communication.

+ Because assertions should not affect the running of a system, adding one is always behavior preserving.

+ Beware of overusing assertions.
+ **Use assertions only to check things that need to be true.**

+ Assertions should be easily removable, so they don’t affect performance in production code.


+ what is that `assert` ? a ruby keyword ?
    + or is it like the guard clause, but instead of return, it raises errors ?
    + I know this keyword in java, not in ruby...
    + Just a method...
```ruby
module Assertions
  class AssertionFailedError < StandardError; end

  def assert(&condition)
    raise AssertionFailedError.new("Assertion Failed") unless condition.call
  end
end

class Employee
  include Assertions

  def expense_limit
    assert { (@expense_limit != NULL_EXPENSE) || (!@primary_project.nil?) }

    # ...
  end
end
```
