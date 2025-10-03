import requests

url = "https://api.sse-server.com/api/events?token=JWT_TOKEN"
with requests.get(url, stream=True) as r:
    for line in r.iter_lines():
        if line:
            print("Received:", line.decode())
