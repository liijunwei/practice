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

# TODO 为什么有错误?
# time ./a.out 1 32
# Error: result=9223372034707292160
# ./a.out 1 32  70.94s user 0.05s system 99% cpu 1:11.04 total
conc_psum-mutex: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_psum-mutex.c
	./a.out
# 	./a.out 1 32

conc_psum-array: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_psum-array.c
	./a.out
# 	./a.out 1 32

.PH
clean:
	rm csapp.o
