// template strings/literals

const recipient = "James"
const sender = "Kaden"
// Create a new variable, sender, and set its value to your name

// Use your sender variable instead of "Per"就是说用这种斜的符号，很直观你空行他也空
const email = `Hey ${sender}!
 How is it going? 
 Cheers Per${ recipient}`

console.log(email)