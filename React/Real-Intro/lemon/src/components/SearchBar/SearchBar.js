import React, {useState} from 'react';
import './SearchBar.css';


// have a collection of items and want to show them one by one to the user

const products = [
    "toothpaste",
    "toothbrush",
    "shampoo",
    "mouthwash",
    "floss",
    "soap",
]



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


    // console.log(
    //     products.map(
    //         (products) => {return products.toUpperCase()}
    //     )
    // )

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


        {/* we render the product here */}

       {/* <ul>you can also wrap this product map with ul</ul> */}
        {products.map(
            (product) => {
                return <div key={product}>{product}</div>
                // take all those products map them to div
                // or you can map them to list
                // key is unique identifier react is using
                // need to have this when you map

            }
        )}
       

    </div>
    // in HTML we take input like 
    // <input type="text" value="123" />


    )
}

export default SearchBar