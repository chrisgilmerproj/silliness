#! /bin/bash

yes "$(jot - 101 107)" | while read i; do printf "\033[${i}m" ; printf "\t%.0s" {1..10}; printf "\n"; sleep .05; done
