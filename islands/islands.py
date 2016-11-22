'''
As input, you'll get a bitmap where 1s represent land, and 0s represent water.
Count the number of islands that exist in that bitmap, where an island is
defined as some continuously connected chunk of 1s.
For example

0 1 1 0
0 1 0 0
0 0 0 0 = 3
1 0 1 0

0 1 1 0
0 1 0 0
0 0 0 0 = 2
1 1 1 0
'''
import random

island3 = [
[0, 1, 1, 0],
[0, 1, 0, 0],
[0, 0, 0, 0],
[1, 0, 1, 0],
]

island2 = [
[0, 1, 1, 0],
[0, 1, 0, 0],
[0, 0, 0, 0],
[1, 1, 1, 0],
]

# more examples
island0 = [
[0, 0, 0, 0],
[0, 0, 0, 0],
[0, 0, 0, 0],
[0, 0, 0, 0],
]

island1 = [
[0, 0, 0, 0],
[0, 1, 0, 0],
[0, 0, 0, 0],
[0, 0, 0, 0],
]

island11 = [
[1, 1, 1, 1],
[1, 1, 1, 1],
[1, 1, 1, 1],
[1, 1, 1, 1],
]

island33 = [
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
[1, 1, 1, 0],
[1, 1, 0, 1],
[1, 0, 1, 1],
[0, 1, 1, 1],
]


# This island is too big to run!!

island_inf = []
for row in range(0, 1234123556):
    island_inf.append([])
    for col in range(0, 9807345):
        island_inf[row].append(random.choice([0, 1]))


def search(island, r, c):
    len_col = len(island[0])
    len_row = len(island)

    if (r - 1) >= 0 and island[r - 1][c] == 1:
        island[r - 1][c] = -1
        search(island, r - 1, c)
    if (r + 1) <= len_row - 1 and island[r + 1][c] == 1:
        island[r + 1][c] = -1
        search(island, r + 1, c)
    if (c - 1) >= 0 and island[r][c - 1] == 1:
        island[r][c - 1] = -1
        search(island, r, c - 1)
    if (c + 1) <= len_col - 1 and island[r][c + 1] == 1:
        island[r][c + 1] = -1
        search(island, r, c + 1)


def find_islands(island):
    num_islands = 0

    for r, row in enumerate(island):
        for c, col in enumerate(row):
            if island[r][c] == 1:
                num_islands += 1
                island[r][c] = -1
                search(island, r, c)

    return num_islands


assert find_islands(island3) == 3
assert find_islands(island2) == 2
assert find_islands(island0) == 0
assert find_islands(island1) == 1
assert find_islands(island11) == 1
assert find_islands(island33) == 9
