#! /usr/local/bin/python

"""
2011/09/23

This is a program which takes a set of characters and generates a
set of morse code characters for them.  You can change the set of
characters or the encoding characers.  It's important that all of
the characters are unique.

Create the class and then you can encode or decode message text.
"""

from math import ceil
from math import log
from pprint import pprint

#FREQ_CHAR = u'abcdefghijklmnopqrstuvwxyz'  # Alphabet
FREQ_CHAR = u'etaoinshrdlucmfwypvbgkqjxz'  # Letter Freq. English
FREQ_CHAR = u'etianmsurwdkgohvflpjbxcyzq'  # Morse Code
MORSE = u'._'


class GenerateMC(object):
    mc = []
    characters = ''
    code = ''
    enc_dict = {' ': ' '}
    dec_dict = {' ': ' '}

    def __init__(self, characters, code):
        self.characters = characters.upper()
        self.code = code
        self.generate()

    def generate_symbols(self, prefix, remaining):
        if remaining == 0:
            self.mc.append(prefix)
            return
        for m in self.code:
            self.generate_symbols(prefix + m, remaining - 1)

    def generate(self):
        for y in range(int(ceil(log(len(self.characters)) / log(2))) - 1):
            self.generate_symbols('', y + 1)
        self.enc_dict.update(dict(zip(self.characters, self.mc)))
        self.dec_dict.update(dict(zip(self.mc, self.characters)))

    def encode(self, message_text):
        return ' '.join(map(lambda x: self.enc_dict[x], message_text.upper()))

    def decode(self, message_text):
        return ''.join(map(lambda x: self.dec_dict[x], message_text.split()))


def main():
    gmc = GenerateMC(FREQ_CHAR, MORSE)
    pprint(gmc.enc_dict)
    #pprint(gmc.dec_dict)

    message_text = 'This is awesome'
    message_enc = gmc.encode(message_text)
    message_dec = gmc.decode(message_enc)
    print message_enc, ' = ', message_dec


if __name__ == '__main__':
    main()
