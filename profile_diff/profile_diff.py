#! /usr/local/bin/python

import difflib
import filecmp
import os


def recursive_diff(dir1, dir2, ext='txt'):

    print '\n' + '=' * 80
    print '\nComparing directories for *.%s files:\n\t%s\n\t%s' % (ext, dir1, dir2)
    # Do a file compare on each directory and print output
    cmp = filecmp.dircmp(dir1, dir2)
    print '\nFiles that are Identical:\n\t%s' % ('\n\t'.join(cmp.same_files) or 'None')
    print '\nFiles with Differences:\n\t%s' % ('\n\t'.join(cmp.diff_files) or 'None')
    print '\nFiles with no match:\n\t%s' % ('\n\t'.join(cmp.funny_files) or 'None')

    # Print the difference between files that differ
    print '\nFile Differences:'
    if cmp.diff_files:
        for file in cmp.diff_files:
            print '\n\t%s:' % (file)
            print '\t' + '@' * 72
            text1 = open(os.path.join(dir1, file), 'r').readlines()
            text2 = open(os.path.join(dir2, file), 'r').readlines()
            result = list(difflib.unified_diff(text1, text2))
            for line in result:
                print '\t%s' % line.strip()
            print '\t' + '@' * 72
    else:
        print '\tNone'

    for common in cmp.common_dirs:
        recursive_diff(os.path.join(dir1, common), os.path.join(dir2, common), ext=ext)


if __name__ == "__main__":
    extension = 'txt'
    dir1 = os.path.abspath('./files1/')
    dir2 = os.path.abspath('./files2/')
    recursive_diff(dir1, dir2, ext=extension)
