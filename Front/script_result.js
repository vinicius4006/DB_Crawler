function show_results(data) {
    console.log(data);
    let container = document.getElementById('container')

    // Cria o html com o conteúdo da pesquisa
    for (let i = 0; i < data.length; i++) {
        // capura o link de cada pagina
        linkcompleto = data[i].Ref.Url
        // extrai o dominio pricipal
        let linkcurto = extractDomain(linkcompleto)

        // descrição e titulo de cada site
        descricao = 'variavel da descrição aqui'
        titulo = 'varialvel do titulo aqui'

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
var base64Data = urlParams.get("q");
var jsonData = atob(base64Data);
var data = JSON.parse(jsonData);

show_results(data)