function seach(input) {
    let texto = input.value.trim();
	if (texto == '') {
		return
	}

	// add aqui o link para o teu servidor de pesquisa
    let link = 'http://127.0.0.1:5500/search?q=' + encodeURIComponent(texto).replace(/%20/g, '+')
    
	// faz a requisição para o servidor 
    fetch(link)
    .then(response => response.json())
    .then(data => {
        console.log(data);

		// tranforma o json em string 
        let jsonData = JSON.stringify(data)

		// transforma a string em base64
        let base64Data = btoa(jsonData)

		// redireciona para o arquivo result.html passando o resultado da presquisa na url
        window.location.href = "result.html?q=" + encodeURIComponent(base64Data)
    })
    .catch(error => {
        console.log(error)
    });    
}

var input = document.getElementById('searchinput');

// ativa a função pesquisa quando presiona enter
input.addEventListener('keydown', function(event) {
    if (event.key === 'Enter' || event.code === 'Enter') {
        seach(input)
    }
  });

// ativa a função quando aperta o botão pesquisar
document.getElementById("enviarJson").addEventListener("click", function() {
    seach(input)
});

