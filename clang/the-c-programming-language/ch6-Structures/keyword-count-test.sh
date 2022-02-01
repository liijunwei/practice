#!/bin/bash

gcc ch6-Structures/keyword-count-demo01.c
cat ch6-Structures/keyword-count-sample.md | ./a.out

echo
echo "pointer to struct"
gcc ch6-Structures/keyword-count-demo02.c
cat ch6-Structures/keyword-count-sample.md | ./a.out

