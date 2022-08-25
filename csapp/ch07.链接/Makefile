clean:
	rm -rf a.out

linkerror: clean
	gcc -Wall -Og linkerror.c

setup_foobar:
	touch foo{1..5}.c
	touch bar{1..5}.c

ex1: clean
	gcc foo1.c bar1.c
	# multiple definition of `main'

ex2: clean
	gcc foo2.c bar2.c
	# multiple definition of `x'

ex3: clean
	gcc foo3.c bar3.c
	./a.out
	# 运行时,函数f将x的值由 15213 改为了 15212, "意外之喜"
	# 注意, 链接器一般不会表明他检测到了多个x的定义

ex4: clean
	gcc foo4.c bar4.c
	./a.out
	# 如果有两个弱定义, 会发生和ex3同样的事情

ex5: clean
	# Linux practice 3.10.0-957.12.2.el7.x86_64 #1 SMP Tue May 14 21:24:32 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux
	gcc foo5.c bar5.c
	./a.out

