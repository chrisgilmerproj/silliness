#! /usr/local/bin/python

import urllib
import urllib2


def main():
    qr_gen = "https://chart.googleapis.com/chart?"
    qr_map = {
              'chs': '150x150',
              'cht': 'qr',
             }

    # Read the words from the text file
    f_words = open('qr_words.txt', 'r')

    # Generate a code for every word
    for word in f_words:

        # Clean up the word
        word = word.strip()

        # Update the GET parameters and construct the url
        qr_map.update({'chl': word})
        url = qr_gen + urllib.urlencode(qr_map)
        print word, url

        # Get the picture data
        opener = urllib2.build_opener()
        page = opener.open(url)
        picture = page.read()

        # Write the file
        filename = word + '.png'
        f = open(filename, "wb")
        f.write(picture)
        f.close()


if __name__ == '__main__':
    main()
