#! /usr/bin/env python
import os
import unittest

from lru_cache import LRUCache


class TestLRUCache(unittest.TestCase):

    def setUp(self):
        self.cache_size_max = 15000
        self.num_urls = 15

        self.base_url = "https://placeholdit.imgix.net/~text?txt=image"

        self.cache_dir = '/tmp/cache'
        self.lru = LRUCache(self.cache_size_max, self.cache_dir)

    def test_lru_cache(self):
        url_count = 0
        out = []
        while url_count < self.num_urls:
            url = "{}{}".format(self.base_url, url_count)
            out.append(self.lru.get(url))
            url_count += 1
        expected = [
            '~text?txt=image11',
            '~text?txt=image12',
            '~text?txt=image13',
            '~text?txt=image14',
        ]
        self.assertEquals(os.listdir(self.cache_dir),
                          expected)
        expected_log = [
            'https://placeholdit.imgix.net/~text?txt=image0 DOWNLOADED 3535',
            'https://placeholdit.imgix.net/~text?txt=image1 DOWNLOADED 3440',
            'https://placeholdit.imgix.net/~text?txt=image2 DOWNLOADED 3538',
            'https://placeholdit.imgix.net/~text?txt=image3 DOWNLOADED 3537',
            'https://placeholdit.imgix.net/~text?txt=image4 DOWNLOADED 3488',
            'https://placeholdit.imgix.net/~text?txt=image5 DOWNLOADED 3511',
            'https://placeholdit.imgix.net/~text?txt=image6 DOWNLOADED 3515',
            'https://placeholdit.imgix.net/~text?txt=image7 DOWNLOADED 3499',
            'https://placeholdit.imgix.net/~text?txt=image8 DOWNLOADED 3522',
            'https://placeholdit.imgix.net/~text?txt=image9 DOWNLOADED 3511',
            'https://placeholdit.imgix.net/~text?txt=image10 DOWNLOADED 3578',
            'https://placeholdit.imgix.net/~text?txt=image11 DOWNLOADED 3459',
            'https://placeholdit.imgix.net/~text?txt=image12 DOWNLOADED 3573',
            'https://placeholdit.imgix.net/~text?txt=image13 DOWNLOADED 3580',
            'https://placeholdit.imgix.net/~text?txt=image14 DOWNLOADED 3526',
        ]
        self.assertEquals(out, expected_log)

    def test_lru_cache_in_cache(self):
        url_count = 0
        url_suffix = [1, 2, 3, 4, 5, 2, 1]
        out = []
        while url_count < len(url_suffix):
            url = "{}{}".format(self.base_url, url_suffix[url_count])
            out.append(self.lru.get(url))
            url_count += 1
        expected = [
            '~text?txt=image1',
            '~text?txt=image2',
            '~text?txt=image4',
            '~text?txt=image5',
        ]
        self.assertEquals(os.listdir(self.cache_dir),
                          expected)
        expected_log = [
            'https://placeholdit.imgix.net/~text?txt=image1 DOWNLOADED 3440',
            'https://placeholdit.imgix.net/~text?txt=image2 DOWNLOADED 3538',
            'https://placeholdit.imgix.net/~text?txt=image3 DOWNLOADED 3537',
            'https://placeholdit.imgix.net/~text?txt=image4 DOWNLOADED 3488',
            'https://placeholdit.imgix.net/~text?txt=image5 DOWNLOADED 3511',
            'https://placeholdit.imgix.net/~text?txt=image2 IN_CACHE 3538',
            'https://placeholdit.imgix.net/~text?txt=image1 DOWNLOADED 3440']
        self.assertEquals(out, expected_log)


if __name__ == '__main__':
    unittest.main()
