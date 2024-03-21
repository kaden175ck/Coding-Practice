// this is the thought process of why we need react?
// why viniall dom or vinaill js is not enough
// // alert("Hello!");
// // console.log(document)


// const incrementBtn = document.getElementById('increment-btn')
// const countDisplay = document.getElementById('count-display')

// // countDisplay.innerText="1"
// // incrementBtn.innerText="+9"

// // console.log(incrementBtn)
// let currentCount = 0 // state!!! actively keeping track of
// // arrow func !!!
// incrementBtn.addEventListener('click',()=>{
//     // alert("you click")
//     currentCount++


//     //update the DOM
//     countDisplay.innerText=currentCount
// })

// const incrementBtn2 = document.getElementById('increment-btn2')
// const countDisplay2 = document.getElementById('count-display2')

// let currentCount2 = 0

// incrementBtn2.addEventListener('click',()=>{
//     currentCount2++
//     countDisplay2.innerText=currentCount2
// })
// // but you double your source code, this is why you
// // need to learn React

// // react allow ou to use this over and over agin





// // here you use the react script tag and redo everything
// alert("Hello!");
// console.log(React);
// // this is like react library, core logic of all this
// console.log(ReactDOM);
// // this allow react to render to the DOM, you also have react native

const reactContentRoot = document.getElementById("root")

ReactDOM.render("Hello World", reactContentRoot)
// the firs arg is what you want to render, 
// the second arg is where you want to render it
// what react do is he comes to the HTML file
// amd he enter the content for the div block,
// he put <div id="root">Helloworld</div>
// ReactDom will render the first arg inside the second arg


// ReactDom can render things onto the DOM,
// but how do we create things
// we use React Api to create things
// React give us this function createElement
// kind of like HTML creating elements
const myFirstElement = React.createElement('li', null,'item1')
// 1 the name of the HTML element you want to create
// 2 the props you want to pass to the element
// 3 the context  or the text

ReactDOM.render(myFirstElement, reactContentRoot)

const myUnorderList = React.createElement('ul', null,
    React.createElement('li', null,'item1'),
)