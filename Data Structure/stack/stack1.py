# Problem Description:

# You're given a string that might contain these parentheses: '(', ')', '{', '}', '[', ']'. 
# Your task is to figure out if the order of these parentheses is correct.

# Example:

# Input: "()[]{}"
# Output: True (The parentheses match correctly)
# Input: "(]"
# Output: False (Mismatched parentheses)

def is_valid(st):
    stack = []
    mapping = { 
        ")" : "(",
        "]" : "[",
        "}" : "{"       
    }

    # check each char in the string
    for char in st:

    # not stack means to check if the stack is empty
    # also check if mapping char equlas to the first thing on the stack(the top of the stack)
    
    # we only run when char in the mapping, if not in mapping, then just return not stack
        if char in mapping:
            if not stack or mapping[char] != stack.pop():
                return False
            else:
                stack.append(char)
            
    return not stack
