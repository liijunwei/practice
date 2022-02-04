#!/bin/bash

git add .

gcc -g ch7-Input-and-Output/EX7-02.c

echo "non-graphic"
echo 'this has a bell \a and a tab \t right here.' | ./a.out

echo "long"
echo 'this is a really long line with no breaks that really needs to wrap at least a little bit.' | ./a.out
