#! /usr/bin/env python

from datetime import datetime
import unittest

from syapse import parse_line
from syapse import timestamp_to_datetime


class TestTimestampToDatetime(unittest.TestCase):

    def test_timestamp_to_datetime(self):
        out = timestamp_to_datetime('01-01-2015 09:00:00')
        expected = datetime(2015, 1, 1, 9, 0, 0)
        self.assertEquals(out, expected)

    def test_timestamp_to_datetime_raises(self):
        with self.assertRaises(Exception) as cm:
            timestamp_to_datetime('01-01-15 09:00:00')
        expected = "Unable to parse timestamp '01-01-15 09:00:00'"
        self.assertEquals(cm.exception.message, expected)


class TestParseLine(unittest.TestCase):

    def setUp(self):
        self.db = {}

    def test_parse_line_begin(self):
        line = '01-01-2015 09:00:00,AAA,BEGIN'
        out = parse_line(self.db, line)
        expected = None
        self.assertEquals(out, expected)

    def test_parse_line_end(self):
        line = '01-01-2015 09:00:00,AAA,BEGIN'
        parse_line(self.db, line)
        line = '01-01-2015 09:02:00,AAA,END'
        out = parse_line(self.db, line)
        expected = 'AAA\t0:02:00'
        self.assertEquals(out, expected)

    def test_parse_line_not_enough_values(self):
        line = '01-01-2015 09:02:00,AAA'
        with self.assertRaises(Exception) as cm:
            parse_line(self.db, line)
        expected = "Unable to parse line '01-01-2015 09:02:00,AAA'"
        self.assertEquals(cm.exception.message, expected)

    def test_parse_line_end_missing_begin(self):
        line = '01-01-2015 09:02:00,AAA,END'
        with self.assertRaises(Exception) as cm:
            parse_line(self.db, line)
        expected = "Task ID 'AAA' has no BEGIN timestamp"
        self.assertEquals(cm.exception.message, expected)

    def test_parse_line_future_task_id(self):
        line = '01-01-2015 09:02:00,AAA,BEGIN'
        parse_line(self.db, line)
        line = '01-01-2015 09:00:00,AAA,END'
        with self.assertRaises(Exception) as cm:
            parse_line(self.db, line)
        expected = "Task ID 'AAA' appears to have started in the future"
        self.assertEquals(cm.exception.message, expected)

    def test_parse_line_not_begin_or_end(self):
        line = '01-01-2015 09:02:00,AAA,BAD'
        with self.assertRaises(Exception) as cm:
            parse_line(self.db, line)
        expected = "Unable to parse line '01-01-2015 09:02:00,AAA,BAD'"
        self.assertEquals(cm.exception.message, expected)


if __name__ == '__main__':
    unittest.main()
