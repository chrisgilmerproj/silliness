#! /usr/local/bin/python

import math
import random


def is_prime(num):
    """
    A simple prime number method using the Sieve of Eratosthenes
    """
    if type(num) != int:
        return False
    if num == 2:
        return True
    if num < 2 or num % 2 == 0:
        return False
    return not any(num % i == 0 for i in range(3, int(math.sqrt(num)) + 1, 2))


class Stream(object):
    """
    A lazy Stream class that returns values from a generator using
    the method popNext().  Once the generator is complete it will
    return the value None
    """

    def generator(self):
        """
        Each subclass must implement this method
        """
        raise NotImplementedError

    def popNext(self):
        """
        Return the next item lazily from the stream or return None
        """
        try:
            return self.generator().next()
        except StopIteration:
            return None

    def __iter__(self):
        return self.generator()


class RandomStream(Stream):
    """
    Stream X random numbers where the value X is given by num
    """

    def __init__(self, num):
        self.num = num
        self.counter = 0

    def generator(self):
        while self.counter < self.num:
            self.counter += 1
            yield random.random()


class PrimeNumberStream(Stream):
    """
    Stream prime numbers up to the specified number
    """
 
    def __init__(self, num):
        self.num = num
        self.counter = 0 

    def generator(self):
        while self.counter < self.num:
            self.counter += 1
            if is_prime(self.counter):
                yield self.counter


class PrimeFactorStream(Stream):
    """
    Stream prime factors of a specified number
    """

    def __init__(self, num):
        self.num = num
        self.counter = 2
        self.last = False
        self.powers = []
        self.limit = (num / 2) + 1

    def generator(self):
        while self.counter <= self.limit:
            while self.num % self.counter == 0:
                # Do not return repeats
                if not self.powers.__contains__(self.counter):
                    self.powers.append(self.counter)
                    val = self.counter
                    yield val
                self.num = self.num / self.counter
            self.counter += 1
            if self.num == 1:
                self.last = True
                break


def main():
    rs = RandomStream(10)
    print map(lambda x: x, rs)

    pns = PrimeNumberStream(14)
    print map(lambda x: x, pns)

    pfs = PrimeFactorStram(90)
    print map(lambda x: x, pfs)

if __name__ == "__main__":
    main()
