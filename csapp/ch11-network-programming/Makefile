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
	./a.out 0x0
	./a.out 0xffffffff
	./a.out 0x7f000001

dd2hex: csapp.o
	$(CC) $(CFLAGS) csapp.o netp_ex_11_3.c
	./a.out 128.2.194.242
	./a.out 205.188.160.121
	./a.out 64.12.149.13
	./a.out 2-5.188.146.23
	./a.out 205.188.146.23

nslookup:
	nslookup baidu.com
	nslookup www.twitter.com
	nslookup twitter.com
	nslookup edu

services_port_map:
	cat /etc/services | less
