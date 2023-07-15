import threading
import requests
from bs4 import BeautifulSoup
from api_data_base import *
from collections import Counter
import re

def criar_vetor_de_termos(soup):
    body = soup.body
    texto = body.get_text()
    texto = re.sub(r'\W+', ' ', texto)
    texto = re.sub(r'\s+', ' ', texto)
    texto = texto.lower()
    termos = texto.split()
    frequencia = Counter(termos)

    return frequencia


def extract_body_content(soup):
    body = soup.find('body')
    if body:
        content = body.get_text()
        content = '\n'.join(line for line in content.splitlines() if line.strip())
        return content
    else:
        return ''
 
def get_meta_tags(soup):
    meta_tags = soup.find_all('meta')
    meta_tag_list = []
    
    for tag in meta_tags:
        meta_tag_list.append(str(tag))
    
    return meta_tag_list

def remove_repeated_items(list_links):
    return list(set(list_links))

def filter_https_links(list_links):
    filtered_links = []
    for link in list_links:
        if link.startswith('https://'):
            filtered_links.append(link)
    return filtered_links

def get_links(soup):
    list_links = []
    for tag_a in soup.find_all('a'):
        link = tag_a.get('href')
        if link:
            list_links.append(link)
    return list_links

def web_crawler(url, profundidade, contador):
    print('\n', contador, url)

    try:
        response = requests.get(url, timeout=100).content
        soup = BeautifulSoup(response, 'html.parser')

    except Exception as e:
        print("Erro ao acessar:", e, url)
        return
    
    try:
        list_meta_tag = get_meta_tags(soup)
        body = extract_body_content(soup)
        vetor_de_termos = criar_vetor_de_termos(soup)
        siteid = post_body(url, body)
        post_meta_tags(list_meta_tag, siteid)
        post_words(siteid, vetor_de_termos)
    except Exception as e:
        print("Erro no Banco de dados:", e, url)
        return
        
    if profundidade == contador:
        print('Chegou na profundidade')
        return

    contador += 1

    try:
        list_links = get_links(soup)
        list_links = remove_repeated_items(list_links)
        list_links = filter_https_links(list_links)

    except Exception as e:
        print("Erro em capturar links", e, url)
        return

    threads = []
    for link in list_links:
        t = threading.Thread(target=web_crawler, args=(link, profundidade, contador))
        threads.append(t)
        t.start()
    
    for t in threads:
        t.join()

    return 'Terminou'


a = web_crawler('https://www.estadao.com.br/', 1, 0)
print(a)
# web_crawler('https://oglobo.globo.com/', 1, 0)
# web_crawler('https://www.folha.uol.com.br/', 1, 0)
# web_crawler('https://www.estadao.com.br/', 1, 0)
# web_crawler('https://valor.globo.com/', 1, 0)
# web_crawler('https://ahnegao.com.br/', 2, 0)
# web_crawler('https://gauchazh.clicrbs.com.br/', 1, 0)
# web_crawler('https://www.jb.com.br/', 1, 0)
# web_crawler('https://www.nsctotal.com.br/diarinho', 2, 0)
# web_crawler('https://www.opovo.com.br/', 1, 0)


