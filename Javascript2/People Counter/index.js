//document.getElementById("counter-el").innerText=5



// let myAge = 25
// console.log(myAge)

// let humanDogRatio = 7
// let myDogAge = humanDogRatio*myAge
// console.log(myDogAge)

// let count = 0
// count =3

// console.log(count)


let countEl = document.getElementById("counter-el")

let saveEl = document.getElementById("save-el")

// console.log(countEl)

let count = 0

// console.log(saveEl)

// count = count + 1

function increment(){
    count = count + 1
    countEl.textContent = count
    // console.log(count)
}


function save(){

    let countStr = count + " - "
    saveEl.textContent += countStr
    //change this to textcontent instead of innertext may work better

    countEl.textContent = 0
    count =0
}










// let welcomeEl = document.getElementById("welcome-el")

// let name = "dasName"
// let greeting = "dsadsaG"

// welcomeEl.innerText = greeting + name

// welcomeEl.innerText += "Emoji"







