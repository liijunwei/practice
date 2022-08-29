LDLIBS = -lpthread

all: ecf_fork

ecf_fork:
	gcc csapp.c ecf_fork.c $(LDLIBS)
	./a.out

ecf_nested_fork:
	gcc csapp.c ecf_nested_fork.c $(LDLIBS)
	./a.out

ecf_forkprob0:
	gcc csapp.c ecf_forkprob0.c $(LDLIBS)
	./a.out
