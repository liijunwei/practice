LDLIBS = -lpthread

CC = gcc
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS)

ecf_fork:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_fork.c
	./a.out

ecf_nested_fork:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_nested_fork.c
	./a.out

ecf_forkprob0:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_forkprob0.c
	./a.out

ecf_forkprob1:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_forkprob1.c
	./a.out

ecf_waitpid1:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_waitpid1.c
	./a.out

ecf_waitpid2:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_waitpid2.c
	./a.out

ecf_snooze:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_snooze.c
	./a.out

ecf_snooze_8_7:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_snooze_8_7.c
	./a.out

ecf_myecho:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_myecho.c
	./a.out a b c de

ecf_shellex_interactive:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_shellex.c
	# /usr/bin/which ls
	# /usr/bin/sleep 3
	./a.out

ecf_shellex_stdin:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_shellex.c
	echo "/usr/bin/ls" | ./a.out
	echo "/usr/bin/pwd" | ./a.out

ecf_kill:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_kill.c
	./a.out

ecf_sigint:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_sigint.c
	./a.out

ecf_sigint_safe:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_sigint_safe.c
	./a.out

ecf_signal1:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_signal1.c
	./a.out
	# ps aux|grep a.out

ecf_signal2:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_signal2.c
	./a.out
	# ps aux|grep a.out

ecf_signalprob0:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_signalprob0.c
	./a.out

ecf_procmask1:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_procmask1.c
	./a.out

ecf_procmask2:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_procmask2.c
	./a.out

ecf_waitforsignal:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_waitforsignal.c
	./a.out

ecf_sigsuspend:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_sigsuspend.c
	./a.out

ecf_setjmp:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_setjmp.c
	./a.out

# kill -9 $(ps aux|grep a.out | awk '{print $2}')
ecf_restart:
	$(CC) $(CFLAGS) $(CSAPP_SRC)/csapp.c ecf_restart.c
	./a.out

