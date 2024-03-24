from multiprocessing.connection import answer_challenge
from question_class import Question


questions_show = [
    "what color are apples?\n(a) Red/Green\n(b) Purple\n(c) Orange\n\n",

    "what color are bananas?\n(a) Teal\n(b) Magenta\n(c) Yellow\n\n",
    
    "what color are strawberries?\n(a) Yellow\n(b) Red\n(c) Blue\n\n"
]


# creat question object
questions = [
    Question(questions_show[0],"a"),
    #this is like Question(prompt, answer)
    Question(questions_show[1],"c"),
    Question(questions_show[2],"b"),
]


def run_test(questions):
    score = 0
    for question in questions:
        user_response = input(question.prompt)
        if user_response == question.answer:
            score +=1
    print("you got "+ str(score)+ "/"+ str(len(questions_show)))
        
            
run_test(questions)