#! /usr/bin/env python

import collections
import os
import shutil

import requests


class LRUCache(object):
    """
    Manages an LRU Cache of downloaded images
    """

    def __init__(self, cache_size_max, cache_dir):
        """
        :param int cache_size_max: The maximum size of the cache in bytes
        :param str cache_dir: The directory to be used for the cache
        """
        # Set the max cache size and the current cache size
        self.cache_size_max = int(cache_size_max)
        self.cache_size = 0
        self.cache_dir = cache_dir

        # Create the cache dir if it does not exist
        if not os.path.isdir(self.cache_dir):
            os.mkdir(self.cache_dir)

        # Clear out the old files from the cache
        for f in os.listdir(self.cache_dir):
            os.remove(os.path.join(self.cache_dir, f))

        # Set the cache as an ordered dict
        # url: size
        self.cache = collections.OrderedDict()

    def _download(self, url):
        """
        Download the url and save to file in cache dir

        :param str url: The url to download and save
        :return: The size in bytes of the file
        """
        r = requests.get(url, stream=True)
        if r.status_code == 200:
            filename = os.path.join(self.cache_dir, os.path.basename(url))
            with open(filename, 'wb') as f:
                shutil.copyfileobj(r.raw, f)
            return os.path.getsize(filename)

    def get(self, url):
        """
        Get a url by first checking the cache and then downloading

        Using LRU this will first remove items from the ordered dict
        until below capacity.
        """
        # Save the state as IN_CACHE or DOWNLOADED
        state = ""

        try:
            size = self.cache.pop(url)
            state = "IN_CACHE"
        except KeyError:
            size = self._download(url)
            state = "DOWNLOADED"
            self.cache_size += size
            while self.cache_size > self.cache_size_max:
                pop_url, pop_size = self.cache.popitem(last=False)
                self.cache_size -= pop_size
                os.remove(os.path.join(self.cache_dir,
                                       os.path.basename(pop_url)))

        self.cache[url] = size
        return "{} {} {}".format(url, state, size)


if __name__ == "__main__":
    # Max size of the cache
    cache_size_max = None
    # Number of urls from the file to process
    num_urls = None
    # Counter for number of processed urls
    url_count = 0

    # Open and read the file out
    with open('sample-files/lru-test-input.txt') as f:

        cache_size_max = int(f.readline().strip())
        num_urls = int(f.readline().strip())

        cache_dir = os.path.abspath(os.path.join(os.getcwd(), 'cache'))
        lru = LRUCache(cache_size_max, cache_dir)

        while url_count < num_urls:
            url = f.readline().strip()
            print(lru.get(url))
            url_count += 1
