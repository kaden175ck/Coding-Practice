var header = document.getElementById("header");
header.innerHTML = "The Java header got Overridden by getelementId"
header.style.color = "red";
header.style.backgroundColor = "blue";

var link = document.getElementById("link");
link.href = "page.html"
document.write("<br>");
document.write("Above link got overridden to page, instead of index")
// this above override the starting href, 
// it was going to index
// but now go to page
document.write("<br>");
link.style = "color: green;"


function handleClick(element){
    element.innerHTML = "Clicked"
    element.style="background-color:blue;"
}


// add hover effect on image all in JS
// function pass on par
var image = document.getElementById("image");
image.addEventListener("mouseover", function(){
    this.style = "box-shadow: 10px 10px 10px yellow";
    this.width = "400"

});

// so that it zoom in and zoom out when mouse hover
image.addEventListener("mouseout", function(){
    this.style = "";
    this.width = "300"

});