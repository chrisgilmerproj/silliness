#! /usr/bin/env python3

from collections import defaultdict

def main():
    s = defaultdict(int)
    with open("popular.txt") as f:
        for line in f:
            w = line.strip()
            if len(w) < 3:
                continue
            s[f"{w[0]}{len(w[1:-2])}{w[-1]}"] += 1

    for item in sorted(zip(s.keys(), s.values()), key=lambda x: x[1], reverse=True)[0:20]:
        print(item)


if __name__ == "__main__":
    main()
