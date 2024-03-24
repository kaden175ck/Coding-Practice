import React from "react";



// but you can always do regular function like
// function CountButton(){}
// here I use arrow fucntion to create a component
const CountButton = () => {

    const currentCount = 10

    return <div>CountButton
        <button>+1</button>
        <div>{currentCount}</div>
        {/* we want this 0 to increment  */}
    </div>
    // all these components function must return JSX



}

export default CountButton