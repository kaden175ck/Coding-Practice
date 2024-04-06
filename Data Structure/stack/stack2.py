# Problem Description:

# Reverse Polish Notation (RPN) is a different way to write math expressions. For example:

# Standard: 10 + 2
# RPN: 10 2 +
# Your job is to get an input in RPN format and calculate the result.

# Example

# Input: ["2", "1", "+", "3", "*"]
# Output: 9 (Explanation: ((2 + 1) * 3) = 9)

def evaluate_rpn(expression):
    stack = []
    for token in expression:
        if token in "+-*/":
            # Pop the top two elements
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
        else:
            # Push the number onto the stack
            stack.append(int(token))
    return stack[0]