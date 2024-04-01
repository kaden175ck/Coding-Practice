import React, { useState, useEffect} from 'react'
import CountButton from './CountButton/CountButton'
// render the countbutton inside the App component
import Button from "./Button/Button"


// everytime you create a component
// you need to import here in APP.js
import SearchBar from './SearchBar/SearchBar'

// const products = 


const App = () => {
    // const myItem = "item3" 

    // create a piece of state:
    // inside usestate you pass the initial/default value
    // in this case just empty array
    const [productsState, setProductsState] = useState([])
    // we create a piece of states called productsState
    // initially just an empty array


    useEffect(()=>{






        fetch('https://fakestoreapi.com/products')
        .then(res=>res.json())
        .then((productsArray)=> {
            // console.log(json)
            // this map function will map items in one array to another array
            const newProductsState = productsArray.map(
                (product)=>{
                    return product.title
                    // so this newProductsstate will be an array
                    // that only contain title
                }
            )
            setProductsState(newProductsState)
        
        })











        // console.log("components mounted")},[]
        setTimeout( () => {
                setProductsState([
                    "tootha",
                    "toothb",
                    "toothc",
                    "brushabc",
                ])

            }, 2000 // after 2000ms, this code will exec
            // when you load the page you will see its blank
            // after 2 sec all these will show up
            // this is just simulating amount of time
            // that you might need to ftech external data
            // maybe large datasets
        )
    },[]
        // I want productstates don't load untill user 
        // actually seeing something
    )


    const hasProdcuts = productsState.length > 0


    return (
        // <ul>
        //     <li>item111155asdas</li>
        //     <li>item222255{myItem.toUpperCase()}</li>
        //     <li>{myItem}</li>
        // </ul>
        <div>
            last button that use child, not so common
            <Button buttonText="My button text"/>
            <Button buttonText="My button text2"/>
            <Button buttonText="My button text3"/>
            <Button>Children button here</Button>
            <Button>Children button here2</Button>
            <Button>Children button here3</Button>
            {/* this is just another type of props called children props */}


            app
            {/* you can also use the traitional HTML way:<CountButton></CountButton> */}
            

            
            {/* <CountButton />  */}
            {/* this is what called standalone component  */}
           
           
            {/* this is how you pass props: */}
            <CountButton incrementBy = {5} buttonColor = {'red'}/> 
            {/* you will see the empty object will 
            have values: incrementBy 5 */}

           
            app
            <CountButton incrementBy = {1} buttonColor = {'yellow'}/> 
            {/* this is how you reuse component easy af,
            and they are independent of each other */}


    {/* when these products come from external source we need useEffect */}


      {/* or you can call this conditonal rendering       */}
{/* this will show loading then 2sec later stuff will rendered done */}
{hasProdcuts ? <SearchBar propsProducts={productsState} />
            : 'Loading'}

            
            {/* we passed the props that takes the products */}
               {/* <SearchBar propsProducts={ */}
            {/* //     productsState
            //     // ["tootha",
            //     // "toothb",
            //     // "toothc",
            //     // "brush",

            //     // ]
            // }
            // /> */}

            <SearchBar 
                propsProducts={[
                "bike",
                "bike lock",
                "bike helmet",
                ]}
            />
        </div>
        
    ) 
}


export default App
// this make the code avaiable to other files in src folder


// App.js just like the body tag in html,
// its like the core content of your project
// this is like a wrapper
// everything we create will be rendered inside here