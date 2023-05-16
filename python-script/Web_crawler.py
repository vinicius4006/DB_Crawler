import threading
import requests
from bs4 import BeautifulSoup
from create_terms import criar_vetor_de_termos
from api_data_base import post_request_with_base64, post_words, post_meta_tags

def get_meta_tags(soup):
    meta_tags = soup.find_all('meta')
    meta_tag_list = []
    
    for tag in meta_tags:
        meta_tag_list.append(str(tag))
    
    return meta_tag_list

def get_body_content(soup):
    return soup.body

def remove_repeated_items(list_links):
    return list(set(list_links))

def filter_https_links(list_links):
    filtered_links = []
    for link in list_links:
        if link.startswith('https://'):
            filtered_links.append(link)
    return filtered_links

def get_links(url, profundidade, contador):
    try:
        response = requests.get(url, timeout=100).content
        soup = BeautifulSoup(response, 'html.parser')
        vetor_de_termos = criar_vetor_de_termos(soup)
        print('\n', contador, url)
        
        listaTemp = dict()
        listaTemp['url'] = url
        list_meta_tag = get_meta_tags(soup)
        body = get_body_content(soup)
        
        siteid = post_request_with_base64(url, body)
        post_meta_tags(list_meta_tag, siteid)
        post_words(siteid, vetor_de_termos)
       
        if profundidade == contador:
            print('chegou na profundidade\n')
            return
    
        contador += 1
        
        list_links = []
        for tag_a in soup.find_all('a'):
            link = tag_a.get('href')
            if link:
                list_links.append(link)
        
        list_links = remove_repeated_items(list_links)
        list_links = filter_https_links(list_links)
        
        threads = []
        for link in list_links:
            t = threading.Thread(target=get_links, args=(link, profundidade, contador))
            threads.append(t)
            t.start()
        
        for t in threads:
            t.join()

    except requests.exceptions.Timeout:
        print("Timeout de conexão. ", url)
        return
    
    except Exception as e:
        print("Ocorreu um erro:", e, url)
        return


get_links('https://recordtv.r7.com/', 1, 0)


