#! /usr/local/bin/python

"""
A Persistent Difference

Pick any four digits from 0-9.  Create two numbers from them,
the maximum and the minimum (M - m).  Reorder the numbers and
continue until reaching the number 6174.
"""

import itertools
import string

num_to_final = {
    3: [495,],
    4: [6174,],
    5: [53955, 61974, 62964],
}

def get_n_list(numbers):
    n_list = []
    for n in numbers:
        n_list.append(n)
    n_list.sort()
    m = n_list
    M = n_list[::-1]
    return M, m


def run(numbers, verbose=False):
    orig = int(numbers)

    sub = 0
    count = 0
    try:
        final = num_to_final[len(numbers)]
    except:
        final = num_to_final[len(numbers) + 1]

    while sub not in final:
        M, m = get_n_list(numbers)
        sub = int("".join(M)) - int("".join(m))
        numbers = str(sub)

        if verbose:
            print "{0} - {1} = {2}".format(M, m, sub)

        if sub == 0:
            count = 0
            break
        else:
            count += 1

    if verbose:
        print "Required {0} steps".format(count)

    return orig, count        


def test(verbose=False):
    paths = {}
    for numbers in itertools.product(string.digits[1:], repeat=5):
        n = str("".join(numbers))
        orig, count = run(n, verbose=verbose)
        if count not in paths:
            paths[count] = []
        paths[count].append(orig)

    for count in paths:
        paths[count].sort()
        print count, len(paths[count])


def main(verbose=False):
    numbers = ''
    while not numbers.isdigit() or len(numbers) < 3:
        numbers = raw_input("Choose any four numbers, 0-9: ")
    orig, count = run(numbers, verbose=verbose)


if __name__ == "__main__":
    main(verbose=True)
    test(verbose=True)
