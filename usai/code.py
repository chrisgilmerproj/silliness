#! /usr/bin/env python3

from math import ceil
from math import log

# FREQ_CHAR = u'etaoinshrdlucmfwypvbgkqjxz'  # Letter Freq. English
FREQ_CHAR = u'etianmsurwdkgohvflpjbxcyzq'  # Morse Code
FREQ_CHAR = u'etianmsurwdkgohvfpljbxzqcy'  # Morse Code
MORSE = u'.-'


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

    cipher = [
        'ncmcvel',
        '282vImmcaIpgI',
        '3yesu',
        '435',
        'jlqpv',
        'ncmcvel',
        '282vImmcaIpgI',
        '3yesu',
        '435',
        'jlqpv',
        'ncmcvel',
        '282vImmcaIpgI',
        '3yesu',
        '435',
        'jlqpv',
        'ncmcvel',
        '282vImmcaIpgI',
        '3yesu',
    ]
    breaker = {
        '1': '',
        '2': 'i',
        '3': 'u',
        '4': 'o',
        '5': 't',
        '6': '',
        '7': '',
        '8': 'n',
        '9': '',
        '0': '',
        'a': 'g',
        'b': 'h',
        'c': 'i',
        'd': 'j',
        'e': 'a',
        'f': '',
        'g': 'c',
        'h': 'd',
        'I': 'e',
        'j': 'f',
        'k': '',
        'l': 'r',
        'm': 'l',
        'n': 'm',
        'o': '',
        'p': 'n',
        'q': 'o',
        'r': '',
        's': 'y',
        't': '',
        'u': 's',
        'v': 't',
        'w': '',
        'x': '',
        'y': 'w',
        'z': '',
    }

    phrase = []
    for word in cipher:
        new_word = []
        for c in word:
            new_word.append(breaker.get(c, ''))
        phrase.append(''.join(new_word))
    print(' '.join(phrase))
    print("=====")

    # military intelligence always out front military intelligence always out frontmilitary
    morse = [
        ['--', '..', '.-..', '..', '-', '.-', '.-.', '-.--'],
        ['..', '-.', '-', '.', '.-..', '.-..', '..', '--.', '.', '-.', '-.-.', '.'],
        ['.-', '.-..', '.--', '.-', '-.--', '...'],
        ['---', '..-', '-'],
        ['..-.', '.-.', '---', '-.', '-'],
        ['--', '..', '.-..', '..', '-', '.-', '.-.', '-.--'],
        ['..', '-.', '-', '.', '.-..', '.-..', '..', '--.', '.', '-.', '-.-.', '.'],
        ['.-', '.-..', '.--', '.-', '-.--', '...'],
        ['---', '..-', '-'],
        ['..-.', '.-.', '---', '-.', '-', '--', '..', '.-..', '..', '-', '.-', '.-.', '-.--'],
    ]

    gmc = GenerateMC(FREQ_CHAR, MORSE)
    phrase = []
    for m in morse:
        message_dec = gmc.decode(' '.join(m))
        phrase.append(message_dec)
    print(' '.join(phrase))
    print("=====")

    # Always Out Front
    binary = [
        '01000001', '01101100', '01110111', '01100001', '01111001', '01110011',
        '00100000', '01001111', '01110101', '01110100', '00100000', '01000110',
        '',         '01110010', '01101111', '01101110', '01110100', '',
    ]

    chars = []
    for val in binary:
        if len(val) == 0:
            continue
        elif len(val) != 8:
            raise Exception
        chars.append(chr(int(val, 2)))
    print(''.join(chars))


if __name__ == "__main__":
    main()
