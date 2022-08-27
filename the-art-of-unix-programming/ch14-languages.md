+ 在现代条件下, 也许最好把C视作使用unix虚拟机的高级汇编器

+ 及时更高级的语言能够满足编程的要求, 我们仍然要学习C, 其中一个最充分的理由就是C能帮助我们学会在硬件体系层次上考虑问题

+ 对于已经是程序员的人来说, 学习C的最好参考和指导仍是 "The C Programming Language" K&R

+ C语言最好的案例分析是Unix内核本身

+ 高效的编译型语言; 对C的相相兼容; 面向对象的平台; STL和泛型等最前沿技术工具---C++试图满足所有人的所有要求, 但代价是C++比任何一个程序员所能处理的复杂度都要高

+ C++的最佳之处是编译效率以及面向对象和泛型编程的结合. 最糟之处是它非常怪异复杂, 往往鼓励过分复杂的设计

+ Shell的最佳之处在于书写小型脚本非常自然快捷. 最糟之处在于大型的shell脚本必须依靠大量辅助命令, 而这些辅助命令不一定在所有目标机器上都表现一致甚至不一定存在

+ Perl 是 增强了的shell, 它的主要缺点在于 某些部分丑陋到无法补救, 某些部分过于复杂, 某些部分必须谨慎的/一成不变的使用以防出错

+ Python的最佳之处在于他鼓励清晰/易读的代码, 易学易用, 又能扩展到大型项目; 最糟糕之处在于 不仅相对于编译语言, 而且相对于其他脚本语言, 他也是效率低下/速度缓慢的

+ Java编程语言的设计目标是: "write once, run anywhere", 并且支持网页中嵌入交互程序(applets), 即可在任何一个浏览器中运行, 但由于sun的一系列技术和战略事务, java没能实现最初的两个设计目标; 但它在系统编程和应用编程方面仍然十分强大, 足以挑战C和C++

+ Java设计非常聪明的抓住了自动管理内存的巨大优势, 也抓住了支持面型对象设计这一虽小确很重要的有点; 保留了大量的类C语法

+ Sun公司在网上提供良好的java文档的工作完成得非常出色(深有体会, 读java的api很舒服)

+ "过度OO分层的陷阱"

+ Emacs Lisp 是一种脚本语言, 用于emacs文本编辑器的 行为编程

+ Emacs Lisp 是一种lisp语言, 它自动管理内存; 它缺少标准的可移植OS规范, 这个老毛病已经由Emacs内核解决, 实际上Emacs内核就是其OS规范

+ Lisp另外有一个老毛病: 狂吃资源. 但在现代机器上也不再是个问题

+ "Emacs Makes A Compouter Slow"
+ "Eventually Munches All Computer Storage"

+ Emacs Lisp的最佳之处在于结合了非常优秀的基础语言Lisp, 其域原语对文本操作非常有效; 最糟之处在于性能较差, 难以和其他程序通讯

+ 当我们寻找方法, 把大程序分解成更简单的协作进程时, 在进程内使用线程应该是最后一招而不是第一招; 通常, 你可以发现避免使用线程是可能的; 如果能够实用有限的共享内存和信号量/使用SIGIO的一步IO, 或者poll(2)/select(2), 而不是使用线程, 就这样做吧, 保持简洁, 在本章所列的技法中优先选择位置更前/复杂度更低的方法.

+ 真实世界里的编程其实就是管理复杂度的问题; 能够管理复杂度的工具都是好东西. 但当这些工具的作用不是控制而是增加复杂度的时候, 最好扔掉, 从零开始; 永远不要忘记这一点, 它是Unix智慧的重要组成部分

