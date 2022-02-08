#!/bin/bash

git add .

gcc ch7-Input-and-Output/EX7-07.c

echo "find main"
./a.out -n main ch7-Input-and-Output/EX7-07.c
echo

echo "find break"
./a.out -n break ch7-Input-and-Output/EX7-07.c
echo

echo "find fclose"
./a.out fclose ch7-Input-and-Output/EX7-07.c
echo

echo "find fclose"
./a.out -n fclose ch7-Input-and-Output/EX7-07.c
echo

echo "find page"
./a.out -n page ch7-Input-and-Output/EX7-07.c
echo

echo "find include"
./a.out -n include ch7-Input-and-Output/EX7-07.c
echo

echo "find MAXLINE"
./a.out -n MAXLINE ch7-Input-and-Output/EX7-07.c
echo
