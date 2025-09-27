//js func
export default function Dosomething () {
    console.log("Hello World");
}

//js arrow func
export const Dosomething2 = () => {
    console.log("Hello World");
}

// react com
const MyComponent = () => {
    return (
        <div> 
            <h1>My Component</h1>
        </div>
    )
}

// 无名函数
<button onClick={()=>{
    console.log("123")
}}>
</button>


// ternary operator
if(true){

} else{

}


let age =10
let name = "H"

let name = age>10 && "H"
let name =age >10 ? "H" : "AS"

// objects/dictionary

const person ={
    name:"das"
    age:21
}

//decontrsuct object
const name= person.name

// deconstruct objects
const {name, age} =person

//....
const person2={...person, name:"Jack"}



//often use function
let names =["123","456","abd"]

names.map((name)=>{
    console.log(name)
})

names.map((name)=>{
    return "joe"
})


names.map((name)=>{
    return name +"1"
})

//react
names.map((name)=>{
    return <h1> {name} </h1>
})

names.filter((name)=>{
    return name !== "123"
})


//async +await +fetch 这三个是现代 JavaScript 处理「异步操作（asynchronous）」的基础，尤其是「向服务器请求数据」时几乎必用。
fetch() 是浏览器内置的一个函数，用来向网络发送 HTTP 请求（例如向某个 API 获取数据）。它返回的是一个 Promise 对象（这是 JS 异步的基础机制）。
// fetch('https://jsonplaceholder.typicode.com/todos/1')fetch(...) → 马上返回一个 Promise（不是结果本身），所以后面的代码不会等它。
  .then(response => response.json()) // 把返回的 JSON 数据解析出来.then() 会在数据回来后才执行。
  .then(data => console.log(data))   // 拿到结果,.then() 会在数据回来后才执行。
  .catch(error => console.error('出错了:', error));



async声明异步函数
  在函数前面加上 async 关键字，表示这个函数里可以用 await，而且这个函数本身会自动返回一个 Promise。
  async function getData() {
  return "Hello"; 
}

getData().then(result => console.log(result));
// 输出: Hello
// getData() 返回的不是字符串本身，而是一个 Promise，所以需要 .then() 来拿到结果



await 必须等待 Promise 完成，拿到结果
// await 只能在 async 函数内部使用。
它的作用就是「暂停」函数的执行，直到 Promise 有了结果，然后再继续往下走。看起来像同步，但实际上是异步。更像同步流程。
async function getTodo() {
  const response = await fetch('https://jsonplaceholder.typicode.com/todos/1');
  const data = await response.json(); // 解析 JSON 也需要时间
  console.log(data);
}

getTodo();
