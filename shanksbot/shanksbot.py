#! /usr/bin/env python

#
# William Shanks Bot
#
# See https://youtu.be/DmfxIhmGPP4
#
# Useful for finding the length of the repeating digits for
# the inverse of a prime number.
#
# As an example:
#
# $ for num in 60013 60017 60029 60037 60041; do ./inverse.py $num; done
# 5001
# 60016
# 60028
# 10006
# 7505
#
# Worst case is that this runs O(N) times where N is the prime being used.
#

import sys


# Generator returns both the modulus and floor division of the number
# starting with the inverse of the number
def decimals(number):
    dividend = 1
    while dividend:
        yield dividend, dividend // number
        dividend = dividend % number * 10


def main():
    num = sys.argv[1]
    # Fast lookup of dividends
    table = {}
    # Count the number of decimal places
    inverse = []

    gen = decimals(int(num))
    for dividend, val in gen:
        if dividend in table and dividend > 0:
            break
        inverse.append(val)
        table[dividend] = True
    # Print the entire decimal up to the repeat point
    # print(
    #     "{}.{}".format(
    #         str(inverse[0]),
    #         "".join(map(str, inverse[1:])),
    #     )
    # )
    # subtract 1 for the zero (0) preceding the decimal place
    print(len(inverse[1:]))


if __name__ == "__main__":
    main()