#!/bin/bash

git add .

gcc -o cross-referencer ch6-Structures/EX6-03.c

echo "file: ch6-Structures/EX6-03.c"
echo "==========================================="
cat ch6-Structures/EX6-03.c | ./cross-referencer
echo

echo "file: ch6-Structures/EX6-03.sample.md"
echo "==========================================="
cat ch6-Structures/EX6-03.sample.md | ./cross-referencer
echo

rm cross-referencer
