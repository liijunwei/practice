clean:
	rm -rf a.out

linkerror: clean
	gcc -Wall -Og linkerror.c

setup_foobar:
	touch foo{1..5}.c
	touch bar{1..5}.c

ex1:
	gcc foo1.c bar1.c
	# 	multiple definition of `main'

ex2:
	gcc foo2.c bar2.c
	# multiple definition of `x'

ex3:
	gcc foo3.c bar3.c
	./a.out
	# 	运行时,函数f将x的值由 15213 改为了 15212, "意外之喜"
	# 	注意, 链接器一般不会表明他检测到了多个x的定义

ex4:
	gcc foo4.c bar4.c

ex5:
	gcc foo5.c bar5.c


