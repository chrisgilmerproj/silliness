#! /usr/bin/env python

from collections import defaultdict
import json
import os
from prettytable import PrettyTable


def print_gifts(gift_winners, gift_types):
    table = PrettyTable()
    table.field_names = ["Gift Num", "Person", "Present", "Previously"]
    for name in table.field_names:
        table.align[name] = "l"
    for key, value in dict(
        sorted(gift_winners.items(), key=lambda x: int(x[0]))
    ).items():
        gift_type = gift_types[key]
        if len(value) == 4:
            key = key + " (LOCKED)"
        row = [key, value[-1], gift_type, value[0:-1]]
        table.add_row(row)
    print(table)


def main():
    # Load any existing data
    gift_winners_file = "gift_winners.json"
    gift_types_file = "gift_types.json"

    gift_winners = defaultdict(list)
    if os.path.exists(gift_winners_file):
        with open(gift_winners_file, "r") as f:
            gift_winners = defaultdict(list, json.loads(f.read()))

    gift_types = {}
    if os.path.exists(gift_types_file):
        with open(gift_types_file, "r") as f:
            gift_types = json.loads(f.read())

    print_gifts(gift_winners, gift_types)

    while True:
        try:
            # Write out the data on each loop
            with open(gift_winners_file, "w") as f:
                f.write(json.dumps(gift_winners, indent=2))
            with open(gift_types_file, "w") as f:
                f.write(json.dumps(gift_types, indent=2, sort_keys=True))

            # Get the data for the game
            gift_num = input("\nEnter the gift number: ")
            if not gift_num.isdigit():
                print("Gift numbers must be numeric!")
                continue

            username = input("Enter the user name: ")

            if gift_num in gift_winners:
                print(
                    f"Gift already taken by {gift_winners[gift_num][-1]}, now it's for {username}"
                )
                if len(gift_winners[gift_num]) == 4:
                    print(f"Choose again {username}!")
                    continue
                else:
                    gift_winners[gift_num].append(username)

            else:
                gift_type = input("What is the gift type: ")
                gift_winners[gift_num].append(username)
                gift_types[gift_num] = gift_type
            print_gifts(gift_winners, gift_types)
        except KeyboardInterrupt:
            break


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        pass
