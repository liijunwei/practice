# https://pythontutor.com/render.html#mode=display

# python3 -m doctest ex008.py
def cascade_v1(n):
  """
  >>> cascade_v1(12345)
  12345
  1234
  123
  12
  1
  12
  123
  1234
  12345
  """
  if n < 10:
    print(n)
  else:
    print(n)
    cascade_v1(n//10)
    print(n)

def cascade_v2(n):
  """
  >>> cascade_v2(12345)
  12345
  1234
  123
  12
  1
  12
  123
  1234
  12345
  """
  print(n)
  if n > 10:
    cascade_v2(n//10)
    print(n)

# prefer v1 over v2

