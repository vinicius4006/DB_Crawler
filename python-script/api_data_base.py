import requests

def post_body(url, body):
    payload = {'url': url, 'body': body}
    response = requests.post('http://localhost:5050/api/sites', json=payload)

    return response.json()
        
def post_meta_tags(list_tag, siteid):
        url='http://localhost:5050/api/metatags'
        
        for tag in list_tag:
            payload = {
                "siteid": siteid,
                "tag": 'meta',
                "content": tag
            }

            requests.post(url, json=payload)

def post_words(siteid, words):
    for word, count in words.items():
        payload = {
            'siteid': siteid,
            'value': word,
            'counter': count
        }

        requests.post('http://localhost:5050/api/words', json=payload)