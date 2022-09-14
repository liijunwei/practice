LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

netp_ex_11_2: csapp.o
	$(CC) $(CFLAGS) csapp.o netp_ex_11_2.c
	./a.out 0x8002c2f2
