#!/bin/bash

git add .

gcc -o cross-referencer ch6-Structures/EX6-03.c

# cat ch6-Structures/EX6-03.c | ./cross-referencer
cat ch6-Structures/EX6-03.sample.md | ./cross-referencer
echo

rm cross-referencer
