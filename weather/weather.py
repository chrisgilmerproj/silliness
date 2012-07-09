#! /usr/local/bin/python

import urllib
import xml.dom.minidom
from xml.parsers.expat import ExpatError


class Weather(object):
    API_URL = 'http://www.google.com/ig/api'

    def current(self, query):
        response = urllib.urlopen('{0}?weather={1}'.format(self.API_URL, urllib.quote_plus(query))).read()

        try:
            dom = xml.dom.minidom.parseString(response)
        except ExpatError:
            raise Exception("Malformed response")

        conditions = self.get_element_from_dom(dom, 'current_conditions')
        temperature = self.get_element_from_dom(conditions, 'temp_f')

        return temperature.getAttribute('data')

    def get_element_from_dom(self, dom, element_name):
        try:
            return dom.getElementsByTagName(element_name)[0]
        except IndexError:
            raise Exception("Unable to parse conditions")


def main():

    weather = Weather()
    query = '94607'

    try:
        conditions = weather.current(query)
        print conditions
    except Exception as e:
        sys.exit(1)


if __name__ == '__main__':
    main()
