#! /usr/local/bin/python

import calendar
import csv
import datetime
import sys
import urllib
import xml.dom.minidom
from xml.parsers.expat import ExpatError


class Weather(object):
    """
    Weather is a class designed to take average temperature
    weather data from a number of cities, contact a weather service
    and return the current temperature and conditions for those cities.

    This class uses the Google API for weather.
    """

    API_URL = 'http://www.google.com/ig/api'

    def __init__(self, filename):
        """
        Initialize the weather class with data from a csv file.

        @param filename: Weather data file
        """
        self.avg_weather = self.parse_data_from_file(filename)

    def parse_data_from_file(self, filename):
        """
        Parse weather data from a file where the columns are:

        City, Month, Avg. Temp

        The parsed data populates a dictionary keyed on city name
        and again by month.

        @param filename: Weather data file
        @returns: A dictionary of parsed weather data
        """
        avg_weather = {}
        weather_dict = csv.DictReader(open(filename, 'rb'), delimiter=',')
        for item in weather_dict:
            city = item['City']
            month = item['Month'].capitalize()
            temp = float(item['Temp'])
            if city not in avg_weather:
                avg_weather[city] = {}
            avg_weather[city][month] = temp
        return avg_weather

    def print_cities(self):
        """
        Print the temperature data for all of the available cities.

        If the current temperature is more than 5.0 degF outside the
        average temperature for a city the condition will be printed
        as 'Unusual'
        """
        month = self.get_month_current()

        for city in self.avg_weather:
            avg_temp = self.get_temp_average(city, month)
            cur_temp = self.get_temp_current(city)

            condition = 'Normal'
            if abs(avg_temp - cur_temp) > 5.0:
                condition = 'Unusual'

            print "City: {0}, Current Month: {1}, Avg Temp {2}, Cur. Temp {3}, Condition: {4}".format(city, month, avg_temp, cur_temp, condition)

    def get_month_current(self):
        """
        Get the current calendar month

        @returns: A calendar month name
        """
        return calendar.month_name[datetime.datetime.now().month]

    def get_temp_average(self, city, month):
        """
        Get the average temperature for a city and a given month

        @param city: The name of a city
        @param month: The calendar month name
        @returns: A temperature as a floating point number
        """
        try:
            return self.avg_weather[city][month]
        except KeyError:
            raise Exception('City "{0}" or Month "{1}" do not exist'.format(city, month))

    def get_temp_current(self, query):
        """
        Get the current temperature for a city by calling a weather service.

        @param query: The query string for the weather api
        @returns: A temperature as a floating point number
        """
        response = urllib.urlopen('{0}?weather={1}'.format(self.API_URL, urllib.quote_plus(query))).read()

        try:
            dom = xml.dom.minidom.parseString(response)
        except ExpatError:
            raise Exception("Malformed response")

        conditions = self.get_element_from_dom(dom, 'current_conditions')
        temperature = self.get_element_from_dom(conditions, 'temp_f')

        return float(temperature.getAttribute('data'))

    def get_element_from_dom(self, dom, element_name):
        """
        Given a dom element return an element

        @param dom: A xml dom element
        @param element_name: The name of a tag element in the dom
        @returns: The dom element or string
        """
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
