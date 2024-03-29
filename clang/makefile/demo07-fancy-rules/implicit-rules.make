# https://makefiletutorial.com/#implicit-rules
#

# man make
# 		make -f implicit-rules.make
# 		make -f implicit-rules.make blah
# 		make -f implicit-rules.make clean

# Make loves c compilation. And every time it expresses its love, things get confusing.(agree)

# Compiling a C program: n.o is made automatically from n.c with a command of the form $(CC) -c $(CPPFLAGS) $(CFLAGS)
# Compiling a C++ program: n.o is made automatically from n.cc or n.cpp with a command of the form $(CXX) -c $(CPPFLAGS) $(CXXFLAGS)
# Linking a single object file: n is made automatically from n.o by running the command $(CC) $(LDFLAGS) n.o $(LOADLIBES) $(LDLIBS)


CC = gcc    # Flag for implicit rules
CFLAGS = -g # Flag for implicit rules. Turn on debug info

# Implicit rule #1: blah is built via the C linker implicit rule
# Implicit rule #2: blah.o is built via the C compilation implicit rule, because blah.c exists
blah: blah.o

blah.c:
	echo "int main() { return 0; }" > blah.c

clean:
	rm -f blah*
