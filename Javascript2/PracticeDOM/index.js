let data = [
    {
        player: "Jane",
        score: 52
    }, 
    {
        player: "Mark",
        score: 41
    }
]

// Fetch the button from the DOM, store it in a variable
// Use addEventListener() to listen for button clicks
// Log Jane's score when the button is clicked (via data)

const logBtn = document.getElementById("log-btn")

logBtn.addEventListener("click",function(){
    console.log(data[0].score)
})





// task

// The generateSentence(desc, arr) takes two parameterer: a description and an array.
// It should return a string based upon the description and array.

// Example 1: if you pass in "largest countries",and ["China", "India", "USA"],
// it should return the string: "The 3 largest countries are China, India, USA"

// Example 2:If you pass in "best fruits" and ["Apples", "Bananas"], it should return:
// "The 2 best fruits are Apples, Bananas"

// Use both a for loop and a template string to solve the challenge
// function generateSentence(desc, arr) {

//     let baseString = `the ${arr.length} ${desc} are `
    
//     for (let i=0;i<arr.length;i++){
        
//         baseString += `${arr[i]} ,`  
//     }
//     这是几乎正确的但是我们还要避免最后的都好，所以要检查最后一位

//     return baseString
    
// }

// const result = generateSentence("largest countries",["China", "India", "USA"])
// console.log(result)
// generateSentence("best fruits",["Apples", "Bananas"])








function generateSentence(desc, arr) {

    let baseString = `the ${arr.length} ${desc} are `
    const lastIndex = arr.length -1
    
    for (let i=0;i<arr.length;i++){
        if(i === lastIndex){
            baseString += `${arr[i]} `
        }
        else{
            baseString += `${arr[i]}, `

        }
    }

    return baseString
    
}

const result = generateSentence("largest countries",["China", "India", "USA"])
console.log(result)
// generateSentence("best fruits",["Apples", "Bananas"])