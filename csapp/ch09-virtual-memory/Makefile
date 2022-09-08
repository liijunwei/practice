LDLIBS = -lpthread

CC = gcc
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

mmapcopy:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c mmapcopy.c
	./a.out ./Makefile
	./a.out ./mmapcopy.c

hard_code_array_size:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c hard_code_array_size.c
	echo 1 1   | ./a.out
	echo 2 1 2 | ./a.out
	echo 99999 | ./a.out

dynamic_array_size:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c dynamic_array_size.c
	echo 1 1   | ./a.out
	echo 2 1 2 | ./a.out
	# echo 99999 | ./a.out

