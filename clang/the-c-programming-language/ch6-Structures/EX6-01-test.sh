#!/bin/bash

gcc -o wordcount ch6-Structures/EX6-01.c

cat ch6-Structures/EX6-01.c | ./wordcount
echo
cat ch6-Structures/keyword-count-sample.md | ./wordcount | sort -r

rm wordcount
