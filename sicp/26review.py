# https://www.bilibili.com/video/BV1nJ41157p6?t=45.9&p=26

# python3 -m doctest 26review.py
def count_group(n):
  """
  >>> count_group(1)
  1
  >>> count_group(2)
  1
  >>> count_group(3)
  2
  >>> count_group(4)
  5
  >>> count_group(5)
  14
  """
  if n==1:
    return 1
  if n==2:
    return 1
  total = 0
  for i in range(1, n):
    total += count_group(i) * count_group(n-i)
  return total
