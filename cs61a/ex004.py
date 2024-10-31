# https://pythontutor.com/render.html#mode=display

# python3 -i ex004.py
def print_sums(n):
  print(n)
  def next_sum(k):
    return print_sums(n+k)
  return next_sum

foo = print_sums(1)(2)(3) # evaluates to a function that have state of 6
foo(4) # 10
