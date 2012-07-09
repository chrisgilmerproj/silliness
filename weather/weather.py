#! /usr/local/bin/python

import calendar
import csv
import datetime
import sys
import urllib
import xml.dom.minidom
from xml.parsers.expat import ExpatError


class Weather(object):
    API_URL = 'http://www.google.com/ig/api'

    def __init__(self, filename):
        self.avg_weather = {}
        self.parse_data_from_file(filename)

    def parse_data_from_file(self, filename):
        weather_dict = csv.DictReader(open(filename, 'rb'), delimiter=',')
        for item in weather_dict:
            city = item['City']
            month = item['Month'].capitalize()
            temp = float(item['Temp'])
            if city not in self.avg_weather:
                self.avg_weather[city] = {}
            self.avg_weather[city][month] = temp

    def print_cities(self):
        month = self.get_month_current()

        for city in self.avg_weather:
            avg_temp = self.get_temp_average(city, month)
            cur_temp = self.get_temp_current(city)

            condition = 'Normal'
            if abs(avg_temp - cur_temp) > 5.0:
                condition = 'Unusual'

            print "City: {0}, Current Month: {1}, Avg Temp {2}, Cur. Temp {3}, Condition: {4}".format(city, month, avg_temp, cur_temp, condition)

    def get_month_current(self):
        return calendar.month_name[datetime.datetime.now().month]

    def get_temp_average(self, city, month):
        try:
            return self.avg_weather[city][month]
        except KeyError:
            raise Exception('City "{0}" or Month "{1}" do not exist'.format(city, month))

    def get_temp_current(self, query):
        response = urllib.urlopen('{0}?weather={1}'.format(self.API_URL, urllib.quote_plus(query))).read()

        try:
            dom = xml.dom.minidom.parseString(response)
        except ExpatError:
            raise Exception("Malformed response")

        conditions = self.get_element_from_dom(dom, 'current_conditions')
        temperature = self.get_element_from_dom(conditions, 'temp_f')

        return float(temperature.getAttribute('data'))

    def get_element_from_dom(self, dom, element_name):
        try:
            return dom.getElementsByTagName(element_name)[0]
        except IndexError:
            raise Exception("Unable to parse conditions")


def main():

    filename = 'weather.csv'
    weather = Weather(filename)
    weather.print_cities()


if __name__ == '__main__':
    main()
