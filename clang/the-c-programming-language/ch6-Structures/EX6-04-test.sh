#!/bin/bash

git add .

gcc -o word-freq ch6-Structures/EX6-04.c

echo "file: ch6-Structures/EX6-03.sample.md"
echo "==========================================="
cat ch6-Structures/EX6-03.sample.md | ./word-freq
echo

rm word-freq
