#!/bin/bash

set -e

check() {
  expected="$1"
  # shift
  actual="$2"

  if [ "$expected" == "$actual" ]; then
    echo "ok"
  else
    echo "not ok"
    echo "==================="
    echo "expected: $expected"
    echo "actual:   $actual"
  fi
}

gcc -o old ch5-Pointers-and-Arrays/dcl-undecl-demo.c
gcc -o new ch5-Pointers-and-Arrays/EX5-19.c

# daytab是一个指针, 它指向一个有13个元素的数组
check 'int (*daytab)[13]' "$(echo 'daytab * [13] int' | ./old)"
check 'int (*daytab)[13]' "$(echo 'daytab * [13] int' | ./new)"

# daytab是一个有13个元素的数组指针, 数组里每个元素分别指向一个int
check 'int (*daytab[13])' "$(echo 'daytab [13] * int' | ./old)"
check 'int *daytab[13]' "$(echo 'daytab [13] * int' | ./new)"

check 'char (*x)' "$(echo 'x * char' | ./old)"
check 'char *x' "$(echo 'x * char' | ./new)"

rm -rf old
rm -rf new
