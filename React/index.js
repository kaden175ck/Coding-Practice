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

// ReactDOM.render("Hello World", reactContentRoot)
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
// const myFirstElement = React.createElement("li", null,"item1")
// 1 the name of the HTML element you want to create
// 2 the props you want to pass to the element
// 3 the context  or the text

// ReactDOM.render(myFirstElement, reactContentRoot)

// const myUnorderList = React.createElement("ul", null,
// [
//     React.createElement("li", null,"item1"),
//     React.createElement("li", null,"item2"),
// ])


// instead of doing the above one, you can do the below one
// react will convert it to the above one
// how does this happen
// this is not sth include in JS library
// this is not sth browser understand
// you use transpiler to convert this to the above one
// just gooogle babel, which basically just a JS compiler


// const myItem = "item3"
// with Babel set up you can write regular js and it embeds well with JSX


// this is the same thing as ReactCreate...
// const myJSXElement = (
//     <ul>
//         <li>item1111</li>
//         <li>item2222{myItem.toUpperCase()}</li>
//         <li>{myItem}</li>
//     </ul>
// )


// here we create a component
// I like arrow function but you can also
// use regular function like function App(){}
// you could also use classes to create components
// but its not used often
const App = () => {
    const myItem = "item3" 

    return (
        <ul>
            <li>item111155</li>
            <li>item222255{myItem.toUpperCase()}</li>
            <li>{myItem}</li>
        </ul>
    )
}


// the way we want to render this component also change
ReactDOM.render(<App />, reactContentRoot)
// this is called component tag? this is not standard JS
// think of calling function,
// you can also write App() instead of <App />



// ReactDOM.render(myJSXElement, reactContentRoot)



// now the question become how do you set up biabel
// most of the time you won't set this up
// and you will use create-react-app
// so esentially you want to get rid of the complex code
// and just use the JSX code

// the simple way is that you add in HTML
// add this to your HTML,
// this will load bibel <script type="text/babel" src="index.js"></script>



