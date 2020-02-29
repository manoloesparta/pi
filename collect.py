import os
import sys
import time
import threading
import requests
from tqdm import tqdm

url = 'https://api.pi.delivery/v1/pi'

def getdigits(part, start, end):

	with open(f'pi/part{part}', 'a+') as f:

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

	print(f'Part {part} confirmed')

def execute():

	threads = []

	for i in range(10):
		t = threading.Thread(target=getdigits, args=[i, i*100000000, (i+1)*100000000])
		t.start()
		threads.append(t)

	[ i.join() for i in threads ]

if __name__ == '__main__':

	start_time = time.time()

	if sys.argv[1] == 'erase':
		files = [ i for i in os.listdir('pi') ]

		for i in files:
			os.remove(os.path.join('pi', i))

	execute()

	print((time.time() - start_time))
