###############################################################################
# This is an example to illustrate how to use                                 #
# the Numista API in Python.                                                  #
#                                                                             #
# The processing is simplified for illustrative                               #
# purpose. It should not be used for real applications                        #
# without considering all possible cases, espcially                           #
# error cases.                                                                #
###############################################################################

import os
import requests
from uuid import uuid4

# Parameters
endpoint = 'https://api.numista.com/api/v2'
api_key = os.getenv("NUMISTA_API_KEY")
client_id = os.getenv("NUMISTA_CLIENT_ID")

###############################################################################

# API call to search coins
search_query = 'Kopecks Siberia'
response = requests.get(
  endpoint + '/coins',
  params={'q': search_query, 'page': 1, 'count': 50, 'lang': 'en'},
  headers={'Numista-API-Key': api_key})
search_results = response.json()

# Display search results
print('== Search results for the query "', search_query, '" ==')
print(search_results['count'], 'results found')
for coin in search_results['coins']:
  print(coin['title'], 'from', coin['issuer']['name'])
print()

###############################################################################

# API call to get details about a coin
coin_type_id = '17970'
response = requests.get(
  endpoint + '/coins/' + coin_type_id,
  params={'lang':'en'},
  headers={'Numista-API-Key': api_key})
coin_details = response.json()

# Display some details about the coin
print('== Details about the coin type #', coin_type_id, '==')
print('URL:', coin_details['url'])
print('Title:', coin_details['title'])
print('Issuer:', coin_details['issuer']['name'])
print('Years:', coin_details['min_year'], '-', coin_details['max_year'])
print('Composition:', coin_details['composition']['text'])
print('Weight:', coin_details['weight'], 'grams')
print()

###############################################################################

# API call to get the years when the coin was minted
response = requests.get(
  endpoint + '/coins/' + coin_type_id + '/issues',
  params={'lang': 'en'},
  headers={'Numista-API-Key': api_key})
years = response.json()

# Display the years
print('== Years ==')
for year in years:
  print(year['year'])
print()

###############################################################################

# API call to get the list of issuers
response = requests.get(
  endpoint + '/issuers',
  params={'lang': 'en'},
  headers={'Numista-API-Key': api_key})
issuers = response.json()

# Display the issuers
print('== Issuers ==')
for i in range(10):
  print(issuers['issuers'][i]['name'])
print('and', issuers['count']-10, 'others...')
print()

###############################################################################

# Authenticate with OAuth2
redirect_uri = 'https://postman-echo.com/get' # Should normally be a URL to your application
state = str(uuid4())
url_template = ('https://{language}.numista.com/api/oauth_authorize.php'
                '?response_type=code'
                '&client_id={client_id}'
                '&redirect_uri={redirect_uri}'
                '&scope={scope}&state={state}')
authorization_url = url_template.format(
  language='en',
  client_id=client_id,
  redirect_uri=redirect_uri,
  scope='view_collection,edit_collection',
  state=state)

print('== Authentication ==')
print('Please open the following URL, authenticate yourself '
      'and provide the data present in the redirection.')
print('Authorization URL:', authorization_url)
authorization_code = input('Code? ')
returned_state = input('State? ')

# Prevent CSFR attacks
if state == returned_state: print ('State is correct.');
else: print('Incorrect state. You should stop the process here.')

# Retrieve access token
response = requests.get(
  endpoint + '/oauth_token',
  params={
    'code': authorization_code,
    'client_id': client_id,
    'client_secret': api_key,
    'redirect_uri': redirect_uri})
authentication_data = response.json()
access_token = authentication_data['access_token']
expires_in = authentication_data['expires_in'] / 3600 / 24
user_id = authentication_data['user_id']
print('The user #', user_id, 'is authenticated for', expires_in, 'days.')
print()

###############################################################################

# API call to get information about a user
response = requests.get(
  endpoint + '/users/' + str(user_id),
  params={'lang': 'en'},
  headers={'Numista-API-Key': api_key})
user_details = response.json()

# Display the issuers
print('== User #', user_id, ' ==')
print('Username:', user_details['username'])
print('Avatar:', user_details['avatar'])
print()

###############################################################################

# API call to get the coins in collection
response = requests.get(
  endpoint + '/users/' + str(user_id) + '/collected_coins',
  params={'lang': 'en'},
  headers={'Numista-API-Key': api_key, 'Authorization': 'Bearer '+access_token})
collection = response.json()

# Display a coin from the collection
print('== Collection ==')
print('The user owns', collection['coin_count'], 'coins.')
print('One of them is:')
coin = collection['collected_coins'][0]
print('Coin:', coin['coin']['title'])
print('Year:', coin['issue']['year'])
print('For swap:', coin['for_swap'])
print()
