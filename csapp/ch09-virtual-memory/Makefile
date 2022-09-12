LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

mmapcopy: csapp.o
	$(CC) $(CFLAGS) csapp.o mmapcopy.c
	./a.out ./Makefile
	./a.out ./mmapcopy.c

hard_code_array_size: csapp.o
	$(CC) $(CFLAGS) csapp.o hard_code_array_size.c
	echo 1 1   | ./a.out
	echo 2 1 2 | ./a.out
	echo 99999 | ./a.out

dynamic_array_size: csapp.o
	$(CC) $(CFLAGS) csapp.o dynamic_array_size.c
	echo 1 1   | ./a.out
	echo 2 1 2 | ./a.out
	# echo 99999 | ./a.out

stackref:
	$(CC) $(CFLAGS) stackref.c
	./a.out

