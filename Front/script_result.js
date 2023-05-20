function show_results(data) {
    history.pushState(null, null, data['linkpesquisa']);
    console.log(data);

    let container = document.getElementById('container')
    for (let index = 0; index < 10; index++) {
        container.innerHTML += 
        '<div id="container">' +
            '<div class="container_site">' +
                '<a href="#">' +
                    '<p class="url">google.com</p>' +
                    '<p class="full_url">google.com/jogos/online/issomermo</p>' +
                    '<h2>Como acessar jogos online</h2>' +
                '</a>' +
                '<p class="description">Todos os arquivos de Assistência Técnica - Documentação básica 1. Assistência Técnica - Documentação básica 10. Assistência Técnica - Documentação básica 10 ...</p>' +
            '</div>' +
        '</div>'
    }
}

var urlParams = new URLSearchParams(window.location.search);
var base64Data = urlParams.get("q");
var jsonData = atob(base64Data);
var data = JSON.parse(jsonData);

show_results(data)