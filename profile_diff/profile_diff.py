#! /usr/local/bin/python

import argparse
import difflib
import filecmp
import os
from colors import red, yellow, green


def recursive_diff(dir1, dir2):
    """ Recursively return the differences between files in directories """

    diff = {
        'diff_files': [],
        'funny_files': [],
        'left_only': [],
        'right_only': [],
    }

    # Do a file compare on each directory and print output
    cmp = filecmp.dircmp(dir1, dir2)

    # Print the difference between files that differ
    for file in cmp.diff_files:
        text1 = open(os.path.join(dir1, file), 'r').readlines()
        text2 = open(os.path.join(dir2, file), 'r').readlines()
        result = list(difflib.unified_diff(text1, text2))
        diff['diff_files'].append(((dir1, dir2), file, result))

    for file in cmp.funny_files:
        diff['funny_files'].append(file)

    for file in cmp.left_only:
        diff['left_only'].append(os.path.join(dir1,file))

    for file in cmp.right_only:
        diff['right_only'].append(os.path.join(dir1,file))

    for common in cmp.common_dirs:
        rdiff = recursive_diff(os.path.join(dir1, common), os.path.join(dir2, common))
        diff['diff_files'] += rdiff['diff_files']
        diff['funny_files'] += rdiff['funny_files']
        diff['left_only'] += rdiff['left_only']
        diff['right_only'] += rdiff['right_only']

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


def print_files(name, files, c):
    if files:
        print '\n' + '=' * 80 + '\n'
        print '%s Files:\n' % name
        for file in files:
            print c(file)


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

    print_files('Left Only', rdiff['left_only'], green)
    print_files('Right Only', rdiff['right_only'], red)
    print_files('Funny', rdiff['funny_files'], yellow)

    print '\n'
