//document.write("Hello World! You are in Script.js file");
//document.write("Never too late");
//alert("this file is working!");

//var num1 = window.prompt("enter a number1");
//var num2 = window.prompt("enter a number2");

// this won't work because prompt box convert it to strings you will be doing string concatenation
// document.write(num1+num2);


//num1=parseInt(num1);//only take int, if enter 5.5 it will be 5
//num2=parseInt(num2);
//document.write(num1+num2);
// so if you are doing 5.5+5, it will give 10
// if you want to work with int, ou do parseFloat



// Working with ARRAYS
var fruits = new Array("Apples", "Oranges", "Peaches");
document.write(fruits);



document.write("<br>");

var sports = ["BB",
                "STR", 
                24, 
                false];
document.write(sports);

document.write("<br>");
document.write("0th element in array:"+ sports[0]);

document.write("<br>");
document.write(sports.length);// give you 4

document.write("<br>");
sports[0]="AA";
document.write("change BB to :"+ sports[0]);


document.write("<br>");
document.write("Never too late");

document.write("<br>");
// this is string var
var people = "Alice, Drake, PeopleRand";
// string function to convert it into arrays
people = people.split(",");
document.write("0th element in people arrays "+people[0]);


// Using functions
function sayHi(firstname1, age){
    var name2 = "Joe";

    document.write("<h1>Hello Hao and " + name2 + " " + firstname1 + "</h1>");
    document.write("<p>You are "+age+"</p>");
    // alert("Page say Hey");
}

//calling the function
sayHi("Biden", 23);
sayHi("Jacker", 32);


// have function return values to you
function addition(numa,numb){
    return numa + numb;

}

var addedNumbers = addition(4,100);


document.write(addedNumbers);

document.write("<br>");
document.write("<br>");
document.write( addition(4,5) );