+ How should performance considerations affect the design process?
+ This chapter discusses how to achieve high performance without sacrificing clean design.
+ **The most important idea is still simplicity: not only does simplicity improve a system’s design, but it usually makes systems faster.**

### How to think about performance

+ The first question to address is “how much should you worry about performance during the normal development process?”

+ The best approach is something between these extremes, where you use basic knowledge of performance to choose design alternatives that are “naturally efficient” yet also clean and simple. The key is to develop an awareness of which operations are fundamentally expensive

不要走极端(开发过程中的每一步都很关注性能 or 完全忽略性能)

应该培养一种意识，能快速识别出程序中哪些操作对性能损耗比较大

比如这些操作的开销是相对更加昂贵的：(Here are a few examples of operations that are relatively expensive today:)
+ 网络请求
+ 和二级存储相关的IO(磁盘/网盘/SSD等)
+ 动态内存分配(Dynamic memory allocation)
+ 缓存未命中(Cache misses)

+ The best way to learn which things are expensive is to run micro-benchmarks(small programs that measure the cost of a single operation in isolation).

+ Once you have a general sense for what is expensive and what is cheap, you can use that information to choose cheap operations whenever possible.

+ If the only way to improve efficiency is by adding complexity, then the choice is more difficult.

+ If the faster design adds a lot of implementation complexity, or if it results in more complicated interfaces, then it may be better to start off with the simpler approach and optimize later if performance turns out to be a problem.
如果一个为了性能考虑的设计引入了复杂性，让接口变得复杂，那么最好先用简单的方式来处理，直到能确定这部分设计导致了性能瓶颈再去考虑优化

+ However, if you have clear evidence that performance will be important in a particular situation, then you might as well implement the faster approach immediately.
当然，如果从一开始就清晰的知道某个部分的性能极其重要，那么还是要尽早的处理

+ In general, simpler code tends to run faster than complex code.

### Measure before modifying(profiling)

+ **It’s tempting to rush off and start making performance tweaks, based on your intuitions about what is slow. Don’t do this!**
+ *Programmers’ intuitions about performance are unreliable. This is true even for experienced developers.*
+ *If you start making changes based on intuition, you’ll waste time on things that don’t actually improve performance, and you’ll probably make the system more complicated in the process.*

+ Before making any changes, measure the system’s existing behavior. This serves two purposes:
    1. identify bottleneck
    2. provide a baseline

+ If the changes didn’t make a measurable difference in performance, then back them out (unless they made the system simpler). There’s no point in retaining complexity unless it provides a significant speedup.

### Design around the critical path




