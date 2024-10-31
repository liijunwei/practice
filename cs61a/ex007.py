# https://pythontutor.com/render.html#mode=display

# the Luhn Algorithm https://en.wikipedia.org/wiki/Luhn_algorithm
# used to verify credit card numbers

# TODO unclear yet

# python3 -i ex007.py
def split(n):
  """
  >>> split(1001)
  (100, 1)
  >>> split(82)
  (8, 2)
  """
  assert n >= 0
  return n // 10, n % 10

def sum_digits(n):
  """
  >>> sum_digits(0)
  0
  >>> sum_digits(10)
  1
  >>> sum_digits(123)
  6
  >>> sum_digits(9527)
  23
  """
  if n < 10:
    return n
  else:
    all_but_last, last = split(n)
    return sum_digits(all_but_last) + last

def luhn_sum(n):
  """
  >>> luhn_sum(2)
  2
  >>> luhn_sum(32)
  8
  """
  if n < 10:
    return n
  else:
    all_but_last, last = split(n)
    return luhn_sum_double(all_but_last) + last

def luhn_sum_double(n):
  all_but_last, last = split(n)
  luhn_digit = sum_digits(2 * last)
  if n < 10:
    return luhn_digit
  else:
    return luhn_sum(all_but_last) + luhn_digit
