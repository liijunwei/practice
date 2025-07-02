#!/usr/local/bin/python3

import sys

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
