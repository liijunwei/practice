#!/usr/local/bin/python3

import sys
from math import pi,sqrt

def f(x, y):
  return g(x)

def g(z):
  return z+x

x = 10
result = f(5, 10)
print(result)
# print(sys.argv)

square1 = lambda x: x * x
print(square1(4))
print(square1)

def square2():
  return x * x

print(square2)

a, b = lambda x: x*x, lambda y: y+y
print(a)
print(a(10))
print(b)
print(b(10))

# a guide to designing function
# 1. give each function exactly one job, but make it apply to many related situations
# DRY

# shift+enter

def area_square(r):
  assert r>0
  return r*r

def are_circle(r):
  assert r>0
  return r*r*pi

def are_hexagon(r):
  assert r>0
  return r*r*3*sqrt(3)/2

print("area_square", area_square(1))
print("are_circle", are_circle(1))
print("are_hexagon", are_hexagon(1))
