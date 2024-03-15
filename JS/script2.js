var header = document.getElementById("header");
header.innerHTML = "The Java header got Overridden"
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