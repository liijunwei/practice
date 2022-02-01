#!/bin/bash

gcc -o wordsort ch6-Structures/EX6-02.c

echo "./wordsort(default 6)"
cat ch6-Structures/EX6-02.c | ./wordsort
echo

echo "./wordsort -3"
cat ch6-Structures/EX6-02.c | ./wordsort -3
echo

echo "./wordsort -4"
cat ch6-Structures/EX6-02.c | ./wordsort -4
echo

echo "./wordsort -8"
cat ch6-Structures/EX6-02.c | ./wordsort -8
echo

echo "./wordsort -9"
cat ch6-Structures/EX6-02.c | ./wordsort -9
echo

echo "./wordsort -10"
cat ch6-Structures/EX6-02.c | ./wordsort -10
echo

rm wordsort
