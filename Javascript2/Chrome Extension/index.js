// function saveLead(){
//     把这块全部取消掉了，html里面也就不需要 onclick那个属性，也取消掉了


//     console.log("save button clicked from onclick attribute")
// }

const saveBtn = document.getElementById("save-btn")

let myLeds = [] //因为数组我们会reassign所以不能用const
//测试保存
// let myLeds = `["awesomepho.cn"]`
// myArrLeads = JSON.parse(myLeds)
// myArrLeads.push("new value")
// myLeadsStr = JSON.stringify(myArrLeads)
// console.log(typeof myLeadsStr)


// you want to save some old leds, that maybe youwant 
//to bring them up sometimes, that's why you
// need to REFACTOR the function render
let oldLeds = []


const inputEl = document.getElementById("input-el")
// 可以直接用let，但是一般用const，const不能被reassign


const unEl = document.getElementById("ul-el")

// const web = "awesome.com"


//测试保存在本地的问题，笔记里有提到
// localStorage.setItem("myLeads","www.kaden.com")
// console.log(localStorage.getItem("myLeads"))
// localStorage.clear()


// get the leads from local storage,Json parse, store it in var
// log out the var
//localStorage.clear()
const localStorageLeds = JSON.parse( localStorage.getItem("myLeds") )//change this to const
//but let also work, its just good habit 
const deleteBtn = document.getElementById("delete-btn")
console.log(localStorageLeds)//这个一开始肯定会显示null如果之前删除了全部的话


// tab btn
const tabBtn = document.getElementById("tab-btn")
// const tabs = [
//     {url: "https://kaden175ck.com"}
//     //测试tab button是否完善,可删
// ]

tabBtn.addEventListener("click",function(){
    console.log("tab button is clicked")

    //我们要保存的是当前页面的url，这里接口chrome API
    chrome.tabs.query({ currentWindow: true, active: true }, function (tabs) {
        //console.log(tabs)

        //console.log(tabs[0].url)// this only print out the url in the dev console


        // save the url in localstorage, first we need to push into my leads array
        myLeds.push(tabs[0].url)

        //这行和下面的一模一样
        localStorage.setItem("myLeds", JSON.stringify(myLeds))//myleds 是key，然后我们把array 变成string
        render(myLeds)
        
        
    })
    


    // 下面这十行代码我们全部放到了上面的函数里
    // console.log(tabs[0])
    // this is essentially a Object, and url is the Key, we want to log out the value only
    //console.log(tabs[0].url)// this only print out the url in the dev console

    // save the url in localstorage, first we need to push into my leads array
    //myLeds.push(tabs[0].url)

    //这行和下面的一模一样
    //localStorage.setItem("myLeds", JSON.stringify(myLeds))//myleds 是key，然后我们把array 变成string
    //render(myLeds)



})




// if(localStorageLeds){
//     myLeds = localStorageLeds
//     renderLeads()
// }


//所以我们这个也要传一下参数
if(localStorageLeds){
    myLeds = localStorageLeds
    render(myLeds)
}



// need to be refactor, what if we need to recall some
// old savings?
// 上面传进的参数myLeds会变成这个leads
//至于这里面的参数名称，根本不重要，叫什么都行，yolo==leads，象叫什么都可以
function render(leads){
    let listItem = ""
    
    for (let i = 0; i<leads.length; i++){

        // listItem  +=  "<li><a href='#'>"  + myLeds[i] +    "</a></li>"
        // #意思是空的，一次一次迭代
        // listItem  +=  "<li><a target = '_blank' href=' " + myLeds[i]   +   "'>"   + myLeds[i] +   "</a></li>"

        listItem  +=  `
        <li>
            <a target = '_blank' href='${leads[i]}'>
                ${leads[i]}
            </a>
        
        
        </li>
        `




        // console.log(listItem)

        // unEl.innerHTML = unEl.innerHTML + "<li>"+ myLeds[i] + "</li> "
        // console.log(myLeds[i])
        // 原本这里是textcontent现在，这两个可以换，反正都是更改界面内容的


        // console.log()


        // 个人认为上面的这个方法更容易理解，但是下面这几个方法也要知道。
        // const li = document.createElement("li")
        // li.textContent = myLeds[i]
        // unEl.append(li)

    }

    unEl.innerHTML = listItem

}







deleteBtn.addEventListener("dblclick",function(){
    console.log("double clicked!")
    localStorage.clear()
    myLeds = []
    render(myLeds)
})

saveBtn.addEventListener("click", function() {
    myLeds.push(inputEl.value)

    // console.log(myLeds)


    // console.log("save button clicked from eventlistner")

    inputEl.value = ""

    //保存实操
    localStorage.setItem("myLeds", JSON.stringify(myLeds))




    render(myLeds)

    //测试保存
    console.log(localStorage.getItem("myLeds"))

})



// need to be refactor, what if we need to recall some
// old savings?
// 现在把他挪上去，传入参数，重新refacor让他reuseable
/*function renderLeads(){
     let listItem = ""
    
     for (let i = 0; i<myLeds.length; i++){

         // listItem  +=  "<li><a href='#'>"  + myLeds[i] +    "</a></li>"
         // #意思是空的，一次一次迭代
         // listItem  +=  "<li><a target = '_blank' href=' " + myLeds[i]   +   "'>"   + myLeds[i] +   "</a></li>"

         listItem  +=  `
         <li>
             <a target = '_blank' href='${myLeds[i]}'>
                 ${myLeds[i]}
             </a>
        
        
         </li>
        `




        // console.log(listItem)

        // unEl.innerHTML = unEl.innerHTML + "<li>"+ myLeds[i] + "</li> "
        // console.log(myLeds[i])
        // 原本这里是textcontent现在，这两个可以换，反正都是更改界面内容的


        // console.log()


        //个人认为上面的这个方法更容易理解，但是下面这几个方法也要知道。
        // const li = document.createElement("li")
        // li.textContent = myLeds[i]
        // unEl.append(li)

     }

     unEl.innerHTML = listItem

} */
