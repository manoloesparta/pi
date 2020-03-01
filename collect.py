import os
import sys
import time
import threading
import requests
from tqdm import tqdm

url = 'https://api.pi.delivery/v1/pi'
batch = 100000000

def getdigits(part, start, end):
	with open(f'num/part{part}', 'a+') as f:
		for i in tqdm(range(start, end, 1000)):
			
			params = {
				'start': i,
				'numberOfDigits': 1000,
			}
			try:
				req = requests.get(url=url, params=params)
			except:
				req = requests.get(url=url, params=params)

			data = req.json()
			f.write(data['content'].strip())

	print(f'Part {part} finished correctly')

def execute():
	threads = []

	for i in range(10):
		t = threading.Thread(target=getdigits, args=[i, i*batch, (i+1)*batch])
		t.start()
		threads.append(t)

	[ i.join() for i in threads ]

if __name__ == '__main__':
	os.mkdir('num')
	execute()