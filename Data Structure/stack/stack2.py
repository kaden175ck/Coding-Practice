# Problem Description:

# Reverse Polish Notation (RPN) is a different way to write math expressions. For example:

# Standard: 10 + 2
# RPN: 10 2 +
# Your job is to get an input in RPN format and calculate the result.

# Example

# Input: ["2", "1", "+", "3", "*"]
# Output: 9 (Explanation: ((2 + 1) * 3) = 9)

from tkinter import E


def evaluate_rpn(expression):
    stack = []
    for token in expression: # check each "token" in the expression

        
        # Operators: When you encounter an operator: Pop the top two numbers off the stack
        if token in "+-*/": # If token is an operator

            # Pop the top two elements, b, a(Order matters!!)
            b, a = stack.pop(), stack.pop()
            
            # Apply the operation
            if token == '+':
                result = a + b
            elif token == '-':
                result = a - b
            elif token == '*':
                result = a * b
            elif token == '/':  # Ensure integer division matches problem statement
                result = int(a / b)
                
            # Push the result back onto the stack
            stack.append(result)



        # Numbers: When you encounter a number, push it onto the top of the stack.
        else:
            # Push the number onto the stack
            stack.append(int(token)) # If not an operator, push the number 


    return stack[0] # The final result is the only item left on the stack


# Example input: `["2", "1", "+", "3", "*"]`

# 1. **"2":** Push 2 onto the stack: `[2]`
# 2. **"1":** Push 1 onto the stack: `[2, 1]`
# 3. **"+":** 
#    - Pop 1 and 2 off the stack.
#    - Calculate 2 + 1 = 3
#    - Push 3 onto the stack: `[3]`
# 4. **"3":** Push 3 onto the stack: `[3, 3]`
# 5. **"*":**
#    - Pop 3 and 3 off the stack.
#    - Calculate 3 * 3 = 9
#    - Push 9 onto the stack: `[9]`

# Finally, the code returns `stack[0]`, which is `9`.
