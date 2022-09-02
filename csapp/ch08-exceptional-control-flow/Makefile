LDLIBS = -lpthread

PROGS = ecf_fork \
				ecf_nested_fork \
				ecf_forkprob0 \
				ecf_waitpid1 \
				ecf_waitpid2

all: $(PROGS)

ecf_fork:
	gcc csapp.c ecf_fork.c $(LDLIBS)
	./a.out

ecf_nested_fork:
	gcc csapp.c ecf_nested_fork.c $(LDLIBS)
	./a.out

ecf_forkprob0:
	gcc csapp.c ecf_forkprob0.c $(LDLIBS)
	./a.out

ecf_forkprob1:
	gcc csapp.c ecf_forkprob1.c $(LDLIBS)
	./a.out

ecf_waitpid1:
	gcc csapp.c ecf_waitpid1.c $(LDLIBS)
	./a.out

ecf_waitpid2:
	gcc csapp.c ecf_waitpid2.c $(LDLIBS)
	./a.out

ecf_snooze:
	gcc csapp.c ecf_snooze.c $(LDLIBS)
	./a.out

ecf_myecho:
	gcc csapp.c ecf_myecho.c $(LDLIBS)
	./a.out a b c de

ecf_shellex_interactive:
	gcc csapp.c ecf_shellex.c $(LDLIBS)
	# /usr/bin/which ls
	# /usr/bin/sleep 3
	./a.out

ecf_shellex_stdin:
	gcc csapp.c ecf_shellex.c $(LDLIBS)
	echo "/usr/bin/ls" | ./a.out
	echo "/usr/bin/pwd" | ./a.out

ecf_kill:
	gcc csapp.c ecf_kill.c $(LDLIBS)
	./a.out

ecf_sigint:
	gcc csapp.c ecf_sigint.c $(LDLIBS)
	./a.out

