// var myArray = [] create an array
// var fruits = ["Apples", ""]
// Object has similar goals, but Object storing pairs
var person = {
    //define key value pairs, like python dictionary
    name: "Hao",
    age: 23,
    isMale: true,
    skill: "code"
    // key must be unique
};


document.write(person);
document.write("<br>");
document.write(person.age);



person.age = 32
document.write("<br>");
document.write(person.age);
document.write("<br>");



document.write(person.isMale);
document.write("<br>");

var phrase = "Strings are cool";
// this string is actually AN OBJECT 
document.write(phrase.length);



// you can also put function inside object
var person2 = {
    //define key value pairs, like python dictionary
    name: "Hao2",
    age: 232,
    isMale: true,
    skill: "code2",
    // key must be unique

    printName: function(){
        document.write("<h1>" + this.name + "</h1>");
    },

    printAge: function(){
        document.write("<h1>" + this.age + "</h1>");
    }

};



person2.printName();
person2.printAge();


// Real world object
document.write("<br>");

var myMovie = {
    title:"The Social Network",
    releaseYear:"2010",
    duration: 2.0,

    // create an array of actors objects
    actors:[
        {
            name:"JE",
            brithday: new Date("July 5, 2000"),
            hometown:"NYU"
        },
        // create another object for another actor
        {
            name:"KD",
            brithday: new Date("July 8, 2000"),
            // the Date constructors don't take "8th or 5th"
            hometown:"UCL"

        },
    ]

}

document.write(myMovie);
document.write("<br>");
document.write(myMovie.title);
document.write("<br>");
document.write(myMovie.duration);
document.write("<br>");
document.write(myMovie.actors[0].name);
document.write("<br>");
document.write(myMovie.actors[1].hometown);
document.write("<br>");
document.write(myMovie.actors[1].brithday);



