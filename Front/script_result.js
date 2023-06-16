async function search(input) {
    var inputHTML = document.getElementById('searchinput');
    inputHTML.value = input;
	// add aqui o link para o teu servidor de pesquisa
    let link = 'http://localhost:5050/cumae?q=' + input
    
	// faz a requisição para o servidor 
    let data = [];
    try{
        let response = await fetch(link);
        data = await response.json();  
    }catch(error){
        console.error(error);
        
    }
   return data;
}

async function show_results(qSearch) {
    const data = await search(qSearch);
    const regex = /\/(\w+(?:-\w+)*)\.html$/;
   

    let container = document.getElementById('container')

    // Cria o html com o conteúdo da pesquisa
    for (let i = 0; i < data.length; i++) {
        // capura o link de cada pagina
        let linkcompleto = data[i].Ref.Url
        // extrai o dominio pricipal
        let linkcurto = extractDomain(linkcompleto)
      
        // descrição e titulo de cada site
        let descricao = data[i].Ref.Body.substring(500, 1000) + "..." 
        const match = linkcompleto.match(regex);
        const titulo = match ? match[1].replace(/-/g, " ").toUpperCase() : "";

        container.innerHTML += 
            '<div class="container_site">' +
                '<a href="'+ linkcompleto +'">' +
                    '<p class="url">'+ linkcurto +'</p>' +
                    '<p class="full_url">'+ linkcompleto +'</p>' +
                    '<h2>'+ titulo +'</h2>' +
                '</a>' +
                '<p class="description">'+ descricao +'</p>' +
            '</div>' 
    }
    container.innerHTML += '<div id="space"></div>'
}

function extractDomain(url) {
    // Remove o protocolo (http://, https://) da URL
    var domain = url.replace(/(^\w+:|^)\/\//, '');
  
    // Remove o caminho após a primeira barra
    domain = domain.split('/')[0];
  
    // Remove www. se estiver presente
    domain = domain.replace(/^www\./, '');
  
    return domain;
}

var urlParams = new URLSearchParams(window.location.search);
var qSearch = urlParams.get("q");
console.log("PEGA", qSearch);
show_results(qSearch)
