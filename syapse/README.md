Syapse
======

Problem Statement
-----------------

Consider a system that executes various tasks, and records task starts and stops in a log file.  The log file has three columns:

Timestamp - the timestamp of the entry, assume any format you like
Task ID - the ID of the task that got started or stopped

A keyword BEGIN for the start of the task, ‘END’ for the end of the task.
Each task has exactly one BEGIN and END entry in the file.  E.g., a valid log file may look like

 Timestamp                      TaskID          BeingOrEnd

01-01-2015 09:00:00      AAA                BEGIN
01-01-2015 09:01:00      BBB                BEGIN
01-01-2015 09:02:00      AAA                END
01-01-2015 09:03:00      CCC                BEGIN
01-01-2015 09:04:00      CCC                END
01-01-2015 09:05:00      BBB                END

You need to print a report indicating how long each task took, e.g., for the example above you need to print roughly the following:

TaskID   Duration
AAA       02:00   
BBB       04:00
CCC      01:00

Please assume the file may be reasonably large.  Other than that, please make any assumptions you want, especially with regard to exact file format.

Running the Program
-------------------

To run the program simply:

```sh
$ ./syapse.py
AAA     0:02:00
CCC     0:01:00
BBB     0:04:00
```

Tests
-----

There are tests for this code.  Run them:

```sh
$ ./test_syapse.py
........
----------------------------------------------------------------------
Ran 8 tests in 0.004s

OK
```
