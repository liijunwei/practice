# Spaces at the end of a line are not stripped, but those at the start are
# To make a variable with a single space, use $(nullstring)

with_spaces = hello   # with_spaces has many spaces after "hello"
after = $(with_spaces)there

nullstring =
space = $(nullstring) # Make a variable with a single space.
two_spaces = $(nullstring)  $(nullstring)

all:
	echo "$(after)"
	echo start"$(space)"end
	echo start"$(two_spaces)"end

# An undefined variable is actually an empty string!
undefined_var:
	echo $(nowhere)

# Use += to append
foo := start
foo += more
foo += and
foo += more
foo += on
foo += the
foo += way
foo += ...

another:
	echo $(foo)
