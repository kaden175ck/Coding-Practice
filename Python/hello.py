print("Hello World")
print("    /| ")
print("   / | ")
print("  /  | ")
print(" /   | ")
print("/____| ")

from math import *
char_name = "John"
char_age = "35"
print("Hello " + char_name + ", ")
print("You are "+char_age+" years old.")
char_age = "25"
print("But Kay is "+char_age+" years old.")

# char_age = 23 # store int
# char_age = 23.5 can also store float but if you want to store string put ""
# is_male = True


# string op
print("HEY\nHEY\nHEY") #escape char, new line: \n
print("Kay's site contains char like \ ")
print("if you want to print special char, you use backslash:\ ")
print("love the book name \"My Space\" ")

some_thing = "Man "
print(some_thing.isupper())
print(some_thing.islower())
print(some_thing.lower())
print(some_thing.upper())
print(some_thing + "what can I say")
print(some_thing.upper().isupper())
print(len(some_thing))
print(some_thing[0])
print(some_thing[1])
print(some_thing[2])

print(some_thing.index("a"))
print(some_thing.index("n"))
print(some_thing.index("M"))
print(some_thing.index(" "))
# if not found will return error print(some_thing.index("apple"))
print(some_thing.replace("Man","Women"))
print(some_thing.replace("n","E"))

# working with Numbers
print(10%3)
print(2*(2+3))
print(2+2*3)
print(-2+3)

my_number1 =5
my_number2 =10
print("My fav number: "+ str(my_number2))

my_negnumber = -99
print(abs(my_negnumber))
print(pow(3,2))
print(max(4,6))
print(round(3.2))
print(round(3.7))


print(floor(3.7))
print(ceil(3.7))
print(sqrt(36))