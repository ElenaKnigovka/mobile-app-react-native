# utils.py

import logging
import os
import json
import sys
import requests
import asyncio
from datetime import datetime
from typing import List, Dict

# Set up logging configuration
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s [%(levelname)s] %(message)s',
    datefmt='%Y-%m-%d %H:%M:%S',
)

class APIClient:
    def __init__(self, api_url: str):
        self.api_url = api_url

    async def get(self, endpoint: str, params: Dict = None) -> Dict:
        try:
            response = await self._make_request('GET', endpoint, params=params)
            return response.json()
        except requests.RequestException as e:
            logging.error(f"Error: {e}")
            raise

    async def post(self, endpoint: str, data: Dict) -> Dict:
        try:
            response = await self._make_request('POST', endpoint, data=data)
            return response.json()
        except requests.RequestException as e:
            logging.error(f"Error: {e}")
            raise

    async def _make_request(self, method: str, endpoint: str, data: Dict = None, params: Dict = None):
        url = f"{self.api_url}{endpoint}"
        headers = {'Content-Type': 'application/json'}
        if data:
            headers['Content-Type'] = 'application/json'
            data = json.dumps(data)
        async with requests.Session() as session:
            if method == 'GET':
                response = await session.get(url, params=params, headers=headers)
            elif method == 'POST':
                response = await session.post(url, data=data, headers=headers)
            else:
                logging.error("Invalid request method")
                raise ValueError("Invalid request method")
            response.raise_for_status()
            return response

def get_current_datetime() -> str:
    return datetime.now().strftime('%Y-%m-%d %H:%M:%S')

def get_current_datetime_epoch() -> int:
    return int(datetime.now().timestamp())

def list_files(directory: str) -> List[str]:
    return [f for f in os.listdir(directory) if os.path.isfile(os.path.join(directory, f))]

def get_process_memory() -> int:
    process = psutil.Process(os.getpid())
    return process.memory_info().rss