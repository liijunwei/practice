# https://www.bilibili.com/video/BV1nJ41157p6?t=1.9&p=18
# tree abstraction
# tree, branches, leaf
def tree(label, branches=[]):
  for branch in branches:
    assert is_tree(branch), 'branches must be trees'
  return [label]+list(branches)

def label(t):
  return t[0]

def branches(t):
  return t[1:]

def is_tree(t):
  """
  >>> demo = tree(3, [tree(1), tree(2, [tree(1), tree(1)])])
  >>> demo
  [3, [1], [2, [1], [1]]]
  >>> is_tree(demo)
  True
  """
  if type(t) != list or len(t) < 1:
    return False
  return all(is_tree(branch) for branch in branches(t))

def is_leaf(t):
  return not branches(t)

def count_nodes(t):
  """
  >>> demo = tree(3, [tree(1), tree(2, [tree(1), tree(1)])])
  >>> demo
  [3, [1], [2, [1], [1]]]
  >>> count_nodes(demo)
  5
  """
  lst = [count_nodes(b) for b in branches(t)]
  return sum(lst, 1)

def fib_tree(n):
  """
  >>> fib_tree(5)
  [5, [2, [1], [1, [0], [1]]], [3, [1, [0], [1]], [2, [1], [1, [0], [1]]]]]
  >>> print_tree(fib_tree(5))
  5
  ....2
  ........1
  ........1
  ............0
  ............1
  ....3
  ........1
  ............0
  ............1
  ........2
  ............1
  ............1
  ................0
  ................1
  """
  if n==0 or n==1:
    return tree(n)
  else:
    left, right = fib_tree(n-2), fib_tree(n-1)
    fib_n = label(left) + label(right)
    return tree(fib_n, [left, right])

def collect_leaves(t):
  """
  >>> ex2 = tree('D', [tree('B', [tree('A'), tree('C')]), tree('F', [tree('E'), tree('H', [tree('G'), tree('I')])])])
  >>> collect_leaves(ex2)
  ['A', 'C', 'E', 'G', 'I']
  """
  if is_leaf(t):
    return [label(t)]
  lst = [collect_leaves(b) for b in branches(t)]
  return sum(lst, [])

def print_tree(t, indent=0):
  if is_leaf(t):
    print('....'*indent+str(label(t)))
  else:
    print('....'*indent+str(label(t)))
    [print_tree(b, indent+1) for b in branches(t)]

ex3 = tree('D', [tree('B', [tree('A'), tree('C')]), tree('F', [tree('E', [tree('M'), tree('N')]), tree('H', [tree('G'), tree('I')])])])
# print_tree(ex3)

def print_calls(name, f):
  def new_f(t):
    print("Name: ", name)
    print("Inputted Tree: ")
    print_tree(t)
    # input()
    ret = f(t)
    print("Returned: ", ret)
    return ret
  return new_f

collect_leaves = print_calls('collect_leaves', collect_leaves)
# collect_leaves(ex3)

def square_tree(t):
  """
  >>> square_tree(tree(2))
  [4]
  >>> demo = tree(2, [tree(3, [tree(1), tree(1)])])
  >>> square_tree(demo)
  [4, [9, [1], [1]]]
  """
  if is_leaf(t):
    return tree(label(t)**2)
  lst = []
  for b in branches(t):
    lst += [square_tree(b)]
  return tree(label(t)**2, lst)
