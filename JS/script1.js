document.write("<h2 style='color: blue;'>Content of JS!</h2>");
document.write("<hr/>")

//variables in JS
var phrase = "To be or not to be, but never give up!";
document.write(phrase);
document.write("<br>");
document.write(phrase);
document.write("<br>");
var phrase = 'Last Dance';


// define few var
var isMale = true;
var age = 23;
var age = 3.4;

// Note: there is a difference between null and undefined
// null means no value!
// undefined means the var doesn't have a value YET!
var flaws = null;
var description = undefined


document.write(phrase);
document.write("<br>");
document.write("Never Give Up");
document.write("<br>");
document.write('single');
document.write("<br>");
var phrase = 'Last Dance Dance?';
// remember space and char will also count as length
document.write(phrase);
document.write("<hr/>");
document.write("<br>");
document.write(phrase.length);
document.write("<br>");

document.write(phrase.toUpperCase());
document.write("<br>");

document.write(phrase.charAt(5));
document.write("<br>");
document.write(phrase.indexOf("L"));
// if repeated symbols, indexOf will show the first one
document.write("<br>");
document.write(phrase.indexOf("D"));

// it will give you -1 if char is not appeared
document.write("<br>");
document.write(phrase.indexOf("Z"));


// this will give you the last time it appeared
// like in phrase, D appeared twice, this method give the last one index
document.write("<br>");
document.write(phrase.lastIndexOf("D"));
document.write("<br>");
document.write("<hr/>");

//--------substring---------------
document.write(phrase);
document.write("<br>");

// it will only print index 012, and stop at 3
// var phrase = 'Last Dance Dance?';
document.write(phrase.substring(0,3));
document.write("<br>");
document.write(  phrase.substring(phrase.indexOf("L"), phrase.length)  );

document.write("<br>");
document.write(phrase.endsWith("D"));
document.write("<br>");
document.write(phrase.endsWith("?"));
document.write("<br>");
document.write(phrase.endsWith("Dance?"));
document.write("<br>");
document.write(phrase.includes("Last"));



document.write("<br>");
document.write("<hr/>");
var text = 'ARE COOL';
document.write(phrase.includes(text));
document.write("<br>");
document.write(text.includes("COOL"));
document.write("<br>");
// concatenation of strings
document.write(phrase+text);
document.write("<br>");
document.write(phrase+ ' ' + text);


/* working with nums*/
document.write("<br>");
document.write(2.2345);

document.write("<br>");
document.write(2*(2+7));




document.write("<br>");
var number = 6;
document.write(number*2);


document.write("<br>");
document.write(number % 5);


document.write("<br>");

var number = -15;
document.write(Math.abs(number));




document.write("<br>");
document.write(Math.max(3,number));






document.write("<br>");
document.write(Math.min(3,number));


document.write("<br>");
var number=0.9;
document.write(Math.round(number));
// standard rounding rules




document.write("<br>");
var number =2;
document.write(Math.pow(number,3));



document.write("<br>");
document.write(Math.sqrt(number));



document.write("<br>");
document.write(Math.random());
// return rand number between 0 and 1, decimal number



document.write("<br>");
document.write(Math.random()*10);

document.write("<br>");
// this will give you rand integer 0-10, zero and ten
document.write(Math.round ( Math.random()*10)  );


//______________________________________________________
document.write("<br>");
var name =window.prompt("Please enter your name: ");
var age =window.prompt("How old are you: ");
// javascript take this as string remember! 
// so if you try to add 2 numbers, it will just give you 2+1=21
document.write("Hey " + name + " How are you?"+"your age is"+age);
document.write("<br>");
//______________________________________________________
// basic calculator
document.write("<br>");
var num1 =window.prompt("enter the first num: ");
var num2 =window.prompt("enter the second num: ");

//this is the conversion you need to do math cal,
// it will convert to integer
// will not work for float
//num1 = parseInt(num1);
//num2 = parseInt(num2);
 num1 = parseFloat(num1);
 num2 = parseFloat(num2);

document.write(num1+num2);
document.write("<br>");
document.write("<br>");
document.write("<br>");
document.write("<br>");
document.write("<br>");
document.write("<br>");
document.write("<br>");
