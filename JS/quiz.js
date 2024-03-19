document.write("Welcome!");
var questions = [
    // this question array is just question object
    // that store question and answer 
    {
        prompt: "what color are apples\n(a)Red\n(b)Green\n(c)Purple",
        answer: "a"
    },

    {
        prompt: "What color are Bananas\n(a)Teal\n\(b)Yellow\n(c)Orange",
        answer: "b"

    },

    {
        prompt: "What color are strawberries\n(a)Red\n\(b)Green\n(c)Blue",
        answer: "a"

    }
]

var score = 0;
for (var i=0; i<questions.length; i++){
    var response = window.prompt(questions[i].prompt);
    if(response == questions[i].answer){
        score++;
        alert("you right!!!");
    }
    else{
        alert("you got it wrong!");
    }
}

alert("you got "+ score + "/"+ questions.length);