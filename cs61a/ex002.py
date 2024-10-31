# https://pythontutor.com/render.html#mode=display

# understand integer (floor) division
assert 23 // 10 == 2
assert 23 / 10 == 2.3
# understand mod operator
assert 23 % 10 == 3
assert 21 % 10 == 1
assert 20 % 10 == 0

# python3 -m doctest ex002.py
# python3 -i ex002.py
def remove(n, digit):
  """
  return all digits of non-negative N that are not DIGIT,
  for some non-negative DIGIT less than 10

  >>> remove(231, 3)
  21
  >>> remove(243132, 2)
  4313
  """
  assert n >= 0
  assert digit >= 0 and digit < 10

  kept, digits = 0, 0
  while n > 0:
    n, last = n // 10, n % 10
    if last != digit:
      kept = kept + last * 10 ** digits
      digits += 1
    print('n:', n, 'last:', last, 'digits:', digits, 'kept', kept, '10**digits', 10 ** digits)

  return kept

