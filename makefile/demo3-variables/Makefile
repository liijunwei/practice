# https://makefiletutorial.com/#variables
# https://makefiletutorial.com/#variables-pt-2

# Variables can only be strings

files := file1 file2

some_file: $(files)
	echo "Look at some variable: " $(files)
	touch some_file

file1:
	touch file1

file2:
	touch file2

clean:
	rm -f file1 file2 some_file

# Recursive variable. This will print "later" below
one = one ${later_variable}
# Simply expanded variable. This will not print "later" below
two := two ${later_variable}

later_variable = later

all:
	echo $(one)
	echo $(two)
