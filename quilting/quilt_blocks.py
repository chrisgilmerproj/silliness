#! /usr/bin/env python3

import random


def main():
    num_colors = 6
    colors = "123456789A"

    color_start = random.randint(0, num_colors)
    # Double the colors so we don't have to wrap
    colors *= 2
    block = []
    while len(block) != num_colors:
        new_color = random.choice(colors[color_start:color_start + num_colors])
        # Add as long as next color is not the same and there's not already two of them
        if len(block) == 0 or (block[-1] != new_color and block.count(new_color) < 2):
            block.append(new_color)
    print("".join(block))


if __name__ == "__main__":
    while True:
        main()
        input("Press enter to continue")
