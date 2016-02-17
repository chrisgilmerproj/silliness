'''

Imagine we have an image where every pixel is white or black. We’ll represent this
image as a simple 2D array (0 = black, 1 = white). The image you get is known to have a
single black rectangle on a white background. Your goal is to find this rectangle and
return its coordinates.

Here’s a sample “image” using JavaScript (feel free to rewrite in your language of
choice):
var image = [
[1, 1, 1, 1, 1, 1, 1],
[1, 1, 1, 1, 1, 1, 1],
[1, 1, 1, 0, 0, 0, 1],
[1, 1, 1, 0, 0, 0, 1],
[1, 1, 1, 1, 1, 1, 1]
];

'''


image = [
    [1, 1, 1, 1, 1, 1, 1],
    [1, 1, 1, 1, 1, 1, 1],
    [1, 1, 1, 0, 0, 0, 1],
    [1, 1, 1, 0, 0, 0, 1],
    [1, 1, 1, 1, 1, 1, 1],
]
image2 = [
    [1, 1, 1, 1, 1, 1, 1],
    [1, 1, 1, 1, 1, 1, 1],
    [1, 1, 1, 0, 0, 0, 1],
    [1, 1, 1, 1, 1, 1, 1],
    [1, 1, 1, 1, 1, 1, 1],
]

# 0 based counting left-down
solution = [
    [2, 3],
    [3, 5],
]
solution2 = [
    [2, 3],
    [2, 5],
]

# Find two things
# Most upper left 0
# Most downward right 0
def find_rect(image):
    solution = []
    upper_left = None
    down_right = None
    for i, row in enumerate(image):
        for j, col in enumerate(row):
            if col == 0:
                if upper_left == None:
                    upper_left = [i, j]
                    down_right = [i, j]
                else:
                    down_right = [i, j]
    return [upper_left, down_right]


def test():
    assert solution == find_rect(image)
    assert solution2 == find_rect(image2)
    

test()


'''

Now there are N solid black rectangles in the image. Find them all.

For example:
var image = [
[1, 1, 1, 1, 1, 1, 1],
[1, 1, 1, 1, 1, 1, 1],
[1, 1, 1, 0, 0, 0, 1],
[1, 0, 1, 0, 0, 0, 1],
[1, 0, 1, 1, 1, 1, 1],
[1, 0, 1, 0, 0, 1, 1],
[1, 1, 1, 0, 0, 1, 1],
[1, 1, 1, 1, 1, 1, 1],
];

'''

image = [
[1, 1, 1, 1, 1, 1, 1],
[1, 1, 1, 1, 1, 1, 1],
[1, 1, 1, 0, 0, 0, 1],
[1, 0, 1, 0, 0, 0, 1],
[1, 0, 1, 1, 1, 1, 1],
[1, 0, 1, 0, 0, 1, 1],
[1, 1, 1, 0, 0, 1, 1],
[1, 1, 1, 1, 1, 1, 1],
]

sq1 = [[2,3], [3,5]]
sq2 = [[3,1], [5,1]]
sq3 = [[5,3], [6,4]]

solution = [sq1, sq2, sq3]

def fill(left, right):
    """Find edges of rect"""
    pass


def in_sol(solutions, left, right):
    """Find if 0 in known solution"""
    for sol in solutions:
        upper_left, down_right = sol
        if upper_left[0] <= left <= down_right[0] and   upper_left[1] <= right <= down_right[1]:
            return True
    return False


def find_mult_rect(image):
    solutions = []
    upper_left = None
    down_right = None
    for i, row in enumerate(image):
        for j, col in enumerate(row):
            if col == 0:
                if not in_sol(solutions, i, j):
                    sol = fill(i, j)
                    solutions.append(sol)
    return solutions


def test_mult():
    assert sorted(solution) == sorted(find_mult_rect(image))

test_mult()
