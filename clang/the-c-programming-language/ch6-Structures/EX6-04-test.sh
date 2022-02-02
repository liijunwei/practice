#!/bin/bash

git add .

gcc -o word-freq ch6-Structures/EX6-04.c
exitcode_compile=$?

echo "file: ch6-Structures/EX6-04.sample.md"
echo "==========================================="
cat ch6-Structures/EX6-04.sample.md | ./word-freq
exitcode_execute=$?
echo

echo "file: ch6-Structures/EX6-04.c"
echo "==========================================="
cat ch6-Structures/EX6-04.c | ./word-freq
exitcode_execute=$?
echo

rm word-freq

echo
echo "-----------------"
echo "compile result: $exitcode_compile"
echo "execute result: $exitcode_execute"
