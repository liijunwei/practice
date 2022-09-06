# https://makefiletutorial.com/#command-line-arguments-and-override

# Overrides command line arguments
override option_one = did_override

# Does not override command line arguments
option_two = not_override
all:
	@echo $(option_one)
	@echo $(option_two)

# make -f command-line-arguments.make option_one=hi
