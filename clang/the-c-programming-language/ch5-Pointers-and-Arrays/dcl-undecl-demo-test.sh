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

gcc ch5-Pointers-and-Arrays/dcl-undecl-demo.c

# x is a function returning a pointer to an array of pointers to functions returning char
# <=>
# x () * [] * () char
check 'char (*(*x())[])()' ./a.out << EOF
x () * [] * () char
EOF

check 'char (*x)' ./a.out << EOF
x * char
EOF
