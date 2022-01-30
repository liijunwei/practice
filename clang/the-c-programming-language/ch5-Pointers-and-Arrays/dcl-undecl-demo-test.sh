#!/bin/bash

check() {
  command=$1
  actual=$(eval $1)
  expected=$2

  if [ "$actual" == "$expected" ]; then
    echo "$command # ok"
  else
    echo "$command # not ok"
  fi
}

gcc ch5-Pointers-and-Arrays/dcl-undecl-demo.c

check 'echo "x () * [] * () char"    | ./a.out' 'char (*(*x())[])()'
