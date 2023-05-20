function seach(input) {
    let texto = input.value.trim();
    let link_enfeite = 'http://127.0.0.1:5500/search?q=' + encodeURIComponent(texto).replace(/%20/g, '+')
    let link = 'http://127.0.0.1:5500/search?q=' + encodeURIComponent(texto).replace(/%20/g, '+')
    
    fetch('https://jsonplaceholder.typicode.com/todos/1')
    .then(response => response.json())
    .then(data => {
        data['linkpesquisa'] = link_enfeite
        let jsonData = JSON.stringify(data);
        let base64Data = btoa(jsonData);
        window.location.href = "result.html?q=" + encodeURIComponent(base64Data);
    })
    .catch(error => {
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