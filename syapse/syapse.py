#! /usr/bin/env python

from datetime import datetime


def timestamp_to_datetime(timestamp):
    """
    Convert a timestamp to a datetime object

    :param str timestamp: A timestamp
    :return: A datetime object
    :rtype: datetime
    :raises Exception: If the timestamp cannot be parsed
    """
    try:
        return datetime.strptime(timestamp, '%d-%m-%Y %H:%M:%S')
    except ValueError:
        raise Exception("Unable to parse timestamp '{}'".format(timestamp))


def parse_line(db, line):
    """
    Take log line and parse it

    :param dict db: A database to store the begining task timestamps
    :param str line: A log line from a csv file
    :return: None or the duration of the task
    :rtype: None or str
    :raises Exception: If the line cannot be parsed
    :raises Exception: If the task ID has no begining entry
    :raises Exception: If the task end timestamp is before the begin timestamp
    :raises Exception: If the value for begin_or_end is not BEGIN or END
    """
    try:
        timestamp, task_id, begin_or_end = line.strip().split(',')
    except ValueError:
        raise Exception("Unable to parse line '{}'".format(line))

    if begin_or_end == 'BEGIN':
        db[task_id] = timestamp
        return None
    elif begin_or_end == 'END':
        try:
            begin = db.pop(task_id)
        except KeyError:
            raise Exception("Task ID '{}' has no BEGIN timestamp".format(task_id))  # nopep8

        # Get the BEGIN and END datetimes
        dt_begin = timestamp_to_datetime(begin)
        dt_end = timestamp_to_datetime(timestamp)

        # Get the timestamp delta
        delta = dt_end - dt_begin
        if delta.total_seconds() < 0:
            raise Exception("Task ID '{}' appears to have started in the future".format(task_id))  # nopep8

        return '{}\t{}'.format(task_id, delta)
    else:
        raise Exception("Unable to parse line '{}'".format(line))


def main(filename):
    ID_TO_DATETIME = {}
    with open(filename, 'r') as f:
        for line in f:
            out = parse_line(ID_TO_DATETIME, line)
            if out:
                print(out)


if __name__ == "__main__":
    main('syapse.log')
