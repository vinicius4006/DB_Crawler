import requests
import base64

def post_request_with_base64(url, body):
    body_base64 = base64.b64encode(body.encode('utf-8')).decode('utf-8')
    payload = {'url': url, 'body': body_base64}

    response = requests.post('http://192.168.2.44:5050/api/sites', json=payload)

    return response.json()


def post_words(siteid, words):
    for word, count in words.items():
        payload = {
            'siteid': siteid,
            'value': word,
            'counter': count
        }
        
        requests.post('http://192.168.2.44:5050/api/words', json=payload)
        
def post_meta_tags(list_tag, siteid):
        url='http://192.168.2.44:5050/api/metatags'
        
        for tag in list_tag:
            tag_base64 = base64.b64encode(tag.encode('utf-8')).decode('utf-8')
            payload = {
                "siteid": siteid,
                "tag": 'meta',
                "content": tag_base64
            }

            requests.post(url, json=payload)
            