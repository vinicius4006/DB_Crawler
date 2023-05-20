function show_results(data) {
    history.pushState(null, null, data['linkpesquisa']);
    console.log(data);
}

var urlParams = new URLSearchParams(window.location.search);
var base64Data = urlParams.get("q");
var jsonData = atob(base64Data);
var data = JSON.parse(jsonData);

show_results(data)