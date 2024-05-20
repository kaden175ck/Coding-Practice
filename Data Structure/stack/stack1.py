# Problem Description:

# You're given a string that might contain these parentheses: '(', ')', '{', '}', '[', ']'. 
# Your task is to figure out if the order of these parentheses is correct.

# Example:

# Input: "()[]{}"
# Output: True (The parentheses match correctly)
# Input: "(]"
# Output: False (Mismatched parentheses)


# How this works with the string `"{[()]}"`:

# 1. `{` is pushed onto the stack.
# 2. `[` is pushed onto the stack.
# 3. `(` is pushed onto the stack.
# 4. `)` is encountered. The dictionary is used to find its corresponding opening symbol (`(`).  The top of the stack is `(`, so it's popped.
# 5. `]` is encountered. The dictionary is used to find its corresponding opening symbol (`[`).  The top of the stack is `[`, so it's popped.
# 6. `}` is encountered. The dictionary is used to find its corresponding opening symbol (`{`). The top of the stack is `{`, so it's popped.
# 7. The stack is now empty. Since the stack is empty at the end, the string is considered valid.





def is_valid(st):
    stack = []    # only store opening symbols, everytime we encounter open, we push to top of stack
    mapping = {   # key value pairs, 
        # closing is key. opening is the value.
        ")" : "(",
        "]" : "[",
        "}" : "{"       
    }

    # check each char in the string
    for char in st:

    # not stack means to check if the stack is empty
    # if stack is empty stack will give false, and not stack will give true
    
    
        if char in mapping: 
            # we only run when legal char is entered(everything in mapping)
            # if not in mapping, then just return the last line: return not stack


            if not stack or mapping[char] != stack.pop():
                # if stack is empty(not stack=true) 
                # or mapping[char] (access the value of the key, corresponding open symbols) doesn't match the the top of the stack(the open symbols)
                return False
                # if stack is empty or the current closing sumbols can't match the top stack return False
            

            else:
                # if stack is not empty and it can match
                stack.append(char)
                # If the current character is an opening symbol,
                # The stack will only store openning symbols
                # (present in the mapping dictionary), it's pushed (appended) onto the stack.
            
    return not stack  
    # return true if the stack is empty
    # return False if the stack is not empty, means some symbols are not closed
    # You could achieve the same result by returning len(stack) == 0 



