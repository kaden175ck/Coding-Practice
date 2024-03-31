import React, { useState, useEffect } from "react";
import "./CountButton.css"


// but you can always do regular function like
// function CountButton(){}
// here I use arrow fucntion to create a component




// remember this isjut normal js function that returtn JSX


// you could also include props const CountButton = (props) => {
// to get more custmoized features

const CountButton = (props) => {
    //console.log(props) // you will see empty object
    // these props are values we can use inside this component
    // create props/pass props in App.js


    // now over here you could say
    // this will just give you numbers not showing the english word

    // console.log(props.incrementBy)


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
        setCurrentCount(currentCount + props.incrementBy)


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


    // now whenever we want to style the component,
    // you need to pass the style object
    const divStyle = {
        color: "blue",
        /* this is an object, 
        as you can see this is not CSS you see often*/
        border: "1px solid black",
        borderRadius:"10px",
    }
    
    const buttonStyles = {
        background: props.buttonColor,
        borderRadius:"10px",
    }


    useEffect(() => {
        console.log("this will only called once say hey is rendered!")
       // the reseaon why this is only called once because no content in array
    }, [])


    

    // its part of the react lifecycle
    useEffect(() => {
        console.log(currentCount)
        if(currentCount === 20){
            alert("Count is 20"	)
            // setCurrentCount(0)
        }
        console.log("useEffect function called and currentcount is update")
        // this will run everytime the component is rendered
    }, [currentCount])
    // this array is called dependency array
    // [currentCount, secondpieceof states...]





    return <div style = {divStyle}>CountButton
        {/* <button onClick={
            () => {alert("Clicked")}
        }>+1</button> here I did arrow function,but you can do whatever */}


        {/* now instead of say +1, you could say plus incrementBy
        <button onClick={handleClick}>+1</button> */}
        {/* in order to let the button not show +1 */}
        <button style = {buttonStyles} onClick={handleClick}>+{props.incrementBy}</button>
        {/* onClick is called props or attribute  */}


        <div className = {'count-display'}>{currentCount}</div>
        {/* we want this 0 to increment  */}
    </div>
    // all these components function must return JSX



}

export default CountButton