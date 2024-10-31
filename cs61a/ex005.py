# https://pythontutor.com/render.html#mode=display

# python3 -m doctest ex005.py
def split(n):
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
    # print('all_but_last', all_but_last, 'last', last)
    return sum_digits(all_but_last) + last

def sum_digits_v2(n):
  """
  >>> sum_digits_v2(0)
  0
  >>> sum_digits_v2(10)
  1
  >>> sum_digits_v2(123)
  6
  >>> sum_digits_v2(9527)
  23
  """
  digits = [int(d) for d in str(n)]
  return sum(digits)
