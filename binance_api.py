import requests
import json

API_KEY = 'YOUR_API_KEY'
SECRET_KEY = 'YOUR_SECRET_KEY'

def list_assets():
    headers = {
        'X-MBX-APIKEY': API_KEY
    }
    response = requests.get('https://api.binance.com/api/v3/account', headers=headers)
    data = response.json()
    return data['balances']

def make_withdrawal(asset, amount, address):
    headers = {
        'X-MBX-APIKEY': API_KEY
    }
    params = {
        'asset': asset,
        'amount': amount,
        'address': address
    }
    response = requests.post('https://api.binance.com/sapi/v1/capital/withdraw/apply', headers=headers, params=params)
    data = response.json()
    return data

# Test the functions
print(list_assets())
print(make_withdrawal('BTC', 0.001, 'YOUR_BTC_ADDRESS'))