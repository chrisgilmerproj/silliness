#! /usr/local/bin/python

import math
import random
import threading


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

    def popN(self, num):
        """
        Return only num elements from stream
        """
        return [self.popNext() for x in xrange(0, num)]

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


def new_map(fn, stream):
    """
    Returns a stream where fn has been applied to each element.
    """
    return [fn(n) for n in stream]


def new_filter(fn, stream):
    """
    Returns a stream containing only the elements of the stream for which
    fn returns True.
    """
    return [n for n in stream if fn(n)]


def zip_with(fn, streamA, streamB):
    """
    Applies a given binary function pairwise to the elements of two given
    lists.
    """
    return [fn(z) for z in zip(streamA, streamB)]


def prefix_reduce(fn, stream, init=None):
    """
    Where fn(x,y) is a function to perform a reduction across the stream,
    returns a stream where the nth element is the result of combining the
    first n elements of the input stream using fn.

    This was particularly hard to understand so I interpreted it as
    returning an iterable where the fn is applied cumulatively up
    the stream.
    """
    numbers = []
    it = iter(stream)
    if init is None:
        try:
            init = next(it)
        except StopIteration:
            raise TypeError
    val = init
    for num in stream:
        val = fn((val, num))
        numbers.append(val)
    return numbers


class StreamThread(threading.Thread):
    """
    A thread class to allow running multiple popN
    methods at the same time.  This should significantly
    increase our ability to process streams and apply
    different higher-order functions.
    """

    def __init__(self, stream):
        self.stream = stream
        super(StreamThread, self).__init__()

    def run(self):
        print self.stream.popN(3)


def main():
    rs = RandomStream(5)
    #print map(lambda x: x, rs)

    pns = PrimeNumberStream(14)
    #print map(lambda x: x, pns)

    pfs = PrimeFactorStream(100)
    #print map(lambda x: x, pfs)

    stream_list = [rs, pns, pfs]
    for stream in stream_list:
        t = StreamThread(stream)
        t.start()


if __name__ == "__main__":
    main()
