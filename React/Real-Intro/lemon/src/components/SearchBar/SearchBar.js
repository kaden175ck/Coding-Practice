import React, {useState} from 'react';
import './SearchBar.css';


// craete out component/function
const SearchBar = () => {

    const[SearchValue, setSearchValue] = useState("")
    // all these info/user input on searchbar
    // will be stored in SearchValue piece of state




    // const SearchValue = "The Search Value"

    const handleInputChnage = (event) => {
        // console.log(event)
        setSearchValue(event.target.value)
        // event.target.value this is the value the user input
        // and we use this user input tto setSearchValue 

        // console.log(event.target.value)
        // alert("Changed")
    }


    const handleClearClick = () => {
        setSearchValue("")
        // setSearchValue("text cleared")
    }


    // we first need a boolen value to tell us whether to show the button or not
    const shouldDisplayButton = SearchValue.length > 0

    console.log(shouldDisplayButton)


    return(

    <div>
        Your Search Bar: 
        <input type="text" value={SearchValue} onChange={handleInputChnage} />

        {/* whenever use hit a key, onChange will exe */}

        {SearchValue} 
        {/* SearchvAlue is stored as piece of state */}


        {/* conditional rendering */}
        {/* this is called short circuiting
        if first part true, then we run second part, vice versa */}
        {shouldDisplayButton && <button onClick={handleClearClick}>Clear</button>}
        {/* we just copy the below code to here,{} store js expression */}

        {/* create clear button to clear all the text */}
        {/* <button onClick={handleClearClick}>Clear</button> */}
        {/* we want to show clear button only when there is text display */}

    </div>
    // in HTML we take input like 
    // <input type="text" value="123" />


    )
}

export default SearchBar