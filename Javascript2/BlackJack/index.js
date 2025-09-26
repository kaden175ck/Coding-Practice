let cards =[]

let player = {
    Name: "Player",
    Chip: 200000,
    sayHello: function(){
        console.log("Hello World!")

    }
}

player.sayHello()

let sum = 0

let hasBlackJack = false
// let isAlive = true //这里要改成false，应为游戏还没有开始，我们要判定玩家没有点击游戏开始，所以玩家是死的暂时
let isAlive = false
let message = ""

let messageEl = document.getElementById("message-el")
console.log(messageEl)

// let sumEl = document.getElementById("sum-el")
// 也可以用
let sumEl = document.querySelector("#sum-el")
/* <p id="sum-el">Sum:</p>这个id也可以改成class，然后这个querryS就用.就可以了.sum-el */
/*getElementById 和 querySelector很像*/

let cardsEl = document.querySelector("#cards-el")

let playerEl = document.getElementById("player-el")
playerEl.textContent = player.Name + " $" + player.Chip

console.log(cards)
// 注意到这里，尽管游戏没开始，但是打开console可以看到牌已经发了，这是不正常的，像是赌场里的作弊
// 所以前面的几个初始化，比如sum，isalive，firstsecond这些变量都不能是全局变量，需要修改到另外一个函数里


function getRandomCard(){
    //math random 会生成0.0000-0.9999之间的任意数，所以一般会乘以如果你是筛子就成一6，扑克牌就成一7，并且还有取floor，不然就是小数
    // return Math.floor( Math.random() *13) + 1
    let randNum = Math.floor( Math.random() *13) + 1
    if (randNum === 1){
        return 11

    }else if(randNum > 10){
        return 10

    }
    else {
        return randNum
    }

}


function startGame(){
    isAlive = true
    let firstCard = getRandomCard()
    let secondCard = getRandomCard()
    cards = [firstCard,secondCard]

    sum = firstCard + secondCard

    renderGame()
}


function renderGame(){

    cardsEl.textContent = "Cards: " 

    // cardsEl.textContent = "Cards: " + cards[0] + "," + cards[1]

    for(let i=0; i<cards.length; i++){
        cardsEl.textContent = cardsEl.textContent + cards[i] + ","
    }


    sumEl.textContent = "Sum: " + sum

    if (sum <= 20){
        message = "Do you want to draw a new card?"
    
    } else if (sum === 21){
        message = "You've got blackjack"
        hasBlackJack = true
        
    } else if (sum > 21){
        message = "you are out of the game!"
        isAlive = false
    }//or you could just do else in the last block


    messageEl.textContent = message



    // console.log(message)

// == 和===不一样
// 100 == "100" True, 所以双等意味着不严格，所以一般不建议用==在JS里面
// 100 === "100" False，所以三等意味着严格，type需要严格match
}


function newCard(){
    if(isAlive === true && hasBlackJack === false){
        let newCard3 = getRandomCard()

        sum = sum + newCard3

        cards.push(newCard3)
        
        console.log(cards)
        renderGame()

    }

    // let newCard3 = getRandomCard()
    // sum = sum + newCard3
    // cards.push(newCard3)
    // console.log(cards)

    

    

    console.log("drawing a new card from the deck")

}




















