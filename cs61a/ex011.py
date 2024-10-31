# https://pythontutor.com/render.html#mode=display

# python3 -m doctest ex011.py

# this is hard for me...
# counting partitions

def count_partitions(n, m):
  """
  >>> count_partitions(6, 4)
  9
  """
  if n == 0:
    return 1
  elif n < 0:
    return 0
  elif m == 0:
    return 0
  else:
    with_m = count_partitions(n-m, m)
    withoug_m = count_partitions(n, m-1)
    return with_m+withoug_m
