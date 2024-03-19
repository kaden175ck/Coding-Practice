for( var i =0; i<4 ; i++){
    for( var j=0; j<3; j++){
        document.write("i = ", + i+" j= "+j+"<br>");
    }
}
document.write("<br>");
document.write("<hr/>");


// 2d arrays
var numberGrid = [
    [1,2,3],
    [4,5,6],
    [7,8,9],
    [0]
];
// how you access
document.write(numberGrid[1][1]);
document.write("<br>");
document.write(numberGrid[0][0]);
document.write("<br>");
document.write(numberGrid[2][2]);
document.write("<br>");
document.write("<hr/>");

for(var x =0; x< numberGrid.length; x++){
    for( var y=0; y < numberGrid[x].length; y++){
        document.write(numberGrid[x][y]);
    }
    document.write("<br>");
}

document.write("<br>");
document.write("<hr/>");

// for each loop is JS

var friends = ["OGA","OGB","OGC"];


for( var i =0; i<friends.length; i++){
    document.write(friends[i]+"<br>")
}
document.write("<br>");
document.write("<hr/>");

// this for function print element three times
friends.forEach(function(){
    document.write("element"+"<br>")
});

document.write("<br>");
document.write("<hr/>");

// this for each gunction do the same thing as the for loop above
friends.forEach(function(element){
    document.write(element+"<br>")
});
// you can name enything you want, not strit elemnet
document.write("<br>");
document.write("<hr/>");

friends.forEach(function(friend){
    document.write(friend+"<br>")
});


//  array of books objects
var books =[
    {
        title: "Harry Potter1",
        author: "JK",
        pages: 100
    },
    {
        title: "Harry Potter2",
        author: "JK",
        pages: 200
    },
    {
        title: "Harry Potter3",
        author: "JK",
        pages: 300
    }
]

books.forEach(function(book){
    document.write(book.pages + "<br>");
    document.write(book.title + "<br>");
});