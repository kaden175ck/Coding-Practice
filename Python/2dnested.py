# this number grid list
number_grid = [
    [1,2,3],
    [4,5,6],
    [7,8,9],
    [0]
]
print(number_grid[0][0])
print(number_grid[0][1])
print(number_grid[1][1])

# the row or col name doesn't matter, python don't care

for row in number_grid:
    for col in row:
        print(col)


for a in number_grid:
    for b in a:
        print(b)

# basic translator
# all vowels become g
# dog --- dgg
# cat --- cgt

def translate(phrase):
    translation = ""
    for letter in phrase:
        if letter in "AEIOUaeiou":
            translation = translation + "g"
        else:# if the letter is not in aeiou,we add that letter to empty translate string
            translation = translation + letter
    return translation#jump out the loop when no letter in phrase


print(translate(input("Enter a phrase: ")))



# here is the imporve version
# def translate(phrase):
#     translation = ""
#     for letter in phrase:
#         if letter.lower() in "aeiou":
#             translation = translation + "g"
#         else:
#             translation = translation + letter
#     return translation
# print(translate(input("Enter a phrase: ")))


# try the "On" exmaple! suppose to give you Gn not gn
# 2nd imporve version
def translate(phrase):
    translation = ""
    for letter in phrase:
        if letter.lower() in "aeiou":
            if letter.isupper(): 
# deal with the situ that first letter is capital
                translation = translation + "G"
            else:
                translation = translation + "g"
        else:
            translation = translation + letter
    return translation

print(translate(input("Enter a phrase: ")))
