import json
import os
import requests
from uuid import uuid4

# Parameters
endpoint = "https://api.numista.com/api/v2"
api_key = os.getenv("NUMISTA_API_KEY")
client_id = os.getenv("NUMISTA_CLIENT_ID")

response = requests.get(
    endpoint + "/oauth_token",
    params={
        "grant_type": "client_credentials",
    },
    headers={"Numista-API-Key": api_key},
)
authentication_data = response.json()
access_token = authentication_data["access_token"]
user_id = authentication_data["user_id"]
# print(authentication_data)

response = requests.get(
    endpoint + "/users/" + str(user_id) + "/collected_coins",
    params={"lang": "en"},
    headers={"Numista-API-Key": api_key, "Authorization": "Bearer " + access_token},
)
collection = json.dumps(response.json())
print(collection)
