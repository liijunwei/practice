# https://makefiletutorial.com/#multiple-targets

all: f1.o f2.o

# When there are multiple targets for a rule, the commands will be run for each target

# $@ is an automatic variable that contains the target name.
# 		https://makefiletutorial.com/#automatic-variables
# 		https://www.gnu.org/software/make/manual/html_node/Automatic-Variables.html
f1.o f2.o:
	echo $@

# Equivalent to:
# f1.o:
#	 echo f1.o
# f2.o:
#	 echo f2.o


hey: one two
	# Outputs "hey", since this is the target name
	echo $@

	# Outputs all prerequisites newer than the target
	echo $?

	# Outputs all prerequisites
	echo $^

	touch hey

one:
	touch one

two:
	touch two

clean:
	rm -f hey one two
