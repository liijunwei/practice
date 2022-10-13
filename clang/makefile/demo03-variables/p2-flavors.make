# https://makefiletutorial.com/#variables-pt-2

# There are two flavors of variables:
#     recursive (use =) - only looks for the variables when the command is used, not when it's defined.
#     simply expanded (use :=) - like normal imperative programming -- only those defined so far get expanded


# Recursive variable. This will print "later" below
one = one ${later_variable}
# Simply expanded variable. This will not print "later" below
two := two ${later_variable}

later_variable = later

all:
	echo $(one)
	echo $(two)

# ?= only sets variables if they have not yet been set

three = hello
three ?= will not be set
four ?= will be set

another:
	echo $(three)
	echo $(four)
