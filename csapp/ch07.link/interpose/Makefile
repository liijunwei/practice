ex_interpose:
	rm mymalloc.o
	gcc -DCOMPILETIME -c mymalloc.c
	gcc -I. int.c mymalloc.o
	./a.out
