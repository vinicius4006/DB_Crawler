function search(input) {
    let texto = input.value.trim();
	if (texto == '') {
		return
	}
    let textoFormatado = encodeURIComponent(texto).replace(/%20/g, '+');
		// redireciona para o arquivo result.html passando o resultado da presquisa na url
        window.location.href = "result.html?q=" + textoFormatado
    
}

var input = document.getElementById('searchinput');

// ativa a função pesquisa quando presiona enter
input.addEventListener('keydown', function(event) {
    if (event.key === 'Enter' || event.code === 'Enter') {
        search(input)
    }
  });

// ativa a função quando aperta o botão pesquisar
document.getElementById("enviarJson").addEventListener("click", function() {
    search(input)
});


