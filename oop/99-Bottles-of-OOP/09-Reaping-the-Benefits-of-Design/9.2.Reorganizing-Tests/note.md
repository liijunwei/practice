# 9.2. Reorganizing Tests

+ Unit tests ought to tell an illuminating story.

单元测试应该把这个类做的事说清楚(故事)

+ They should demonstrate and confirm the class’s direct responsibilities, and do nothing else.

他们应该证明并确认这个类的直接指责, 不作任何其他多余的事

## 9.2.1. Gathering BottleVerse Tests

move tests from `BottlesTest` to `BottleVerseTest`

## 9.2.2. Revealing Intent

+ The problem is circular. You can’t write tests until you understand the problem, but you can’t understand the problem without writing some tests. No wonder it’s hard to get started.

+ Even so, here is where writing tests first really shines. It’s far better to struggle with a test that you don’t understand than to write code that you don’t understand.

写测试代码的闪光之处在于: 写出让人看不明白的测试代码 要比 写出让人看不明白的应用程序的代码要好得多

+ Tests force you to clarify your intentions because they make explicit assertions. Code has no such pressure, and can be left a confusing mess forever.

因为测试代码(用例)迫使你澄清你的意图, 因为每个测试用例里面都要做出明确的断言

应用程序的代码则没有这种压力, 他们即使意图不清, 难于理解, 可能也不会被关注到, 因为他们

