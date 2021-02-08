#%%
import requests
import json
APIUrl = "http://140.112.251.124:9090/"
# %%
query = {'articleTitle': '同做', 'totalTasks': 20, 'totalAnswered': 10}
queryList = {"dataList": [query]}
r = requests.post(APIUrl+'saveArticles', json=queryList)
response = r.json()
# %%
query = {'articleId': 'articleId2cc6edb8'}
r = requests.post(APIUrl+'articles', json=query)
response = r.json()
# %%
query = {'articleTitle': '......地做', 'totalTasks': 20, 'totalAnswered': 10}
queryList = {"dataList": [query]}
data = [
    {
        'url': '/articleTitle',
        'params': {'articleTitle': '同地做', 'totalTasks': 20, 'totalAnswered': 10},
        'method': 'post',
    }
]
r = requests.post(APIUrl+'testArticles', json=[query])
response = r.json()
# %%
