#! /usr/bin/env python3

class bcolors(object):
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

remember = {}

for i in range(1,26):
    row = []
    for j in range(1, 26):
        val = i * j
        if i*j in remember:
            remember[val] += 1
        else:
            remember[val] = 0
        start = bcolors.ENDC
        end = bcolors.ENDC
        if remember[val] > 0:
            start = bcolors.OKBLUE
        row.append(start + str(val).rjust(3) + end)
    print(" ".join(row))
