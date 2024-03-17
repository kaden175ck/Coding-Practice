document.write("<br>");

document.write("<hr/>");


var isMale = true;
var isTall = true;

// you could also use || or.

if(isMale && isTall){
    document.write("you are male and tall");
}

// the!... is like not(...) in python
else if(isMale && !isTall){
    document.write("you are short male")
}

else if(!isMale && isTall){
    document.write("you are not male, but you tall")
}

else{
    document.write("you are not male and not tall")
}


document.write("<br>");

document.write("<hr/>");


// could use == to check same value, != not equal to
function max(num1, num2, num3){
    if(num1 >= num2 && num1 >= num3){
        return num1
    }

    else if(num2 >= num1 && num2 >= num3){
        return num2;

    }

    else{
        return num3;
    }
}

document.write(max(3,4,5));


document.write("<br>");

document.write("<hr/>");


if("phrase" == "phrase"){
    document.write("they are equal");
}



document.write("<br>");

document.write("<hr/>");

document.write("switch statement");
document.write("<br>");

document.write("<hr/>");

// function getDayName(dayNum){
//     var day;
//     if(dayNum == 0){
//         day = "Sunday";
//     }
//     else if(dayNum = 1){
//         day = "Monday";
//     }
//     return day;
// }

function getDayName(dayNum){
    var day;
    switch(dayNum){
        case 0:
            day = "Sunday";
            break;
        case 1:
            day = "Monday";
            break;
        case 2:
            day = "Tuesday";
            break;
        case 3:
            day = "Wednesday";
            break;
        case 4:
            day = "Thursday";
            break;
        case 5:
            day = "Friday";
            break;
        case 6:
            day = "Saturday";
            break;

        default:
            day = "Day doesn't exists " + dayNum + " isn't a valid";
    }

    return day;
}

document.write(getDayName(90));