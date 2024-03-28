import React, {useState} from 'react';
import './SearchBar.css';


// craete out component/function
const SearchBar = () => {

    const[SearchValue, setSearchValue] = useState("The Search Value")
    // all these info/user input on searchbar
    // will be stored in SearchValue piece of state




    // const SearchValue = "The Search Value"

    const handleInputChnage = (event) => {
        setSearchValue(event.target.value)

        // console.log(event.target.value)
        // alert("Changed")
    }

    return(

    <div>
        Your Search Bar: 
        <input type="text" value={SearchValue} onChange={handleInputChnage} />

        {SearchValue} 
        {/* SearchvAlue is stored as piece of state */}

    </div>
    // in HTML we take input like 
    // <input type="text" value="123" />


    )
}

export default SearchBar