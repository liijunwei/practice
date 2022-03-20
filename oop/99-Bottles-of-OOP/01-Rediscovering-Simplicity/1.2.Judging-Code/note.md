+ You judge code constantly.

+ Writing code requires making choices; the choices you make reflect personal, internalized criteria.

+ You intend to write "good" code and if, in your estimation, you’ve written "bad" code, you are clearly aware that you’ve done so.

+ Regardless of how implicit, unachievable, or unhelpful they may be, you already have rules about code.

+ Answering the question "What makes code good?" thus requires defining goodness in concrete and actionable ways. This is harder than one might think.

> I like my code to be elegant and efficient.

> Clean code is full of crisp abstractions

> Clean code was written by someone who cares.

+ **Any pile of code can be made to work; good code not only works, but is also simple, understandable, expressive and changeable.**

+ The problem with these definitions is that although they accurately describe how good code looks once it’s written, **they give no help with achieving this state, and provide little guidance for choosing between competing solutions. The attributes they use to describe good code are qualitative, not quantitative.**

+ Metrics are fallible but human opinion is no more precise. Checking metrics regularly will keep you humble and improve your code.

代码的质量应该多用指标工具去检测

+ As programmers grow, they get better at solving challenging problems, and become comfortable with complexity.
+ This higher level of comfort sometimes leads to the belief that complexity is inevitable, as if it’s the natural, inescapable state of all finished code.
+ However, there’s something beyond complexity --- a higher level of simplicity.
+ **Infinitely experienced programmers do not write infinitely complex code; they write code that’s blindingly simple.**

复杂 是可以避免的

"更高层次的简洁" 是 "复杂" 下一阶段, 或者目标

有经验的程序员不会写出很复杂难懂的代码, 他们会把代码变得很简单, 清晰易懂, 同时能解决问题

+ **Shameless Green is defined as the solution which quickly reaches green while prioritizing understandability over changeability.**

"Shameless Green"的解决方案是那种 能快速实现解决方案, 并且平衡了易读性/易于改变特性的 的方案

+ It uses tests to drive comprehension, and patiently accumulates concrete examples while awaiting insight into underlying abstractions.

"Shameless Green"的解决方案使用测试驱动达到让代码易于理解的目的, 然后耐心的积累实际经验, 直到有一天认识到隐藏在代码背后的抽象(???)

+ It doesn’t dispute that DRY is good, rather it believes that it is cheaper to manage temporary duplication than to recover from incorrect abstractions.

"Shameless Green"不怀疑"DRY"的价值, 但是它认为暂时的代码冗余是廉价的, 并且可以接受, 它认为这比太早做出不合适的抽象要好得多




