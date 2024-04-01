import React from 'react'

const Button = (props) =>{

    // console.log(props)

    return <div style={{color: "blue", background: "green",

    borderRadius: "5px"}}>
    
    {props.buttonText}
    {props.children}</div>



}

export default Button