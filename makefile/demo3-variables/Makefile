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
