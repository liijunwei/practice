# https://jyywiki.cn/OS/2023/build/lect9.ipynb

import sys

BATCH_SIZE = 1000000

num_of_nest = int(sys.argv[1])
count = 0
checked = 0

while True:
    for ch in sys.stdin.read(BATCH_SIZE):
        match ch:
            case '(': count += 1
            case ')': count -= 1
            case _: assert 0
        assert 0 <= count <= num_of_nest
    checked += BATCH_SIZE
    print(f'{checked} OK')
