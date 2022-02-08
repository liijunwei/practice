#!/bin/bash

git add .

gcc ch7-Input-and-Output/EX7-07.c

echo "find main"
./a.out main ch7-Input-and-Output/EX7-07.c
echo

echo "find break"
./a.out break ch7-Input-and-Output/EX7-07.c
echo

echo "find fclose"
./a.out fclose ch7-Input-and-Output/EX7-07.c
echo
