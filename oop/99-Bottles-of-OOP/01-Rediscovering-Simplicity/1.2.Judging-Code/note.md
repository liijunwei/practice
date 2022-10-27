+ You judge code constantly.

+ Writing code requires making choices; the choices you make reflect personal, internalized criteria.

+ You intend to write "good" code and if, in your estimation, you’ve written "bad" code, you are clearly aware that you’ve done so.

+ Regardless of how implicit, unachievable, or unhelpful they may be, you already have rules about code.

+ Answering the question "What makes code good?" thus requires defining goodness in concrete and actionable ways. This is harder than one might think.

+ **I like my code to be elegant and efficient; Clean code is full of crisp abstractions; Clean code was written by someone who cares.**

+ **Any pile of code can be made to work; good code not only works, but is also simple, understandable, expressive and changeable.**

+ The problem with these definitions is that although they accurately describe how good code looks once it’s written, **they give no help with achieving this state, and provide little guidance for choosing between competing solutions. The attributes they use to describe good code are qualitative, not quantitative.**

+ Metrics are fallible but human opinion is no more precise. Checking metrics regularly will keep you humble and improve your code.

代码的质量应该多用指标工具去检测, 不该仅仅靠感觉(感谢 flog/flay 这类工具)

+ As programmers grow, they get better at solving challenging problems, and become comfortable with complexity.
+ This higher level of comfort sometimes leads to the belief that complexity is inevitable, as if it’s the natural, inescapable state of all finished code.
    + 业务的复杂度 是代码复杂度的下限
+ However, there’s something beyond complexity --- a higher level of simplicity.
+ **Infinitely experienced programmers do not write infinitely complex code; they write code that’s blindingly simple.**

有些复杂性 是可以避免的

"更高层次的简洁" 是 "复杂" 下一阶段, 或者目标

有经验的程序员不会写出很复杂难懂的代码, 他们会把代码变得很简单, 清晰易懂, 同时能解决问题

(make it work, make it right, make it fast; don't stop at the intermediate state; https://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast)

+ **Shameless Green is defined as the solution which quickly reaches green while prioritizing understandability over changeability.**

"Shameless Green"是那种 能快速实现功能 并且平衡了易读性/易扩展(修改)性 的方案

+ It uses tests to drive comprehension, and patiently accumulates concrete examples while awaiting insight into underlying abstractions.

"Shameless Green"的解决方案使用测试驱动达到让代码易于理解的目的, 然后耐心的积累实际经验, 直到有一天认识到隐藏在代码背后的抽象

重构代码, 对于帮助理解业务很有帮助, 前提是 测试用例写得比较靠谱...

+ It doesn’t dispute that DRY is good, rather it believes that it is cheaper to manage temporary duplication than to recover from incorrect abstractions.

"Shameless Green"不怀疑"DRY"的价值, 但是它认为暂时的代码冗余是廉价的, 并且可以接受, 它认为这比太早做出不合适的抽象要好得多

+ Writing Shameless Green is fast, and the resulting code might be "good enough."
+ Most programmers find it embarrassingly duplicative, and the code is certainly not very object-oriented.
+ However, if nothing ever changes, the most cost-effective strategy is to deploy this code and walk away.

写出"Shameless Green"的代码不需要花很长时间, 很多精力, 就能达到足够好的效果

但是很多程序员觉得代码重复, 没有用"面向对象"的方法 有点儿可耻的

但是, 如果没有任何(需求)的变化, 那么 最具性价比的代码就是把代码部署好, 然后再也不动它了
(心理上的)

+ The challenge comes when a change request arrives. Code that’s good enough when nothing ever changes may well be code that’s not good enough when things do.

当需求变化时，挑战就来了。当什么都没有改变时就足够好的代码 很可能是 当改变发生时还不够好的代码。

