LDLIBS = -lpthread

CC = gcc
CSAPP_SRC = ../csapp.c

mmapcopy:
	$(CC) csapp.c mmapcopy.c $(LDLIBS)
	./a.out ./Makefile
	./a.out ./mmapcopy.c

hard_code_array_size:
	$(CC) csapp.c hard_code_array_size.c $(LDLIBS)
	echo 1 1   | ./a.out
	echo 2 1 2 | ./a.out
	echo 99999 | ./a.out

dynamic_array_size:
	$(CC) csapp.c dynamic_array_size.c $(LDLIBS)
	echo 1 1   | ./a.out
	echo 2 1 2 | ./a.out
	# echo 99999 | ./a.out

