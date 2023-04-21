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

### ch20.3 Design around the critical path

+ **The best way to improve its performance is with a “fundamental” change**, such as introducing a cache, or using a different algorithmic approach (balanced tree vs. list, for instance).
+ If you can identify a fundamental fix, then you can implement it using the design techniques discussed in previous chapters.

+ core issue: how to redesign an existing piece of code so that it runs faster.
    + This should be your last resort, and it shouldn’t happen often, but there are cases where it can make a big difference.
    + The key idea is to design the code around the critical path.
    不过这应该是你最后的手段，我们应该尽可能避免这种情况

+ Start off by asking yourself what is the smallest amount of code that must be executed to carry out the desired task in the common case.
+ consider only the data needed for the critical path, and assume whatever data structure is most convenient for the critical path.
+ Assume that you could completely redesign the system in order to minimize the code that must be executed for the critical path. Let’s call this code “the ideal.”

+ The next step is to look for a new design that comes as close as possible to the ideal while still having a clean structure.
+ "In my experience it’s almost always possible to find a design that is clean and simple, yet comes very close to the ideal."
+ One of the most important things that happens in this process is to remove special cases from the critical path.
+ When redesigning for performance, try to minimize the number of special cases you must check.
+ Performance isn’t as important for special cases, so you can structure the special- case code for simplicity rather than performance.

+ **The most important overall lesson from this chapter is that clean design and high performance are compatible.**

+ Complicated code tends to be slow because it does extraneous or redundant work.
+ On the other hand, if you write clean, simple code, your system will probably be fast enough that you don’t have to worry much about performance in the first place.

+ *In the few cases where you do need to optimize performance, the key is simplicity again: find the critical paths that are most important for performance and make them as simple as possible.*

