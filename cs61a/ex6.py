# https://pythontutor.com/render.html#mode=display

# python3 -i ex6.py
def factorial(n):
  """
  >>> factorial(1)
  1
  >>> factorial(2)
  2
  >>> factorial(3)
  6
  >>> factorial(4)
  24
  >>> factorial(5)
  120
  """
  if n == 0:
    return 1
  else:
    return n * factorial(n-1)

factorial(3)
