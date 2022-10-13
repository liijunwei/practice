# https://makefiletutorial.com/#automatic-variables-and-wildcards

# I suggest that you always wrap it in the wildcard function, because otherwise you may fall into a common pitfall described below.

# Print out file information about every .c file
print: $(wildcard *.c)
	ls -lah $?

# thing_wrong := *.o # Don't do this! '*' will not get expanded
thing_right := $(wildcard *.o)

# all: one two three four
all: three four

# Fails, because $(thing_wrong) is the string "*.o"
# one: $(thing_wrong)

# Stays as *.o if there are no files that match this pattern :(
two: *.o

# Works as you would expect! In this case, it does nothing.
three: $(thing_right)

# Same as rule three
four: $(wildcard *.o)

create_fixtures:
	touch blah{1..3}.c

clean:
	rm blah*
