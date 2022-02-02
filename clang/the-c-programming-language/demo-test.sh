#!/bin/bash

git add .

gcc -o demo demo.c
exitcode_compile=$?

./a.out
exitcode_execute=$?
echo

rm demo

echo
echo "-----------------"
echo "compile result: $exitcode_compile"
echo "execute result: $exitcode_execute"
