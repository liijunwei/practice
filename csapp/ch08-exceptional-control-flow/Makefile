LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

csapp.o:
	$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

ecf_fork: csapp.o
	$(CC) $(CFLAGS) csapp.o ecf_fork.c
	./a.out

ecf_nested_fork: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_nested_fork.c
	./a.out

ecf_forkprob0: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_forkprob0.c
	./a.out

ecf_forkprob1: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_forkprob1.c
	./a.out

ecf_waitpid1: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_waitpid1.c
	./a.out

ecf_waitpid2: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_waitpid2.c
	./a.out

ecf_snooze: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_snooze.c
	./a.out

ecf_snooze_8_7: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_snooze_8_7.c
	./a.out

ecf_myecho: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_myecho.c
	./a.out a b c de

ecf_shellex_interactive: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_shellex.c
	# /usr/bin/which ls
	# /usr/bin/sleep 3
	./a.out

ecf_shellex_stdin: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_shellex.c
	echo "/usr/bin/ls" | ./a.out
	echo "/usr/bin/pwd" | ./a.out

ecf_kill: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_kill.c
	./a.out

ecf_sigint: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_sigint.c
	./a.out

ecf_sigint_safe: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_sigint_safe.c
	./a.out

ecf_signal1: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_signal1.c
	./a.out
	# ps aux|grep a.out

ecf_signal2: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_signal2.c
	./a.out
	# ps aux|grep a.out

ecf_signalprob0: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_signalprob0.c
	./a.out

ecf_procmask1: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_procmask1.c
	./a.out

ecf_procmask2: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_procmask2.c
	./a.out

ecf_waitforsignal: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_waitforsignal.c
	./a.out

ecf_sigsuspend: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_sigsuspend.c
	./a.out

ecf_setjmp: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_setjmp.c
	./a.out

# kill -9 $(ps aux|grep a.out | awk '{print $2}')
ecf_restart: csapp.o
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_restart.c
	./a.out

