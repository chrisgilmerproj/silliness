#! /usr/bin/env python3

import string


def ascii_to_bin(s=''):
    return [bin(ord(item))[2:].zfill(8) for item in s]


def bin_to_ascii(data=None):
    return ''.join([chr(int(item, 2)) for item in data])


def invert_bin(data=None):
    inverted = []
    for item in data:
        new = ''.join(['1' if char == '0' else '0' for char in item])
        inverted.append(new)
    return inverted


def caesar(plaintext, shift):
    alphabet = string.ascii_lowercase
    shifted_alphabet = alphabet[shift:] + alphabet[:shift]
    table = str.maketrans(alphabet, shifted_alphabet)
    return plaintext.translate(table)


def caesar_all(plaintext):
    out = []
    for i in range(1, 27):
        out.append((i, caesar(plaintext, i)))
    return out
