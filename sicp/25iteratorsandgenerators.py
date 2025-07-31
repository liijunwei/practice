# https://www.bilibili.com/video/BV1nJ41157p6?t=2.0&p=25

def dict_iter():
  d = {1: "one", 2: "two"}
  for i in d:
    print(i, d[i])

# 1. build the whole list
# 2. turn the list into an iterator
def fib_iter(n):
  prev, curr = 0, 1
  lst = [prev, curr]
  index = 2
  while index < n:
    prev, curr = curr, prev+curr
    lst += [curr]
    index += 1
  return iter(lst)
