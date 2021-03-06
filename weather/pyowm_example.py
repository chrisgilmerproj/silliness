import pyowm
from settings_secret import OWM_API_KEY

LOCATION = 'Oakland,CA'

owm = pyowm.OWM(OWM_API_KEY)

# Will it be sunny tomorrow at this time in this location?
forecast = owm.daily_forecast(LOCATION)
tomorrow = pyowm.timeutils.tomorrow()
forecast.will_be_sunny_at(tomorrow)

# Search for current weather in location
observation = owm.weather_at_place(LOCATION)
w = observation.get_weather()
print(w)                      # <Weather - reference time=2013-12-18 09:20, 
                              # status=Clouds>

# Weather details
w.get_wind()                  # {'speed': 4.6, 'deg': 330}
w.get_humidity()              # 87
w.get_temperature('celsius')  # {'temp_max': 10.5, 'temp': 9.7, 'temp_min': 9.0}

# Search current weather observations in the surroundings of 
# lat=22.57W, lon=43.12S (Rio de Janeiro, BR)
observation_list = owm.weather_around_coords(-22.57, -43.12)
