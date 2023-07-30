# https://stackoverflow.com/questions/14226563/how-to-read-and-execute-gdb-commands-from-a-file
# gdb -x test1.gdb

layout next
layout next
break 8
run < input.txt # read input from file
