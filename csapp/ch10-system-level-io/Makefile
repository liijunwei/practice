LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

io_ex_10_1: csapp.o
	$(CC) $(CFLAGS) csapp.o io_ex_10_1.c
	./a.out

io_cpstdin: csapp.o
	$(CC) $(CFLAGS) csapp.o io_cpstdin.c
	echo 12345 | ./a.out
	echo ningbotobeijing | ./a.out

io_statcheck: csapp.o
	$(CC) $(CFLAGS) csapp.o io_statcheck.c
	./a.out io_statcheck.c
	./a.out a.out
	./a.out csapp.o
	./a.out .
	./a.out ../
