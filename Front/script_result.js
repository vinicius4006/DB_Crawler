function show_results(data) {
    history.pushState(null, null, data['linkpesquisa']);
    console.log(data);

    let container = document.getElementById('container')
    for (let i = 0; i < data.length; i++) {
        linkcompleto = data[i].Ref.Url
        let linkcurto = extractDomain(linkcompleto)
        container.innerHTML += 
        '<div id="container">' +
            '<div class="container_site">' +
                '<a href="'+ linkcompleto +'">' +
                    '<p class="url">'+ linkcurto +'</p>' +
                    '<p class="full_url">'+ linkcompleto +'</p>' +
                    '<h2>Como acessar jogos online</h2>' +
                '</a>' +
                '<p class="description">Todos os arquivos de Assistência Técnica - Documentação básica 1. Assistência Técnica - Documentação básica 10. Assistência Técnica - Documentação básica 10 ...</p>' +
            '</div>' +
        '</div>'
    }
    container.innerHTML += '<div id="space"></div>'
    revovePreload()
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

function revovePreload() {
    // document.getElementById('preload').style.transform = 'translate(0%, -100%)'
}

var urlParams = new URLSearchParams(window.location.search);
var base64Data = urlParams.get("q");
var jsonData = atob(base64Data);
var data = JSON.parse(jsonData);

show_results(data)