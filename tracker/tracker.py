#! /usr/local/bin/python

import heapq
import random


class Singleton(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton,
                                        cls).__call__(*args, **kwargs)
        return cls._instances[cls]


class LargestTracker(object):
    __metaclass__ = Singleton
    _entries = []

    @staticmethod
    def getInstance():
        """
        guarantees the creation of a single instance across the virtual
        machine. assumed to be called very frequently.

        @return an instance of largesttracker
        """
        return LargestTracker()

    def getNLargest(self, numberOfTopLargestElements):
        """
        Returns a list in O(n log m) time OR BETTER where n is the number of
        entries added to LargestTracker and m is numberOfTopLargestElements.
        Duplicates are allowed

        @param numberOfTopLargestElements
                    the number of top-most-elements to return
        @return the top-most-elements in the tracker sorted in ascending order
                    as a list of integers
        """
        # Docs say for 1 use max()
        if numberOfTopLargestElements == 1:
            return max(self._entries)
        # Docs say heapq.nlargest() only beats sort() for small number of N
        elif numberOfTopLargestElements < 1000:
            return heapq.nlargest(numberOfTopLargestElements,
                                  self._entries)[::-1]
        # Use sort() for large numbers of elements
        else:
            return sorted(self._entries)[-numberOfTopLargestElements:]

    def add(self, anEntry):
        """
        Adds an entry to the tracker. This method must operate in O(log n) time
        OR BETTER.

        @param anEntry
                   the entry to add to the tracker. Entries need not be unique.
        """
        heapq.heappush(self._entries, anEntry)

    def clear(self):
        """
        Removes all the entries from the tracker. This should return in
        constant time.
        """
        del self._entries[:]


def test():
    largeNumber = 100000

    # Test singleton
    tracker = LargestTracker()
    new_tracker = tracker.getInstance()
    assert tracker == new_tracker

    # Test adding elements
    for i in xrange(0, largeNumber):
        elem = random.randint(0, largeNumber)
        tracker.add(elem)
    assert len(tracker._entries) == largeNumber

    # Test getting largest elements
    largest = tracker.getNLargest(1)
    largest = tracker.getNLargest(50)
    assert len(largest) == 50
    assert largest[0] < largest[49]
    largest = tracker.getNLargest(1000)

    # Test clear
    tracker.clear()
    assert len(tracker._entries) == 0


if __name__ == "__main__":
    test()
