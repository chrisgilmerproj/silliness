#! /usr/bin/env python

#
# Cabinet Project Nov 11, 2022
# Many cuts are similar so they need to be grouped for easy reference
#

from pprint import pprint
from typing import DefaultDict


def main():
    dimensions = {
        "v1": 53.5,
        "v2": 16.25,
        "v3": 37,
        "v4": 16.25,
        "v5": 35,
        "v6": 18.25,
        "v7": 22.25,
        "v8": 21.5,
        "v9": 22.25,
        "v10": 31,
        "v11": 22.25,
        "v12": 18.25,
        "v13": 35,
        "v14": 53.5,
        "v15": 53.5,
        "h1": 24.5,
        "h2": 43.25,
        "h3": 43.25,
        "h4": 68.75,
        "h5": 15.5,
        "h6": 34.25,
        "h7": 15.5,
        "h8": 44,
        "h9": 15.5,
        "h10": 18.5,
        "h11": 24.5,
        "h12": 24.5,
        "t1": 43.875,
        "t2": 53.25,
        "t3": 59.625,
    }

    sorted_dimensions = DefaultDict(list)

    for name, length in dimensions.items():
        sorted_dimensions[length].append(name)
    pprint(sorted_dimensions)

if __name__ == "__main__":
    main()