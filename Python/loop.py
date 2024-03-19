i = 1
while i<=10:
    print(i)
    i+=1


print("done with loop")

print("now you enter the guessing game")

secret_word = "OG"
guess = ""
guess_count = 0
guess_limit = 3
out_of_guesses = False # we set this to False first

while guess != secret_word and not(out_of_guesses):
    # not(False) = True, so both have to be true


    if guess_count < guess_limit:
        guess = input("Enter Guess: ")
        guess_count += 1
    else:
        out_of_guesses = True

if out_of_guesses:
    print("you are out of guesses!")
else:
    print("You win!")


# for loops
for letter in "OGGoat":
    print(letter)

friends = ["A","B","C","D"]
for friend in friends:
    print(friend)

OGfriends = ["OGA","OGB","OGC","OGD"]
for rank in OGfriends:
    print(rank)

for index in range(10):
    print(index)
# from 0 to 9!print 0-9

for index2 in range(20,25):# this print 20-24
    print(index2)

for index3 in range(len(friends)):#len(array) will give a number
    #so you are running range(3)
    print(index3)
    #this will print 0123

for index4 in range(len(friends)):#len(array) will give a number
    #so you are running range(3)
    print(friends[index4])
    # this will print friends[0],fri[1]...
    # this will give the element for all



for index5 in range(5):
    if index5 == 0:
        print("This is first iteration")
    else:
        print("not first")


