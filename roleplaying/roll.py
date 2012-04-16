#! /usr/local/bin/python

import argparse
import random
import re
import sys

REXP_DICE = r"^(?P<number>\d?)[d|D]{1}(?P<dice>[0-9]+)(?P<oper>[+|-]{0,1})(?P<modifier>[\d]{0,2})"


def roll(dice):
    """ Roll all the dice """
    for roll in dice.split():
        match = re.match(REXP_DICE, roll)
        if match:
            group = match.groupdict()
            number = int(group['number']) if group['number'] else 1
            dice = group['dice']
            oper = group['oper']
            modifier = group['modifier']

            rolls = []
            for i in xrange(number):
                rolls.append(random.randint(1, int(dice)))
            if modifier:
                if oper == '-':
                    modifier = -1 * int(modifier)
                rolls.append(int(modifier))
            total = sum(rolls)
            print "\t%s: %s = %d" % (roll, ' + '.join([str(r) for r in rolls]), total)
        else:
            print "You've entered a bad dice roll: '%s'" % (args.roll)


if __name__ == "__main__":

    parser = argparse.ArgumentParser(description='Roll Dice')
    parser.add_argument('-r', '--roll', dest='roll', help='Roll a dice')
    args = parser.parse_args()
    if not args.roll:
        print "Enter a roll"
        sys.exit()
    roll(args.roll)
