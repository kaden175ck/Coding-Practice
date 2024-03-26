import React from 'react'
import CountButton from './CountButton'
// render the countbutton inside the App component


const App = () => {
    // const myItem = "item3" 

    return (
        // <ul>
        //     <li>item111155asdas</li>
        //     <li>item222255{myItem.toUpperCase()}</li>
        //     <li>{myItem}</li>
        // </ul>
        <div>
            app
            {/* you can also use the traitional HTML way:<CountButton></CountButton> */}
            

            
            {/* <CountButton />  */}
            {/* this is what called standalone component  */}
           
           
            {/* this is how you pass props: */}
            <CountButton incrementBy = {5} /> 
            {/* you will see the empty object will 
            have values: incrementBy 5 */}

           
            app
            <CountButton incrementBy = {1}/> 
            {/* this is how you reuse component easy af,
            and they are independent of each other */}

        </div>
    )
}


export default App
// this make the code avaiable to other files in src folder


// App.js just like the body tag in html,
// its like the core content of your project
// this is like a wrapper
// everything we create will be rendered inside here