#! /usr/bin/env python3

from collections import Counter


def main():
    s = Counter(
        map(
            lambda w: f"{w[0]}{len(w)-2}{w[-1]}",
            map(
                str.strip,
                filter(
                    lambda x: len(x) > 1,
                    open("popular.txt"),
                ),
            ),
        )
    )

    for item in sorted(s.items(), key=lambda x: x[1], reverse=True)[:20]:
        print(item)


if __name__ == "__main__":
    main()
