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
var text = "ARE COOL";
document.write(phrase.includes(text));
