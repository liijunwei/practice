+ If you find the duplication distressing, consider the alternatives. Your choices are:
    + Assert that the expected output matches that of some other method.(???)
    + Assert that the expected output matches a dynamically generated string.
    + Assert that the expected output matches a hard-coded string.

+ Of these three choices, only the third is independent of the current implementation and so guaranteed to survive changes to Bottles .

上面的这三个选择中, 只有第三个是独立于Bottles的实现的, 并且它能在Bottles的实现改变时也能保持不变


