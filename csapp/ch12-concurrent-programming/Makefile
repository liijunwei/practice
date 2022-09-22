LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

conc_badcnt: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_badcnt.c
	./a.out 1000000

