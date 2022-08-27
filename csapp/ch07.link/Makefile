clean:
	rm -rf a.out

main: clean
	gcc -Og main.c sum.c
	./a.out

verbose_main: clean
	gcc -Og -v main.c sum.c
	./a.out

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

setup_ar_demo:
	touch addvec.c
	touch multvec.c
	touch main2.c

ex_ar:
# 	sudo yum install glibc-static -y # if needed
	gcc -c addvec.c multvec.c
	ar rcs libvector.a addvec.o multvec.o
	gcc -c main2.c
	gcc -static -o a.out main2.o ./libvector.a
	./a.out
	objdump -f a.out

# tldr objdump

ex_ar_fail:
	gcc -c addvec.c multvec.c
	ar rcs libvector.a addvec.o multvec.o
	gcc -static ./libvector.a main2.c
	objdump -f a.out

make_libvector.so:
	gcc -shared -fpic -o libvector.so addvec.c multvec.c

ex_dlf: clean make_libvector.so
# 	gcc dll.c
	gcc -rdynamic -o a.out dll.c -ldl
	./a.out

ex_interpose_compile:
	gcc -DCOMPILETIME -c ./interpose/mymalloc-compile.c
	gcc -I. ./interpose/int.c mymalloc-compile.o
	./a.out

ex_interpose_link:


