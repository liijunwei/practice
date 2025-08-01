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
    lst.append(curr)
    index += 1
  return iter(lst)

# generator: an iterator created automatically by calling a generator function
# generator function: a function that contains the keyword yield anywhere in the body
# yielding values are the same as returning values, expect yield REMEMBERS where it left off

def plusminus(x):
  yield x
  yield -x

gen1 = plusminus(10)
next(gen1)
next(gen1)

def another_gen(x):
  x = x*x
  yield x
  x = x**2
  yield x

def gen_list(lst):
  yield lst

# def gen_list(lst):
#   [yield x for x in lst]

def naturals():
  x = 0
  while True:
    yield x
    x += 1

d = naturals()
[next(d) for _ in range(1,10)]

def gen():
  start = 0
  while start != 3:
    yield start
    start += 1

d = gen()

def a_then_b1(a, b):
  for x in a:
    yield x
  for x in b:
    yield x

d1 = a_then_b1([1,2,3], [4,5])

def a_then_b2(a, b):
  yield from a
  yield from b
d2 = a_then_b2([1,2,3], [4,5])

def countdown(k):
  if k==0:
    yield "blast off"
  else:
    yield k
    yield from countdown(k-1)
