#!/bin/bash

set -e

check() {
  expected=$1
  shift
  actual=$(eval "$@")

  if [ "$actual" == "$expected" ]; then
    echo "ok"
  else
    echo "not ok"
  fi
}

gcc -o old ch5-Pointers-and-Arrays/dcl-undecl-demo.c
gcc -o new ch5-Pointers-and-Arrays/EX5-19.c


# daytab是一个指针, 它指向一个有13个元素的数组
check 'int (*daytab)[13]' ./old << EOF
daytab * [13] int
EOF
check 'int (*daytab)[13]' ./new << EOF
daytab * [13] int
EOF

# daytab是一个有13个元素的数组指针, 数组里每个元素分别指向一个int
check 'int (*daytab[13])' ./old << EOF
daytab [13] * int
EOF
check 'int *daytab[13]' ./new << EOF
daytab [13] * int
EOF

check 'char (*x)' ./old << EOF
x * char
EOF
check 'char *x' ./new << EOF
x * char
EOF

rm -rf old
rm -rf new
