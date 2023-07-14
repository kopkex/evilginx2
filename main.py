binance_console_app.py

import requests
import json

# Define Binance API endpoint URLs
BASE_URL = "https://api.binance.com"
API_KEY = "YOUR_API_KEY"
API_SECRET = "YOUR_API_SECRET"

def get_active_assets():
    """
    Retrieve user's active assets from Binance API.
    """
    response = requests.get(f"{BASE_URL}/api/v3/account", headers={"X-MBX-APIKEY": API_KEY})
    assets = json.loads(response.text)["balances"]
    active_assets = [asset for asset in assets if float(asset["free"]) > 0]
    return active_assets

def display_active_assets(assets):
    """
    Display user's active assets.
    """
    print("Active Assets:")
    for asset in assets:
        print(f"{asset['asset']}: {asset['free']}")

def select_asset_for_withdrawal(assets):
    """
    Allow user to select an asset for withdrawal.
    """
    print("Select an asset for withdrawal:")
    for i, asset in enumerate(assets):
        print(f"{i+1}. {asset['asset']}")
    choice = int(input("Enter the number of the asset: "))
    selected_asset = assets[choice-1]["asset"]
    return selected_asset

def withdraw_asset(asset):
    """
    Handle withdrawal of selected asset.
    """
    amount = float(input(f"Enter the amount of {asset} to withdraw: "))
    # Perform withdrawal using Binance API
    ...

def sell_all_assets_and_buy_xmr():
    """
    Sell all active assets and buy XMR Monero.
    """
    # Sell all active assets using Binance API
    ...
    # Buy XMR Monero using Binance API
    ...

# Main program
assets = get_active_assets()
display_active_assets(assets)
selected_asset = select_asset_for_withdrawal(assets)
withdraw_asset(selected_asset)
sell_all_assets_and_buy_xmr()
