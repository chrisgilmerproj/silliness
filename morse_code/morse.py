#! /usr/local/bin/python

from math import ceil, log
from pprint import pprint

#FREQ_CHAR = 'abcdefghijklmnopqrstuvwxyz'
FREQ_CHAR = 'etaoinshrdlucmfwypvbgkqjxz'
MORSE = '._'


class GenerateMC(object):
    mc = []
    characters = ''
    code = ''
    dictionary = {}

    def __init__(self, characters, code):
        self.characters = characters
        self.code = code
        self.generate()

    def generate_symbols(self, prefix, remaining):
        if remaining == 0:
            self.mc.append(prefix)
            return
        for m in self.code:
            self.generate_symbols(prefix + m, remaining - 1)

    def generate(self):
        for y in range(int(ceil(log(len(self.characters))/log(2)))-1):
            self.generate_symbols('', y+1)
        self.dictionary = dict(zip(self.characters, self.mc))

    def encode(message_text):
        pass

    def decode(code_text):
        pass

def main():
    gmc = GenerateMC(FREQ_CHAR, MORSE)
    pprint(gmc.dictionary)

if __name__ == '__main__':
    main()
