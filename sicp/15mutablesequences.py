# https://www.bilibili.com/video/BV1nJ41157p6?t=3.2&p=19

total = 0
def count(func):
  def counted_func(*args):
    global total
    total += 1
    return func(*args)
  return counted_func

@count
def fact(n):
  assert n>=0
  if n<=1:
    return 1
  return n * fact(n-1)

@count
def fib(n):
  assert n>=0
  if n<=1:
    return n
  return fib(n-1)+fib(n-2)

@count
def fib_iter(n):
  assert n>=0
  curr, next, i = 0, 1, 0
  while i<n:
    curr, next = next, curr+next
    i+=1
  return curr

def reset():
  global total
  total = 0

data = []
data.append(("desc", "result", "recursive_iterations"))
data.append(("result(fact)", fact(30), total)); reset()
data.append(("result(fib)", fib(30), total)); reset()
data.append(("result(fib_iter)", fib_iter(30), total)); reset()

col_widths = [max(len(str(row[i])) for row in data) for i in range(len(data[0]))]

for row in data:
    print("  ".join(f"{str(cell):<{col_widths[i]}}" for i, cell in enumerate(row)))
