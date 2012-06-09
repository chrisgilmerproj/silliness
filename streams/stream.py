#! /usr/local/bin/python

import random


class Stream(object):

    def generator(self):
        raise NotImplementedError

    def popNext(self):
        return self.generator().next()


class RandomStream(Stream):

    def generator(self):
        yield random.random()


class PrimeNumberStream(Stream):

    def generator(self):
        yield 1


class PrimeFactorStream(Stream):

    def generator(self):
        yield 1


def main():
    #s = RandomStream()
    s = PrimeNumberStream()
    #s = PrimeFactorStream()
    for i in xrange(5):
        print s.popNext()


if __name__ == "__main__":
    main()
