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
query = {'userId': '117245596358325515598'}
r = requests.post(APIUrl+'projects', json=query)
response = r.json()
response

# %%
query = {'projectId': "1"}
r = requests.post(APIUrl+'users', json=query)
response = r.json()
response
# %%
