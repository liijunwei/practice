#!/usr/local/bin/python3

import sys
from math import pi,sqrt,pow

# exact representation as fractions
# a pair of integers
# as soon as division occurs, the exact representation may be lost
# assume we can compose and decompose rathional numbers:
  # [constructor] rational(n, d) returns a rational number x
  # [selector]    numer(x) return the numerator of x
  # [selector]    denom(x) return the denominator of x

# compound values combine other values together
  # date: year/month/day
  # geo position: latitude/longitude
  # ...
# data abstraction lets us manipulate compound values as units
# isolate two parts of any program that uses data:
  # how data are represented(as parts)
  # how data are manipulated(as units)
# data abstraction: a methodology by which functions enforce an abstraction barrier between representation and use

# using data
def mul_rational(x, y):
  return rational(
    numer(x)*numer(y),
    denom(x)*denom(y)
  )

def add_rational(x, y):
  nx, dx = numer(x), denom(x)
  ny, dy = numer(y), denom(y)
  return rational(nx*dy+ny*dx, dx*dy)

def print_rational(x):
  print(numer(x), '/', denom(x))

def rationals_are_equal(x, y):
  return numer(x) * denom(y) == numer(y) * denom(x)

# defining data
# constructor and selectors
def rational(n, d):
  """
  a representation of the rational number N/D
  """
  return [n, d]

def numer(x):
  assert len(x) == 2
  return x[0]

def denom(x):
  assert len(x) == 2
  return x[1]

print_rational(rational(10,20))
print_rational(add_rational(rational(1,4), rational(1,4)))
