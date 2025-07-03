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
print("area_circle", are_circle(1))
print("area_hexagon", are_hexagon(1))

# python3 -m doctest -v demo.py
# generalizing over computational processes

def summation(n, term):
  total, k = 0, 1
  while k<=n:
    total, k = total+term(k), k+1
  return total

def sum_naturals(n):
  """sum of the first N natural numbers
  >>> sum_naturals(5)
  15
  """
  return summation(n, lambda x: x)

def sum_cubes(n):
  """
  >>> sum_cubes(5)
  225.0
  """
  return summation(n, lambda x: pow(x, 3))

def make_adder(n):
  def adder(k):
    return k+n
  return adder

add_two = make_adder(2)
add_two(3)

"""
functions defined within other function bodies are bound to names in a local frame
"""

def compose1(f, g):
  def h(x):
    return f(g(x))
  return h

# make_adder(2)(3) => 5
# square(5) => 25
print(compose1(square1, make_adder(2))(3))

# returning a function using its own name
def print_sums(n):
  print(n)
  def next_sum(k):
    return print_sums(n+k)
  return next_sum

print(print_sums(1)(3)(5))

# TODO p06
