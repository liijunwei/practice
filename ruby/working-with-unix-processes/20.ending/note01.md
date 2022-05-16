# [Ending](https://workingwithruby.com/wwup/ending/)

+ Working with processes in Unix is about two things:
    1. abstraction
    2. communication

## Abstraction

+ The kernel has an extremely abstract (and simple) view of its processes.

内核对其进程有一个极其抽象(且简单)的视图

+ We are masters of many programming languages, using each for different purposes. We couldn't possibly write memory-efficient code in a language with a garbage collector, we'll have to use C. But we need objects, let's use C++. On and on.

我们掌握了很多编程语言, 为了不同的目的, 我们使用不同的语言

我们很难用没有GC的语言写出内存效率很高的代码, 我们得用C语言写这种代码

我们得用面向对象抽象世界, 所以我们使用C++ ....

+ But if you ask the kernel it all looks the same.

但是对内核来说, 编程语言都是他的上层应用, 在它看来他们没有什么区别

> like Rick can't tell the difference between Morty and Summer in S02E01 08:40 lol

+ In the end, all of our code is compiled down to something simple that the kernel can understand.

最后, 所有的上层代码都会被翻译成内核能够理解的语言

+ Using Unix programming lets you do things that you can't accomplish when working at the programming language level

使用unix编程, 可以让你实现在编程语言层次无法实现或者难以实现的事情

+ Unix programming is programming language agnostic([æɡˈnɑːstɪk]).

unix编程和语言无关

+ It lets you interface your Ruby script with a C program, and vice versa.

他能让你的ruby脚本和C程序交互, 反之亦然

+ It also lets you reuse its concepts across programming languages.

他还能让你跨编程语言重用它的概念

+ The Unix Programming skills that you get from Ruby will be just as applicable in Python, or node.js, or C.

从ruby获得的unix编程技能同样适用于python/node/js/C...

+ These are skills that are about programming in general.

这些是关于编程的一般技能

## Communication

+ Besides the basic act of creating new processes, almost everything else we talked about was regarding communication.

+ Using signals any two processes on the system can communicate with each other.

+ By naming your processes you can communicate with any user who is inspecting your program on the command line.

+ Using exit codes you can send success/failure messages to any process that's looking after your own.

## Farewell, But Not Goodbye

> That's the end! Congratulations for making it here! **Believe it or not, you now know more than most programmers about the inner workings of Unix processes.**

hhhhh

+ Now that you know the fundamentals you can go out apply your newfound knowledge to anything that you work on.
+ Things are going to start making more sense for you.
+ And the more you apply your newfound knowledge: the clearer things will become.
+ There's no stopping you now.
