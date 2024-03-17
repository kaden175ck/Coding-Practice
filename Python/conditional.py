is_male = True
is_tall = True

if is_male and is_tall: 
    print("you are a male and tall")
    print("    /| ")
    print("   / | ")
    print("  /  | ")
    print(" /   | ")
    print("/____| ")

elif is_male and not(is_tall):
    print("you are a short male")

elif not(is_male) and is_tall:
    print("you are not male but tall")


else:
    print("you are not male and you are not tall")





# Comparisons
    
print("    /| ")
print("   / | ")
print("  /  | ")
print(" /   | ")
print("/____| ")

# == , !=, all comparisons op
def max_num(num1,num2,num3):
    if num1 >= num2 and num1 >= num3:
        return num1
    elif num2 >= num1 and num2 >= num3:
        return num2
    else:
        return num3

print(max_num(3,4,5))
             
