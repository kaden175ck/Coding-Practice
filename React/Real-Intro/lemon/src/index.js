

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

