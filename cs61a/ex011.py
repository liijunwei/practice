# https://pythontutor.com/render.html#mode=display

# python3 -m doctest ex011.py

# this is hard for me... (TODO)
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

from functools import lru_cache

@lru_cache(maxsize=None)
def partition(number):
  answer = set()
  answer.add((number, ))
  for x in range(1, number):
      for y in partition(number - x):
          answer.add(tuple(sorted((x, ) + y)))
  return answer

@lru_cache(maxsize=None)
def partition_with_max(number, max):
  answer = set()
  answer.add((number, ))
  for x in range(max, number):
      for y in partition_with_max(number - x, number-max):
          answer.add(tuple(sorted((x, ) + y)))
  return answer


def part(n, k):
  """
  >>> part(5, 1)
  [[5]]
  >>> part(5, 2)
  [[4, 1], [3, 2]]
  >>> part(5, 3)
  [[3, 1, 1], [2, 2, 1]]
  >>> part(5, 4)
  [[2, 1, 1, 1]]
  >>> part(5, 5)
  [[1, 1, 1, 1, 1]]
  >>> part(6, 3)
  [[4, 1, 1], [3, 2, 1], [2, 2, 2]]
  """
  def _part(n, k, pre):
    if n <= 0:
        return []
    if k == 1:
        if n <= pre:
            return [[n]]
        return []
    ret = []
    for i in range(min(pre, n), 0, -1):
        ret += [[i] + sub for sub in _part(n-i, k-1, i)]
    return ret
  return _part(n, k, n)

