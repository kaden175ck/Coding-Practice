import React, { useState } from "react";



// but you can always do regular function like
// function CountButton(){}
// here I use arrow fucntion to create a component
const CountButton = () => {

    // console.log(useState(0))
    // after we call usefstate we get an array back [0,f]
    // a 0 which is our var
    // a function f

    // let currentCount = 0 // this is a value directly link to UI, piece of state
    // whenever we have this piece of state we want to use react hooks
    // so you import react hook, this is given by react library by default


    // we make this a const again
    
    const [currentCount, setCurrentCount] = useState(10)
    // function f will be stored in setCurrentCount
    // setCurrent is a function which we tell React to update the current count
    // 0 stored in currentCount
    // if we do let currentCount = 0, react just think this is a var, and won't do anything
    // but if you do the above one, basically just register currentCount with React
    // you are saying: Hey react, currentCount is a piece of state that I want to keep track


    const handleClick = () => {
        // currentCount++
        // now instead of doing ++,you just call setCuurentCount fucntion
        // setCurrentCount, a way to update this piece of state
        setCurrentCount(currentCount + 1)


        console.log(currentCount)
        // alert("Woeking")
    }
    // this is updating in the back but not showing on UI
    // we want to tell react directly
    // react isn't updating the UI, only do that when we explicitly tell so
    // this is what React Hook come in play!!!!!
    // you tell hook I update the coruentCount and I want you to update on drwoswer


    // usestate



    console.log("components re-rendered")
    // everytime you update the state, this whole page(this whole file will get execute agin)

    return <div>CountButton
        {/* <button onClick={
            () => {alert("Clicked")}
        }>+1</button> here I did arrow function,but you can do whatever */}

        <button onClick={handleClick}>+1</button>
        {/* onClick is called props or attribute  */}


        <div>{currentCount}</div>
        {/* we want this 0 to increment  */}
    </div>
    // all these components function must return JSX



}

export default CountButton