#%%
import requests
import json
APIUrl = "http://127.0.0.1:9090/"
# %%
query = {'articleTitle': '同做', 'totalTasks': 20, 'answered': 10}
queryList = {"dataList": [query]}
r = requests.post(APIUrl+'saveArticles', json=queryList)
response = r.json()
# %%
query = {'userId': 'userId2cc6edb8'}
r = requests.post(APIUrl+'articles', json=query)
response = r.json()
response
# %%
query = {'articleTitle': '......地做', 'totalTasks': 20, 'answered': 10}
queryList = {"dataList": [query]}
data = [
    {
        'url': '/articleTitle',
        'params': {'articleTitle': '同地做', 'totalTasks': 20, 'answered': 10},
        'method': 'post',
    }
]
r = requests.post(APIUrl+'testArticles', json=[query])
response = r.json()
response
# %%
query = {'articleId': 'articleId2cc6edb8'}
r = requests.post(APIUrl+'tasks', json=query)
response = r.json()
response
# %%
query = {'ProductNumber': 0, 'quantity': 2}
testURL = "https://localhost:44311/ProductDetail/Buy"
r = requests.post(testURL, data=query, verify=False)
response = r.json()
response
# %%
