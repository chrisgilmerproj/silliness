#! /usr/local/bin/python

import argparse


def main(usr_tld=['com'], dictionary=None, min_len=1, max_len=6):

    tld_dict = {}

    # Get a word list based on the user provided tld
    word_list = [word.strip().lower() for word in open(dictionary, 'r')]

    # Iterate through words, adding them to dictionary based on tld
    for word in word_list:
        for tld in usr_tld:
            if tld not in tld_dict:
                tld_dict[tld] = []
            if len(word) >= len(tld) + min_len and \
               len(word) <= len(tld) + max_len and \
               word.strip().lower()[-len(tld):] == tld:
                tld_dict[tld].append(word)

    for tld in usr_tld:
        print tld, tld_dict[tld]


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Get domain names based on provided tld.')
    parser.add_argument('tld', metavar='TLD', type=str, nargs='+',
                        help='give a top level domain')
    parser.add_argument('--dict', dest='dict', type=str, action='store',
                        default='/usr/share/dict/web2',
                        help='location of the dictionary file (default=%(default)s)')
    parser.add_argument('--min', dest='min', type=int, action='store',
                        default=1,
                        help='min domain length = tld+min (default=%(default)s)')
    parser.add_argument('--max', dest='max', type=int, action='store',
                        default=6,
                        help='max domain length = tld+max (default=%(default)s)')
    args = parser.parse_args()

    # Check that the tlds are in the available list
    tld_master_list = [tld.strip().lower() for tld in open('tlds-alpha-by-domain.txt', 'r')]
    tld_list = [tld.lower() for tld in args.tld if tld in tld_master_list]

    main(usr_tld=tld_list, dictionary=args.dict, min_len=args.min, max_len=args.max)
