#! /usr/bin/env python3

import random


def main():
    colors = (
        "color 1",
        "color 2",
        "color 3",
        "color 4",
        "color 5",
        "color 6",
        "color 7",
        "color 8",
        "color 9",
        "color 10",
    )

    block = []
    while len(block) != 6:
        new_color = random.choice(colors)
        if len(block) == 0 or block[-1] != new_color:
            block.append(new_color)
    print(block)


if __name__ == "__main__":
    main()
