def say_hi(name,age):
    print("Hello "+name+", you are "+age)
# you can also pass int, but in this case
    # if you pass int, you do str(age)
    # to covert it into string
print("top")
say_hi("Hao","23")
print("bottom")
say_hi("Kaden","20")

print("    /| ")
print("   / | ")
print("  /  | ")
print(" /   | ")
print("/____| ")

def cube(num):
    num*num*num

cube(3)
print(cube(3))
print("    /| ")
print("   / | ")
print("  /  | ")
print(" /   | ")
print("/____| ")

def cube(num):
    return num*num*num
    print("hello")# this will never be reached

cube(3)
print(cube(3))


def cube(num):
    return num*num*num

result = cube(4)
print(result)
