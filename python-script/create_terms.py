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
        

