// rememer before we include those react script in html
// when we want to simulate the react app
// now the pro way after you have react env installed
// you want to import library here

import React from "react"
import ReactDOM from "react-dom"


// you see the old good friend root div in public/index.html
// so you already know we did exact same thing in vanilla js
const reactContentRoot = document.getElementById("root")


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


ReactDOM.render(<App />, reactContentRoot)

