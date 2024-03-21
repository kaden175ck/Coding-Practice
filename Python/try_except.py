print("comments are fun")

"""
yes this code will not show up in the output
"""

# here are example of how to use try and except

# value = 10/0 
# this will cause an error because you can't divide by zero

try:
    value = 10/0
    number = int(input("Enter a number: "))
    print(number)
except ZeroDivisionError:
    print("Divided by zero")
except ValueError:
    print("Invalid input")

# except:
#     print("All errors")
#  this is not good practice
# if you only use except:...it will catch all error
    
# # you can also store the error as a variable
# except ZeroDivisionError as err:
#     print(err)