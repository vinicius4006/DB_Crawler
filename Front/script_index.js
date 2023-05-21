let resposta = [
	{
		"ID": 219536,
		"CreatedAt": "2023-05-19T17:16:04.312771-03:00",
		"UpdatedAt": "2023-05-19T17:16:04.312771-03:00",
		"DeletedAt": null,
		"SiteID": 1479,
		"Ref": {
			"ID": 1479,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://adorofutebol.com.br",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 1
	},
	{
		"ID": 268874,
		"CreatedAt": "2023-05-19T17:25:05.08377-03:00",
		"UpdatedAt": "2023-05-19T17:25:05.08377-03:00",
		"DeletedAt": null,
		"SiteID": 1669,
		"Ref": {
			"ID": 1669,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://gda.com/",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 3
	},
	{
		"ID": 358579,
		"CreatedAt": "2023-05-19T17:37:56.10388-03:00",
		"UpdatedAt": "2023-05-19T17:37:56.10388-03:00",
		"DeletedAt": null,
		"SiteID": 2156,
		"Ref": {
			"ID": 2156,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://adorofutebol.com.br",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 1
	},
	{
		"ID": 372241,
		"CreatedAt": "2023-05-19T18:04:39.281663-03:00",
		"UpdatedAt": "2023-05-19T18:04:39.281663-03:00",
		"DeletedAt": null,
		"SiteID": 2173,
		"Ref": {
			"ID": 2173,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://www.mongehumorista.com/",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 1
	},
	{
		"ID": 377620,
		"CreatedAt": "2023-05-19T18:05:19.713143-03:00",
		"UpdatedAt": "2023-05-19T18:05:19.713143-03:00",
		"DeletedAt": null,
		"SiteID": 2190,
		"Ref": {
			"ID": 2190,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://www.mensagensdebomdia.com.br/",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 1
	},
	{
		"ID": 378785,
		"CreatedAt": "2023-05-19T18:08:55.10239-03:00",
		"UpdatedAt": "2023-05-19T18:08:55.10239-03:00",
		"DeletedAt": null,
		"SiteID": 2217,
		"Ref": {
			"ID": 2217,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://www.mongehumorista.com/",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 1
	},
	{
		"ID": 387220,
		"CreatedAt": "2023-05-19T18:09:25.885555-03:00",
		"UpdatedAt": "2023-05-19T18:09:25.885555-03:00",
		"DeletedAt": null,
		"SiteID": 2199,
		"Ref": {
			"ID": 2199,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"Url": "https://www.mensagensdebomdia.com.br/",
			"Body": ""
		},
		"Value": "inicio",
		"Counter": 1
	}
]

function seach(input) {
    let texto = input.value.trim();
    let link_enfeite = 'http://127.0.0.1:5500/search?q=' + encodeURIComponent(texto).replace(/%20/g, '+')
    let link = 'http://127.0.0.1:5500/search?q=' + encodeURIComponent(texto).replace(/%20/g, '+')
    
    fetch('https://jsonplaceholder.typicode.com/todos/1')
    .then(response => response.json())
    .then(data => {
        resposta['linkpesquisa'] = link_enfeite
        console.log(resposta);
        let jsonData = JSON.stringify(resposta)
        let base64Data = btoa(jsonData)
        window.location.href = "result.html?q=" + encodeURIComponent(base64Data)
    })
    .catch(error => {
        console.log(error)
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

