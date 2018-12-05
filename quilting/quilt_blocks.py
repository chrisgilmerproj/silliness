#! /usr/bin/env python3

import random


def main():
    colors = (
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
    )

    block = []
    while len(block) != 6:
        new_color = random.choice(colors)
        if len(block) == 0 or block[-1] != new_color:
            block.append(new_color)
    print(block)


if __name__ == "__main__":
    main()
