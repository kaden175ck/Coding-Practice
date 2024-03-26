// rememer before we include those react script in html
// when we want to simulate the react app
// now the pro way after you have react env installed
// you want to import library here

import React from "react"
import ReactDOM from "react-dom"
import App from "./components/App"

import "./styles/global.css"
// the external stylesheet
// two ways to style one in file one external

// you see the old good friend root div in public/index.html
// so you already know we did exact same thing in vanilla js
const reactContentRoot = document.getElementById("root")

console.log(App)
// put below code in App.js in components folder
// this is industry standard
// const App = () => {
//     const myItem = "item3" 

//     return (
//         <ul>
//             <li>item111155</li>
//             <li>item222255{myItem.toUpperCase()}</li>
//             <li>{myItem}</li>
//         </ul>
//     )
// }


// reactdomrender render our app
// and inside the app we are able to render smalled components
ReactDOM.render(<App />, reactContentRoot)


// ReactDOM.render(<App />,document.getElementById("root"))
// a lot of people would just use the combination here
// like the above is more readable