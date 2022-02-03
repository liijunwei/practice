#!/bin/bash

git add .

gcc -o define-processor ch6-Structures/EX6-06.c
exitcode_compile=$?

echo "file: ch6-Structures/EX6-06.sample.c"
echo "==========================================="
cat ch6-Structures/EX6-06.sample.c | ./define-processor
exitcode_execute=$?
echo

rm define-processor

echo
echo "-----------------"
echo "compile result: $exitcode_compile"
echo "execute result: $exitcode_execute"
