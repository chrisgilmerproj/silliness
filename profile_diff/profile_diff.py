#! /usr/local/bin/python

import argparse
import difflib
import filecmp
import os
from colors import red, yellow, green


def recursive_diff(dir1, dir2):
    """ Recursively return the differences between files in directories """

    diff = {'diff_files': []}

    # Do a file compare on each directory and print output
    cmp = filecmp.dircmp(dir1, dir2)

    # Print the difference between files that differ
    for file in cmp.diff_files:
        text1 = open(os.path.join(dir1, file), 'r').readlines()
        text2 = open(os.path.join(dir2, file), 'r').readlines()
        result = list(difflib.unified_diff(text1, text2))
        diff['diff_files'].append(((dir1, dir2), file, result))

    for common in cmp.common_dirs:
        rdiff = recursive_diff(os.path.join(dir1, common), os.path.join(dir2, common))
        diff['diff_files'] += rdiff['diff_files']

    return diff


def print_diff(diffs):
    """ Pretty print the diffs between files """

    for diff in diffs:
        print '\n' + '=' * 80 + '\n'

        dirs = diff[0]
        file = diff[1]

        print file
        print '---', os.path.join(dirs[0], file)
        print '+++', os.path.join(dirs[1], file)

        result = diff[2][2:]
        for line in result:
            line = line.strip()
            if len(line):
                if line[0] == '-':
                    line = green(line)
                elif line[0] == '+':
                    line = red(line)
                elif line[0] == '?':
                    line = yellow(line)
            print '\t%s' % line


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Recursively diff two directories')
    parser.add_argument('dirs', nargs=2,
                        help='Directory paths')
    
    args = parser.parse_args()

    dir1 = os.path.abspath(args.dirs[0])
    dir2 = os.path.abspath(args.dirs[1])

    rdiff = recursive_diff(dir1, dir2)
    diffs = sorted(rdiff['diff_files'])
    print_diff(diffs)
