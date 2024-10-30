# https://pythontutor.com/render.html#mode=display

# decorators

def trace(fn):
  def traced(x):
    print('calling', fn, 'on argument', x)
    return fn(x)
  return traced

# python3 -m doctest ex3.py
# python3 -i ex3.py
@trace
def square(x):
  return x * x

# square = trace(square)

@trace
def sum_squares_up_to(n):
  """
  >>> sum_squares_up_to(3)
  14
  >>> sum_squares_up_to(4)
  30
  """
  k = 1
  total = 0
  while k <= n:
    total, k = total+square(k), k+1
  return total

# sum_squares_up_to = trace(sum_squares_up_to)
