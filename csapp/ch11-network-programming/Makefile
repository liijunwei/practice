LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

hex2dd: csapp.o
	$(CC) $(CFLAGS) csapp.o netp_ex_11_2.c
	./a.out 0x8002c2f2

dd2hex: csapp.o
	$(CC) $(CFLAGS) csapp.o netp_ex_11_3.c
	./a.out 128.2.194.242
