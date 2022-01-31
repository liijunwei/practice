#!/bin/bash

gcc ch6-Structures/keyword-count-demo0.c

cat ch6-Structures/keyword-count-demo0.c   | ./a.out
echo
cat ch6-Structures/keyword-count-sample.md | ./a.out | sort -r
