var passwd = "og";
var response;
var entryCount = 0;
var error = false;
var entryLimit = 3;

while(response != passwd && !error){
    if(entryCount < entryLimit){
        response = window.prompt("enter passwd: ");
        entryCount++;
    }

    else{
        error = true;
    }

}

if(error){
    alert("too many try!");
}

else{
    alert("you got it right! ");
}

