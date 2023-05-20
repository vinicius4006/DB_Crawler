function seach(input) {
    let texto = input.value.trim();
    let link = 'http://127.0.0.1:5500/search?q=' + encodeURIComponent(texto).replace(/%20/g, '+')
    
    console.log(link);

    fetch('https://jsonplaceholder.typicode.com/todos/1')
    .then(response => response.json())
    .then(data => {
        data['linkpesquisa'] = link
        let jsonData = JSON.stringify(data);
        // Codifica a string em base64
        let base64Data = btoa(jsonData);
        // Abre a página Y com o JSON como parâmetro na URL
        window.location.href = "result.html?q=" + encodeURIComponent(base64Data);
    })
    .catch(error => {
      // Em caso de erro na requisição
      console.log(error);
    });    
}

var input = document.getElementById('searchinput');
input.addEventListener('keydown', function(event) {
    if (event.key === 'Enter' || event.code === 'Enter') {
        seach(input)
    }
  });

document.getElementById("enviarJson").addEventListener("click", function() {
    seach(input)
});