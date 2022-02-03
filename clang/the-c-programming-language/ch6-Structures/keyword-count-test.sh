#!/bin/bash

gcc ch6-Structures/keyword-count-demo01.c
cat ch6-Structures/keyword-count-sample.md | ./a.out

echo
echo "pointer to struct"
gcc ch6-Structures/keyword-count-demo02.c
cat ch6-Structures/keyword-count-sample.md | ./a.out

echo
echo "word freq"
gcc ch6-Structures/keyword-count-demo03.c
echo "now is the time for all good men to come the aid of their party" | ./a.out

echo
echo "ch6-Structures/typedef-demo02.c"
gcc ch6-Structures/typedef-demo02.c
echo "now is the time for all good men to come the aid of their party" | ./a.out

