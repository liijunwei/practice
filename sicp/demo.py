#!/usr/local/bin/python3

import sys
from math import pi,sqrt,pow

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

def area(r, shape_constant):
  assert r>0
  return r*r*shape_constant

def area_square(r):
  return area(r, 1)

def are_circle(r):
  return area(r, pi)

def are_hexagon(r):
  return area(r, 3*sqrt(3)/2)

print("area_square", area_square(1))
print("are_circle", are_circle(1))
print("are_hexagon", are_hexagon(1))

# python3 -m doctest demo.py
# generalizing over computational processes
def sum_naturals(n):
  """sum of the first N natural numbers
  >>> sum_naturals(5)
  15
  """
  total, k = 0, 1
  while k<=n:
    total, k = total+k, k+1
  return total

def sum_cubes(n):
  """
  >>> sum_cubes(5)
  225.0
  """
  total, k = 0, 1
  while k<=n:
    total, k = total+pow(k,3), k+1
  return total
