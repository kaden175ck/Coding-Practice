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

