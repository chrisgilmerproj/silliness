#! /usr/local/bin/python

import difflib
import filecmp
import os
from colors import red, green


def recursive_diff(dir1, dir2, ext='txt'):
    """ Recursively return the differences between files in directories """

    diff = []

    # Do a file compare on each directory and print output
    cmp = filecmp.dircmp(dir1, dir2)

    # Print the difference between files that differ
    for file in cmp.diff_files:
        text1 = open(os.path.join(dir1, file), 'r').readlines()
        text2 = open(os.path.join(dir2, file), 'r').readlines()
        result = list(difflib.unified_diff(text1, text2))
        diff.append(((dir1, dir2), file, result))

    for common in cmp.common_dirs:
        diff += recursive_diff(os.path.join(dir1, common), os.path.join(dir2, common), ext=ext)

    return diff


def print_diff(diffs):
    """ Pretty print the diffs between files """

    for diff in diffs:
        print '\n' + '=' * 80 + '\n'

        dirs = diff[0]
        print '---', os.path.join(dirs[0], diff[1])
        print '+++', os.path.join(dirs[1], diff[1])

        result = diff[2][2:]
        for line in result:
            line = line.strip()
            if len(line):
                if line[0] == '-':
                    line = red(line)
                elif line[0] == '+':
                    line = green(line)
            print '\t%s' % line


if __name__ == "__main__":
    extension = 'txt'
    dir1 = os.path.abspath('./files1/')
    dir2 = os.path.abspath('./files2/')

    diffs = sorted(recursive_diff(dir1, dir2, ext=extension))
    print_diff(diffs)
