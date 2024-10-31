# https://pythontutor.com/render.html#mode=display

# python3 -m doctest ex009.py

def inverse_cascade(n):
  """
  >>> inverse_cascade(1234)
  1
  12
  123
  1234
  123
  12
  1
  """
  grow(n)
  print(n)
  shrink(n)

def f_then_g(f, g, n):
  if n != 0:
    f(n)
    g(n)

grow   = lambda n: f_then_g(grow, print, n//10)
shrink = lambda n: f_then_g(print, shrink, n//10)
