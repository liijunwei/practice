#!/bin/bash

# bash ch7-Input-and-Output/EX7-07-debug.sh "page 145"
# bash ch7-Input-and-Output/EX7-07-debug.sh Pointers
# bash ch7-Input-and-Output/EX7-07-debug.sh "修改第5章的模式查找程序"
# bash ch7-Input-and-Output/EX7-07-debug.sh "当发现一个匹配行时" # error here
# bash ch7-Input-and-Output/EX7-07-debug.sh "fixbug"
# bash ch7-Input-and-Output/EX7-07-debug.sh "stdio"
# bash ch7-Input-and-Output/EX7-07-debug.sh "MAXLINE"
# bash ch7-Input-and-Output/EX7-07-debug.sh main

keyword=$1

echo "expected"
# ag "$keyword" ch7-Input-and-Output/EX7-07.c
grep -nr "$keyword" ch7-Input-and-Output/EX7-07.c
echo

git add .

gcc ch7-Input-and-Output/EX7-07.c

echo "actual"
./a.out -n "$keyword" ch7-Input-and-Output/EX7-07.c
echo
