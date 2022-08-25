clean:
	rm -rf a.out linkerror

linkerror: clean
	gcc -Wall -Og -o linkerror linkerror.c

setup_foobar:
	touch foo{1..5}.c
	touch bar{1..5}.c

ex1:
	gcc foo1.c bar1.c
	# 	multiple definition of `main'

ex2:
	gcc foo2.c bar2.c

ex3:
	gcc foo3.c bar3.c

ex4:
	gcc foo4.c bar4.c

ex5:
	gcc foo5.c bar5.c


