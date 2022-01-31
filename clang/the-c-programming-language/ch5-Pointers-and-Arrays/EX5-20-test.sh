#!/bin/bash

set -e

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

gcc ch5-Pointers-and-Arrays/EX5-20.c
echo

check 'echo "int (*daytab)[13]"    | ./a.out' 'daytab:  pointer to array[13] of int'
check 'echo "int *daytab[13]"      | ./a.out' 'daytab:  array[13] of pointer to int'
check 'echo "void *comp()"         | ./a.out' 'comp:  function returning pointer to void'
check 'echo "void (*comp)()"       | ./a.out' 'comp:  pointer to function returning void'
check 'echo "char (*(*x())[])()"   | ./a.out' 'x:  function returning pointer to array[] of pointer to function returning char'
check 'echo "char (*(*x[3])())[5]" | ./a.out' 'x:  array[3] of pointer to function returning pointer to array[5] of char'

check 'echo "void *(*comp)(int *, char *, int (*fnc)())" | ./a.out' 'comp:  pointer to function expecting pointer to int, pointer to char, pointer to function returning int and returning pointer to void'
# TODO
check 'echo "char const demo" | ./a.out' 'x:  array[3] of pointer to function returning pointer to array[5] of char'
