# stack = []  # Create an empty stack

# stack.append(10)  
# stack.append(20)  
# stack.append(30)  

# print(stack.pop())  # Output: 30
# print(stack[-1])    # Output: 20 (peek) 
# stack[-1] uses negative indexing to access the last element in the list. 
# Since the last (and now top) element is 20, this is the value printed.


# print(len(stack))   # Output: 2 because 30 is poped




class Stack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if not self.items:  # If stack is empty, then raise Exception. 
            # Empty means self.items is False,
            raise Exception("Stack is empty")
        return self.items.pop()

    def peek(self):
        if not self.items:
            raise Exception("Stack is empty")
        return self.items[-1]

    def isEmpty(self):
        return self.items == []

    def size(self):
        return len(self.items)


# Create a stack and add elements
my_stack = Stack() 
my_stack.push(5)
my_stack.push(10)
my_stack.push("hello") 

# Remove and retrieve the top item
top_item = my_stack.pop()
print("Popped item:", top_item)  # Output: Popped item: hello

# Check the current top item 
current_top = my_stack.peek()
print("Current top item:", current_top) # Output: Current top item: 10

# Get the stack size
stack_size = my_stack.size()
print("Stack size:", stack_size)  # Output: Stack size: 2

