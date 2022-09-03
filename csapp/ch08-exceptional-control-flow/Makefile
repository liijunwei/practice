LDLIBS = -lpthread

CC = gcc
CSAPP_SRC = ../csapp.c

PROGS = ecf_fork \
				ecf_fork \
				ecf_nested_fork \
				ecf_forkprob0 \
				ecf_forkprob1 \
				ecf_waitpid1 \
				ecf_waitpid2 \
				ecf_snooze \
				ecf_snooze_8_7 \
				ecf_myecho \
				ecf_shellex_interactive \
				ecf_shellex_stdin \
				ecf_kill \
				ecf_sigint

all: $(PROGS)

ecf_fork:
	$(CC) csapp.c ecf_fork.c $(LDLIBS)
	./a.out

ecf_nested_fork:
	$(CC) csapp.c ecf_nested_fork.c $(LDLIBS)
	./a.out

ecf_forkprob0:
	$(CC) csapp.c ecf_forkprob0.c $(LDLIBS)
	./a.out

ecf_forkprob1:
	$(CC) csapp.c ecf_forkprob1.c $(LDLIBS)
	./a.out

ecf_waitpid1:
	$(CC) csapp.c ecf_waitpid1.c $(LDLIBS)
	./a.out

ecf_waitpid2:
	$(CC) csapp.c ecf_waitpid2.c $(LDLIBS)
	./a.out

ecf_snooze:
	$(CC) csapp.c ecf_snooze.c $(LDLIBS)
	./a.out

ecf_snooze_8_7:
	$(CC) csapp.c ecf_snooze_8_7.c $(LDLIBS)
	./a.out

ecf_myecho:
	$(CC) csapp.c ecf_myecho.c $(LDLIBS)
	./a.out a b c de

ecf_shellex_interactive:
	$(CC) csapp.c ecf_shellex.c $(LDLIBS)
	# /usr/bin/which ls
	# /usr/bin/sleep 3
	./a.out

ecf_shellex_stdin:
	$(CC) csapp.c ecf_shellex.c $(LDLIBS)
	echo "/usr/bin/ls" | ./a.out
	echo "/usr/bin/pwd" | ./a.out

ecf_kill:
	$(CC) csapp.c ecf_kill.c $(LDLIBS)
	./a.out

ecf_sigint:
	$(CC) csapp.c ecf_sigint.c $(LDLIBS)
	./a.out

ecf_sigint_safe:
	$(CC) csapp.c ecf_sigint_safe.c $(LDLIBS)
	./a.out

ecf_signal1:
	$(CC) csapp.c ecf_signal1.c $(LDLIBS)
	./a.out
	# ps aux|grep a.out

ecf_signal2:
	$(CC) csapp.c ecf_signal2.c $(LDLIBS)
	./a.out
	# ps aux|grep a.out

