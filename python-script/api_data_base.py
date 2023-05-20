import requests

def post_body(url, body):
    payload = {'url': url, 'body': body}
    response = requests.post('http://192.168.43.110:5050/api/sites', json=payload)

    return response.json()
        
def post_meta_tags(list_tag, siteid):
        url='http://192.168.43.110:5050/api/metatags'
        
        for tag in list_tag:
            payload = {
                "siteid": siteid,
                "tag": 'meta',
                "content": tag
            }

            requests.post(url, json=payload)
