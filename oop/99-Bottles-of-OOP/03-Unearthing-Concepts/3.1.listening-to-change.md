+ However, in the real world, requirements do change, and when that happens, the standards for code rise.

真实世界不是一成不变的, 代码必须经历需求增加或需求改变的洗礼

+ This chapter defines a new requirement, which triggers a deeper look at the structure of the code.

本章引入了一个新的需求, 这个需求将带领我深入理解代码

+ The rules are simple, but they allow complex behavior to emerge(出现,产生,呈现).
+ By the end of this chapter, you’ll have begun to unearth(挖掘,发掘,发现) concepts that are currently hidden in the code.

学习并本章后, 你将发掘隐藏在代码里的概念

+ Code is expensive. Writing it costs time or money. It therefore behooves(理应,必须) you to be as efficient as possible.

代码不是免费的, 它甚至是昂贵的

所以代码理应越高效越好

aesthetics 美学

+ A good way to know that you’re using limited time wisely is to be driven by changes in requirements. The arrival of a new requirement tells you two things, one very specific, the other more general.

为满足需求里所做的改动, 能识别出你是否很好的利用了有限的时间

新需求会告诉你两件事: 一件具体, 一件抽象(通用)

+ Specifically, a new requirement tells you exactly how the code should change.
+ The requirement reveals exactly how you should have initially arranged the code.

具体的是, 需求会告诉你, 代码应该想什么方向演进, 应该修改什么地方

+ More generally, the need for change imposes higher standards on the affected code.

抽象(通用)的是, 需求的改变会对受影响的代码提出更高的标准

+ Code that needs to be changed must be changeable. Thus, a new requirement for the 99 Bottles problem will drive you to improve the code.

需要改动的代码, 必须是可变的, 可扩展的, 修改起来容易且风险低的

+ **Here’s that new requirement**: users have requested that you alter the 99 Bottles code to output "1 six-pack" in each place where it currently says "6 bottles."

需求: 把所有输出 "6 bottles." 的地方全部替换为 "1 six-pack"

+ 一种实现方式

```ruby
def verse(number)
  case number
  when 0
    "No more bottles of beer on the wall, " +
    # ...
  when 1
    "1 bottle of beer on the wall, " +
    # ...
  when 2
    "2 bottles of beer on the wall, " +
    # ...
  when 6
    "1 six-pack of beer on the wall, " +
    "1 six-pack of beer.\n" +
    "Take one down and pass it around, " +
    "5 bottles of beer on the wall.\n"
  when 7
    "7 bottles of beer on the wall, " +
    "7 bottles of beer.\n" +
    "Take one down and pass it around, " +
    "1 six-pack of beer on the wall.\n"
  else
    "#{number} bottles of beer on the wall, " +
    # ...
  end
end
```

+ The verse case statement initially contained four branches, and in the code above the number of branches has ballooned to six.

最初case的分支只有4个, 并且里面还有不少的代码冗余, 那时候还尚可接受

这种实现方式一下让分支膨胀到了6个

+ This is unacceptable. Conditionals breed, and now that this one has started reproducing, you must do something to stop it.

这种修改代码的方式不能让人满意, 需要想办法遏制这种趋势

