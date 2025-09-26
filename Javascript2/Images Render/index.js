// Create a function that renders the three team images
// Use a for loop, template strings (``), plus equals (+=)
// .innerHTML to solve the challenge.

const imgs = [
    "images/hip1.jpg",
    "images/hip2.jpg",
    "images/hip3.jpg"
]

const container = document.getElementById("container")

function renderImg(){
    let imgsDom = ""
    for (let i=0;i<imgs.length;i++){
        imgsDom += `<img alt="pics" class="team-img" src="${imgs[i]}">`
        //container.innerHTML += `<img class="team-img" src="${imgs[i]}">`
        //但是这里我们用了innerHMTML三次
        //频繁操作DOM，让整体性能下降，尽量少操作DOM,所以创建新变量html

    }
    container.innerHTML = imgsDom
}

renderImg()