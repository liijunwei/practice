# https://jyywiki.cn/OS/2023/build/lect9.ipynb

import sys

batch_size = int(sys.argv[2])
nesting_size = int(sys.argv[1])
count = 0
checked = 0

while True:
    for ch in sys.stdin.read(batch_size):
        match ch:
            case '(': count += 1
            case ')': count -= 1
            case _: assert 0
        assert 0 <= count <= nesting_size
    checked += batch_size
    print(f'{checked} OK')
