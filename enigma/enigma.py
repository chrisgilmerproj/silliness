#! /usr/local/bin/python

"""
Modified from:
http://www.stealthcopter.com/blog/2011/05/recreating-the-enigma-in-python/
"""

import argparse
from copy import copy
from random import choice
from random import randint
from random import shuffle

alphabet = range(0, 26)


class Cog(object):
    """ Simple substitution cipher for each cog """

    def create(self):
        self.transformation = copy(alphabet)
        shuffle(self.transformation)
        return

    def passthrough(self, i):
        return self.transformation[i]

    def passthroughrev(self, i):
        return self.transformation.index(i)

    def rotate(self):
        self.transformation = self.shift(self.transformation, 1)

    def setcog(self, a):
        self.transformation = a

    def shift(self, l, n):
        """ Method to rotate arrays/cogs """
        return l[n:] + l[:n]


class Enigma(object):  # Enigma class
    def __init__(self, num_cogs, specialchars):
        self.specialchars = specialchars
        self.num_cogs = num_cogs
        self.cogs = []
        self.oCogs = []  # Create backup of original cog positions for reset

        for i in range(0, self.num_cogs):  # Create cogs
            self.cogs.append(Cog())
            self.cogs[i].create()
            self.oCogs.append(self.cogs[i].transformation)

        # Create reflector
        refabet = copy(alphabet)
        self.reflector = copy(alphabet)
        while len(refabet) > 0:
            a = choice(refabet)
            refabet.remove(a)
            b = choice(refabet)
            refabet.remove(b)
            self.reflector[a] = b
            self.reflector[b] = a

    def print_setup(self):
        """ To print the enigma setup for debugging/replication """
        print "Enigma Setup:\nCogs: ", self.num_cogs, "\nCog arrangement:"
        for i in range(0, self.num_cogs):
            print self.cogs[i].transformation
        print "Reflector arrangement:\n", self.reflector

    def reset(self):
        """ Reset all cogs to original position """
        for i in range(0, self.num_cogs):
            self.cogs[i].setcog(self.oCogs[i])

    def encode(self, text):
        """ Encode the text """
        ln = 0
        ciphertext = ""
        for l in text.lower():
            num = ord(l) % 97
            if (num > 25 or num < 0):
                if (self.specialchars):  # readability
                    ciphertext += l
                else:
                    pass  # security
            else:
                ln += 1
                for i in range(0, self.num_cogs):  # Move thru cogs forward...
                    num = self.cogs[i].passthrough(num)

                num = self.reflector[num]  # Pass thru reflector

                for i in range(0, self.num_cogs):  # Move back thru cogs...
                    num = self.cogs[self.num_cogs - i - 1].passthroughrev(num)
                # add encrypted letter to ciphertext
                ciphertext += "" + chr(97 + num)

                for i in range(0, self.num_cogs):  # Rotate cogs...
                    if (ln % ((i * 6) + 1) == 0):  # in a ticker clock style
                        self.cogs[i].rotate()
        return ciphertext


def main(plaintext, num_cogs, specialchars, display):
    """ The main cipher method """

    x = Enigma(num_cogs, specialchars)
    ciphertext = x.encode(plaintext)

    if display:
        x.print_setup()

        print "\nPlaintext:"
        print plaintext
        print "\nCiphertext:"

    print ciphertext

    if display:
        # To proove that encoding and decoding are symmetrical
        # we reset the enigma to starting conditions and enter
        # the ciphertext, and get out the plaintext
        x.reset()
        plaintext = x.encode(ciphertext)
        print "\nPlaintext:\n" + plaintext + "\n"


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Engima Encode/Decode')
    parser.add_argument('-f', '--file', dest='file',
                       help='Plain text file')
    parser.add_argument('-t', '--text', dest='text',
                       help='Plain text')
    parser.add_argument('-n', '--num', dest='num_cogs', type=int, default=4,
                       help='Number of cogs to use')
    parser.add_argument('-s', '--special', dest='special', action='store_true',
                       help='Keep special characters')
    parser.add_argument('-p', '--print', dest='display', action='store_true',
                       help='Display extra details')
    args = parser.parse_args()

    if args.file:
        plaintext = open(args.file, 'r').read().lower().replace('\n', ' ')
    elif args.text:
        plaintext = args.text
    else:
        plaintext = "The quick brown fox jumps over the lazy dog"

    main(plaintext, args.num_cogs, args.special, args.display)
