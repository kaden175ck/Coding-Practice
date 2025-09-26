// Use .innerHTML to render a Buy! button inside the div conta

const container = document.getElementById("container")

container.innerHTML = "<button onclick= 'buy()' > BUY! </button>"

function buy(){
    container.innerHTML += "<p>Thanks For Buying!</p>"
}