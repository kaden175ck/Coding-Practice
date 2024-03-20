// alert("Hello!");
// console.log(document)


const incrementBtn = document.getElementById('increment-btn')
const countDisplay = document.getElementById('count-display')

// countDisplay.innerText="1"
// incrementBtn.innerText="+9"

// console.log(incrementBtn)
let currentCount = 0 // state!!! actively keeping track of
// arrow func !!!
incrementBtn.addEventListener('click',()=>{
    // alert("you click")
    currentCount++


    //update the DOM
    countDisplay.innerText=currentCount
})











