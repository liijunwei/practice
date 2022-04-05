+ This section will provide background on some key concepts used in the book.

这一节将介绍本书的背景和一些关键概念

+ It’s definitely recommended that you read this before moving on to the meatier chapters.

强烈建议先读这节, 再读其他

## Why Care?

为什么要学unix编程?

+ The Unix programming model has existed, in some form, since 1970.

unix编程1970年就存在了

+ These techniques transcend programming languages.

这些技巧超越了编程语言

+ This stuff has existed, largely unchanged, for decades. Smart programmers have been using Unix programming to solve tough problems with a multitude of programming languages for the last 40 years, and they will continue to do so for the next 40 years.

这些技巧, 有几十年没有发生过变化了

聪明的程序员已经用他们解决了很多困难的问题, 使用时长超过了40年

并且这些技巧很可能再有40年不会有变化

## Harness the Power!

驾驭力量

+ I’ll warn you now, the concepts and techniques described in this book can bring you great power.

这些概念和技巧将赋予你强大的力量

+ With this power you can create new software, understand complex software that is already out there, even use this knowledge to advance your career to the next level.

使用这份力量, 你可以创建新的软件, 理解已有的复杂软件, 甚至使用他们帮助你的职业进阶

+ Just remember, with great power comes great responsibility.

能力越大, 责任越大

+ Read on and I’ll tell you everything you need to know to gain the power and avoid the pitfalls.

继续读吧, 我会将你所需要的所有知识都传授给你, 并且帮你避开陷阱

## System Calls

+ To understand system calls first requires a quick explanation of the components of a Unix system, specifically userland vs. the kernel.

为了理解系统调用, 需要首先了解unix系统改动组成部分中的两个: 用户态(userland) 和 内核态(kernel)

+ The kernel of your Unix system sits atop the hardware of your computer.

kernel是电脑硬件上的一个软件

+ It’s a middleman for any interactions that need to happen with the hardware.

和硬件的交互需要先经过kernel这个中间人

+ This includes things like writing/reading from the filesystem, sending data over the network, allocating memory, or playing audio over the speakers.

和硬件的交互包含 文件系统的读写/通过网卡收发数据/内存分配/播放音频和视频等

+ Given its power, programs are not allowed direct access to the kernel. Any communication is done via system calls.

由于kernel功能很强大, 应用程序不能直接使用内核的功能; 因此所有和内核的交互, 都由系统调用完成

