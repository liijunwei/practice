# https://pythontutor.com/render.html#mode=display

# python3 -m doctest ex010.py

def fib(n):
  """
  >>> fib(0)
  0
  >>> fib(1)
  1
  >>> fib(2)
  1
  >>> fib(3)
  2
  >>> fib(5)
  5
  >>> fib(30)
  832040
  """
  if n == 0:
    return 0
  elif n == 1:
    return 1
  else:
    return fib(n-2)+fib(n-1)

cache = {}

def cached_fib(n):
  """
  >>> cached_fib(0)
  0
  >>> cached_fib(1)
  1
  >>> cached_fib(2)
  1
  >>> cached_fib(3)
  2
  >>> cached_fib(5)
  5
  >>> cached_fib(30)
  832040
  >>> cached_fib(350)
  6254449428820551641549772190170184190608177514674331726439961915653414425
  """
  if n in cache:
    return cache[n]

  if n == 0:
    return 0
  elif n == 1:
    return 1
  else:
    cache[n] = cached_fib(n-2)+cached_fib(n-1)
    return cache[n]
