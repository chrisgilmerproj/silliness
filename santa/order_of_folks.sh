#! /usr/bin/env bash

#
# Randomly order the names in names.txt based on the contents of message.txt
#
# Each name in names.txt is on a separate line with no whitespace surrounding it
# The message in message.txt should not have any whitespace surrounding it
#

set -eu -o pipefail

# If message.txt does not exist exit
if [ ! -f message.txt ]; then
    echo "message.txt does not exist"
    exit 1
fi

# If names.txt does not exist exit
if [ ! -f names.txt ]; then
    echo "names.txt does not exist"
    exit 1
fi

sort -R --random-source=<(sha256sum message.txt | cut -d' ' -f1 ) names.txt | cat -n
