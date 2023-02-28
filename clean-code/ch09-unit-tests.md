+ once I got a suite of tests to pass, I would make sure that those tests were convenient to run for anyone else who needed to work with the code.
+ I would ensure that the tests and the code were checked in together into the same source package

+ The three lays of TDD:
    1. you may not write production code until you have written a failing unit test
    2. you may not write more of a unit test than is sufficient to fail, and not compiling is failing
    3. you may not write more production code than is sufficient to pass the currently failing test

+ "Drive"

+ "your code and test should go to the opposite direction: code -> more general, test -> more specific"

+ having dirty tests is equivalent to, if not worse than, haveing no tests
    + the problem is that tests must change as the production code evolves
    + the dirtier the tests, the harder they are to change

