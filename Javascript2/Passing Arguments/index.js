// Task1 
const welcomeEl = document.getElementById("welcome-el")

// Give the function a parameter, greeting, that replaces "Welcome back"
function greetUser(greeting, name) {
    //welcomeEl.textContent = greeting + ", " + " 👋"  +name  
    //把这个改成templete literal，
    welcomeEl.textContent = `${greeting}, ${name}👋`
    //
}

greetUser("Welcome To The Page For Template Strings!","People")





// 后面这两个全部log在dev console里了
// Task2
// Create a function, add(), that adds two numbers together and returns the sum

function add(num1, num2){
    let sum = num1 + num2
    return sum
}
console.log("These are the results for sum functions: ", add(3,4), add(9, 102))
console.log( add(3,4)    ) // should log 7
console.log( add(9, 102) ) // should log 111




// Task 3 array as argument
// Create a function, getFirst(arr), that returns the first item in the array

function getFirst(arrPar){
    return arrPar[0]
}

let result = getFirst([2,4,6])


console.log("These are the results for passing array as an argument!")

console.log(result)


// Call it with an array as an argument to verify that it works