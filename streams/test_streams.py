#! /usr/local/bin/python

import unittest

from streams import RandomStream, PrimeNumberStream, PrimeFactorStream


class TestRandomStream(unittest.TestCase):

    def setUp(self):
        self.stream = RandomStream

    def test_random_numbers_map(self):
        size = 4
        number_list = map(lambda x: x, self.stream(size))
        self.assertEquals(len(number_list), size)
        self.assertEquals(len(set(number_list)), size)

    def test_random_numbers_filter(self):
        size = 4
        number_list = filter(lambda x: x < 0.5, self.stream(size))
        self.assertTrue(all([num < 0.5 for num in number_list]))

    def test_random_numbers_are_unique(self):
        size = 4
        stream = self.stream(size)
        number_list = []
        for i in xrange(size):
            number_list.append(stream.popNext())

        number_list = [num for num in number_list if num != None]
        self.assertEquals(len(set(number_list)), size)

    def test_random_numbers_return_only_set_number(self):
        size = 4
        stream = self.stream(size)
        number_list = []
        length = 10
        for i in xrange(length):
            number_list.append(stream.popNext())

        self.assertEquals(len([num for num in number_list if num != None]), size)
        self.assertEquals(len([num for num in number_list if num == None]), length - size)
        

class TestPrimeNumberStream(unittest.TestCase):

    def setUp(self):
        self.stream = PrimeNumberStream

    def test_prime_numbers_map(self):
        size = 12
        number_list = map(lambda x: x, self.stream(size))
        expected = [2, 3, 5, 7, 11]
        self.assertEquals(number_list, expected)
        self.assertEquals(len(set(number_list)), len(expected))

    def test_prime_numbers_filter(self):
        size = 12
        number_list = filter(lambda x: x < 5, self.stream(size))
        self.assertTrue(all([num < 5 for num in number_list]))

    def test_prime_numbers_are_unique(self):
        size = 12
        stream = self.stream(size)
        number_list = []
        length = 10
        for i in xrange(length):
            number_list.append(stream.popNext())

        number_list = [num for num in number_list if num != None]
        expected = [2, 3, 5, 7, 11]
        self.assertEquals(len(set(number_list)), len(expected))

    def test_prime_numbers_return_only_set_number(self):
        size = 12
        stream = self.stream(size)
        number_list = []
        length = 10
        for i in xrange(length):
            number_list.append(stream.popNext())

        expected = [2, 3, 5, 7, 11]
        self.assertEquals(len([num for num in number_list if num != None]), len(expected))
        self.assertEquals(len([num for num in number_list if num == None]), length - len(expected))

    def test_prime_numbers_are_correct(self):
        size = 12
        stream = self.stream(size)
        number_list = []
        length = 10
        for i in xrange(length):
            number_list.append(stream.popNext())

        number_list = [num for num in number_list if num != None]
        expected = [2, 3, 5, 7, 11]
        self.assertEquals(number_list, expected)


class TestPrimeFactorStream(unittest.TestCase):

    def setUp(self):
        self.stream = PrimeFactorStream

    def test_prime_factors_map(self):
        size = 90
        number_list = map(lambda x: x, self.stream(size))
        expected = [2, 3, 5]
        self.assertEquals(number_list, expected)
        self.assertEquals(len(set(number_list)), len(expected))

    def test_prime_factors_filter(self):
        size = 90
        number_list = filter(lambda x: x < 5, self.stream(size))
        self.assertTrue(all([num < 5 for num in number_list]))

    def test_prime_factors_are_unique(self):
        size = 90
        stream = self.stream(size)
        number_list = []
        length = 5
        for i in xrange(length):
            number_list.append(stream.popNext())

        number_list = [num for num in number_list if num != None]
        expected = [2, 3, 5]
        self.assertEquals(len(set(number_list)), len(expected))

    def test_prime_factors_are_correct(self):
        size = 90
        stream = self.stream(size)
        number_list = []
        length = 5
        for i in xrange(length):
            number_list.append(stream.popNext())

        number_list = [num for num in number_list if num != None]
        expected = [2, 3, 5]
        self.assertEquals(number_list, expected)


if __name__ == '__main__':
    unittest.main()
