gcc ch5-Pointers-and-Arrays/dcl-demo.c

echo "char **argv"       | ./a.out
echo "int (*daytab)[13]" | ./a.out
echo "int *daytab[13]"   | ./a.out
echo "void *comp()"      | ./a.out
echo "void (*comp)()"    | ./a.out
echo "char (*(*x())[])()"   | ./a.out
echo "char (*(*x[3])())[5]" | ./a.out
