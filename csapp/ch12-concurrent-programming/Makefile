LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS) -Wall -g

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

conc_badcnt: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_badcnt.c
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000

conc_goodcnt: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_goodcnt.c
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000

conc_sharing: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_sharing.c
	./a.out

conc_psum-mutex: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_psum-mutex.c
	./a.out
