#!/bin/bash

git add .

gcc -g ch7-Input-and-Output/EX7-01.c -o lower
echo "tolower"
echo "THE C PROGRAMING LANGUAGE" | ./lower
echo
rm lower

gcc -g ch7-Input-and-Output/EX7-01.c -o upper
echo "toupper"
echo "the c programing language" | ./upper
echo
rm upper

gcc -g ch7-Input-and-Output/EX7-01.c -o unsupported
echo "unsupported"
echo "the c programing language" | ./unsupported
echo
rm unsupported

